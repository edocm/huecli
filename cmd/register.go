package cmd

import (
	"encoding/json"
	"log"

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
		registerApp()
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)

	registerCmd.Flags().StringVarP(&bridge, "bridge", "b", "", "Determine which room should be turned on.")
	registerCmd.MarkFlagRequired("bridge")
}

func registerApp() {

	type registerMessage struct {
		devicetype        string
		generateclientkey bool
	}

	requestBody, err := json.Marshal(registerMessage{
		devicetype:        "huecli",
		generateclientkey: true,
	})
	if err != nil {
		log.Fatal(err)
	}

	res := api.POST("https://"+bridge+"/api", requestBody)

}
