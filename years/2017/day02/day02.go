package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/p4radi53/aoc/internal"
)

func main() {
	path := internal.GetPath(2017, 2)
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file:", err)
	}
	input := strings.TrimSpace(string(data))
	lines := strings.Split(input, "\n")

	fmt.Printf("Result part 1: %v\n", partOne(lines))
	fmt.Printf("Result part 2: %v\n", partTwo(lines))
}

func partOne(lines []string) int {
	result := 0
	for _, line := range lines {
		numbers, err := parseNumbers(line)
		if err != nil {
			fmt.Println("parsing error!")
		}

		minimum, maximum := findExtremes(numbers)
		diff := maximum - minimum
		result += diff
	}
	return result

}

func partTwo(lines []string) int {
	result := 0

Outerloop:
	for _, line := range lines {
		numbers, err := parseNumbers(line)
		if err != nil {
			fmt.Println("parsing error!")
		}

		for _, first_iterator := range numbers {
			for _, second_iterator := range numbers {
				larger := max(second_iterator, first_iterator)
				smaller := min(second_iterator, first_iterator)

				if larger%smaller == 0 && larger != smaller {
					result += larger / smaller
					continue Outerloop
				}
			}

		}
	}
	return result
}

func parseNumbers(line string) ([]int, error) {
	numbers := strings.Fields(line)
	var intNumbers []int

	for _, num := range numbers {
		n, err := strconv.Atoi(num)
		if err != nil {
			return intNumbers, fmt.Errorf("invalid number: %v", err)
		}
		intNumbers = append(intNumbers, n)
	}

	return intNumbers, nil
}

func findExtremes(numbers []int) (int, int) {
	maximum := slices.Max(numbers)
	minimum := slices.Min(numbers)

	return minimum, maximum
}
