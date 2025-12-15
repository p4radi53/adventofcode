package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/p4radi53/aoc/internal"
)

func main() {
	path := internal.GetPath(2025, 9)
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file:", err)
	}
	input := strings.TrimSpace(string(data))
	lines := strings.Split(input, "\n")
	coordinates := parseCoordinates(lines)
	// part one
	maxArea := findMaxArea(coordinates)
	fmt.Println("Max area:", maxArea)

	// part two
	fmt.Println("Len of coordinates before adding sides:", len(coordinates))
	coordinates = addSides(coordinates)
	fmt.Println("Len of coordinates after adding sides:", len(coordinates))

}


func addSides(coordinates []Point) []Point {
	for i := 0; i < len(coordinates); i++ {
		var p1, p2 Point
		if i == len(coordinates)-1 {
			p1 = coordinates[i]
			p2 = coordinates[0]
		} else {
			p1 = coordinates[i]
			p2 = coordinates[i+1]
		}

		if p1.x == p2.x {
			// vertical line
			minY := p1.y
			maxY := p2.y
			if minY > maxY {
				minY, maxY = maxY, minY
			}
			for y := minY + 1; y < maxY; y++ {
				newPoint := Point{x: p1.x, y: y}
				coordinates = append(coordinates, newPoint)
			}
		} else if p1.y == p2.y {
			// horizontal line
			minX := p1.x
			maxX := p2.x
			if minX > maxX {
				minX, maxX = maxX, minX
			}
			for x := minX + 1; x < maxX; x++ {
				newPoint := Point{x: x, y: p1.y}
				coordinates = append(coordinates, newPoint)
			}
		}

	}
	return coordinates

}

func findMaxArea(coordinates []Point) int {
	maxArea := 0
	for i := range coordinates {
		for j := i + 1; j < len(coordinates); j++ {
			area := rectangleArea(coordinates[i], coordinates[j])
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}

type Point struct {
	x int
	y int
}

func rectangleArea(p1, p2 Point) int {
	width := p2.x - p1.x
	if width < 0 {
		width = -width
	}
	width += 1
	height := p2.y - p1.y
	if height < 0 {
		height = -height
	}
	height += 1
	return width * height
}

func parseCoordinates(lines []string) []Point {
	result := make([]Point, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
		y, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
		result[i] = Point{x: x, y: y}
	}
	return result
}
