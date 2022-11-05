package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/edocm/huecli/api"
	"github.com/edocm/huecli/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	bridge string
)

type ErrorMessage struct {
	Error struct {
		Type        int    `json:"type" binding:"required"`
		Address     string `json:"address" binding:"required"`
		Description string `json:"description" binding:"required"`
	} `json:"error" binding:"required"`
}

type SuccessMessage struct {
	Success struct {
		Username  string `json:"username" binding:"required"`
		Clientkey string `json:"clientkey" binding:"required"`
	} `json:"success" binding:"required"`
}

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "A brief description of your command",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		if config.Exists {
			fmt.Println("Huecli is already registered.")
			return
		}
		pressButton()
		registerApp()
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)

	registerCmd.Flags().StringVarP(&bridge, "bridge", "b", "", "Determine which room should be turned on.")
	registerCmd.MarkFlagRequired("bridge")
}

// TODO: add logging with log libary and change see where log fatal is useful

func registerApp() {

	type RegisterMessage struct {
		Devicetype        string `json:"devicetype"`
		Generateclientkey bool   `json:"generateclientkey"`
	}

	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	requestBody, err := json.Marshal(RegisterMessage{
		Devicetype:        "huecli#" + hostname,
		Generateclientkey: true,
	})
	if err != nil {
		log.Fatal(err)
	}

	var successMessage SuccessMessage
	var errorMessage ErrorMessage

	res, err := api.Request("POST", "https://"+bridge+"/api", requestBody)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal([]byte(strings.Trim(string(res), "[]")), &successMessage); (err != nil || successMessage == SuccessMessage{}) {
		if err := json.Unmarshal([]byte(strings.Trim(string(res), "[]")), &errorMessage); (err != nil || errorMessage == ErrorMessage{}) {
			log.Fatal(err)
		}
		if errorMessage.Error.Type == 101 {
			fmt.Println("You were too slow.")
			return
		} else {
			fmt.Println("An unexpected error occurred.")
			return
		}
	}

	viper.Set("bridge", bridge)
	viper.Set("username", successMessage.Success.Username)
	viper.Set("clientkey", successMessage.Success.Clientkey)
	if err := viper.WriteConfigAs("./config.yaml"); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Huecli is registered successful.")
}

func pressButton() {
	fmt.Println("Please press the button on your bridge. You have 10 seconds.")
	time.Sleep(10 * time.Second)
}
