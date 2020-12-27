package storage

import (
	"fmt"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	f, err := excelize.OpenFile("detailed_report.xlsx")
	if err != nil {
		fmt.Printf("Cannot open excel file %v : %v\n", f, err)
	}

	var sumHours float64

	for i := 1; i < 100; i++ {
		cellPosition := "I" + strconv.Itoa(i)
		cellValue, err := f.GetCellValue("Detailed report", cellPosition)
		if err != nil {
			fmt.Errorf("Cannot get cell value: %v", err)
		}

		if cellValue == "Hours" || cellValue == "" {
			continue
		}

		convertedNum, convToIntErr := strconv.ParseFloat(cellValue, 64)
		if convToIntErr != nil {
			fmt.Errorf("Cannot conver %v to Int: %v", cellValue, convToIntErr)
		}

		sumHours += convertedNum
		fmt.Printf("Line: %v ; Value: %v\n", cellPosition, cellValue)
		fmt.Println(sumHours)
	}

	//stringifiedNum := fmt.Sprintf("%.2f%%", sumHours)
	fmt.Printf("Total hours: %v", sumHours)
}
