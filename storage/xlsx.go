package storage

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	log "github.com/sirupsen/logrus"
)

// ParseTrackingFromExcel - take excel path with individual tracking report; return total tracked hours
func ParseTrackingFromExcel(xlsFilePath string, xlsWorksheet string) (trackedTotal string, parseExcelErr error) {
	// Open Excel File by path
	f, openExcelErr := excelize.OpenFile(xlsFilePath)
	if openExcelErr != nil {
		log.Errorf("Cannot open excel file %v : %v\n", f, openExcelErr)
		return "", openExcelErr
	}

	var sumHours float64

	// Go through each "I" cell (where tracking hours are stored) and sum the total hours
	for i := 1; i < 100; i++ {
		cellPosition := "I" + strconv.Itoa(i)
		// TODO: parameterize worksheet name
		cellValue := f.GetCellValue(xlsWorksheet, cellPosition)

		if cellValue == "Hours" || cellValue == "" {
			continue
		}

		convertedNum, convToFloatErr := strconv.ParseFloat(cellValue, 64)
		if convToFloatErr != nil {
			log.Errorf("Cannot convert %v to Int: %v", cellValue, convToFloatErr)
			return "", convToFloatErr
		}

		sumHours += convertedNum
	}

	stringifiedNum := fmt.Sprintf("%.2f", sumHours)
	return stringifiedNum, nil
}

// ParsePeriodFromExcel - take excel path with individual tracking report; return month and year of the report
func ParsePeriodFromExcel(xlsFilePath string, xlsWorksheet string) (trackingMonth string, trackingYear string, parseExcelErr error) {
	// Open Excel File by path
	f, openExcelErr := excelize.OpenFile(xlsFilePath)
	if openExcelErr != nil {
		log.Errorf("Cannot open excel file %v : %v\n", f, openExcelErr)
		return "", "", openExcelErr
	}

	// TODO: parameterize worksheet name
	// TODO: verify if there is a better field to get from this excel
	cellValue := f.GetCellValue(xlsWorksheet, "A5")

	// Parse Date to proper format
	date, dateParseErr := time.Parse("2-Jan-06", cellValue)
	if dateParseErr != nil {
		log.Errorf("Cannot parse date to standard format")
	}

	// Conver to string
	parsedYear := fmt.Sprintf("%v", date.Year())
	parsedMonth := fmt.Sprintf("%v", date.Month())

	return parsedMonth, parsedYear, nil
}

// CalculateTracking - calculate the % of tracked hours - with 160 being 100%
func CalculateTrackingPercent(totalHours string) (percentUtilization string, err error) {
	convertedHours, convertErr := strconv.ParseFloat(totalHours, 64)
	if convertErr != nil {
		log.Errorf("Cannot convert string to float64: %v", convertErr)
	}

	if convertedHours < 0 {
		hoursValueErr := errors.New("Total hours should be a positive number")
		log.Errorf("Total hours should be a positive number, provided %v", convertedHours)
		return "", hoursValueErr
	}

	var fullFte float64 = 160

	// Calculate remaining percent to fullFte
	percentUtil := (convertedHours) / (fullFte) * 100
	fmtTrackingPercent := fmt.Sprintf("%.2f%%", percentUtil)

	return fmtTrackingPercent, nil
}
