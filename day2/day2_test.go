package day2

import "testing"

func TestPart1(t *testing.T) {
	input := `A Y
B X
C Z`

	exp := 15
	got := Part1(input)

	if exp != got {
		t.Errorf("expected %d, got %d", exp, got)
	}
}
