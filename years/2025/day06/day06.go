package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/p4radi53/aoc/internal"
)

func main() {
	path := internal.GetPath(2025, 6)
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file:", err)
	}
	input := strings.TrimSpace(string(data))
	lines := strings.Split(input, "\n")
	// numbers := parseNumbers(lines[0:(len(lines) - 1)])
	operationSigns := parseOperationSigns(lines[len(lines)-1])
	numbers := parseNumbers(lines[:len(lines)-1])
	partOneResult := partOne(numbers, operationSigns)
	fmt.Println("part one:", partOneResult)

	// part two requires completely different parsing
	grid := internal.ConvertToGrid(lines[:len(lines)-1])
	numbersGrid := transposeGrid(grid)
	newLines := convertGridToStrings(numbersGrid)
	numbersGridInts := parseNumbers(newLines)

	partTwoResult := partOne(numbersGridInts, operationSigns)
	fmt.Println("part two result:", partTwoResult)
}

func convertGridToStrings(grid [][]rune) []string {
	result := make([]string, 0, len(grid))
	for i := range grid {
		result = append(result, string(grid[i]))
	}
	return result
}

func transposeGrid(grid [][]rune) [][]rune {
	if len(grid) == 0 {
		return [][]rune{}
	}
	numRows := len(grid)
	numCols :=
		len(grid[0])
	transposed := make([][]rune, numCols)
	for i := range numCols {
		transposed[i] = make([]rune, numRows)
		for j := range numRows {
			transposed[i][j] = grid[j][i]
		}
	}
	return transposed
}

var intRe = regexp.MustCompile(`-?\d+`)

func parseNumbers(lines []string) [][]int {
	numInts := make([][]int, 0, len(lines))
	for _, line := range lines {
		matches := intRe.FindAllString(line, -1)
		numRow := make([]int, 0, len(matches))
		for _, m := range matches {
			v, err := strconv.Atoi(m)
			if err != nil {
				panic(err)
			}
			numRow = append(numRow, v)
		}
		numInts = append(numInts, numRow)
	}
	return numInts
}

func parseOperationSigns(line string) []rune {
	signs := make([]rune, 0, len(line))
	for _, char := range line {
		if char == '+' || char == '*' {
			signs = append(signs, char)
		}
	}
	return signs
}

func partOne(numbers [][]int, operationSigns []rune) int {
	results := make([]int, 0, len(numbers))
	for _, sign := range operationSigns {
		var result int
		switch sign {
		case '+':
			result = 0
		case '*':
			result = 1
		}
		results = append(results, result)
	}

	for i := range numbers {
		for j := range numbers[i] {
			switch operationSigns[j] {
			case '+':
				results[j] += numbers[i][j]
			case '*':
				results[j] *= numbers[i][j]
			}
		}

	}

	resultSum := 0
	for _, v := range results {
		resultSum += v
	}
	return resultSum
}
