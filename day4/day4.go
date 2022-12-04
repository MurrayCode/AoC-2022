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

		range1 := makeRange(rangeOneMin, rangeOneMax)
		range2 := makeRange(rangeTwoMin, rangeTwoMax)

		if contains(range2, range1[0]) && contains(range2, range1[len(range1)-1]) {
			res += 1
		} else if contains(range1, range2[0]) && contains(range1, range2[len(range2)-1]) {
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

		range1 := makeRange(rangeOneMin, rangeOneMax)
		range2 := makeRange(rangeTwoMin, rangeTwoMax)

		if contains(range2, range1[0]) || contains(range2, range1[len(range1)-1]) {
			res += 1
		} else if contains(range1, range2[0]) || contains(range1, range2[len(range2)-1]) {
			res += 1
		}

	}
	return res
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func contains(s []int, num int) bool {
	for _, v := range s {
		if v == num {
			return true
		}
	}

	return false
}
