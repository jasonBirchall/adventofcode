package main

import (
	"fmt"
	"strconv"
	"strings"

	_ "embed"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println("Part 1:", part1())
	fmt.Println("Part 2:", part2())
}

func part2() int {
	ids := strings.Split(strings.TrimSpace(input), ",")

	var sum int
	for _, id := range ids {
		// check for hyphen, e.g. number-number
		if strings.Count(id, "-") != 1 {
			fmt.Println("Invalid ID format:", id)
		}
		// seprate by hyphen
		first, _ := strconv.Atoi(strings.SplitN(id, "-", 2)[0])
		last, _ := strconv.Atoi(strings.SplitN(id, "-", 2)[1])

		// validate the id, you can find the invalid IDs by looking
		// for any ID which is made only of some sequence of digits repeated twice.
		// So, 55 (5 twice), 6464 (64 twice), and 123123 (123 twice) would all be invalid IDs.
		sum += isStillInvalidID(first, last)
		// add the validated ids to the sum
	}

	return sum
}

func part1() int {
	ids := strings.Split(strings.TrimSpace(input), ",")

	var sum int
	for _, id := range ids {
		// check for hyphen, e.g. number-number
		if strings.Count(id, "-") != 1 {
			fmt.Println("Invalid ID format:", id)
		}
		// seprate by hyphen
		first, _ := strconv.Atoi(strings.SplitN(id, "-", 2)[0])
		last, _ := strconv.Atoi(strings.SplitN(id, "-", 2)[1])

		// validate the id, you can find the invalid IDs by looking
		// for any ID which is made only of some sequence of digits repeated twice.
		// So, 55 (5 twice), 6464 (64 twice), and 123123 (123 twice) would all be invalid IDs.
		sum += isInvalidID(first, last)
		// add the validated ids to the sum
	}
	return sum
}

func isStillInvalidID(first, last int) int {
	sum := 0
	a := makeRange(first, last)
	for _, n := range a {
		// also removes leading zeros
		if check(n) {
			sum += n
		}
	}
	return sum
}

func check(n int) bool {
	s := strconv.Itoa(n)
	L := len(s)

	// block size from 1 up to L/2 (if block > L/2, you can't have at least 2 repeats)
	for blockSize := 1; blockSize <= L/2; blockSize++ {
		if L%blockSize != 0 {
			continue // can't tile string with this block size
		}

		block := s[:blockSize]
		valid := true

		// start from the second block
		for i := blockSize; i < L; i += blockSize {
			if s[i:i+blockSize] != block {
				valid = false
				break
			}
		}

		if valid {
			// we know there are at least 2 blocks because blockSize <= L/2
			return true
		}
	}

	return false
}

func isInvalidID(first int, last int) int {
	sum := 0
	a := makeRange(first, last)
	for _, n := range a {
		// also removes leading zeros
		strN := strconv.Itoa(n)
		// if it's odd length, skip
		if len(strN)%2 != 0 {
			continue
		}

		half := len(strN) / 2
		if strN[:half] == strN[half:] {
			invalid, _ := strconv.Atoi(strN)
			sum += invalid
		}
	}
	return sum
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}
