package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed day1.txt
var input string

func main() {
	fmt.Println("The Elf carrying the most amount of calories is", part1(input))

	fmt.Println("The top three elves are carrying", part2(input), "calories")
}

func part1(input string) int {
	var highest int
	sacks := strings.Split(input, "\n\n")

	for _, sack := range sacks {
		var total int
		lines := strings.Split(sack, "\n")
		for _, line := range lines {
			if line == "" {
				continue
			}
			lineInt, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			total += lineInt
		}
		if total > highest {
			highest = total
		}
	}
	return highest
}

func part2(input string) int {
	// create a collection of all totals
	var totals []int

	// grab the top three
	sacks := strings.Split(input, "\n\n")
	for _, sack := range sacks {
		var total int
		lines := strings.Split(sack, "\n")
		for _, line := range lines {
			if line == "" {
				continue
			}
			lineInt, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			total += lineInt
		}
		totals = append(totals, total)
	}

	sort.Ints(totals)
	return totals[len(totals)-1] + totals[len(totals)-2] + totals[len(totals)-3]
}
