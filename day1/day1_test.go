package day1

import "testing"

func TestPart1(t *testing.T) {
	input := `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`

	exp := 24000
	got := Part1(input)
	if got != exp {
		t.Errorf("expected %d, got %d", exp, got)
	}
}

func TestPart2(t *testing.T) {
	input := `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`

	exp := 45000
	got := Part2(input)
	if got != exp {
		t.Errorf("expected %d, got %d", exp, got)
	}
}
