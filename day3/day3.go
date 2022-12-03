package day3

import (
	"strings"
	"unicode"
)

func Part1(input string) int {
	arr := strings.Split(input, "\n")
	var res int
	for _, i := range arr {
		one := strings.Split(i[0:len(i)/2], "")
		two := strings.Split(i[len(i)/2:], "")
		for x, j := range one {
			if strings.Contains(strings.Join(two, ""), string(j)) {
				res += GetNum(string(one[x]))
				break
			}
		}
	}
	return res
}

func GetNum(letter string) int {
	if unicode.IsUpper([]rune(letter)[0]) {
		return int([]byte(letter)[0]) - 38
	}
	return int([]byte(letter)[0]) - 96
}
