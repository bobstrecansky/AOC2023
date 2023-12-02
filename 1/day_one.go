package day_one

import (
	"fmt"
	"strconv"
	"strings"
)

func calibrationSum(calibrationDocument []string) int {
	var res int
	charInt := map[int]string{1: "one", 2: "two", 3: "three", 4: "four", 5: "five", 6: "six", 7: "seven", 8: "eight", 9: "nine"}

	// Iterate through all strings in the input file
	for i := range calibrationDocument {
		nums := calibrationDocument[i]
		for k := 1; k < 10; k++ {
			// start PART2 add
			// if the string value of 0-9 is present, write NUM[0-9]NUM
			if strings.Contains(calibrationDocument[i], charInt[k]) {
				nums = strings.Replace(nums, charInt[k], charInt[k]+strconv.Itoa(k)+charInt[k], -1)
			}
			// end PART2 add
		}

		var sb strings.Builder

		// If the character in the string is between 0 and 9, write it as a string to the string builder
		for j := range nums {
			if nums[j] > 47 && nums[j] < 58 {
				sb.WriteString(string(nums[j]))
			}
		}

		firstChar := sb.String()[0]
		lastChar := sb.String()[len(sb.String())-1]
		calibrationValueStr := fmt.Sprint(string(firstChar), string(lastChar))
		calibrationValueInt, _ := strconv.Atoi(calibrationValueStr)
		res += calibrationValueInt
	}

	return res
}
