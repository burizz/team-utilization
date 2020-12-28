package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"

	"github.com/burizz/team-utilization/config"
	"github.com/burizz/team-utilization/storage"
	"github.com/burizz/team-utilization/teams"
)

func main() {
	envType := os.Getenv("ENV_TYPE")

	// Load .env config from project root and see import files
	var envFile string = config.ProjectRootPath + "/.env"
	var seedDataJSON string = config.ProjectRootPath + "/seed/initial_seed_data.json"
	var excelReport string = config.ProjectRootPath + "/seed/detailed_report.xlsx"

	var pEnvFile *string = &envFile
	var pSeedDataJSON *string = &seedDataJSON
	var pExcelReport *string = &excelReport

	if envType == "Docker" || envType == "DOCKER" {
		// Configure local paths for Docker env
		*pEnvFile = ".env"
		*pSeedDataJSON = "seed/initial_seed_data.json"
		*pExcelReport = "seed/detailed_report.xlsx"
	}

	if loadDotEnvErr := godotenv.Load(*pEnvFile); loadDotEnvErr != nil {
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

	// Parse JSON file into team struct
	if jsonParseErr := storage.ParseJSON(*pSeedDataJSON, &itgixTeams); jsonParseErr != nil {
		log.Fatalf("Parse JSON Err: %v", jsonParseErr) // exit if json cannot be parsed
	}

	// Parse tracked hours from Excel
	trackedTotal, excelParseErr := storage.ParseTrackingFromExcel(*pExcelReport)
	if excelParseErr != nil {
		log.Errorf("Parse Excel Err: %v", excelParseErr)
	}

	// TODO: Test this
	// Parse month and year from tracking report
	trackingMonth, trackingYear, excelParseErr := storage.ParsePeriodFromExcel(*pExcelReport)
	if excelParseErr != nil {
		log.Errorf("Parse Excel Err: %v", excelParseErr)
	}

	// Calculate % tracking utilization - 100% is 160 hrs
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
