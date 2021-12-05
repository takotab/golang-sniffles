package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Meter struct {
	readings [][]string
}

func (m *Meter) getFrequecies(position int) (map[string]int, error) {
	frequencies := make(map[string]int)
	for _, reading := range m.readings {
		frequencies[reading[position]]++
	}
	return frequencies, nil
}

func (m *Meter) removeReadingsByPositionValue(position int, removeValue string) {
	for _, reading := range m.readings {
		if reading[position] == removeValue {
			m.readings = append(m.readings[:position], m.readings[position+1:]...)
		}
	}
}

func (m *Meter) iterateTillOneOver(removeMostCommon bool) {
	numPositons := len(m.readings[0])
	for position := 0; position < numPositons; position++ {
		if len(m.readings) == 0 {
			break
		}
		frequencies, err := m.getFrequecies(position)
		if err != nil {
			log.Fatal(err)
		}

		if removeMostCommon { // Co2
			if frequencies["1"] >= frequencies["0"] {
				fmt.Println("Removing 1")
				m.removeReadingsByPositionValue(position, "1")
			} else {
				fmt.Println("Removing 0")
				m.removeReadingsByPositionValue(position, "0")
			}
		} else { // Oxigen
			if frequencies["1"] >= frequencies["0"] {
				fmt.Println("Removing 0")
				m.removeReadingsByPositionValue(position, "0")
			} else {
				fmt.Println("Removing 1")
				m.removeReadingsByPositionValue(position, "1")
			}
		}
		fmt.Println("len m.readings", len(m.readings), "after position ", position)
	}
}

func main() {
	filename := "input_test.txt"
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	var oxigenOptions Meter
	var co2Options Meter
	for scanner.Scan() {
		parsedCommands := strings.Split(scanner.Text(), "")
		oxigenOptions.readings = append(oxigenOptions.readings, parsedCommands)
		co2Options.readings = append(co2Options.readings, parsedCommands)
	}
	fmt.Println("len oxigenOptions", len(oxigenOptions.readings))
	fmt.Println("len co2Options", len(co2Options.readings))
	co2Options.iterateTillOneOver(true)
	print(co2Options.readings)
	// output, err := strconv.ParseInt(code, 2, 64)
	// fmt.Println(report.positions)
	// fmt.Println(report.powerConsumption())
}
