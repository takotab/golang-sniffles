package main

import (
	"log"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

func write() {

	f := excelize.NewFile()

	f.SetCellValue("Sheet1", "A1", "name")
	f.SetCellValue("Sheet1", "B1", "value x")
	f.SetCellValue("Sheet1", "C1", "value y")

	f.SetCellValue("Sheet1", "A2", "a")
	f.SetCellValue("Sheet1", "B2", 3.1)
	f.SetCellValue("Sheet1", "C2", 1.0)

	f.SetCellValue("Sheet1", "A3", "b")
	f.SetCellValue("Sheet1", "B3", 2.2)
	f.SetCellValue("Sheet1", "C3", 1.9)

	f.SetCellValue("Sheet1", "A4", "c")
	f.SetCellValue("Sheet1", "B4", 0.9)
	f.SetCellValue("Sheet1", "C4", 3.8)

	if err := f.SaveAs("simple.xlsx"); err != nil {
		log.Fatal(err)
	}
}
