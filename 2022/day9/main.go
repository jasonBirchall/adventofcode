package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed day9.txt
var input string

type coordinates struct {
	x, y int
}

func main() {
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input string) int {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")

	head := coordinates{0, 0}
	tail := coordinates{0, 0}
	grid := make(map[coordinates]bool)
	grid[head] = true
	for _, line := range lines {
		str := strings.Split(line, " ")
		direction := str[0]
		distance, err := strconv.Atoi(str[1])
		if err != nil {
			fmt.Printf("error: %v", err)
		}
		for i := 0; i < distance; i++ {
			switch direction {
			case "U":
				head.y++
			case "R":
				head.x++
			case "L":
				head.x--
			case "D":
				head.y--
			}

			if abs(head.x-tail.x) > 1 || abs(head.y-tail.y) > 1 {
				tail = coordinates{tail.x + sign(head.x-tail.x), tail.y + sign(head.y-tail.y)}
			}

			grid[tail] = true
		}
	}

	return len(grid)

}

func part2(input string) int {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")

	const totalKnots = 10
	knots := make([]coordinates, totalKnots)
	grid := make(map[coordinates]bool)
	grid[coordinates{0, 0}] = true

	for _, line := range lines {
		str := strings.Split(line, " ")
		direction := str[0]
		distance, err := strconv.Atoi(str[1])
		if err != nil {
			fmt.Printf("error: %v", err)
		}
		for i := 0; i < distance; i++ {
			switch direction {
			case "U":
				knots[0].y++
			case "R":
				knots[0].x++
			case "L":
				knots[0].x--
			case "D":
				knots[0].y--
			}

			for i := 1; i < totalKnots; i++ {
				if abs(knots[i-1].x-knots[i].x) > 1 || abs(knots[i-1].y-knots[i].y) > 1 {
					knots[i] = coordinates{knots[i].x + sign(knots[i-1].x-knots[i].x), knots[i].y + sign(knots[i-1].y-knots[i].y)}
				}
			}

			grid[knots[9]] = true
		}
	}

	return len(grid)
}

func sign(x int) int {
	if x < 0 {
		return -1
	}
	if x == 0 {
		return 0
	}
	return 1
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
