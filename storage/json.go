package storage

import (
	"encoding/json"
	"io/ioutil"

	log "github.com/sirupsen/logrus"

	"github.com/burizz/team-utilization/teams"
)

func ParseJSON(filePath string, team *teams.ItgixTeams) error {
	// TODO: Refactor this to take the JSON from URL where it can be uploaded automatically
	byteValue, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Errorf("parseJSON: Cannot open file: %v - %v", filePath, err)
		return err
	}

	// Unmarshal byte array into team struct
	json.Unmarshal(byteValue, &team)
	if err != nil {
		log.Errorf("parseJSON: Cannot unmarshal JSON to team struct: %v", err)
		return err
	}

	log.Infof("JSON parsed successfully")
	return nil
}
