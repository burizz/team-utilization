package main

import (
	"fmt"
)

func init() {
	// DB connection
	// Seed data
}

func main() {
	var trackedHours int = 160
	utilizationPercent, err := calculateUtilization(trackedHours)
	if err != nil {
		fmt.Printf("Error calculating utilization percent: %v\n", err)
	}

	fmt.Println(fmt.Sprintf("%.2f", utilizationPercent) + "%")
}

func calculateUtilization(trackedHours int) (float32, error) {
	if trackedHours < 0 {
		// check if negative number
		return float32(trackedHours), fmt.Errorf("Tracked hours should be a positive number")
	}

	// Calculate remaining percent to fullFte
	fullFte := 160
	percentUtil := (float32(trackedHours) / float32(fullFte)) * 100
	return percentUtil, nil
}
