package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listflag bool

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

func init() {
	roomCmd.AddCommand(onCmd)
	roomCmd.AddCommand(offCmd)

	roomCmd.Flags().BoolVarP(&listflag, "list", "l", true, "Determine if the user wants to list all available rooms.")
}
