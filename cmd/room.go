package cmd

import (
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/edocm/huecli/errors"
	"github.com/edocm/huecli/hue"
	"github.com/spf13/cobra"
)

var (
	name       string
	brightness int
	color      string
)

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
		var err error
		if cmd.Flags().Changed("color") {
			err = hue.ChangeRoomLightColor(name, color, brightness)
		} else {
			err = hue.TurnRoomLightOn(name, brightness)
		}
		if err != nil {
			if err == errors.ErrRoomNotAvailable {
				printRoomNotAvailable(name)
			} else {
				log.Fatal(err)
			}
		} else {
			printSuccessMessageOn(name, brightness, color)
		}
	},
}

var offCmd = &cobra.Command{
	Use:   "off",
	Short: "Turn the lights in a specific room off",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		if err := hue.TurnRoomLightOff(name); err != nil {
			if err == errors.ErrRoomNotAvailable {
				printRoomNotAvailable(name)
			} else {
				log.Fatal(err)
			}
		} else {
			printSuccessMessageOff(name)
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
	onCmd.Flags().IntVarP(&brightness, "brightness", "b", 50, "Determine how bright the room should be")
	onCmd.Flags().StringVarP(&color, "color", "c", "", "Determine the color")
	onCmd.MarkFlagRequired("name")

	offCmd.Flags().StringVarP(&name, "name", "n", "", "Determine which room should be turned off")
	offCmd.MarkFlagRequired("name")
}

func printRoomNotAvailable(roomName string) {
	fmt.Printf("The room %s is not available. Please try again. \n", roomName)
}

func printSuccessMessageOn(roomName string, brightness int, color string) {
	if color == "" {
		fmt.Printf("The lights in room %s are on and brightness is set to %d \n", roomName, brightness)
	} else {
		fmt.Printf("The lights in room %s are on. The color is set to %s and brightness to %d \n", roomName, color, brightness)
	}
}

func printSuccessMessageOff(roomName string) {
	fmt.Printf("The lights in room %s are off.  \n", roomName)
}

func printRoomList(roomList map[string]string) {
	fmt.Println("These are the available rooms:")
	for roomName := range roomList {
		fmt.Println("-", roomName)
	}
}
