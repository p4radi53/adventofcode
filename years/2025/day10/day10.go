package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/p4radi53/aoc/internal"
)

func main() {
	path := internal.GetPath(2025, 10)
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file:", err)
	}
	input := strings.TrimSpace(string(data))
	lines := strings.Split(input, "\n")
	// part one
	machines := make([]Machine, 0)
	for _, line := range lines {
		m := fromInput(line)
		machines = append(machines, m)
	}
	fmt.Println("machine example:", machines[0])

}

type Machine struct {
	buttons            [][]bool
	lightConfiguration []bool
	joltage            []int
}

func fromInput(line string) Machine {
	// example input
	//[.##.#..#.] (2,3,4,6) (0,1,3,4,5,6) (1,2,3,7,8) (0,1,3,5,6,7,8) (0,1,3,4,5,6,8) (0,1,4,7,8) (2,4,6,7,8) {41,56,34,37,46,22,41,67,71}
	parts := strings.Split(line, " ")
	lightConfiguration := make([]bool, 0)
	for _, c := range strings.Trim(parts[0], "[]") {
		if c == '#' {
			lightConfiguration = append(lightConfiguration, true)
		} else {
			lightConfiguration = append(lightConfiguration, false)
		}
	}
	// buttons
	buttons := make([][]bool, 0)
	for _, b := range parts[1 : len(parts)-1] {
		b = strings.Trim(b, "()")
		numStrs := strings.Split(b, ",")
		nums := make([]bool, len(lightConfiguration))
		for _, ns := range numStrs {
			n, _ := strconv.Atoi(ns)
			nums[n] = true
		}
		buttons = append(buttons, nums)
	}
	// joltage
	joltageStrs := strings.Trim(parts[len(parts)-1], "{}")
	joltageParts := strings.Split(joltageStrs, ",")
	joltage := make([]int, 0)
	for _, js := range joltageParts {
		j, _ := strconv.Atoi(js)
		joltage = append(joltage, j)
	}
	return Machine{
		buttons:            buttons,
		lightConfiguration: lightConfiguration,
		joltage:            joltage,
	}
}
