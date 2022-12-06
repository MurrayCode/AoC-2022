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
		one := strings.ReplaceAll(str, string(str[0]), "")
		two := strings.ReplaceAll(str, string(str[1]), "")
		three := strings.ReplaceAll(str, string(str[2]), "")
		four := strings.ReplaceAll(str, string(str[3]), "")
		five := strings.ReplaceAll(str, string(str[4]), "")
		six := strings.ReplaceAll(str, string(str[5]), "")
		seven := strings.ReplaceAll(str, string(str[6]), "")
		eight := strings.ReplaceAll(str, string(str[7]), "")
		nine := strings.ReplaceAll(str, string(str[8]), "")
		ten := strings.ReplaceAll(str, string(str[9]), "")
		eleven := strings.ReplaceAll(str, string(str[10]), "")
		twelve := strings.ReplaceAll(str, string(str[11]), "")
		thirteen := strings.ReplaceAll(str, string(str[12]), "")
		if len(one) == 13 && len(two) == 13 &&
			len(three) == 13 && len(four) == 13 &&
			len(five) == 13 && len(six) == 13 &&
			len(seven) == 13 && len(eight) == 13 &&
			len(nine) == 13 && len(ten) == 13 &&
			len(eleven) == 13 && len(twelve) == 13 &&
			len(thirteen) == 13 {
			return i + 14
		} else {
			continue
		}
	}
	return 0
}
