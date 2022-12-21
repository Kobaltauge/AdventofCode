package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1() {
	monkeys := map[string]int{}
	stack := map[string][]string{}

	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		raw := strings.Fields(fileScanner.Text())
		mn := strings.Trim(strings.Split(raw[0], ":")[0], "")
		if len(raw) == 2 {
			my, _ := strconv.Atoi(strings.Split(raw[1], ":")[0])
			monkeys[mn] = my
		} else {
			stack[mn] = []string{raw[1], raw[2], raw[3]}
		}
	}

	for len(stack) > 0 {
		for mn, my := range stack {
			m1 := my[0]
			m2 := my[2]

			_, ok := monkeys[m1]
			// If the key exists
			if ok {
				_, ok := monkeys[m2]
				if ok {
					switch my[1] {
					case "+":
						monkeys[mn] = monkeys[m1] + monkeys[m2]
					case "-":
						monkeys[mn] = monkeys[m1] - monkeys[m2]
					case "*":
						monkeys[mn] = monkeys[m1] * monkeys[m2]
					case "/":
						monkeys[mn] = monkeys[m1] / monkeys[m2]
					}
					delete(stack, mn)
				}
			} else {
				continue
			}
		}
	}
	fmt.Println("Part 1:", monkeys["root"])
}

func diff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

// https://stackoverflow.com/a/30765350

func part2() {
	monkeys := map[string]int{}
	stack := map[string][]string{}

	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		raw := strings.Fields(fileScanner.Text())
		mn := strings.Trim(strings.Split(raw[0], ":")[0], "")
		if len(raw) == 2 {
			my, _ := strconv.Atoi(strings.Split(raw[1], ":")[0])
			monkeys[mn] = my
		} else {
			stack[mn] = []string{raw[1], raw[2], raw[3]}
		}
	}

	stack["root"][1] = "="
	stackstore := make(map[string][]string)
	for k, v := range stack {
		stackstore[k] = v
	}
	monkeysstore := make(map[string]int)
	for k, v := range monkeys {
		monkeysstore[k] = v
	}
	result := 0
out:
	for i := 0; true; i = i + 100000 {

		stack := make(map[string][]string)
		for k, v := range stackstore {
			stack[k] = v
		}
		monkeys := make(map[string]int)
		for k, v := range monkeysstore {
			monkeys[k] = v
		}

		monkeys["humn"] = i
		// fmt.Println(monkeys["humn"])
		for len(stack) > 0 {
			for mn, my := range stack {
				m1 := my[0]
				m2 := my[2]

				_, ok := monkeys[m1]
				// If the key exists
				if ok {
					_, ok := monkeys[m2]
					if ok {
						switch my[1] {
						case "+":
							monkeys[mn] = monkeys[m1] + monkeys[m2]
						case "-":
							monkeys[mn] = monkeys[m1] - monkeys[m2]
						case "*":
							monkeys[mn] = monkeys[m1] * monkeys[m2]
						case "/":
							monkeys[mn] = monkeys[m1] / monkeys[m2]
						case "=":
							fmt.Println(monkeys[m1], monkeys[m2])
							if monkeys[m1] == monkeys[m2] {
								result = i
								break out
							}
						}
						delete(stack, mn)
					}
				} else {
					continue
				}
			}
		}
	}
	fmt.Println("Part 2:", result)
}

func main() {
	part1()
	part2()
}
