package day1

import (
	"sort"
	"strconv"
	"strings"
)

func Part1(input string) int {
	elfs := strings.Split(input, "\n\n")
	var results []int

	for _, i := range elfs {
		calories := strings.Split(i, "\n")
		var total int
		for _, j := range calories {
			num, _ := strconv.Atoi(j)
			total += num
		}
		results = append(results, total)
	}
	sort.Slice(results, func(i, j int) bool {
		return results[i] > results[j]
	})
	return results[0]
}

func Part2(input string) int {
	elfs := strings.Split(input, "\n\n")
	var results []int

	for _, i := range elfs {
		calories := strings.Split(i, "\n")
		var total int
		for _, j := range calories {
			num, _ := strconv.Atoi(j)
			total += num
		}
		results = append(results, total)
	}
	sort.Slice(results, func(i, j int) bool {
		return results[i] > results[j]
	})

	res := results[0] + results[1] + results[2]
	return res
}
