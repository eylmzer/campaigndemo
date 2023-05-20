package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/eylmzer/campaingdemo/cmd/commands"
	"github.com/eylmzer/campaingdemo/pkg/campaingscenario"
	"gopkg.in/yaml.v2"
)

func main() {
	// Create a logger for logging messages
	logger := log.Default()

	data, err := ioutil.ReadFile("commands/commands.yml")
	if err != nil {
		log.Fatalf("Failed to read YAML file: %v", err)
	}

	var c []commands.Command
	var cs = campaingscenario.NewCampaingScenario(logger)
	err = yaml.Unmarshal(data, &c)
	if err != nil {
		log.Fatalf("Failed to unmarshal YAML: %v", err)
	}

	for _, cmd := range c {
		commandString := cmd.Name + " " + strings.Join(cmd.Args, " ")
		output, err := commands.ExecuteCommand(commandString, cs)
		if err != nil {
			logger.Printf("Error executing command '%s': %v", commandString, err)
			continue
		}
		fmt.Println(output)
	}
}
