package hue

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/edocm/huecli/api"
	"github.com/edocm/huecli/errors"
	"github.com/spf13/viper"
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

type RegisterMessage struct {
	Devicetype        string `json:"devicetype"`
	Generateclientkey bool   `json:"generateclientkey"`
}

func Register(bridge string) error {

	hostname, err := os.Hostname()
	if err != nil {
		return fmt.Errorf("error while retrieving hostname: %v", err)
	}

	requestBody, err := json.Marshal(RegisterMessage{
		Devicetype:        "huecli#" + hostname,
		Generateclientkey: true,
	})
	if err != nil {
		return fmt.Errorf("error while marshling RegisterMessage: %v", err)
	}

	var successMessage SuccessMessage
	var errorMessage ErrorMessage

	res, err := api.Request(http.MethodPost, "https://"+bridge+"/api", requestBody)
	if err != nil {
		return fmt.Errorf("error while sending post to hue bridge: %v", err)
	}

	if err := json.Unmarshal([]byte(strings.Trim(string(res), "[]")), &successMessage); (err != nil || successMessage == SuccessMessage{}) {
		if err := json.Unmarshal([]byte(strings.Trim(string(res), "[]")), &errorMessage); (err != nil || errorMessage == ErrorMessage{}) {
			return err //TODO: what if error == nil or first err while unmarshaling?
		}
		if errorMessage.Error.Type == 101 { //TODO: create error package to define error types
			return errors.ErrNoRegisterValidation
		} else {
			return fmt.Errorf("unexpected error from hue bridge for register request: %v", errorMessage.Error.Type)
		}
	}

	viper.Set("bridge", bridge)
	viper.Set("username", successMessage.Success.Username)
	viper.Set("clientkey", successMessage.Success.Clientkey)
	if err := viper.WriteConfigAs("./config.yaml"); err != nil {
		return fmt.Errorf("error while saving authentication parameters to config.yaml: %v", err)
	}

	log.Debugf("Application successfully registered at %s with username %s", bridge, successMessage.Success.Username)
	return nil
}
