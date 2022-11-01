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
}
