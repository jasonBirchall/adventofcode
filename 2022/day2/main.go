package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"os"
)

func main() {
	part1Scores := map[string]int{
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

	fmt.Println("Part 1:", solve("day2.txt", part1Scores))

	part2Scores := map[string]int{
		"B X": 1,
		"C X": 2,
		"A X": 3,
		"A Y": 4,
		"B Y": 5,
		"C Y": 6,
		"C Z": 7,
		"A Z": 8,
		"B Z": 9,
	}
	fmt.Println("Part 2:", solve("day2.txt", part2Scores))
}

func solve(input string, scores map[string]int) (finalScore int) {
	f, err := os.Open(input)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	scan := bufio.NewScanner(f)

	for scan.Scan() {
		finalScore += scores[scan.Text()]
	}

	return
}
