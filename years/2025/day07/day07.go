package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/p4radi53/aoc/internal"
)

func main() {
	path := internal.GetPath(2025, 7)
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file:", err)
	}
	input := strings.TrimSpace(string(data))
	lines := strings.Split(input, "\n")
	grid := internal.ConvertToGrid(lines)
	fmt.Println(partOne(grid))
	fmt.Println(partTwo(grid))
}

type Beam struct {
	beamXs   []int
	currentY int
	grid     [][]rune
}

func NewBeam(grid [][]rune) *Beam {
	beamStartX := -1
	for i := range grid[0] {
		if grid[0][i] == 'S' {
			beamStartX = i
			break
		}
	}
	if beamStartX == -1 {
		panic("no start found")
	}
	beamXs := make([]int, len(grid[0]))
	beamXs[beamStartX] = 1
	return &Beam{
		beamXs:   beamXs,
		currentY: 0,
		grid:     grid,
	}
}

func (beam *Beam) Move() int {
	beam.currentY++
	splits := 0
	for beamX, beamTimelineCount := range beam.beamXs {
		if beamTimelineCount >= 1 && beam.grid[beam.currentY][beamX] == '^' {
			beam.beamXs[beamX-1] = beam.beamXs[beamX-1] + beamTimelineCount
			beam.beamXs[beamX+1] = beam.beamXs[beamX+1] + beamTimelineCount
			beam.beamXs[beamX] = 0
			splits++
		}
	}
	return splits
}

func partOne(grid [][]rune) int {
	beam := NewBeam(grid)
	splitterCount := 0
	for beam.currentY < len(grid)-1 {
		splitterCount += beam.Move()
	}
	fmt.Println("splitter count:", splitterCount)
	return splitterCount
}
func partTwo(grid [][]rune) int {
	beam := NewBeam(grid)
	for beam.currentY < len(grid)-1 {
		beam.Move()
	}
	countBeamValues := 0
	for _, v := range beam.beamXs {
		countBeamValues += v
	}

	return countBeamValues

}
