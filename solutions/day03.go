package solutions

import "strconv"

func SolveDay03(lines []string) int {
	var s int
	for _, line := range lines {
		var maxJoltage int
		for i := 0; i < len(line)-1; i++ {
			for j := i + 1; j < len(line); j++ {
				joltage, _ := strconv.Atoi(string(line[i]) + string(line[j]))
				if joltage > maxJoltage {
					maxJoltage = joltage
				}
			}
		}
		s += maxJoltage
	}
	return s
}
