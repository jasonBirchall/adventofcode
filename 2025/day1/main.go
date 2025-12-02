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
	println("Part 2:", part2())
}

func rotateDial(pos int, direction, distance string) (int, int) {
	dialSize := 100
	if direction != "L" && direction != "R" {
		fmt.Errorf("invalid direction %q in %q", direction, distance)
	}

	newDistance, err := strconv.Atoi(distance)
	if err != nil {
		fmt.Errorf("invalid distance %q in %q: %v", direction, distance, err)
	}

	hits := 0
	if newDistance > 0 {
		if pos == 0 {
			// it has started at zero
			hits = newDistance / dialSize
		} else {
			var firstHit int
			switch direction {
			case "L":
				firstHit = pos
			case "R":
				firstHit = dialSize - pos
			}

			if newDistance >= firstHit {
				hits = 1 + (newDistance-firstHit)/dialSize
			}
		}
	}

	newDistance = newDistance % dialSize
	println("Rotating from position", pos, "to the", direction, "by", newDistance)
	switch direction {
	case "L":
		pos = (pos - newDistance + dialSize) % dialSize
	case "R":
		pos = (pos + newDistance) % dialSize
	}

	return pos, hits
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

		position, _ = rotateDial(position, direction, distance)
		if position == 0 {
			numberOfZeros++
		}
	}

	return numberOfZeros
}

func part2() int {
	position := 50
	numberOfClicks := 0

	instructions := strings.Split(input, "\n")
	for _, inst := range instructions {
		if inst == "" {
			continue
		}
		direction := inst[:1]
		distance := inst[1:]

		var hits int
		position, hits = rotateDial(position, direction, distance)
		numberOfClicks += hits
	}
	return numberOfClicks
}
