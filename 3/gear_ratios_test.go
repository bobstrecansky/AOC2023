package gear_ratios

import "testing"

type test struct {
	input                  string
	result_one, result_two int
}

func TestGearRatios(t *testing.T) {
	tests := []test{
		{input: "day_three_data.txt", result_one: 539590, result_two: 80703636},
	}

	for _, tc := range tests {
		gotPartOne := partNumberSum(tc.input)
		gotPartTwo := gearRatio(tc.input)

		if gotPartOne != tc.result_one {
			t.Fatalf("input: %v, expected: %v, got: %v", tc.input, tc.result_one, gotPartOne)
		}

		if gotPartTwo != tc.result_two {
			t.Fatalf("input %v, expected: %v, got: %v", tc.input, tc.result_two, gotPartTwo)
		}
	}
}
