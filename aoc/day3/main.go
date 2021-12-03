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
	for scanner.Scan() {
		command_parsed := strings.Split(scanner.Text(), "")
		for position_idx, amount_str := range command_parsed {
			amount, err := strconv.Atoi(amount_str)
			if err != nil {
				log.Fatal(err)
			}
			if amount == 0 {
				if len(report.positions) <= position_idx {
					report.positions = append(report.positions, Position{frequency0: 1})
				} else {
					report.positions[position_idx].frequency0 += 1
				}
			} else if amount == 1 {
				if len(report.positions) <= position_idx {
					report.positions = append(report.positions, Position{})
				} else {
					report.positions[position_idx].frequency1 += 1
				}
			}

		}

	}
	fmt.Println(report.positions)
	fmt.Println(report.powerConsumption())
}
