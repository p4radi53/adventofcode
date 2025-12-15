package main

import "testing"

func equalSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func Test_findVoltagePartTwo(t *testing.T) {
	tests := []struct {
		name      string // description of this test case
		batteries []int
		want      []int
	}{
		{
			name:      "Test case 1",
			batteries: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3},
			want:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3},
		},
		{
			name:      "Test case 2",
			batteries: []int{9, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3},
			want:      []int{9, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3},
		},
		{
			name:      "Test case 3",
			batteries: []int{9, 9, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3},
			want:      []int{9, 9, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findVoltagePartTwo(tt.batteries)
			if !equalSlices(got, tt.want) {
				t.Errorf("findVoltagePartTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}
