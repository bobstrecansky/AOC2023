package gear_ratios

import (
	"strconv"
	"unicode"

	"github.com/bobstrecansky/AOC2023/internal"
)

type coordinates struct {
	x int
	y int
}

type Part struct {
	number  int
	engines []coordinates
}

func partNumberSum(file string) int {
	matrix := createMatrix(file)
	return calculatePartsSum(matrix)
}

func gearRatio(file string) int {
	matrix := createMatrix(file)
	parts, engines := getPartsAndEngines(matrix)
	return calculateGearsRatio(parts, engines, matrix)
}

func createMatrix(file string) [][]rune {
	lines := internal.ReadInput(file)
	var matrix [][]rune
	for _, line := range lines {
		var row []rune
		for _, char := range line {
			row = append(row, char)
		}
		matrix = append(matrix, row)
	}
	return matrix
}

func checkForSpecialChars(c []coordinates, matrix [][]rune) bool {
	for _, coord := range c {
		char := matrix[coord.x][coord.y]
		if char != '.' && !unicode.IsDigit(char) {
			return true
		}
	}
	return false
}

func coordsToCheck(r, start, end int, matrix [][]rune) []coordinates {
	rMax, cMax := len(matrix)-1, len(matrix[0])-1
	var l []coordinates
	for col := start - 1; col <= end+1; col++ {
		l = append(
			l,
			coordinates{x: r - 1, y: col},
			coordinates{x: r + 1, y: col},
		)
	}
	l = append(l, coordinates{x: r, y: start - 1}, coordinates{x: r, y: end + 1})
	var parsedList []coordinates
	for _, v := range l {
		row, col := v.x, v.y
		if row >= 0 && row <= rMax && col >= 0 && col <= cMax {
			parsedList = append(parsedList, v)
		}
	}
	return parsedList
}

func calculatePartsSum(matrix [][]rune) int {
	var res, start, end int
	var number string
	var parsingNumber bool
	for i, row := range matrix {
		start = 0
		number, parsingNumber = "", false
		for j, char := range row {
			if unicode.IsDigit(char) {
				number += string(char)
			}
			if unicode.IsDigit(char) && !parsingNumber {
				parsingNumber = true
				start = j
			}
			if !unicode.IsDigit(char) && parsingNumber {
				end = j - 1
				// check
				coordsToCheck := coordsToCheck(i, start, end, matrix)
				if checkForSpecialChars(coordsToCheck, matrix) {
					n, _ := strconv.Atoi(number)
					res += n
				}
				number, start, parsingNumber = "", 0, false
			}
		}
		if parsingNumber {
			coordsToCheck := coordsToCheck(i, start, len(row)-1, matrix)
			if checkForSpecialChars(coordsToCheck, matrix) {
				n, _ := strconv.Atoi(number)
				res += n
			}
		}
	}
	return res
}

func calculateGearsRatio(parts [][]Part, engines [][]coordinates, matrix [][]rune) int {
	var sum int
	for i, row := range engines {
		for _, engine := range row {
			var engineParts []Part
			if i > 0 {
				engineParts = append(engineParts, getEngineParts(parts[i-1], engine)...)
			}
			engineParts = append(engineParts, getEngineParts(parts[i], engine)...)
			if i < (len(engines) - 1) {
				engineParts = append(engineParts, getEngineParts(parts[i+1], engine)...)
			}
			if len(engineParts) == 2 {
				sum += engineParts[0].number * engineParts[1].number
			}
		}
	}
	return sum
}

func getPartsAndEngines(matrix [][]rune) ([][]Part, [][]coordinates) {
	parts := make([][]Part, len(matrix))
	engines := make([][]coordinates, len(matrix))
	var start, end int
	var number string
	var parsingNumber bool
	for i, row := range matrix {
		start = 0
		number, parsingNumber = "", false
		for j, char := range row {
			if char == '*' {
				engines[i] = append(engines[i], coordinates{x: i, y: j})
			}
			if unicode.IsDigit(char) {
				number += string(char)
			}
			if unicode.IsDigit(char) && !parsingNumber {
				parsingNumber = true
				start = j
			}
			if !unicode.IsDigit(char) && parsingNumber {
				end = j - 1
				// check
				enginesAround := getEnginesAround(i, start, end, matrix)
				n, _ := strconv.Atoi(number)
				part := Part{number: n, engines: enginesAround}
				parts[i] = append(parts[i], part)
				number, start, parsingNumber = "", 0, false
			}
		}
		if parsingNumber {
			enginesAround := getEnginesAround(i, start, len(row)-1, matrix)
			n, _ := strconv.Atoi(number)
			part := Part{number: n, engines: enginesAround}
			parts[i] = append(parts[i], part)
		}
	}
	return parts, engines
}

func getEnginesAround(rowIndex, from, to int, matrix [][]rune) []coordinates {
	var engines []coordinates
	coordToCheck := coordsToCheck(rowIndex, from, to, matrix)
	for _, coord := range coordToCheck {
		char := matrix[coord.x][coord.y]
		if char == '*' {
			engines = append(engines, coord)
		}
	}
	return engines
}

func getEngineParts(parts []Part, engine coordinates) []Part {
	var engineParts []Part
	for _, part := range parts {
		for _, e := range part.engines {
			if e.x == engine.x && e.y == engine.y {
				engineParts = append(engineParts, part)
			}
		}
	}
	return engineParts
}
