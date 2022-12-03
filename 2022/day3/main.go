package main

import (
	"fmt"
	"os"
	"strings"
)

var (
	alphabet = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
)

func main() {
	fmt.Println("answer to part 1 is: ", part1("day3.txt"))

	fmt.Println("answer to part 2 is: ", part2("day3.txt"))
}

func part1(input string) int {
	i, err := os.ReadFile(input)
	if err != nil {
		panic(err)
	}

	// split the input into a slice of strings
	lines := strings.Split(string(i), "\n")

	// for each line split the line in half and store in a slice
	var sum int
	for _, line := range lines {
		// split the line in half
		half := len(line) / 2
		// store the first half in a slice
		firstHalf := line[:half]
		// store the second half in a slice
		secondHalf := line[half:]

		// if a char exists in the first half and the second print it
		for _, char := range firstHalf {
			if strings.Contains(secondHalf, string(char)) {
				// find the char in the priority map and print it
				// add the char to the sum
				sum += getPriority(string(char))
				break
			}
		}
	}
	return sum
}

func part2(input string) int {
	i, err := os.ReadFile(input)
	if err != nil {
		panic(err)
	}

	// split the input into a slice of strings
	lines := strings.Split(string(i), "\n")

	// for each line split the line in half and store in a slice
	var sum int
	for l := 0; l < len(lines)-3; l += 3 {
		// find the intersection of the three elves
		sum += getPriority(intersection(lines[l], lines[l+1], lines[l+2]))
	}
	return sum
}

func intersection(elf1, elf2, elf3 string) string {
	for _, char := range elf1 {
		if strings.Contains(elf2, string(char)) && strings.Contains(elf3, string(char)) {
			return string(char)
		}
	}
	return ""
}

func getPriority(character string) int {
	for i, char := range alphabet {
		if char == character {
			return i + 1
		}
	}
	return 0
}
