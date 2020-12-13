package teams

import (
	"fmt"
)

// ItgixTeam - list of devops engineers within a team
type ItgixTeam struct {
	Team      string           `json:"team"`
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

// GetTeam method - returns a list of all team members
func (t ItgixTeam) GetTeam(teamName string) []DevOpsEngineer {
	return t.Engineers
}

// GetEngineer method - returns all details of an individual engineer
func (e DevOpsEngineer) GetEngineer(engineerName string) string {
	// TODO: return all fields properly
	return fmt.Sprintf("%v %v %v", e.Firstname, e.Lastname, e.Level)
}
