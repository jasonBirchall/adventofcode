package main

import (
	"fmt"
	"os"
	"strings"
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
				sum += priority[string(char)]
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
	// fmt.Println(lines[0])

	// for each line split the line in half and store in a slice
	var sum int
	for l := 0; l < len(lines)-3; l += 3 {
		// find the intersection of the three elves
		// intersection := intersection(lines[l], lines[l+1], lines[l+2])
		sum += priority[intersection(lines[l], lines[l+1], lines[l+2])]
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

// define priority of the items in the input
var (
	priority = map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 6,
		"g": 7,
		"h": 8,
		"i": 9,
		"j": 10,
		"k": 11,
		"l": 12,
		"m": 13,
		"n": 14,
		"o": 15,
		"p": 16,
		"q": 17,
		"r": 18,
		"s": 19,
		"t": 20,
		"u": 21,
		"v": 22,
		"w": 23,
		"x": 24,
		"y": 25,
		"z": 26,
		"A": 27,
		"B": 28,
		"C": 29,
		"D": 30,
		"E": 31,
		"F": 32,
		"G": 33,
		"H": 34,
		"I": 35,
		"J": 36,
		"K": 37,
		"L": 38,
		"M": 39,
		"N": 40,
		"O": 41,
		"P": 42,
		"Q": 43,
		"R": 44,
		"S": 45,
		"T": 46,
		"U": 47,
		"V": 48,
		"W": 49,
		"X": 50,
		"Y": 51,
		"Z": 52,
	}
)
