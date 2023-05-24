package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/log"
	"github.com/edocm/huecli/api"
	"github.com/edocm/huecli/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	bridge         string
	shouldResponse bool
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

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Change huecli configuration",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Flags().Changed("responses") {
			setResponses()
		}
	},
}

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register huecli at your hue bridge",
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
	rootCmd.AddCommand(configCmd)

	configCmd.AddCommand(registerCmd)

	configCmd.Flags().BoolVarP(&shouldResponse, "responses", "r", true, "Determine if huecli should give you responses.")

	registerCmd.Flags().StringVarP(&bridge, "bridge", "b", "", "Determine which bridge should be registered.")
	registerCmd.MarkFlagRequired("bridge")
}

// TODO: add logging with log library and change see where log fatal is useful

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

	res, err := api.Request(http.MethodPost, "https://"+bridge+"/api", requestBody)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal([]byte(strings.Trim(string(res), "[]")), &successMessage); (err != nil || successMessage == SuccessMessage{}) {
		if err := json.Unmarshal([]byte(strings.Trim(string(res), "[]")), &errorMessage); (err != nil || errorMessage == ErrorMessage{}) {
			log.Fatal(err)
		}
		if errorMessage.Error.Type == 101 { //TODO: create error package to define error types
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
	log.Infof("Application successfully registered at %s with username %s", bridge, successMessage.Success.Username)
	fmt.Println("Huecli is registered successful.")
}

func pressButton() {
	fmt.Println("Please press the button on your bridge. You have 10 seconds.")
	time.Sleep(10 * time.Second)
}

func setResponses() {
	viper.Set("responses", shouldResponse)
	if err := viper.WriteConfigAs("./config.yaml"); err != nil {
		log.Fatal(err)
	}
	log.Infof("Config for responses successfully changed to %b", shouldResponse)
}
