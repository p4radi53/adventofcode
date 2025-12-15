package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/p4radi53/aoc/internal"
)

func main() {
	path := internal.GetPath(2025, 3)
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file:", err)
	}
	input := strings.TrimSpace(string(data))
	lines := strings.Split(input, "\n")
	fmt.Println(partOne(lines))
	fmt.Println(partTwo(lines))
}

func findVoltage(batteries []int) (int, int) {
	maxBattery := 0
	secondMaxBattery := 0

	for i := 0; i < len(batteries)-1; i++ {
		if batteries[i] > maxBattery {
			maxBattery = batteries[i]
			secondMaxBattery = batteries[i+1]
		} else if batteries[i] > secondMaxBattery {
			secondMaxBattery = batteries[i]
		}
	}
	// handle last element
	if batteries[len(batteries)-1] > secondMaxBattery {
		secondMaxBattery = batteries[len(batteries)-1]
	}

	return maxBattery, secondMaxBattery
}

func parseLine(line string) []int {
	// assume line only contains digits
	newList := make([]int, len(line))
	for i, char := range line {
		newList[i] = int(char - '0')
	}
	return newList
}

func partOne(lines []string) int {
	totalVoltage := 0
	for _, line := range lines {
		numbers := parseLine(line)
		maxBattery, secondMaxBattery := findVoltage(numbers)

		volatage := maxBattery*10 + secondMaxBattery
		totalVoltage += volatage
	}
	return totalVoltage
}

func findMax(ints []int) (int, int) {
	max := ints[0]
	maxIx := 0
	for i, v := range ints {
		if v > max {
			max = v
			maxIx = i
		}
	}
	return maxIx, max
}

func findVoltagePartTwo(batteries []int) []int {
	result := make([]int, 12)
	startIx := 0
	var tempIx int
	for i := range 12 {
		slice := batteries[(startIx) : len(batteries)-(11-i)]
		tempIx, result[i] = findMax(slice)
		startIx += tempIx + 1
	}

	return result
}
func pow(base, exp int) int {
	result := 1
	for range exp {
		result *= base
	}
	return result
}

func batteriesToVoltage(batteries []int) int {
	voltage := 0
	for i, v := range batteries {
		voltage += v * pow(10, 11-i)
	}
	return voltage
}

func partTwo(lines []string) int {
	totalVoltage := 0
	for _, line := range lines {
		numbers := parseLine(line)
		selectedBatteries := findVoltagePartTwo(numbers)
		volatage := batteriesToVoltage(selectedBatteries)
		totalVoltage += volatage
	}
	return totalVoltage
}
// 169123684480212
// 167302518850275
