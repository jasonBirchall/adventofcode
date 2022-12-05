package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed day4.txt
var input string

func main() {
	fmt.Println("the number of duplicate assignments is", part1(input))
	fmt.Println("the number of non-overlapping assignments is", part2(input))
}

func part1(input string) (result int) {
	// split the input into a slice of strings
	trim := strings.TrimSuffix(input, "\n")
	lines := strings.Split(trim, "\n")

	for _, line := range lines {
		var firstStart, firstEnd, secondStart, secondEnd int
		fmt.Sscanf(line, "%d-%d,%d-%d", &firstStart, &firstEnd, &secondStart, &secondEnd)
		if (firstStart <= secondStart && secondEnd <= firstEnd) || (secondStart <= firstStart && firstEnd <= secondEnd) {
			fmt.Println(line)
			fmt.Println("overlap")
			result++
		}
	}

	return
}

func part2(input string) (result int) {
	// split the input into a slice of strings
	trim := strings.TrimSuffix(input, "\n")
	lines := strings.Split(trim, "\n")

	for _, line := range lines {
		var firstStart, firstEnd, secondStart, secondEnd int
		fmt.Sscanf(line, "%d-%d,%d-%d", &firstStart, &firstEnd, &secondStart, &secondEnd)
		if !(firstEnd < secondStart || secondEnd < firstStart) {
			fmt.Println(line)
			fmt.Println("overlap")
			result++
		}
	}

	return

}
