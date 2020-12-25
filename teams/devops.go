package teams

import (
	"encoding/json"
	"fmt"

	log "github.com/sirupsen/logrus"
)

// ItgixTeams - list of devops teams within company
type ItgixTeams struct {
	Teams []Team `json:"itgixteams"`
}

// Team - list of Cells, their function and team members
type Team struct {
	Name      string           `json:"team"`
	Engineers []DevOpsEngineer `json:"engineers"`
}

// DevOpsEngineer - engineer definition
type DevOpsEngineer struct {
	Firstname     string `json:"firstname"`
	Lastname      string `json:"lastname"`
	Level         string `json:"level"`
	TrackedHours  int    `json:"trackedHours"`
	TrackingMonth string `json:"trackingMonth"`
	TrackingYear  string `json:"trackingYear"`
}

// TODO: Test methods !
// GetTeam method - returns a slice of all team members - type []DevOpsEngineer
func (t Team) GetTeam(teamName string) []DevOpsEngineer {
	return t.Engineers
}

// GetTeamMarshalled - returns a []byte slice of all team members
func (t Team) GetTeamMarshalled(teamName string) []byte {
	parsedEngineers, jsonMarshalErr := json.Marshal(t.Engineers)
	if jsonMarshalErr != nil {
		log.Errorf("Cannot marshal JSON to string %v", jsonMarshalErr)
	}
	return parsedEngineers
}

func (t Team) GetTeamString(teamName string) string {
	parsedEngineers, jsonMarshalErr := json.Marshal(t.Engineers)
	if jsonMarshalErr != nil {
		log.Errorf("Cannot marshal JSON to string %v", jsonMarshalErr)
	}

	return string(parsedEngineers)
}

// GetEngineer method - returns all details of an individual engineer
func (e DevOpsEngineer) GetEngineer(engineerName string) string {
	// TODO: Fix this
	return fmt.Sprintf("%v %v %v", e.Firstname, e.Lastname, e.Level)
}

func (e DevOpsEngineer) GetName(engineerName string) string {
	// TODO: Fix this
	return fmt.Sprintf("%v %v", e.Firstname, e.Lastname)
}

func (e DevOpsEngineer) GetTracking(engineerName string) []byte {
	// TODO: Fix this
	return fmt.Sprintf("Hours: %v - %v %v", e.TrackedHours, e.TrackingMonth, e.TrackingYear)
}
