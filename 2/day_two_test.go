package day_two

import (
	"testing"
)

type test struct {
	input                            string
	result_part_one, result_part_two int
}

func TestCubeConundrum(t *testing.T) {
	partOneInput := "day_two_input.txt"
	tests := []test{
		{input: partOneInput, result_part_one: 2771, result_part_two: 70924},
	}

	for _, tc := range tests {
		got_part_one := partOne(tc.input)
		got_part_two := partTwo(tc.input)

		if got_part_one != tc.result_part_one {
			t.Fatalf("input: %v, expected: %v, got: %v", tc.input, tc.result_part_one, got_part_one)
		}
		if got_part_two != tc.result_part_two {
			t.Fatalf("input: %v, expected: %v, got: %v", tc.input, tc.result_part_two, got_part_two)
		}
	}
}
