package main

import (
	_ "embed"
	"fmt"
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
	fmt.Println(monkeySet)

	for i := 1; i <= 20; i++ {
		fmt.Println("Round", i)
		fmt.Println("----")
		for j := range monkeySet {
			inspectedMonkey := &monkeySet[j]
			fmt.Println("Monkey", inspectedMonkey.id, ":", inspectedMonkey.items.startingItems)
			for _, item := range inspectedMonkey.items.startingItems {
				fmt.Println("	Monkey inspects an item with a worry level of", item)
				new := inspectedMonkey.operator(item)
				fmt.Println("		Worry level increases to", new)
				// After each monkey inspects an item but before it tests your worry level, your relief that the monkey's inspection didn't damage the item causes your worry level to be divided by three and rounded down to the nearest integer.
				new = new / 3
				fmt.Println("		Monkey gets bored with item. Worry level is now:", new)
				if new%inspectedMonkey.test.divisableBy == 0 {
					fmt.Println("		Current worry level is divisible by", inspectedMonkey.test.divisableBy, "so monkey", inspectedMonkey.id, "sends item to monkey", inspectedMonkey.test.trueReceiver)
					monkeySet[inspectedMonkey.test.trueReceiver].items.startingItems = append(monkeySet[inspectedMonkey.test.trueReceiver].items.startingItems, new)
					fmt.Println("		monkey:", monkeySet[inspectedMonkey.test.trueReceiver].id, "now has items:", monkeySet[inspectedMonkey.test.trueReceiver].items.startingItems)
				} else {
					fmt.Println("		Current worry level is not divisible by", inspectedMonkey.test.divisableBy, "so monkey", inspectedMonkey.id, "sends item to monkey", inspectedMonkey.test.falseReceiver)
					monkeySet[inspectedMonkey.test.falseReceiver].items.startingItems = append(monkeySet[inspectedMonkey.test.falseReceiver].items.startingItems, new)
					fmt.Println("		monkey:", monkeySet[inspectedMonkey.test.falseReceiver].id, "now has items:", monkeySet[inspectedMonkey.test.falseReceiver].items.startingItems)
				}
				_, inspectedMonkey.items.startingItems = inspectedMonkey.items.startingItems[len(inspectedMonkey.items.startingItems)-1], inspectedMonkey.items.startingItems[:len(inspectedMonkey.items.startingItems)-1]
				fmt.Println("	monkey:", inspectedMonkey.id, "now has items:", inspectedMonkey.items.startingItems)
				inspectedMonkey.items.total++
			}
		}
	}
	fmt.Println("After round 1, the monkeys are holding items with worry levels of")
	var totalPasses []int
	for _, monkey := range monkeySet {
		fmt.Println("	monkey", monkey.id, ":", monkey.items.startingItems)
		totalPasses = append(totalPasses, monkey.items.total)
	}
	for _, monkey := range monkeySet {
		fmt.Println("	monkey", monkey.id, "inspected items", monkey.items.total, "times")
	}

	sort.IntSlice(totalPasses).Sort()

	return totalPasses[len(totalPasses)-1] * totalPasses[len(totalPasses)-2]
}

func part2(input string) int {
	observation := strings.Split(strings.TrimSuffix(input, "\n"), "\n\n")

	monkeySet := []monkey{}
	for _, monkey := range observation {
		monkeySet = append(monkeySet, newMonkey(monkey))
	}

	for i := 1; i <= 10000; i++ {
		for j := range monkeySet {
			inspectedMonkey := &monkeySet[j]
			for _, item := range inspectedMonkey.items.startingItems {
				new := inspectedMonkey.operator(item)
				// After each monkey inspects an item but before it tests your worry level, your relief that the monkey's inspection didn't damage the item causes your worry level to be divided by three and rounded down to the nearest integer.
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
		fmt.Println("	monkey", monkey.id, ":", monkey.items.startingItems)
		totalPasses = append(totalPasses, monkey.items.total)
	}
	for _, monkey := range monkeySet {
		fmt.Println("	monkey", monkey.id, "inspected items", monkey.items.total, "times")
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
