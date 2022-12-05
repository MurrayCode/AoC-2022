package day5

import "testing"

// func TestPart1(t *testing.T) {
// 	input := `    [D]
// [N] [C]
// [Z] [M] [P]
//  1   2   3

// move 1 from 2 to 1
// move 3 from 1 to 3
// move 2 from 2 to 1
// move 1 from 1 to 2`

// 	got := Part1(input)
// 	exp := "CMZ"

// 	if got != exp {
// 		t.Errorf("expected %s, got %s", exp, got)
// 	}
// }

func TestPart2(t *testing.T) {
	input := `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`

	got := Part2(input)
	exp := "MCD"

	if got != exp {
		t.Errorf("expected %s, got %s", exp, got)
	}
}
