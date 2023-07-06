package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/edocm/huecli/config"
	"github.com/edocm/huecli/errors"
	"github.com/edocm/huecli/hue"
	"github.com/spf13/cobra"
)

var bridge string

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register huecli at your hue bridge",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		if config.Exists {
			fmt.Println("Huecli is already registered.")
			return
		}
		printPressButton()
		if err := hue.Register(bridge); err != nil {
			if err == errors.ErrNoRegisterValidation {
				fmt.Println("You were too slow.")
			} else {
				log.Fatal(err)
			}
		}
		fmt.Println("Huecli is registered successful.")
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)

	registerCmd.Flags().StringVarP(&bridge, "bridge", "b", "", "Determine which bridge should be registered.")
	registerCmd.MarkFlagRequired("bridge")
}

func printPressButton() {
	fmt.Println("Please press the button on your bridge. You have 10 seconds.")
	time.Sleep(10 * time.Second)
}
