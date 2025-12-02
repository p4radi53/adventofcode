package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/p4radi53/aoc/internal"
)

func main() {
	path := internal.GetPath(2017, 4)
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file:", err)
	}
	input := strings.TrimSpace(string(data))
	lines := strings.Split(input, "\n")
	fmt.Println(partOne(lines))
}

func partOne(lines []string) int {
	result := 0
OuterLoop:
	for _, line := range lines {
		words := strings.Split(line, " ")

		usedWords := make(map[string]struct{})
		for _, word := range words {
			if _, exists := usedWords[word]; exists {
				continue OuterLoop
			}
			usedWords[word] = struct{}{}
		}
		result += 1
	}

	return result
}

func partTwo(lines []string) int {
	result := 0

	for _, line := range lines {
		words := strings.Split(line, " ")

		usedWords := make(map[string]struct{})
		for _, word := range words {

			if _, exists := usedWords[word]; exists {

			}

		}
	}

	return result
}

func isAnagram(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	aRunes := strings.Split(a, "")
	bRunes := strings.Split(b, "")

	sort.Strings(aRunes)
	sort.Strings(bRunes)

	for i := range aRunes {
		if aRunes[i] != bRunes[i] {
			return false
		}
	}
	return true
}
