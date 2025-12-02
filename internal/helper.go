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
