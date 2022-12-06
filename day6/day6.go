package day6

import (
	"strings"
)

func Part1(input string) int {
	for i := 0; i < len(input)-3; i++ {
		str := input[i : i+4]
		one := strings.ReplaceAll(str, string(str[0]), "")
		two := strings.ReplaceAll(str, string(str[1]), "")
		three := strings.ReplaceAll(str, string(str[2]), "")
		if len(one) == 3 && len(two) == 3 && len(three) == 3 {
			return i + 4
		} else {
			continue
		}
	}
	return 0
}

func Part2(input string) int {
	for i := 0; i < len(input)-13; i++ {
		str := input[i : i+14]
		counter := 0
		for j := 0; j < 14; j++ {
			check := strings.ReplaceAll(str, string(str[j]), "")
			if len(check) == 13 {
				counter++
			}
			if counter == 13 {
				return i + 14
			}
		}
	}
	return 0
}
