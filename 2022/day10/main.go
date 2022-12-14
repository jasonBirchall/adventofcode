package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed day10.txt
var input string

func main() {
	fmt.Println("Answer to part 1:\n", part1(input))
	fmt.Println("Answer to part 2:\n", part2(input))
}

func part1(input string) int {
	var (
		x     = 1
		cycle = 1
		total = 0
		n     string
	)

	instructions := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	for _, instruction := range instructions {
		instruction := strings.Split(instruction, " ")
		i := instruction[0]
		if i != "noop" {
			n = instruction[1]
		}

		if cycle%40 == 20 {
			total += cycle * x
		}
		cycle++

		if i == "addx" {
			if cycle%40 == 20 {
				total += cycle * x
			}
			cycle++
			n, _ := strconv.Atoi(n)
			x += n
		}

	}
	return total
}

func part2(input string) [6]string {
	var (
		x        = 1
		cycle    = 1
		rows     = [6]string{}
		rowIndex = 0
		colIndex = 0
	)

	instructions := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	for _, instruction := range instructions {
		num := 0
		iterations := 1
		if instruction != "noop" {
			iterations = 2
			num, _ = strconv.Atoi(strings.Split(instruction, " ")[1])
		}
		for i := 0; i < iterations; i++ {
			// draw the pixel
			char := "."
			if colIndex == x || colIndex == x-1 || colIndex == x+1 {
				char = "#"
			}
			rows[rowIndex] += char
			cycle++
			colIndex++

			if colIndex > 39 {
				colIndex = 0
				rowIndex++
			}
		}
		x += num
	}
	return rows
}
