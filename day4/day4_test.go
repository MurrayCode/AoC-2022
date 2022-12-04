package day4

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`

	exp := 2
	got := Part1(input)

	if exp != got {
		t.Errorf("expected %d, got %d", exp, got)
	}
}
func TestPart2(t *testing.T) {
	input := `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`

	exp := 4
	got := Part2(input)

	if exp != got {
		t.Errorf("expected %d, got %d", exp, got)
	}
}
