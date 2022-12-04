package day4

import (
	"strconv"
	"strings"
)

func Part1(input string) int {
	arr := strings.Split(input, "\n")
	var res int
	for _, i := range arr {
		nums := strings.Split(i, ",")
		one := strings.Split(nums[0], "-")
		two := strings.Split(nums[1], "-")
		rangeOneMin, _ := strconv.Atoi(one[0])
		rangeOneMax, _ := strconv.Atoi(one[1])

		rangeTwoMin, _ := strconv.Atoi(two[0])
		rangeTwoMax, _ := strconv.Atoi(two[1])

		if rangeOneMin >= rangeTwoMin && rangeOneMax <= rangeTwoMax {
			res += 1
		} else if rangeTwoMin >= rangeOneMin && rangeTwoMax <= rangeOneMax {
			res += 1
		}
	}
	return res
}
func Part2(input string) int {
	arr := strings.Split(input, "\n")
	var res int
	for _, i := range arr {
		nums := strings.Split(i, ",")
		one := strings.Split(nums[0], "-")
		two := strings.Split(nums[1], "-")
		rangeOneMin, _ := strconv.Atoi(one[0])
		rangeOneMax, _ := strconv.Atoi(one[1])

		rangeTwoMin, _ := strconv.Atoi(two[0])
		rangeTwoMax, _ := strconv.Atoi(two[1])

		if rangeOneMin >= rangeTwoMin && rangeOneMin <= rangeTwoMax || rangeOneMax <= rangeTwoMax && rangeOneMax >= rangeTwoMin {
			res += 1
		} else if rangeTwoMin >= rangeOneMin && rangeTwoMin <= rangeOneMax || rangeTwoMax <= rangeOneMax && rangeTwoMax >= rangeOneMin {
			res += 1
		}
	}
	return res
}
