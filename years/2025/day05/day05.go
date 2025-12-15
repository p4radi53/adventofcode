package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/p4radi53/aoc/internal"
)

func parseInput(input string) ([][2]int, []int) {
	inputs := strings.SplitN(input, "\n\n", 2)
	freshRangesInput := inputs[0]
	ingredientsInput := inputs[1]
	freshRangesLines := strings.Split(strings.TrimSpace(freshRangesInput), "\n")
	ingredientsLines := strings.Split(strings.TrimSpace(ingredientsInput), "\n")

	freshRanges := make([][2]int, len(freshRangesLines))
	for i, line := range freshRangesLines {
		var start, end int
		fmt.Sscanf(line, "%d-%d", &start, &end)
		freshRanges[i] = [2]int{start, end}
	}

	ingredients := make([]int, len(ingredientsLines))
	for i, line := range ingredientsLines {
		var ingredient int
		fmt.Sscanf(line, "%d", &ingredient)
		ingredients[i] = ingredient
	}
	return freshRanges, ingredients
}

func mergeRanges(freshRanges [][2]int) [][2]int {
	if len(freshRanges) == 0 {
		return freshRanges
	}
	merged := make([][2]int, 0)
	current := freshRanges[0]

	for i := 1; i < len(freshRanges); i++ {
		if freshRanges[i][0] <= current[1]+1 {
			if freshRanges[i][1] > current[1] {
				current[1] = freshRanges[i][1]
			}
		} else {
			merged = append(merged, current)
			current = freshRanges[i]
		}
	}
	merged = append(merged, current)
	return merged
}

func checkFreshness(freshRanges [][2]int, ingredient int) bool {
	for _, fr := range freshRanges {
		if ingredient >= fr[0] && ingredient <= fr[1] {
			return true
		}
	}
	return false
}

func main() {
	path := internal.GetPath(2025, 5)
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file:", err)
	}
	input := strings.TrimSpace(string(data))
	freshRanges, ingredients := parseInput(input)
	sort.Slice(freshRanges, func(i, j int) bool {
		return (freshRanges[i][0] < freshRanges[j][0]) || (freshRanges[i][0] == freshRanges[j][0] && freshRanges[i][1] < freshRanges[j][1])
	})
	freshRanges = mergeRanges(freshRanges)
	fmt.Println(partOne(freshRanges, ingredients))
	fmt.Println(partTwo(freshRanges))
}

func partOne(freshRanges [][2]int, ingredients []int) int {
	count := 0
	for _, ingredient := range ingredients {
		if checkFreshness(freshRanges, ingredient) {
			count++
		}
	}
	return count
}

func partTwo(freshRanges [][2]int) int {
	sum := 0
	for _, fr := range freshRanges {
		sum += fr[1] - fr[0] + 1
	}
	return sum
}
