package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/edocm/huecli/api"
	"github.com/spf13/cobra"
)

var (
	bridge string
)

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "A brief description of your command",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		pressButton()
		registerApp()
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)

	registerCmd.Flags().StringVarP(&bridge, "bridge", "b", "", "Determine which room should be turned on.")
	registerCmd.MarkFlagRequired("bridge")
}

func registerApp() {
	type RegisterMessage struct {
		devicetype        string
		generateclientkey bool
	}

	var registerResponse map[string]any

	requestBody, err := json.Marshal(RegisterMessage{
		devicetype:        "huecli#" + fmt.Sprint(os.Hostname()),
		generateclientkey: true,
	})
	if err != nil {
		log.Fatal(err)
	}

	res := api.POST("https://"+bridge+"/api", requestBody)

	if err := json.Unmarshal(res, &registerResponse); err != nil {
		log.Fatal(err)
	}

	// TODO: evaluate if request was successful and save keys from response in config
}

func pressButton() {
	fmt.Println("Please press the button on your bridge.")
	timer := time.NewTimer(30 * time.Second)
	<-timer.C
}
