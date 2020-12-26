package main

import (
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

	var seedDataJSON string = config.ProjectRootPath + "/seed/sample_input_data.json"

	// Parse JSON file into team struct
	if jsonParseErr := storage.ParseJSON(seedDataJSON, &itgixTeams); jsonParseErr != nil {
		log.Fatalf("Err: %v", jsonParseErr) // exit if json cannot be parsed
	}

	// Initial run - populate KV store with team and seed data
	for _, team := range itgixTeams.Teams {
		// Update engineers in Consul KV Store
		for _, engineer := range team.Engineers {
			// Build KV path - Team/Engineer
			keyPath := team.Name + "/" + engineer.GetName()
			// Build KV contents for each engineer
			engDef := "Team: " + team.Name + "\n" + engineer.GetLevel() + "\n" + engineer.GetTracking()

			// Put in Consul KV
			if setSeedKvPairErr := storage.SetConsulKV(kv, keyPath, engDef); setSeedKvPairErr != nil {
				log.Fatalf("Err: %v", setSeedKvPairErr)
			}

			// TODO: figure out how to get the new months tracking here
			newMonthTracking := "September 2020 - 142 hrs"

			// Take current tracking and append latest month
			latestVal, getKvPairErr := storage.GetConsulKV(kv, keyPath)
			if getKvPairErr != nil {
				log.Errorf("Err: %v", getKvPairErr)
			}

			// Set both in Consul KV
			updatedVal := latestVal + "\n" + "  - " + newMonthTracking
			if updateKvPairErr := storage.SetConsulKV(kv, keyPath, updatedVal); updateKvPairErr != nil {
				log.Fatalf("Err: %v", updateKvPairErr)
			}

			// TODO: Add calculated utilization % for each month and average overall
			// ...
			// Tracking:
			//		- September 2020 - 142 hrs - 87%
			// Average Utilization: 83%
		}
	}
}
