package hue

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/edocm/huecli/api"
	"github.com/edocm/huecli/errors"
	"github.com/spf13/viper"
	"golang.org/x/exp/maps"
)

type RoomListResponse struct {
	Errors []struct {
	} `json:"errors"`
	Data []struct {
		Id       string `json:"id"`
		IdV1     string `json:"id_v1"`
		Children []struct {
			RId   string `json:"rid"`
			RType string `json:"rtype"`
		} `json:"children"`
		Services []struct {
			RId   string `json:"rid"`
			RType string `json:"rtype"`
		} `json:"services"`
		Metadata struct {
			Name      string `json:"name"`
			ArcheType string `json:"archetype"`
		} `json:"metadata"`
		Type string `json:"type"`
	} `json:"data"`
}

func GetRoomList() (map[string]string, error) {
	var roomListResponse RoomListResponse
	roomList := make(map[string]string)

	log.Debug("Request roomList")
	res, err := api.Request(http.MethodGet, "https://"+viper.GetString("bridge")+"/clip/v2/resource/room", nil)
	if err != nil {
		return nil, fmt.Errorf("error while request room list: %v", err)
	}
	if err := json.Unmarshal(res, &roomListResponse); err != nil {
		return nil, fmt.Errorf("error while parsing room list response: %v", err)
	}
	for _, room := range roomListResponse.Data {
		for _, service := range room.Services {
			if service.RType == "grouped_light" {
				roomList[room.Metadata.Name] = service.RId
				log.Infof("id: %v", service.RId)
			}
		}
	}
	log.Debugf("Successfully requested roomList: %v", maps.Keys(roomList))
	return roomList, nil
}

func TurnRoomLightOn(roomName string, brightness int) error {
	roomId, err := getRoomId(roomName)
	if err != nil {
		return fmt.Errorf("error while request roomId: %v", err)
	}

	message := GroupLightOnMessage{
		On: OnProperty{
			On: true,
		},
		Dimming: DimmingProperty{
			Brightness: brightness,
		},
	}

	if err := ChangeGroupedLight(roomId, message); err != nil {
		return fmt.Errorf("error while request for changing roomlight status: %v", err)
	}

	return nil
}

func TurnRoomLightOff(roomName string) error {
	roomId, err := getRoomId(roomName)
	if err != nil {
		return fmt.Errorf("error while request roomId: %v", err)
	}

	message := GroupLightOffMessage{
		On: OnProperty{
			On: false,
		},
	}

	if err := ChangeGroupedLight(roomId, message); err != nil {
		return fmt.Errorf("error while request for changing roomlight status: %v", err)
	}

	return nil
}

func ChangeRoomLightColor(roomName string, color string, brightness int) error {
	// TODO: implement color change
	log.Info("Change Color!")
	return nil
}

func getRoomId(roomName string) (string, error) {
	roomList, err := GetRoomList()
	if err != nil {
		return "", fmt.Errorf("error while requesting room list: %v", err)
	}
	roomId, ok := roomList[roomName]
	if ok {
		return roomId, nil
	} else {
		return "", errors.ErrRoomNotAvailable
	}
}
