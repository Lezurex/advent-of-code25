package solutions

import (
	"testing"

	"dev.lezurex.advent-of-code25/utils"
)

func TestSolveDay01(t *testing.T) {
	i, _ := utils.ReadFile("./inputs/day01.txt")

	s := SolveDay01(i)

	t.Logf("Solution Day 01: %d\n", s)
}

func TestSolveDay01Sample(t *testing.T) {
	i := []string{"L68", "L30", "R48", "L5", "R60", "L55", "L1", "L99", "R14", "L82"}
	want := 6

	s := SolveDay01(i)

	if s != want {
		t.Errorf("solution incorrect, got: %d, want: %d.", s, want)
	}
}
