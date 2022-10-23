package main

import (
	"encoding/json"
	"fmt"

	"github.com/edocm/huecli/cmd"
)

func main() {
	cmd.Execute()

	var test map[string]any

	testJson := `{"birds":{"pigeon":"likes to perch on rocks","eagle":"bird of prey"},"animals":"none"}`

	json.Unmarshal([]byte(testJson), &test)

	println(test)

	for key := range test {
		fmt.Println(fmt.Sprint(test[key]))
	}

}
