package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	// "github.com/go-gota/gota/dataframe"
)

func check(e error) {
	if e != nil {
		panic(e)
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
	check(filepath.Walk(".",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if filepath.Ext(path) == ".xlsx" {
				fmt.Println(path, info.Size())
				files_found <- path
			}
			return nil
		}))
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
