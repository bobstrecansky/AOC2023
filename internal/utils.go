package internal

import (
	"os"
	"strings"
)

func ReadInput(input string) []string {
	content, err := os.ReadFile(input)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\n")
	return lines
}
