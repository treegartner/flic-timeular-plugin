package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/treegartner/timeular-api-client"
)

type configuration struct {
	TimeularKey     string
	TimeularSecret  string
	TimeularBaseURL string
}

func main() {
	name := flag.String("name", "Meeting", "Activity Name (Case Sensitive!) to start | stop | toggle")
	action := flag.String("action", "toggle", "Action to process: start | stop | toggle")
	confFile := flag.String("config", "config.toml", "Configuration file")
	flag.Parse()

	var conf configuration
	file, _ := os.Open(*confFile)
	defer file.Close()
	_, err := toml.DecodeFile(*confFile, &conf)
	if err != nil {
		fmt.Println("error:", err)
	}

	api, err := timeular.NewAPI(
		conf.TimeularBaseURL,
		conf.TimeularKey,
		conf.TimeularSecret,
	)
	if err != nil {
		log.Fatal("Error while creating API: ", err)
	}

	switch strings.ToLower(*action) {
	case "start":
		_, err = api.StartTracking(*name)
		if err != nil {
			log.Fatal("Error Toggle: ", err)
		}
	case "stop":
		_, err = api.StopTracking()
		if err != nil {
			log.Fatal("Error Stopping Activity: ", err)
		}
	case "toggle":
		err = api.ToggleTracking(*name)
		if err != nil {
			log.Fatal("Error Toggle: ", err)
		}
	}
}
