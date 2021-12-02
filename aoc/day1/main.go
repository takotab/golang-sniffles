package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type stack []int

func (s stack) Push(v int) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, int) {
	l := len(s)
	return s[1:], s[l-1]
}

func (s stack) sum() int {
	sum := 0
	for _, v := range s {
		sum += v
	}
	return sum
}

type Position struct {
	amount_increases int
	last             int
	last3            stack
}

var position = Position{amount_increases: 0, last3: stack{}}

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
		amount, err := strconv.Atoi(command_parsed[0])
		if err != nil {
			log.Fatal(err)
		}
		position.last3 = position.last3.Push(amount)
		fmt.Println(position.last3, position.last3.sum())
		if len(position.last3) > 3 {
			position.last3, _ = position.last3.Pop()
			if position.last == 0 {
				position.last = amount
			} else if position.last3.sum() > position.last {
				position.amount_increases += 1
				fmt.Println(amount, position, "amount_increases: ")
			}
		}
		position.last = position.last3.sum()

	}
	fmt.Println(position.amount_increases)
}
