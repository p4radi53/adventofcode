package main

import "testing"

func Test_partTwo(t *testing.T) {
	// starting value is 50
	tests := []struct {
		name  string // description of this test case
		lines []string
		want  int
	}{
		{
			name:  "Test case 1",
			lines: []string{"R14", "L7", "R86", "L43"},
			want:  2,
		},
		{
			name:  "Test case 2",
			lines: []string{"L60", "R50", "L30", "R20"},
			want:  2,
		},
		{
			name:  "One rotation with multiple crossings",
			lines: []string{"R250"},
			want:  3,
		},
		{
			name:  "Multiple rotations with crossings",
			lines: []string{"L350", "R450"},
			want:  8,
		},
		{
			name:  "No crossings",
			lines: []string{"R10", "L20", "R30"},
			want:  0,
		},
		{
			name:  "Exact multiple of 100",
			lines: []string{"R100", "L200", "R300"},
			want:  6,
		},
		{
			name:  "multiple landings on zero",
			lines: []string{"R100", "R100", "R100", "R100", "R100"},
			want:  5,
		},
		{
			name:  "Start at 0 and move Left (Leaving 0)",
			lines: []string{"R50", "L10"},
			want:  1,
		},
		{
			name:  "Full rotation starting exactly at 0",
			lines: []string{"R50", "L100"},
			want:  2,
		},
		{
			name:  "Zig-zag landing and leaving 0",
			lines: []string{"L50", "L1", "R1", "R1"},
			want:  2,
		},
		{
			name:  "Large exact rotation on 0",
			lines: []string{"R50", "R200"},
			want:  3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := partTwo(tt.lines)
			if got != tt.want {
				t.Errorf("partTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}
