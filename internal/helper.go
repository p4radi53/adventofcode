package internal

import (
	"fmt"
	"path/filepath"
	"strconv"
)

const filename = "input.txt"

func GetPath(year int, day int) string {
	formatted := fmt.Sprintf("day%02d", day)
	cwd := filepath.Join("years", strconv.Itoa(year), formatted, filename)
	return cwd
}


func ConvertToGrid(lines []string) [][]rune {
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


