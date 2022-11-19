package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/edocm/huecli/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
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
		fmt.Println("on called")
	},
}

var offCmd = &cobra.Command{
	Use:   "off",
	Short: "Turn the lights in a specific room off",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("off called")
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

	res, err := api.Request("GET", "https://"+viper.GetString("bridge")+"/clip/v2/resource/room", nil)
	if err != nil {
		return nil, fmt.Errorf("error while request room list: %v", err)
	}
	if err := json.Unmarshal(res, &roomListResponse); err != nil {
		return nil, fmt.Errorf("error while parsing room list response: %v", err)
	}
	for _, data := range roomListResponse.Data {
		roomList[data.Metadata.Name] = data.Id
	}
	return roomList, nil
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

func turnRoomOn(roomName string) {

}

func turnRoomOff(roomName string) {

}
