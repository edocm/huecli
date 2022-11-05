package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/edocm/huecli/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

var id string

var roomCmd = &cobra.Command{
	Use:   "room",
	Short: "Control a room via hue api.",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var onCmd = &cobra.Command{
	Use:   "on",
	Short: "Turn the lights in a specific room on.",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("on called")
	},
}

var offCmd = &cobra.Command{
	Use:   "off",
	Short: "Turn the lights in a specific room off.",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("off called")
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List up all rooms.",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		getRoomList()
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

	onCmd.Flags().StringVar(&id, "id", "", "Determine which room should be turned on.")
	onCmd.MarkFlagRequired("id")

	offCmd.Flags().StringVar(&id, "id", "", "Determine which room should be turned off.")
	offCmd.MarkFlagRequired("id")

}

func getRoomId(roomName string) {

}

func getRoomList() {
	var roomListResponse RoomListResponse

	res, err := api.Request("GET", "https://"+viper.GetString("bridge")+"/clip/v2/resource/room", nil)
	if err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(res, &roomListResponse); err != nil {
		log.Fatal(err)
	}
	for _, data := range roomListResponse.Data {
		fmt.Println(data.Metadata.Name)
	}
}

func turnRoomOn(roomId string) {

}

func turnRoomOff(roomId string) {

}
