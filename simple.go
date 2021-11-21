package main

import (
	"log"
	"sync"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

func write(fn string, files_made chan string, wg *sync.WaitGroup) {
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
	files_made <- fn
}
