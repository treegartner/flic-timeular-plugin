package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/BurntSushi/toml"
	"github.com/treegartner/timeular-api-client"
)

type flicPluginConfig struct {
	PluginName        string       `json:"pluginName"`
	PluginDescription string       `json:"pluginDescription"`
	PluginReadMore    string       `json:"pluginReadMore"`
	ProtocolVersion   int          `json:"protocolVersion"`
	FlicActions       []flicAction `json:"actions"`
}

type flicAction struct {
	ActionName string `json:"actionName"`
	FileName   string `json:"fileName"`
}

type scriptSetup struct {
	Param1 string
	Action string
	Param2 string
	Name   string
}

type configuration struct {
	TimeularKey     string
	TimeularSecret  string
	TimeularBaseURL string
}

func main() {
	name := flag.String("name", "", "Activity Name (Case Sensitive!) to start | stop | toggle")
	action := flag.String("action", "", "Action to process: start | stop | toggle; use 'setup' to create actions out of existing activities")
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
	case "setup":
		activities, err := api.ReadActivities()
		if err != nil {
			log.Fatal("Error fetching activities: ", err)
		}
		createShellScripts(activities, []string{"toggle", "start"})
	}
}

// Create shell script files out of activities
func createShellScripts(activities *timeular.Activities, actions []string) error {
	tmplStr := `#!/bin/bash
WORKINGDIR=$( dirname "${BASH_SOURCE[0]}" )
test=$("$WORKINGDIR/flic-timeular-plugin" -config "$WORKINGDIR/config.toml" {{.Param1}} {{.Action}} {{.Param2}} {{.Name}})`
	scriptActs := []scriptSetup{
		{
			Param1: "-action",
			Action: "stop",
			Param2: "-name",
			Name:   "Activity",
		},
	}
	for _, activity := range activities.Activities {
		for _, action := range actions {
			scriptAct := scriptSetup{
				Param1: "-action",
				Action: action,
				Param2: "-name",
				Name:   activity.Name,
			}
			scriptActs = append(scriptActs, scriptAct)
		}
	}

	// create shell scripts AND flicPlugin actions
	flicActs := []flicAction{}
	for _, act := range scriptActs {
		filename := strings.Join([]string{act.Action, act.Name, ".sh"}, "")
		tmpl, err := template.New(filename).Parse(tmplStr)
		if err != nil {
			return fmt.Errorf("Error parsing template: %v", err)
		}

		w, err := os.Create(filename)
		if err != nil {
			return fmt.Errorf("Error writing shellscript %s: %v\n: ", filename, err)
		}
		err = tmpl.Execute(w, act)
		if err != nil {
			panic(err)
		}

		// flicActions
		flicAct := flicAction{
			ActionName: strings.Join([]string{act.Action, act.Name}, " "),
			FileName:   filename,
		}
		flicActs = append(flicActs, flicAct)
	}

	// create flic plugin configuration file config.json
	pluginConfig := flicPluginConfig{
		PluginName:        "Timeular",
		PluginDescription: "Track time with timeular",
		PluginReadMore:    "This action lets track activities from timeular with the push of a button.",
		ProtocolVersion:   1,
		FlicActions:       flicActs,
	}

	file, _ := json.MarshalIndent(pluginConfig, "", " ")
	err := ioutil.WriteFile("config.json", file, 0644)
	if err != nil {
		return fmt.Errorf("Error writing config.json: %v", err)
	}

	return nil
}
