package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/p4radi53/aoc/internal"
)

func main() {
	path := internal.GetPath(2025, 11)
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file:", err)
	}
	input := strings.TrimSpace(string(data))
	lines := strings.Split(input, "\n")
	devices := make(map[string]Device)
	for _, line := range lines {
		d := fromInput(line)
		devices[d.id] = d
	}
	// part one
	pathsYouToOut := traverse(devices, "you", "out")
	fmt.Println("Node 'out' visited times:", pathsYouToOut)

	// part two
	// all paths from 'svr' to 'out' that pass through 'fft' and 'dvc' in any order
	pathsSvrToOutViaFftDvc := 0
	pathsSvrToFft := traverse(devices, "svr", "fft")
	pathsSvrToDvc := traverse(devices, "svr", "dvc")
	pathsFftToOut := traverse(devices, "fft", "out")
	pathsDvcToOut := traverse(devices, "dvc", "out")

	pathsSvrToOutViaFftDvc += pathsSvrToFft * pathsDvcToOut
	pathsSvrToOutViaFftDvc += pathsSvrToDvc * pathsFftToOut

	fmt.Println("Paths from 'svr' to 'out' via 'fft' and 'dvc':", pathsSvrToOutViaFftDvc)
}

type Device struct {
	id      string
	outputs []string
}

func fromInput(line string) Device {
	parts := strings.Split(line, ": ")
	id := parts[0]
	outputs := strings.Split(parts[1], " ")
	return Device{
		id:      id,
		outputs: outputs,
	}
}

func traverse(devices map[string]Device, startId string, targetId string) int {
	visited := make(map[string]int)

	var traversalHelper func(currentId string) int
	traversalHelper = func(currentId string) int {
		if v, ok := visited[currentId]; ok {
			return v
		}

		visited[currentId]++
		if currentId == targetId {
			return 1
		}

		total := 0
		for _, outputId := range devices[currentId].outputs {
			total += traversalHelper(outputId)
		}

		return total
	}

	return traversalHelper(startId)
}
