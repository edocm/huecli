package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/edocm/huecli/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/exp/maps"
)

var name string

var roomCmd = &cobra.Command{
	Use:   "room",
	Short: "Control a room via hue api",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var onCmd = &cobra.Command{
	Use:   "on",
	Short: "Turn the lights in a specific room on",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		changeRoomLightStatus(name, true)
	},
}

var offCmd = &cobra.Command{
	Use:   "off",
	Short: "Turn the lights in a specific room off",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		changeRoomLightStatus(name, false)
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List up all rooms",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		printRoomList()
	},
}

type RoomListResponse struct {
	Errors []struct {
	} `json:"errors"`
	Data []struct {
		Id       string `json:"id"`
		IdV1     string `json:"id_v1"`
		Children []struct {
			RId   string `json:"rid"`
			RType string `json:"rtype"`
		} `json:"children"`
		Services []struct {
			RId   string `json:"rid"`
			RType string `json:"rtype"`
		} `json:"services"`
		Metadata struct {
			Name      string `json:"name"`
			ArcheType string `json:"archetype"`
		} `json:"metadata"`
		Type string `json:"type"`
	} `json:"data"`
}

type GroupLightMessage struct {
	On OnProperty `json:"on"`
}

type OnProperty struct {
	On bool `json:"on"`
}

func init() {
	rootCmd.AddCommand(roomCmd)

	roomCmd.AddCommand(onCmd)
	roomCmd.AddCommand(offCmd)
	roomCmd.AddCommand(listCmd)

	onCmd.Flags().StringVarP(&name, "name", "n", "", "Determine which room should be turned on")
	onCmd.MarkFlagRequired("name")

	offCmd.Flags().StringVarP(&name, "name", "n", "", "Determine which room should be turned off")
	offCmd.MarkFlagRequired("name")
}

func getRoomList() (map[string]string, error) {
	var roomListResponse RoomListResponse
	roomList := make(map[string]string)

	log.Debug("Request roomList")
	res, err := api.Request(http.MethodGet, "https://"+viper.GetString("bridge")+"/clip/v2/resource/room", nil)
	if err != nil {
		return nil, fmt.Errorf("error while request room list: %v", err)
	}
	if err := json.Unmarshal(res, &roomListResponse); err != nil {
		return nil, fmt.Errorf("error while parsing room list response: %v", err)
	}
	for _, room := range roomListResponse.Data {
		for _, service := range room.Services {
			if service.RType == "grouped_light" {
				roomList[room.Metadata.Name] = service.RId
			}
		}
	}
	log.Debugf("Successfully requested roomList: %v", maps.Keys(roomList))
	return roomList, nil
}

func changeRoomLightStatus(roomName string, status bool) {
	roomList, err := getRoomList()
	if err != nil {
		log.Fatal(err)
	}
	roomId, ok := roomList[roomName]
	if ok {
		requestBody, err := json.Marshal(GroupLightMessage{
			On: OnProperty{
				On: status,
			},
		})
		if err != nil {
			log.Fatal(fmt.Errorf("error while build request for changing roomlight status: %v", err))
		}
		_, err = api.Request(http.MethodPut, "https://"+viper.GetString("bridge")+"/clip/v2/resource/grouped_light/"+roomId, requestBody)
		if err != nil {
			log.Fatal(fmt.Errorf("error while request for changing roomlight status: %v", err))
		}
		printSuccessMessage(roomName, status)
	} else {
		printRoomNotAvailable(roomName)
	}
}

//TODO: add package for responses which is turned off by config

func printRoomNotAvailable(roomName string) {
	fmt.Printf("The room %s is not available. Please try again. \n", roomName)
}

func printSuccessMessage(roomName string, lightsOn bool) {
	if lightsOn {
		fmt.Printf("The lights in room %s are on. \n", roomName)
	} else {
		fmt.Printf("The lights in room %s are off.  \n", roomName)
	}
}

func printRoomList() {
	roomList, err := getRoomList()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("These are the available rooms:")
	for roomName := range roomList {
		fmt.Println("-", roomName)
	}
}
