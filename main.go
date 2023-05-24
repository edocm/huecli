package main

import (
	"os"

	"github.com/charmbracelet/log"

	"github.com/edocm/huecli/cmd"
	"github.com/edocm/huecli/config"
)

var Logger *log.Logger // TODO: research go global logging

func init() {
	config.LoadConfig()
}

func main() {
	Logger = log.NewWithOptions(os.Stderr, log.Options{
		Level: log.DebugLevel,
	})

	cmd.Execute()
}
