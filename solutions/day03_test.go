package solutions

import (
	"testing"

	"dev.lezurex.advent-of-code25/utils"
)

func TestSolveDay03(t *testing.T) {
	i, _ := utils.ReadFile("./inputs/day03.txt")

	s := SolveDay03(i)

	t.Logf("Solution Day 03: %d\n", s)
}

func TestSolveDay03Sample(t *testing.T) {
	i := []string{
		"987654321111111",
		"811111111111119",
		"234234234234278",
		"818181911112111",
	}
	want := 357

	s := SolveDay03(i)

	if want != s {
		t.Errorf("solution incorrect, got: %d, want %d.", s, want)
	}
}
