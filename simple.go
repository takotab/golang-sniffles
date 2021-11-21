package main

import (
	"fmt"
	"log"
	"strconv"
	"sync"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

func write(fn string, wg *sync.WaitGroup) {
	defer wg.Done()
	f := excelize.NewFile()

	f.SetCellValue("Sheet1", "A1", "name")
	f.SetCellValue("Sheet1", "B1", "value x")
	f.SetCellValue("Sheet1", "C1", "value y")

	f.SetCellValue("Sheet1", "A2", "a")
	f.SetCellValue("Sheet1", "B2", 3.1)
	f.SetCellValue("Sheet1", "C2", 1.0)
	f.SetCellValue("Sheet1", "D2", 1.0)

	f.SetCellValue("Sheet1", "A3", "b")
	f.SetCellValue("Sheet1", "B3", 2.2)
	f.SetCellValue("Sheet1", "C3", 1.9)
	f.SetCellValue("Sheet1", "D3", 1.9)

	f.SetCellValue("Sheet1", "A4", "c")
	f.SetCellValue("Sheet1", "B4", 0.9)
	f.SetCellValue("Sheet1", "C4", 3.8)
	f.SetCellValue("Sheet1", "D4", 1.9)

	if err := f.SaveAs(fn); err != nil {
		log.Fatal(err)
	}
}

type Row struct {
	max    float64
	values []string
}

func get_max_rows(fn string, column_numbers []int, return_chan chan []string, wg *sync.WaitGroup) {
	defer wg.Done()
	f, err := excelize.OpenFile(fn)

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
	for _, column_num := range column_numbers {
		ColumnsToGetMax[column_num] = Row{}
	}
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
	for _, columnMax := range ColumnsToGetMax {
		return_chan <- columnMax.values
	}
}
