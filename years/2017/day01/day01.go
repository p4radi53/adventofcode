package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/p4radi53/aoc/internal"
)

func main() {
	path := internal.GetPath(2017, 1)

	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	data_str := strings.TrimSpace(string(data))

	digits, err := parseDigits(data_str)
	if err != nil {
		fmt.Println("Error parsing digts:", err)
		return
	}

	result := part1(digits)
	fmt.Println(result)

	result2 := part2(digits)
	fmt.Println(result2)
}

func parseDigits(input string) ([]int, error) {
	digits := make([]int, len(input))
	for i, ch := range input {
		digit, err := strconv.Atoi(string(ch))
		if err != nil {
			return nil, fmt.Errorf("invalid character at position %d: %v", i, err)
		}
		digits[i] = digit
	}
	return digits, nil
}

func part1(digits []int) int {
	firstDigit := digits[0]
	previousDigit := firstDigit

	sum := 0
	for _, d := range digits[1:] {
		if d == previousDigit {
			sum += d
		}
		previousDigit = d
	}

	lastDigit := previousDigit
	if lastDigit == firstDigit {
		sum += firstDigit
	}

	return sum
}

func part2(digits []int) int {
	result := 0
	length := len(digits)

	for i, d := range digits {
		var comparisonIndex int
		if i > (length/2 - 1) {
			comparisonIndex = i - length/2
		} else {
			comparisonIndex = i + length/2
		}

		if digits[comparisonIndex] == d {
			result += d
		}
	}
	return result
}
