package main

import (
	"encoding/json"
	"fmt"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"

	"github.com/burizz/team-utilization/config"
	"github.com/burizz/team-utilization/storage"
	"github.com/burizz/team-utilization/teams"
)

func main() {
	// Load .env config from project root
	if loadDotEnvErr := godotenv.Load(config.ProjectRootPath + "/.env"); loadDotEnvErr != nil {
		log.Fatalf("Error loading config: %v", loadDotEnvErr)
	} else {
		// Set loglevel, format and output stream
		config.LoggingConfig()

		log.Debugf(".env file loaded successfully")
	}

	// Create Consul client
	kv, consulInitClientErr := config.ConsulConfig()
	if consulInitClientErr != nil {
		log.Fatalf("Consul config: %v", consulInitClientErr)
	}

	var itgixTeams teams.ItgixTeams
	var teamVar string = "Team: "

	var seedDataJSON string = config.ProjectRootPath + "/seed/sample_input_data.json"

	// TODO: Move JSON func into separate package and leave just the function call
	// Parse JSON file into team struct
	if jsonParseErr := storage.ParseJSON(seedDataJSON, &itgixTeams); jsonParseErr != nil {
		log.Fatalf("Err: %v", jsonParseErr) // exit if json cannot be parsed
	}

	// TODO: Not sure if I need this, probably makes more sense to parse individual parts of the individual teams
	teamValue, err := json.Marshal(itgixTeams)
	if err != nil {
		log.Errorf("Err : %v", err)
	}

	// TODO: Parse each team into separate consul kv pair
	// TODO: Parse each team member into separate kv / pair
	for _, value := range itgixTeams.Teams {
		fmt.Println(value)
	}

	// Put a new KV pair
	if setKvPairErr := storage.SetConsulKV(kv, teamVar, teamValue); setKvPairErr != nil {
		log.Fatalf("Err: %v", setKvPairErr) // exit if Consul KV pair cannot be set
	}

	// Lookup KV pair in Consul
	if _, getKvPairErr := storage.GetConsulKV(kv, teamVar); getKvPairErr != nil {
		log.Errorf("Err: %v", getKvPairErr)
	}
}
