package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"sync"

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

func main() {
	wg := sync.WaitGroup{}
	files := []string{"test.xlsx", "test2.xlsx", "test3.xlsx"}
	for _, file := range files {
		wg.Add(1)
		go write(file, &wg)
	}
	wg.Wait()

	files_found := make(chan string, 1000)
	err := filepath.Walk(".",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if filepath.Ext(path) == ".xlsx" {
				fmt.Println(path, info.Size())
				files_found <- path
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}
	close(files_found)
	fmt.Println("done")

	maxRows := make(chan []string, 100)
	for file := range files_found {
		wg.Add(1)
		go get_max_rows(file, []int{1, 2}, maxRows, &wg)
	}

	wg.Wait()
	close(maxRows)
	for rowPrint := range maxRows {
		fmt.Println(rowPrint)
	}

}
