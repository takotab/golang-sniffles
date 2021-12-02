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
	horizontal, depth int
}

var position = Position{horizontal: 0, depth: 0}

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
		command_parsed := strings.Split(scanner.Text(), " ")
		amount, err := strconv.Atoi(command_parsed[1])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(command_parsed[0], amount)

		switch command_parsed[0] {
		case "forward":
			position.horizontal += amount
		case "up":
			position.depth -= amount
		case "down":
			position.depth += amount
		}

	}

	fmt.Println(position)
	fmt.Println(position.horizontal * position.depth)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
