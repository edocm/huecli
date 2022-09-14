package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
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
		fmt.Println("list called")
	},
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
