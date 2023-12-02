package day_one

import (
	"testing"

	"github.com/bobstrecansky/AOC2023/internal"
)

type test struct {
	input  []string
	result int
}

func TestCalibrationSum(t *testing.T) {
	lines := internal.ReadInput("day_one_data.txt")
	tests := []test{
		{input: lines, result: 54719},
	}

	for _, tc := range tests {
		got := calibrationSum(tc.input)
		want := tc.result
		if got != want {
			t.Fatalf("input: %v, expected: %v, got: %v", tc.input, tc.result, got)
		}
	}
}
