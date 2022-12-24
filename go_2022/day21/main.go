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

// https://stackoverflow.com/a/30765350

func part2() {
	imin := 0
	imax := int(1e16)
	i := int(imax / 2)
	for {
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
		x1 := stack["root"][0]
		x2 := stack["root"][2]
		monkeys["humn"] = i

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
							if monkeys[m1] == monkeys[m2] {
								fmt.Println(monkeys[x1], monkeys[x2], i, imin, imax)
								// -1 because of the int() round error
								fmt.Println("Part 2:", monkeys["humn"]-1)
								os.Exit(0)
							}
						}
						delete(stack, mn)
					}
				} else {
					continue
				}
			}
		}
		if monkeys[x1] < monkeys[x2] {
			imax = i
			i = int((imax-imin)/2 + imin)
			fmt.Println(monkeys[x1], monkeys[x2], i, imin, imax)
		} else {
			imin = i
			i = int((imax-imin)/2 + imin)
			fmt.Println(monkeys[x1], monkeys[x2], i, imin, imax)
		}
	}
}

func main() {
	part1()
	part2()
}
