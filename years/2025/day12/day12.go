package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/p4radi53/aoc/internal"
)

func main() {
	path := internal.GetPath(2025, 12)
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file:", err)
	}
	input := strings.TrimSpace(string(data))
	presents := parsePresents(input)
	for _, p := range presents {
		fmt.Println("Present area:", p.area)
	}
	regions := parseRegions(input)
	partOneAnswer := partOne(presents, regions)
	fmt.Println("Part one answer:", partOneAnswer)
	// no part two :) christmas gift
}
func partOne(presents []Present, regions []Region) int {
	result := 0

	for _, r := range regions {
		regionArea := r.Area()
		polygonsArea := 0
		for i, count := range r.presentCounts {
			polygonsArea += presents[i].area * count
		}

		if polygonsArea <= regionArea {
			result++
		}
	}
	return result
}

type Present struct {
	area int
}

func parsePresents(input string) []Present {
	presentInputStringPart := strings.Split(input, "\n\n")[:6]
	presents := make([]Present, 6)
	for i, presentString := range presentInputStringPart {
		splits := strings.Split(presentString, "\n")
		area := 0
		for _, line := range splits[1:] {
			for _, char := range line {
				if char == '#' {
					area++
				}
			}
		}
		presents[i] = Present{area: area}
	}

	return presents
}

type Region struct {
	width         int
	height        int
	presentCounts []int
}

func (r *Region) Area() int {
	return r.width * r.height
}

func parseRegions(input string) []Region {
	regionsInputStringPart := strings.Split(input, "\n\n")[6]
	regions := make([]Region, 0)
	splits := strings.SplitSeq(regionsInputStringPart, "\n")
	for regionString := range splits {
		// example: 47x49: 60 48 64 49 52 78
		parts := strings.Split(regionString, ": ")
		dimensions := strings.Split(parts[0], "x")
		width, _ := strconv.Atoi(dimensions[0])
		height, _ := strconv.Atoi(dimensions[1])
		countStrs := strings.Split(parts[1], " ")
		presentCounts := make([]int, len(countStrs))
		for j, cs := range countStrs {
			count, _ := strconv.Atoi(cs)
			presentCounts[j] = count
		}
		regions = append(regions, Region{
			width:         width,
			height:        height,
			presentCounts: presentCounts,
		})
	}
	return regions
}
