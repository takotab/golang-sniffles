package main

import (
	"fmt"
	"log"
"strconv"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	// "github.com/go-gota/gota/dataframe"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	write()
	f, err := excelize.OpenFile("simple.xlsx")

	if err != nil {
		log.Fatal(err)
	}
	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	ColumnsToGetMax := make(map[int][]float64)
	ColumnsToGetMax[1] = []float32(0.0,0.0)
	ColumnsToGetMax[2] = []float32(0.0,0.0)
	for row_index, row := range rows {
		for col_index, colCell := range row {
			if s, err := strconv.ParseFloat(colCell, 32); err == nil {
				if ColumnsToGetMax[col_index][0] < s {				
					ColumnsToGetMax[col_index][0] := s				
					ColumnsToGetMax[col_index][1] := row_index
					fmt.Println("high", col_index, s)
				}
			}
			
		}
		fmt.Println("nn")
	}
	csvStr := ""
	for _, row := range rows {
		for _, colCell := range row {
			csvStr += colCell + ","
			fmt.Println(colCell)
		}
		csvStr += "\n"
		fmt.Println("nn")
	}
	fmt.Println(csvStr)


}
