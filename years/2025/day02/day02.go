package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/p4radi53/aoc/internal"
)

func main() {
	path := internal.GetPath(2025, 2)
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file:", err)
	}
	input := strings.TrimSpace(string(data))
	lines := strings.Split(input, ",")
	fmt.Println(partOne(lines))
	fmt.Println(partTwo(lines))
}

func parseLine(line string) (int, int) {
	split := strings.Split(line, "-")
	i1, err := strconv.Atoi(split[0])
	if err != nil {
		panic(err)
	}
	i2, err := strconv.Atoi(split[1])
	if err != nil {
		panic(err)
	}
	return i1, i2
}
func lenItoa(i int) int {
	return len(strconv.Itoa(i))
}
func pow10(n int) int {
	result := 1
	for range n {
		result *= 10
	}
	return result
}

func partOne(lines []string) int {
	count := 0
	for _, line := range lines {
		lower, upper := parseLine(line)
		for i := lower; i <= upper; i++ {
			if lenItoa(i)%2 != 0 {
				continue
			}
			power := pow10(lenItoa(i) / 2)
			if i%power == i/power {
				count += i
			}
		}

	}
	return count
}
func findDividers(i int) []int {
	dividers := []int{}
	for d := 2; d < i; d++ {
		if i%d == 0 {
			dividers = append(dividers, d)
		}
	}
	return dividers
}

func partTwo(lines []string) int {
	count := 0
	for _, line := range lines {
		lower, upper := parseLine(line)
		for i := lower; i <= upper; i++ {
			for _, d := range findDividers(i) {
				power := pow10(lenItoa(i) / d)
				if i%power == i/power {
					count += i
				}
			}
		}

	}
	return count
}
