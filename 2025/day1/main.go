package main

import (
	"fmt"
	"strconv"
	"strings"

	_ "embed"
)

//go:embed input.txt
var input string

func main() {
	println("Part 1:", part1())
}

func rotateDial(pos int, direction, distance string) int {
	dialSize := 100
	if direction != "L" && direction != "R" {
		fmt.Errorf("invalid direction %q in %q", direction, distance)
	}

	newDistance, err := strconv.Atoi(distance)
	if err != nil {
		fmt.Errorf("invalid distance %q in %q: %v", direction, distance, err)
		return 100
	}
	newDistance = newDistance % dialSize

	switch direction {
	case "L":
		pos = (pos - newDistance + dialSize) % dialSize
	case "R":
		pos = (pos + newDistance) % dialSize
	}

	return pos
}

func part1() int {
	position := 50
	numberOfZeros := 0

	instructions := strings.Split(input, "\n")
	for _, inst := range instructions {
		if inst == "" {
			continue
		}
		direction := inst[:1]
		distance := inst[1:]

		position = rotateDial(position, direction, distance)
		if position == 0 {
			numberOfZeros++
		}
	}

	return numberOfZeros
}
