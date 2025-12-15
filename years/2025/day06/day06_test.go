package main

import "testing"

func compareGrids[T comparable](a, b [][]T) bool {

	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}

	return true
}
func Test_transposeGrid(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		grid  [][]rune
		want  [][]rune
		want2 [][]int
	}{
		{
			name: "Test case 1",
			grid: [][]rune{
				{'1', '2', '3'},
				{' ', '4', '5'},
				{' ', ' ', '6'},
			},
			want: [][]rune{
				{'1', ' ', ' '},
				{'2', '4', ' '},
				{'3', '5', '6'},
			},
			want2: [][]int{
				{1},
				{24},
				{356},
			},
		},
		{
			name: "more numbers",
			grid: [][]rune{
				{'7', '8', '9', '0', ' ', '3', '2', '8'},
				{' ', '4', '5', '6', ' ', '6', '4', ' '},
				{' ', ' ', '1', '2', ' ', '1', '0', ' '},
			},
			want: [][]rune{
		{'7', ' ', ' '},
		{'8', '4', ' '},
		{'9', '5', '1'},
		{'0', '6', '2'},
		{' ', ' ', ' '},
		{'3', '6', '1'},
		{'2', '4', '0'},
		{'8', ' ', ' '},
			},
			want2: [][]int{

		}
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := transposeGrid(tt.grid)
			if compareGrids(got, tt.want) == false {
				t.Errorf("transposeGrid() = %v, want %v", got, tt.want)
			}
			lines := convertGridToStrings(got)

			parsed := parseNumbers(lines)
			if !compareGrids(parsed, tt.want2) {
				t.Errorf("parseNumbers() = %v, want %v", parsed, tt.want2)
			}
		})
	}
}
