package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Part 1:", part1("day2.txt"))
}

func part1(input string) int {
	f, err := os.Open(input)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	scan := bufio.NewScanner(f)

	var finalScore int
	scores := map[string]int{
		"B X": 1,
		"C Y": 2,
		"A Z": 3,
		"A X": 4,
		"B Y": 5,
		"C Z": 6,
		"C X": 7,
		"A Y": 8,
		"B Z": 9,
	}

	for scan.Scan() {
		finalScore += scores[scan.Text()]
	}

	return finalScore
}
