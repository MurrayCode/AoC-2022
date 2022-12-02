package day2

import (
	"strings"
)

func Part1(input string) int {
	rules := map[string]int{
		"AY": 8,
		"AX": 4,
		"AZ": 3,
		"BY": 5,
		"BZ": 9,
		"BX": 1,
		"CY": 2,
		"CX": 7,
		"CZ": 6,
	}

	turns := strings.Split(input, "\n")
	var result int
	for _, i := range turns {
		i = strings.ReplaceAll(i, " ", "")
		result += rules[i]
	}
	return result
}
