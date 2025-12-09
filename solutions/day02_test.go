package solutions

import (
	"fmt"
	"testing"

	"dev.lezurex.advent-of-code25/utils"
)

func TestSolveDay02(t *testing.T) {
	i, _ := utils.ReadFile("./inputs/day02.txt")

	s := SolveDay02(i[0])

	t.Logf("Solution Day 02: %d\n", s)
}

func TestSolveDay02Sample(t *testing.T) {
	i := "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"
	want := 4174379265

	s := SolveDay02(i)

	if want != s {
		t.Errorf("solution incorrect, got: %d, want %d.", s, want)
	}
}

func TestIsValidId(t *testing.T) {
	tcs := []struct {
		id   int
		want bool
	}{
		{1234, true},
		{12345, true},
		{12131415, true},

		{12341234, false},
		{123123123, false},
		{1212121212, false},
		{1111111, false},

		{824824888, true},
	}
	for _, tc := range tcs {
		t.Run(fmt.Sprintf("%d", tc.id), func(t *testing.T) {
			if got := isValidId(tc.id); got != tc.want {
				t.Errorf("isValidId(%d) = %t, want %t", tc.id, got, tc.want)
			}
		})
	}
}
