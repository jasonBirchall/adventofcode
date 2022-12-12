package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed day8.txt
var input string

func main() {
	fmt.Println("answer to part 1 is:", part1(strings.Split(strings.TrimSuffix(input, "\n"), "\n")))
	fmt.Println("answer to part 2 is:", part2(strings.Split(strings.TrimSuffix(input, "\n"), "\n")))

}
func part1(trees []string) int {
	visibleTrees := map[string]struct{}{}

	for row := 1; row < len(trees)-1; row++ {

		// find tallest from left to right
		tallestTree, _ := strconv.Atoi(string(trees[row][0]))
		for col := 1; col < len(trees[row])-1; col++ {
			currentTree, _ := strconv.Atoi(string(trees[row][col]))
			if currentTree > tallestTree {
				tallestTree = currentTree
				visibleTrees[strconv.Itoa(row)+" "+strconv.Itoa(col)] = struct{}{}
			}
		}

		// find tallest from right to left
		tallestTree, _ = strconv.Atoi(string(trees[row][len(trees[row])-1]))
		for col := len(trees[row]) - 2; col > 0; col-- {
			currentTree, _ := strconv.Atoi(string(trees[row][col]))
			if currentTree > tallestTree {
				tallestTree = currentTree
				visibleTrees[strconv.Itoa(row)+" "+strconv.Itoa(col)] = struct{}{}
			}
		}
	}

	for col := 1; col < len(trees)-1; col++ {
		// find tallest from top to bottom
		tallestTree, _ := strconv.Atoi(string(trees[0][col]))
		for row := 1; row < len(trees)-1; row++ {
			currentTree, _ := strconv.Atoi(string(trees[row][col]))
			if currentTree > tallestTree {
				tallestTree = currentTree
				visibleTrees[strconv.Itoa(row)+" "+strconv.Itoa(col)] = struct{}{}
			}
		}

		// find tallest from bottom to top
		tallestTree, _ = strconv.Atoi(string(trees[len(trees)-1][col]))
		for row := len(trees) - 2; row > 0; row-- {
			currentTree, _ := strconv.Atoi(string(trees[row][col]))
			if currentTree > tallestTree {
				tallestTree = currentTree
				visibleTrees[strconv.Itoa(row)+" "+strconv.Itoa(col)] = struct{}{}
			}
		}
	}

	return (len(trees) * 2) + (len(trees[0]) * 2) + len(visibleTrees) - 4
}

func part2(trees []string) int {
	bestScore := 0
	for row := 0; row < len(trees)-1; row++ {
		for col := 0; col < len(trees[row])-1; col++ {
			currentScore := score(trees, row, col)
			if currentScore > bestScore {
				bestScore = currentScore
			}
		}
	}

	return bestScore
}

func score(trees []string, row int, col int) int {
	treeHeight, _ := strconv.Atoi(string(trees[row][col]))

	// check top
	topScore := 0
	for r := row - 1; ; r-- {
		if r < 0 {
			break
		}

		currentTree, _ := strconv.Atoi(string(trees[r][col]))
		if currentTree <= treeHeight {
			topScore++
		}

		if currentTree >= treeHeight {
			break
		}
	}

	// check bottom
	bottomScore := 0
	for r := row + 1; ; r++ {
		if r == len(trees) {
			break
		}

		currentTree, _ := strconv.Atoi(string(trees[r][col]))
		if currentTree <= treeHeight {
			bottomScore++
		}

		if currentTree >= treeHeight {
			break
		}
	}

	// check right
	rightScore := 0
	for c := col + 1; ; c++ {
		if c == len(trees[row]) {
			break
		}

		currentTree, _ := strconv.Atoi(string(trees[row][c]))
		if currentTree <= treeHeight {
			rightScore++
		}

		if currentTree >= treeHeight {
			break
		}
	}

	// check left
	leftScore := 0
	for c := col - 1; ; c-- {
		if c < 0 {
			break
		}

		currentTree, _ := strconv.Atoi(string(trees[row][c]))
		if currentTree <= treeHeight {
			leftScore++
		}

		if currentTree >= treeHeight {
			break
		}
	}

	return topScore * bottomScore * rightScore * leftScore
}
