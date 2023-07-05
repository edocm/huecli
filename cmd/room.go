package cmd

import (
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/edocm/huecli/errors"
	"github.com/edocm/huecli/hue"
	"github.com/spf13/cobra"
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
		if err := hue.ChangeRoomLightStatus(name, true); err != nil {
			if err == errors.ErrRoomNotAvailable {
				printRoomNotAvailable(name)
			} else {
				log.Fatal(err)
			}
		} else {
			printSuccessMessage(name, true)
		}
	},
}

var offCmd = &cobra.Command{
	Use:   "off",
	Short: "Turn the lights in a specific room off",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		if err := hue.ChangeRoomLightStatus(name, false); err != nil {
			if err == errors.ErrRoomNotAvailable {
				printRoomNotAvailable(name)
			} else {
				log.Fatal(err)
			}
		} else {
			printSuccessMessage(name, false)
		}
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List up all rooms",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		roomList, err := hue.GetRoomList()
		if err != nil {
			log.Fatal(err)
		}
		printRoomList(roomList)
	},
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

func printRoomList(roomList map[string]string) {
	fmt.Println("These are the available rooms:")
	for roomName := range roomList {
		fmt.Println("-", roomName)
	}
}
