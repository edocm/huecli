package hue

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/edocm/huecli/api"
	"github.com/spf13/viper"
)

type GroupLightOnMessage struct {
	On      OnProperty      `json:"on"`
	Dimming DimmingProperty `json:"dimming"`
}

type GroupLightOffMessage struct {
	On OnProperty `json:"on"`
}

type GroupLightColorMessage struct {
	Color ColorProperty `json:"color"`
}

type OnProperty struct {
	On bool `json:"on"`
}

type ColorProperty struct {
	Xy struct {
		X float32 `json:"x"`
		Y float32 `json:"y"`
	} `json:"xy"`
}

type DimmingProperty struct {
	Brightness int `json:"brightness"`
}

func ChangeGroupedLight(id string, groupLightMessage any) error {
	requestBody, err := json.Marshal(groupLightMessage)
	if err != nil {
		return fmt.Errorf("error while parsing requestBody: %v", err)
	}

	log.Infof("requestBody: %v", string(requestBody))

	_, err = api.Request(http.MethodPut, "https://"+viper.GetString("bridge")+"/clip/v2/resource/grouped_light/"+id, requestBody)
	if err != nil {
		return fmt.Errorf("error while sending request to chnage grouped light: %v", err)
	}

	return nil
}
