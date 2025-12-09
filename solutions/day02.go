package solutions

import (
	"strconv"
	"strings"
)

func SolveDay02(input string) int {
	var sum int
	ranges := strings.Split(input, ",")

	for _, r := range ranges {
		rmin, _ := strconv.Atoi(strings.Split(r, "-")[0])
		rmax, _ := strconv.Atoi(strings.Split(r, "-")[1])

		for i := rmin; i <= rmax; i++ {
			if !isValidId(i) {
				sum += i
			}
		}
	}

	return sum
}

func isValidId(id int) bool {
	is := strconv.Itoa(id)
	l := len(strconv.Itoa(id))
	for i := 0; i < l; i++ {
		segmentLength := i + 1
		if l%segmentLength != 0 {
			continue
		}
		s := is[0:segmentLength]
		if len(s) > l/2 {
			return true
		}
		match := true
		for j := 0; j < l; j += segmentLength {
			if s != is[j:j+segmentLength] {
				match = false
			}
		}
		if match {
			return false
		}
	}
	return true
}
