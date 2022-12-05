package day5

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Stack struct {
	items []rune
}

func (s *Stack) Push(item rune) []rune {
	s.items = append(s.items, item)
	return s.items
}

func (s *Stack) Pop() (r rune) {
	r = s.items[0]
	s.items = s.items[1:len(s.items)]
	return
}

func (s *Stack) addToBottom(r rune) {
	s.items = append([]rune{r}, s.items...)
}

func PopulateStacks(input string) [9]Stack {
	items := strings.Split(input, "\n\n")[0]
	itemsArr := strings.Split(items, "\n")
	stacks := [9]Stack{}
	iter := -1
	for _, item := range itemsArr {
		iter = -1
		if item == " 1   2   3   4   5   6   7   8   9 " {
			break
		}
		for i := 1; i <= len(item)-2; i += 4 {
			iter++
			if item[i] != ' ' {

				stacks[iter].Push(rune(item[i]))
			}
			if iter == 8 {
				iter = -1
			}
		}
	}
	return stacks
}

func Part1(input string) string {
	stacks := PopulateStacks(input)
	instructions := strings.Split(input, "\n\n")[1]
	lines := strings.Split(instructions, "\n")
	for _, line := range lines {
		line = regexp.MustCompile(`[^0-9]+`).ReplaceAllString(line, " ")
		nums := strings.Split(line, " ")
		amount, _ := strconv.Atoi(nums[1])
		from, _ := strconv.Atoi(nums[2])
		to, _ := strconv.Atoi(nums[3])
		for i := 0; i < amount; i++ {
			moving := stacks[from-1].Pop()
			stacks[to-1].addToBottom(moving)
		}
	}
	var res string
	for _, stack := range stacks {
		res += fmt.Sprintf("%c", stack.items[0])

	}
	return res
}
