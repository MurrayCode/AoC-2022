package day3

import "testing"

func TestPart1(t *testing.T) {
	input := `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDwA`

	exp := 157
	got := Part1(input)

	if exp != got {
		t.Errorf("expected %d, got %d", exp, got)
	}
}

func TestGetNum(t *testing.T) {
	var tests = []struct {
		name  string
		input string
		ans   int
	}{
		{
			"Capital A returns 27",
			"A",
			27,
		},
		{
			"Lowercase a returns 1",
			"a",
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetNum(tt.input)
			exp := tt.ans
			if got != exp {
				t.Errorf("got %d, expected %d", got, exp)
			}
		})

	}
}
