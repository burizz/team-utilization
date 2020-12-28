package main

import (
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
	// TODO: Better place for configuring seed file path
	var seedDataJSON string = config.ProjectRootPath + "/seed/initial_seed_data.json"
	var excelReport string = config.ProjectRootPath + "/seed/detailed_report.xlsx"

	// Parse JSON file into team struct
	if jsonParseErr := storage.ParseJSON(seedDataJSON, &itgixTeams); jsonParseErr != nil {
		log.Fatalf("Parse JSON Err: %v", jsonParseErr) // exit if json cannot be parsed
	}

	trackedTotal, excelParseErr := storage.ParseTrackingFromExcel(excelReport)
	if excelParseErr != nil {
		log.Errorf("Parse Excel Err: %v", excelParseErr)
	}

	// TODO: Test this
	trackingMonth, trackingYear, excelParseErr := storage.ParsePeriodFromExcel(excelReport)
	if excelParseErr != nil {
		log.Errorf("Parse Excel Err: %v", excelParseErr)
	}

	percentUtil, trackingCalcErr := storage.CalculateTrackingPercent(trackedTotal)
	if trackingCalcErr != nil {
		log.Errorf("Calculate Tracking Err: %v", trackingCalcErr)
	}

	// TODO: Test this
	newMonthTracking := fmt.Sprintf("%v %v - %v hrs - %v", trackingMonth, trackingYear, trackedTotal, percentUtil)

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
		}
	}
}
