package main

import "testing"

func Test_partOne(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		freshRanges [][2]int
		ingredients []int
		want        int
	}{
		{
			name:        "Test case 1",
			freshRanges: [][2]int{{1, 5}, {10, 15}, {20, 25}},
			ingredients: []int{3, 12, 18, 22, 30},
			want:        3,
		},
		{
			name:        "Overlapping ranges",
			freshRanges: [][2]int{{1, 10}, {5, 15}, {10, 20}},
			ingredients: []int{7, 12, 17, 25},
			want:        3,
		},
		{
			name:        "No fresh ingredients",
			freshRanges: [][2]int{{30, 40}, {50, 60}},
			ingredients: []int{10, 20, 25},
			want:        0,
		},

		{
			name:        "Overlapping ranges with edge cases",
			freshRanges: [][2]int{{1, 10}, {10, 20}, {20, 30}},
			ingredients: []int{10, 20, 30, 31},
			want:        3,
		},
		{
			name:        "Overlapping ranges with same starts",
			freshRanges: [][2]int{{1, 5}, {1, 10}, {1, 15}},
			ingredients: []int{3, 7, 12},
			want:        3,
		},
		{
			name:        "Single range covering all ingredients",
			freshRanges: [][2]int{{1, 100}},
			ingredients: []int{10, 20, 30, 40, 50},
			want:        5,
		},
		{
			name:        "No ranges",
			freshRanges: [][2]int{},
			ingredients: []int{10, 20, 30},
			want:        0,
		},
		{
			name:        "No ingredients",
			freshRanges: [][2]int{{1, 10}, {20, 30}},
			ingredients: []int{},
			want:        0,
		},
		{
			name: "Smaller range then larger range including smaller",
			freshRanges: [][2]int{{5, 10}, {1, 15}},
			ingredients: []int{6, 11, 14},
			want:        3,
		},
		{
			name: "Larger rane then smaller range including smaller",
			freshRanges: [][2]int{{1, 15}, {5, 10}},
			ingredients: []int{6, 11, 14},
			want:        3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := partOne(tt.freshRanges, tt.ingredients)
			if got != tt.want {
				t.Errorf("partOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
