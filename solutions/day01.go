package solutions

import "strconv"

func SolveDay01(lines []string) int {
	dial := Dial{position: 50}

	for _, l := range lines {
		direction := string(l[0])
		distance, _ := strconv.Atoi(l[1:])
		for range distance {
			dial.Click(direction)
		}
	}

	return dial.count0
}

type Dial struct {
	position int
	count0   int
}

func (d *Dial) Click(direction string) {
	if direction == "L" {
		if d.position == 0 {
			d.position = 100
		}
		d.position--
	} else if direction == "R" {
		if d.position == 99 {
			d.position = -1
		}
		d.position++
	}
	if d.position == 0 {
		d.count0++
	}
}
