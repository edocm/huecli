package main

import (
	"github.com/edocm/huecli/cmd"
	"github.com/edocm/huecli/config"
)

func init() {
	config.LoadConfig()
}

func main() {
	cmd.Execute()

	// var test map[string]any

	// type Test []struct {
	// 	Error struct {
	// 		Type        int    `json:"type"`
	// 		Address     string `json:"address"`
	// 		Description string `json:"description"`
	// 	} `json:"error"`
	// }

	// var message Test

	// testJson := `[{"error":{"type":101,"address":"","description":"link button not pressed"}}]`

	// json.Unmarshal([]byte(testJson), &message)

	// fmt.Println(message)

	// for key := range test {
	// 	fmt.Println(fmt.Sprint(test[key]))
	// }

}
