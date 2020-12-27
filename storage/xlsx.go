package storage

import (
	"fmt"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
)

// TODO: Finish this - errors, returns, prints
func ParseTrackingFromExcel(xlsFilePath string) (trackedTotal string err error) {
	f, err := excelize.OpenFile(xlsFilePath)
	if err != nil {
		fmt.Printf("Cannot open excel file %v : %v\n", f, err)
		return "", err 
	}

	var sumHours float64

	for i := 1; i < 100; i++ {
		cellPosition := "I" + strconv.Itoa(i)
		cellValue, err := f.GetCellValue("Detailed report", cellPosition)
		if err != nil {
			fmt.Errorf("Cannot get cell value: %v", err)
			return "", err
		}

		if cellValue == "Hours" || cellValue == "" {
			continue
		}

		convertedNum, convToFloatErr := strconv.ParseFloat(cellValue, 64)
		if convToFloatErr != nil {
			fmt.Errorf("Cannot convert %v to Int: %v", cellValue, convToFloatErr)
			return "", convToFloatErr 
		}

		sumHours += convertedNum
		fmt.Printf("Line: %v ; Value: %v\n", cellPosition, cellValue)
		fmt.Println(sumHours)
	}

	//stringifiedNum := fmt.Sprintf("%.2f%%", sumHours)
	fmt.Printf("Total hours: %v", sumHours)
}
