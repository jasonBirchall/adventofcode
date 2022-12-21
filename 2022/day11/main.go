package main

import (
	_ "embed"
	"sort"
	"strconv"
	"strings"
)

//go:embed day11.txt
var input string

type monkey struct {
	id        int
	items     items
	operation string
	test      test
}

type items struct {
	startingItems []int
	total         int
}

type test struct {
	divisableBy   int
	trueReceiver  int
	falseReceiver int
}

func main() {
	println("Part 1:", part1(input))
	println("Part 2:", part2(input))
}

func part1(input string) int {
	observation := strings.Split(strings.TrimSuffix(input, "\n"), "\n\n")

	monkeySet := []monkey{}
	for _, monkey := range observation {
		monkeySet = append(monkeySet, newMonkey(monkey))
	}

	for i := 1; i <= 20; i++ {
		for j := range monkeySet {
			inspectedMonkey := &monkeySet[j]
			for _, item := range inspectedMonkey.items.startingItems {
				new := inspectedMonkey.operator(item)
				// After each monkey inspects an item but before it tests your worry level, your relief that the monkey's inspection didn't damage the item causes your worry level to be divided by three and rounded down to the nearest integer.
				new = new / 3
				if new%inspectedMonkey.test.divisableBy == 0 {
					monkeySet[inspectedMonkey.test.trueReceiver].items.startingItems = append(monkeySet[inspectedMonkey.test.trueReceiver].items.startingItems, new)
				} else {
					monkeySet[inspectedMonkey.test.falseReceiver].items.startingItems = append(monkeySet[inspectedMonkey.test.falseReceiver].items.startingItems, new)
				}
				_, inspectedMonkey.items.startingItems = inspectedMonkey.items.startingItems[len(inspectedMonkey.items.startingItems)-1], inspectedMonkey.items.startingItems[:len(inspectedMonkey.items.startingItems)-1]
				inspectedMonkey.items.total++
			}
		}
	}
	var totalPasses []int
	for _, monkey := range monkeySet {
		totalPasses = append(totalPasses, monkey.items.total)
	}

	sort.IntSlice(totalPasses).Sort()

	return totalPasses[len(totalPasses)-1] * totalPasses[len(totalPasses)-2]
}

func part2(input string) int {
	observation := strings.Split(strings.TrimSuffix(input, "\n"), "\n\n")

	monkeySet := []monkey{}
	bigMod := 1
	for _, monkey := range observation {
		m := newMonkey(monkey)
		bigMod *= m.test.divisableBy
		monkeySet = append(monkeySet, m)
	}

	for i := 1; i <= 10000; i++ {
		for j := range monkeySet {
			inspectedMonkey := &monkeySet[j]
			for _, item := range inspectedMonkey.items.startingItems {
				new := inspectedMonkey.operator(item)
				new %= bigMod
				if new%inspectedMonkey.test.divisableBy == 0 {
					monkeySet[inspectedMonkey.test.trueReceiver].items.startingItems = append(monkeySet[inspectedMonkey.test.trueReceiver].items.startingItems, new)
				} else {
					monkeySet[inspectedMonkey.test.falseReceiver].items.startingItems = append(monkeySet[inspectedMonkey.test.falseReceiver].items.startingItems, new)
				}
				_, inspectedMonkey.items.startingItems = inspectedMonkey.items.startingItems[len(inspectedMonkey.items.startingItems)-1], inspectedMonkey.items.startingItems[:len(inspectedMonkey.items.startingItems)-1]
				inspectedMonkey.items.total++
			}
		}
	}
	var totalPasses []int
	for _, monkey := range monkeySet {
		totalPasses = append(totalPasses, monkey.items.total)
	}

	sort.IntSlice(totalPasses).Sort()

	return totalPasses[len(totalPasses)-1] * totalPasses[len(totalPasses)-2]
}

func (m monkey) operator(item int) int {
	operator := strings.Split(m.operation, " ")
	if operator[1] == "old" {
		operator[1] = strconv.Itoa(item)
	}

	switch operator[0] {
	case "+":
		o, _ := strconv.Atoi(operator[1])
		return item + o
	case "*":
		o, _ := strconv.Atoi(operator[1])
		return item * o
	}
	return 0
}

func newMonkey(monkeyNotes string) monkey {
	m := monkey{}

	name := strings.TrimRight(strings.Split(monkeyNotes, ":")[0], " ")
	m.id, _ = strconv.Atoi(strings.Split(name, " ")[1])

	for _, line := range strings.Split(monkeyNotes, "\n")[1:] {
		if strings.Contains(line, "items") {
			collection := strings.SplitAfter(line, "items:")[1]
			for _, i := range strings.Split(collection, ",") {
				conv, _ := strconv.Atoi(strings.TrimSpace(i))
				m.items.startingItems = append(m.items.startingItems, conv)
			}
		}

		if strings.Contains(line, "Operation") {
			m.operation = strings.SplitAfter(line, "old ")[1]
		}

		if strings.Contains(line, "Test") {
			m.test.divisableBy, _ = strconv.Atoi(strings.SplitAfter(line, "by ")[1])
		}

		if strings.Contains(line, "If true:") {
			m.test.trueReceiver, _ = strconv.Atoi(strings.SplitAfter(line, "to monkey ")[1])
		}
		if strings.Contains(line, "If false:") {
			m.test.falseReceiver, _ = strconv.Atoi(strings.SplitAfter(line, "to monkey ")[1])
		}
	}

	return m
}
