package main

import (
	_ "embed"
	"fmt"
)

//go:embed day6.txt
var day6 string

func main() {
	// part 1
	fmt.Println("Part 1: The starter sequence is", findStarterMessage(day6, 4))
	// part 2
	fmt.Println("Part 2: The starter sequence is", findStarterMessage(day6, 14))

}

func findStarterMessage(input string, length int) int {
	var slide string
	var count int
	m := make(map[string]int)

	for i := 0; i < len(input); i++ {
		slide = string(input[i])

		if v, found := m[slide]; found {
			i = v
			m = make(map[string]int)
			count = 0

			continue
		}

		i := i
		m[slide] = i
		count++

		if count > length-1 {
			return i + 1
		}
	}
	return -1

}
