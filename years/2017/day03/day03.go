package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/p4radi53/aoc/internal"
)

func main() {
	path := internal.GetPath(2017, 3)
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file:", err)
	}
	input, err := strconv.Atoi(strings.TrimSpace(string(data)))
	if err != nil {
		fmt.Println("Error parsing")
		return
	}

	insideSquareSide := 1
	var insideNumberCount int
	for {
		insideNumberCount = insideSquareSide * insideSquareSide
		if insideNumberCount > input {
			insideSquareSide -= 2
			insideNumberCount = insideSquareSide * insideSquareSide
			break
		}
		insideSquareSide += 2
	}

	outsideSquareSide := insideSquareSide + 2
	outsideNumberCount := input - insideNumberCount

	fullDiff := (outsideSquareSide - 1) / 2
	partialDiff := outsideNumberCount - (outsideSquareSide - 1) - fullDiff // -1 because it starts not in the corner but right above it

	fmt.Println("outsideSquareSide", outsideSquareSide)
	fmt.Println("outsideNumberCount", outsideNumberCount)
	fmt.Println(fullDiff)
	fmt.Println(partialDiff)
}
