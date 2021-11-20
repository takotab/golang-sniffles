package main

import (
	"fmt"
	"log"
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

func get_max_rows(fn string, column_numbers []int, return_chan chan []string) {

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
	write()
	wg := sync.WaitGroup{}
	wg.Add(1)
	maxRows := make(chan []string, 100)
	go func() {
		get_max_rows("simple.xlsx", []int{1, 2}, maxRows)
		wg.Done()
	}()
	for i := 0; i < 2; i++ {
		fmt.Println(<-maxRows)
	}

	wg.Wait()

}
