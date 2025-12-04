package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/p4radi53/aoc/internal"
)

func main() {
	path := internal.GetPath(2025, 4)
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file:", err)
	}
	input := strings.TrimSpace(string(data))
	lines := strings.Split(input, "\n")
	fmt.Println(partOne(lines))
	fmt.Println(partTwo(lines))
}

func convertToGrid(lines []string) [][]rune {
	result := make([][]rune, len(lines))
	for i := len(lines) - 1; i >= 0; i-- {
		row := make([]rune, len(lines[i]))
		for j, char := range lines[i] {
			row[j] = char
		}
		result[i] = row
	}
	return result
}

func checkNeighbors(x int, y int, grid [][]rune, gridSize int) bool {
	if grid[x][y] != '@' {
		return false
	}
	incorrectOnes := 0
	combinations := []struct{ dx, dy int }{
		{1, 1}, {1, -1}, {-1, 1}, {-1, -1}, {-1, 0}, {0, -1}, {0, 1}, {1, 0},
	}
	for i := range combinations {
		newX := x + combinations[i].dx
		newY := y + combinations[i].dy
		if newX >= 0 && newX < gridSize && newY >= 0 && newY < gridSize {
			if grid[newX][newY] == '@' {
				incorrectOnes++
			}
		}
	}
	if incorrectOnes < 4 {
		return true
	}
	return false
}

// Part One
func partOne(lines []string) int {
	grid := convertToGrid(lines)
	size := len(grid)
	removableRollsOfPaper := 0
	for x := range size {
		for y := range size {
			if checkNeighbors(x, y, grid, size) {
				removableRollsOfPaper++
			}
		}
	}
	return removableRollsOfPaper
}

func partTwo(lines []string) int {
	grid := convertToGrid(lines)
	size := len(grid)
	count := 0
	maxIterations := 100
	iterations := 0
	for {
		stopFlag := true
		for x := range size {
			for y := range size {
				if checkNeighbors(x, y, grid, size) {
					grid[x][y] = 'X'
					count++
					stopFlag = false
				}
			}
		}
		if stopFlag || iterations >= maxIterations {
			fmt.Println("Stopped after iterations:", iterations)
			break
		}
		iterations++
	}
	return count
}
