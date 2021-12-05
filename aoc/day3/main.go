package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	frequency0 int
	frequency1 int
}
type Report struct {
	positions []Position
}

func (r Report) getCode(if0Isfrequent string, if1isFrequent string) string {
	var code string
	for _, position := range r.positions {
		if position.frequency0 > position.frequency1 {
			code += if0Isfrequent //"0"
		} else {
			code += if1isFrequent // 1
		}
	}
	return code
}
func (r Report) gamma() (int, error) {
	code := r.getCode("0", "1")
	output, err := strconv.ParseInt(code, 2, 64)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return int(output), nil
}
func (r Report) epsilon() (int, error) {
	code := r.getCode("1", "0")
	output, err := strconv.ParseInt(code, 2, 64)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return int(output), nil
}
func (r Report) powerConsumption() int {
	gamme, err := r.gamma()
	if err != nil {
		fmt.Println(err)
	}
	eps, err := r.epsilon()
	if err != nil {
		fmt.Println(err)
	}
	return gamme * eps
}

var report = Report{}

func main() {
	filename := "input.txt"
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	var oxigenOptions [][]string
	var co2Options [][]string
	for scanner.Scan() {
		parsedCommands := strings.Split(scanner.Text(), "")
		oxigenOptions = append(oxigenOptions, parsedCommands)
		co2Options = append(co2Options, parsedCommands)
	}
	fmt.Println("len oxigenOptions", len(oxigenOptions))
	fmt.Println("len co2Options", len(co2Options))

	for position := 0; position < len(oxigenOptions[0]); position++ {
		oxigenFrequency := make(map[string]int)
		co2Frequency int
		for _, oxigenOption := range oxigenOptions {
			oxigenFrequency[oxigenOption[position]]++
		}
		for _, co2Option := range co2Options {
			if co2Option[position] == "1" {
				co2Frequency++
			}
		}
		report.positions = append(report.positions, Position{oxigenFrequency, co2Frequency})
	}

	output, err := strconv.ParseInt(code, 2, 64)
	fmt.Println(report.positions)
	fmt.Println(report.powerConsumption())
}
