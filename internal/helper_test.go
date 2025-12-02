package internal

import (
	"path/filepath"
	"testing"
)

func TestGetPath(t *testing.T) {
	tests := []struct {
		year     int
		day      int
		expected string
	}{
		{2017, 1, filepath.Join("years", "2017", "day01", "day01.go", filename)},
		{2020, 10, filepath.Join("years", "2020", "day10", "day10.go", filename)},
		{2023, 9, filepath.Join("years", "2023", "day09", "day09.go", filename)},
	}

	for _, tt := range tests {
		got := GetPath(tt.year, tt.day)
		if got != tt.expected {
			t.Errorf("GetPath(%d, %d) = %q, want %q", tt.year, tt.day, got, tt.expected)
		}
	}
}
