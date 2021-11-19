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

type Row struct {
	max    float64
	values []string
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
	ColumnsToGetMax := make(map[int]Row)
	ColumnsToGetMax[1] = Row{0, []string{}}
	ColumnsToGetMax[2] = Row{0, []string{}}
	for _, row := range rows {
		for col_index, colCell := range row {
			if CurrentMaxRow, ok := ColumnsToGetMax[col_index]; ok {
				if s, err := strconv.ParseFloat(colCell, 32); err == nil {
					if CurrentMaxRow.max < s {
						ColumnsToGetMax[col_index] = Row{s, row}
						// fmt.Println("high", col_index, s)
					}
				}

			}
		}
		fmt.Println(row)
	}
	fmt.Println(ColumnsToGetMax)

}
