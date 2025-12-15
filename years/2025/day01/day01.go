package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/p4radi53/aoc/internal"
)

func main() {
	path := internal.GetPath(2025, 1)
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file:", err)
	}
	input := strings.TrimSpace(string(data))
	lines := strings.Split(input, "\n")
	fmt.Println(partOne(lines))
	fmt.Println(partTwo(lines))
}
func parseLine(line string) (byte, int) {
	// example line: "R14"
	number, err := strconv.Atoi(line[1:])
	if err != nil {
		panic(err)
	}
	return line[0], number
}

func partOne(lines []string) int {
	currentValue := 50
	times0Visited := 0
	for _, line := range lines {
		sign, value := parseLine(line)
		switch sign {
		case 'R':
			currentValue = (currentValue + value) % 100
		case 'L':
			currentValue = (currentValue - value + 100) % 100
		}
		if currentValue == 0 {
			times0Visited++
		}
	}

	return times0Visited
}

func partTwo(lines []string) int {
	currentValue := 50
	times0Visited := 0
	for _, line := range lines {
		sign, turnValue := parseLine(line)
		times0Visited += turnValue / 100
		turnValue = turnValue % 100
		switch sign {
		case 'R':
			if currentValue+turnValue >= 100 {
				times0Visited++
			}
			currentValue = (currentValue + turnValue) % 100
		case 'L':
			if currentValue-turnValue <= 0 && currentValue > 0  {
				times0Visited++
			}
			currentValue = (currentValue - turnValue + 100) % 100
		}

	}

	return times0Visited
}
