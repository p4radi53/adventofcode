package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strings"

	"github.com/p4radi53/aoc/internal"
)

func main() {
	path := internal.GetPath(2025, 8)
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file:", err)
	}
	input := strings.TrimSpace(string(data))
	lines := strings.Split(input, "\n")
	junctionBoxes := parseInputToJunctionBoxes(lines)
	distances := getAllDistances(junctionBoxes)
	sortDistancesByValue(distances)
	fmt.Println("Part One answer", partOne(junctionBoxes, distances))

	// Reset connections for part two
	for i := range junctionBoxes {
		junctionBoxes[i].connectedTo = []*JunctionBox{}
	}
	fmt.Println("Part Two answer:", partTwo(junctionBoxes, distances))
}

func partOne(junctionBoxes []JunctionBox, distances []Distance) int {
	fmt.Println("Starting part one...")
	for i, distance := range distances {
		circuit := junctionBoxes[distance.from].TraverseConnections(make(map[int]bool))
		if _, exists := circuit[distance.to]; exists {
			// already connected
			continue
		}

		junctionBoxes[distance.from].connectedTo = append(junctionBoxes[distance.from].connectedTo, &junctionBoxes[distance.to])
		junctionBoxes[distance.to].connectedTo = append(junctionBoxes[distance.to].connectedTo, &junctionBoxes[distance.from])

		// Only 1000 connections needed for part one
		if i > 1000 {
			break
		}
	}

	allVisited := make(map[int]bool)
	circuitLengths := make([]int, 0)
	for i := range junctionBoxes {
		if allVisited[i] {
			continue
		}
		visited := junctionBoxes[i].TraverseConnections(make(map[int]bool))
		for k := range visited {
			allVisited[k] = true
		}
		circuitLengths = append(circuitLengths, len(visited))
	}
	sort.Ints(circuitLengths)
	fmt.Println("Total circuits:", len(circuitLengths))
	product := 1
	for _, x := range circuitLengths[len(circuitLengths)-3:] {
		product *= x
	}
	fmt.Println("Product of top 3 largest circuits:", product)
	return product
}

func partTwo(junctionBoxes []JunctionBox, distances []Distance) int {
	fmt.Println("Starting part two...")
	for i, distance := range distances {
		circuit := junctionBoxes[distance.from].TraverseConnections(make(map[int]bool))
		if len(circuit) == len(junctionBoxes) {
			fmt.Println("All junction boxes connected.")
			fmt.Println("It took", i+1, "iterations.")
			if i == 0 {
				fmt.Println("Only one distance found, cannot multiply last two X coordinates.")
				break
			}
			lastDistance := distances[i-1]
			multiplied := junctionBoxes[lastDistance.from].x * junctionBoxes[lastDistance.to].x
			fmt.Println("Multiplied X coordinates of last two connected junction boxes:", multiplied)
			return multiplied
		}
		if _, exists := circuit[distance.to]; exists {
			// already connected
			continue
		}

		junctionBoxes[distance.from].connectedTo = append(junctionBoxes[distance.from].connectedTo, &junctionBoxes[distance.to])
		junctionBoxes[distance.to].connectedTo = append(junctionBoxes[distance.to].connectedTo, &junctionBoxes[distance.from])

	}

	fmt.Println("Could not connect all junction boxes.")
	return -1
}

type JunctionBox struct {
	id          int
	x           int
	y           int
	z           int
	connectedTo []*JunctionBox
}

func (jb *JunctionBox) TraverseConnections(visited map[int]bool) map[int]bool {
	if visited[jb.id] {
		return visited
	}
	visited[jb.id] = true

	for _, conn := range jb.connectedTo {
		conn.TraverseConnections(visited)
	}

	return visited
}

type Distance struct {
	from  int
	to    int
	value float64
}

func getDistanceBetweenBoxes(jb *JunctionBox, jb2 *JunctionBox) Distance {
	dx := jb2.x - jb.x
	dy := jb2.y - jb.y
	dz := jb2.z - jb.z
	sum := float64(dx*dx + dy*dy + dz*dz)
	return Distance{from: jb.id, to: jb2.id, value: math.Sqrt(sum)}
}

func getAllDistances(junctionBoxes []JunctionBox) []Distance {
	distances := make([]Distance, 0, len(junctionBoxes)*len(junctionBoxes))
	for i := range junctionBoxes {
		for j := i; j < len(junctionBoxes); j++ {
			if i != j {
				x := getDistanceBetweenBoxes(&junctionBoxes[i], &junctionBoxes[j])
				distances = append(distances, x)
			}
		}
	}
	return distances
}

func sortDistancesByValue(distances []Distance) {
	sort.Slice(distances, func(i, j int) bool {
		return distances[i].value < distances[j].value
	})
}

func parseInputToJunctionBoxes(lines []string) []JunctionBox {
	junctionBoxes := make([]JunctionBox, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, ",")
		if len(parts) != 3 {
			panic("invalid input line: " + line)
		}
		var x, y, z int
		fmt.Sscanf(parts[0], "%d", &x)
		fmt.Sscanf(parts[1], "%d", &y)
		fmt.Sscanf(parts[2], "%d", &z)
		junctionBoxes[i] = JunctionBox{id: i, x: x, y: y, z: z, connectedTo: []*JunctionBox{}}
	}
	return junctionBoxes
}
