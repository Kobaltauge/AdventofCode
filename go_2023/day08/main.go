package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type knot struct {
	left, right string
}

func part1() {
	knods := make(map[string]knot)
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	steps := 0

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)
	fileScanner.Scan()
	way := []rune(fileScanner.Text())
	fileScanner.Scan()
	for fileScanner.Scan() {
		x := fileScanner.Text()

		line := strings.FieldsFunc(x, func(r rune) bool {
			if r == '=' || r == '(' || r == ')' || r == ',' || r == ' ' {
				return true
			}
			return false
		})
		knods[line[0]] = knot{left: line[1], right: line[2]}

	}
	f.Close()
	n := "AAA"
	for {
		for _, w := range way {
			if w == 76 {
				n = knods[n].left
				steps++
			} else {
				n = knods[n].right
				steps++
			}
			if n == "ZZZ" {
				break
			}
		}
		if n == "ZZZ" {
			break
		}
	}
	fmt.Println(steps)
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func part2() {
	knods := make(map[string]knot)
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	starts := make([]string, 0)

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)
	fileScanner.Scan()
	way := []rune(fileScanner.Text())
	fileScanner.Scan()
	for fileScanner.Scan() {
		x := fileScanner.Text()

		line := strings.FieldsFunc(x, func(r rune) bool {
			if r == '=' || r == '(' || r == ')' || r == ',' || r == ' ' {
				return true
			}
			return false
		})
		knods[line[0]] = knot{left: line[1], right: line[2]}
		if strings.HasSuffix(line[0], "A") {
			starts = append(starts, line[0])
		}

	}
	f.Close()
	res := make([]int, 0)
	for _, n := range starts {
		steps := 0
		for {
			for _, w := range way {
				if w == 76 {
					n = knods[n].left
					steps++
				} else {
					n = knods[n].right
					steps++
				}
				if strings.HasSuffix(n, "Z") {
					break
				}
			}
			if strings.HasSuffix(n, "Z") {
				res = append(res, steps)
				break
			}
		}
	}
	result := LCM(res[0], res[1], res...)
	fmt.Println(result)
}

func main() {
	fmt.Println("Part 1")
	part1()
	fmt.Println("Part 2")
	part2()
}
