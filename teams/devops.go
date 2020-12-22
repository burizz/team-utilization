package teams

import (
	"fmt"
)

// ItgixTeams - list of devops teams within company
type ItgixTeams struct {
	Teams []Cell `json:"itgixteams"`
}

// Cell - list of Cells, their function and team members
type Cell struct {
	Cell      string           `json:"team"`
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
// GetTeam method - returns a list of all team members
func (t Cell) GetTeam(teamName string) []DevOpsEngineer {
	return t.Engineers
}

// GetEngineer method - returns all details of an individual engineer
func (e DevOpsEngineer) GetEngineer(engineerName string) string {
	// TODO: return all fields properly
	return fmt.Sprintf("%v %v %v", e.Firstname, e.Lastname, e.Level)
}
