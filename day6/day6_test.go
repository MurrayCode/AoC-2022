package day6

import "testing"

func TestPart1(t *testing.T) {
	var tests = []struct {
		name   string
		inputs string
		ans    int
	}{
		{
			"first",
			"bvwbjplbgvbhsrlpgdmjqwftvncz",
			5,
		},
		{
			"second",
			"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			10,
		},
		{
			"third",
			"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Part1(tt.inputs)
			exp := tt.ans
			if got != exp {
				t.Errorf("got %d, expected %d", got, exp)
			}
		})

	}
}

func TestPart2(t *testing.T) {
	var tests = []struct {
		name   string
		inputs string
		ans    int
	}{
		{
			"first",
			"bvwbjplbgvbhsrlpgdmjqwftvncz",
			23,
		},
		{
			"second",
			"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			29,
		},
		{
			"third",
			"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			26,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Part2(tt.inputs)
			exp := tt.ans
			if got != exp {
				t.Errorf("got %d, expected %d", got, exp)
			}
		})

	}
}
