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

// Team - list of Cells, their function, and team members
type Team struct {
	Name      string           `json:"team"`
	Engineers []DevOpsEngineer `json:"engineers"`
}

// DevOpsEngineer - defines devops engineer
type DevOpsEngineer struct {
	Firstname     string `json:"firstname"`
	Lastname      string `json:"lastname"`
	Level         string `json:"level"`
	TrackedHours  int    `json:"trackedHours"`
	TrackingMonth string `json:"trackingMonth"`
	TrackingYear  string `json:"trackingYear"`
}

// GetTeam method - returns a slice of all team members - type []DevOpsEngineer
func (t Team) GetTeam() []DevOpsEngineer {
	return t.Engineers
}

// GetTeamMarshalled - returns a []byte slice of all team members
func (t Team) GetTeamMarshalled() string {
	parsedEngineers, jsonMarshalErr := json.Marshal(t.Engineers)
	if jsonMarshalErr != nil {
		log.Errorf("Cannot marshal JSON to string %v", jsonMarshalErr)
	}
	return string(parsedEngineers)
}

// GetTeamString - returns a []byte slice of all team members
func (t Team) GetTeamString() string {
	parsedEngineers, jsonMarshalErr := json.Marshal(t.Engineers)
	if jsonMarshalErr != nil {
		log.Errorf("Cannot marshal JSON to string %v", jsonMarshalErr)
	}
	return string(parsedEngineers)
}

// GetEngineer - returns all details of an engineer
func (e DevOpsEngineer) GetEngineer() string {
	return fmt.Sprintf("%v %v %v", e.Firstname, e.Lastname, e.Level)
}

// GetName - returns Firstname and Lastname of an engineer
func (e DevOpsEngineer) GetName() string {
	return fmt.Sprintf("%v %v", e.Firstname, e.Lastname)
}

func (e DevOpsEngineer) GetLevel() string {
	return fmt.Sprintf("Level: %v", e.Level)
}

// GetTracking - returns tracking information - hours, month, year
func (e DevOpsEngineer) GetTracking() string {
	if e.TrackedHours < 0 {
		log.Errorf("Tracked hours should be a positive number, provided %v", e.TrackedHours)
	}

	fullFte := 160

	// Calculate remaining percent to fullFte
	percentUtil := (float32(e.TrackedHours) / float32(fullFte)) * 100
	// result of this format looks like 67.50%
	fmtTrackingPercent := fmt.Sprintf("%.2f%%", percentUtil)

	return fmt.Sprintf("Tracking: \n  - %v %v - %v hrs - %v", e.TrackingMonth, e.TrackingYear, e.TrackedHours, fmtTrackingPercent)
}

// TODO: Implment this - should calculate average % utilization for previous year
func (e DevOpsEngineer) AverageUtilization(yearTracking string) (averageUtilization string, err error) {
	someSlice := []int{1, 2, 3, 4}
	sizeOfSlice := 4
	sumTotal := 0
	for i := 0; i < sizeOfSlice; i++ {
		sumTotal += someSlice[i]
	}
	avg := (float64(sumTotal) / (float64(sizeOfSlice)))
	fmt.Println(avg)
	return "", nil
}
