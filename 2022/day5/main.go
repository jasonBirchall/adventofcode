package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed day5.txt
var input string

func main() {
	// amend input to get just the table data
	i := strings.Split(input, "\n\n")

	fmt.Println("Part 1:", part1(buildStack(i[0]), i[1]))
	fmt.Println("Part 2:", part2(buildStack(i[0]), i[1]))
}

func part1(stacks map[int][]string, instructions string) []string {
	for _, s := range strings.Split(instructions, "\n") {
		var howmany, from, to int

		fmt.Sscanf(s, "move %d from %d to %d", &howmany, &from, &to)
		// decrease values by 1 to get the correct index
		from--
		to--

		for i := 0; i < howmany; i++ {
			stacks[to] = append(stacks[to], stacks[from][len(stacks[from])-1])
			stacks[from] = stacks[from][:len(stacks[from])-1]
		}

	}

	// add the last element of each stack to the result
	var result []string
	for i := 0; i < len(stacks); i++ {
		result = append(result, stacks[i][len(stacks[i])-1])
	}

	return result
}

func part2(stacks map[int][]string, instructions string) []string {
	for _, s := range strings.Split(instructions, "\n") {
		var howmany, from, to int

		fmt.Sscanf(s, "move %d from %d to %d", &howmany, &from, &to)
		// decrease values by 1 to get the correct index
		from--
		to--

		for _, stack := range stacks[from][len(stacks[from])-howmany:] {
			stacks[to] = append(stacks[to], stack)
			stacks[from] = stacks[from][:len(stacks[from])-1]
		}

	}

	// add the last element of each stack to the result
	var result []string
	for i := 0; i < len(stacks); i++ {
		result = append(result, stacks[i][len(stacks[i])-1])
	}

	return result
}

func buildStack(input string) map[int][]string {
	input = strings.ReplaceAll(input, "    ", "[-] ")
	input = strings.ReplaceAll(input, "  ", " ")
	input = strings.ReplaceAll(input, "][", "] [")

	stackInput := strings.Split(input, "\n")
	n := strings.Count(stackInput[0], "]")

	// create a map of stacks
	stacks := make(map[int][]string, n)

	for _, stack := range stackInput[:len(stackInput)-1] {
		s := strings.Trim(strings.TrimSpace(stack), "[]")
		r := strings.Split(s, "] [")
		for i, v := range r {
			if v != "-" {
				stacks[i] = append([]string{v}, stacks[i]...)
			}
		}
	}

	return stacks
}
