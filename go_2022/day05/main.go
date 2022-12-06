package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// {G]                 [D] [R]
// [W]         [V]     [C] [T] [M]
// [L]         [P] [Z] [Q] [F] [V]
// [J]         [S] [D] [J] [M] [T] [V]
// [B]     [M] [H] [L] [Z] [J] [B] [S]
// [R] [C] [T] [C] [T] [R] [D] [R] [D]
// [T] [W] [Z] [T] [P] [B] [B] [H] [P]
// [D] [S] [R] [D] [G] [F] [S] [L] [Q]
//  1   2   3   4   5   6   7   8   9

func part1() {
	var m = map[int][]string{0: {"x"}, 1: {"D", "T", "R", "B", "J", "L", "W", "G"}, 2: {"S", "W", "C"}, 3: {"R", "Z", "T", "M"}, 4: {"D", "T", "C", "H", "S", "P", "V"},
		5: {"G", "P", "T", "L", "D", "Z"}, 6: {"F", "B", "R", "Z", "J", "Q", "C", "D"}, 7: {"S", "B", "D", "J", "M", "F", "T", "R"}, 8: {"L", "H", "R", "B", "T", "V", "M"},
		9: {"Q", "P", "D", "S", "V"}}

	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		x := strings.Fields(fileScanner.Text())
		quant, _ := strconv.Atoi(x[1])
		from, _ := strconv.Atoi(x[3])
		to, _ := strconv.Atoi(x[5])
		for idx := 1; idx <= quant; idx++ {
			el := m[from][len(m[from])-1]
			m[to] = append(m[to], el)
			m[from] = m[from][:len(m[from])-1]
		}
	}
	fmt.Println("Part 1:")
	for idx := 1; idx < len(m); idx++ {
		fmt.Print(m[idx][len(m[idx])-1])
	}
	fmt.Println("")
}

func part2() {
	var m = map[int][]string{0: {"x"}, 1: {"D", "T", "R", "B", "J", "L", "W", "G"}, 2: {"S", "W", "C"}, 3: {"R", "Z", "T", "M"}, 4: {"D", "T", "C", "H", "S", "P", "V"},
		5: {"G", "P", "T", "L", "D", "Z"}, 6: {"F", "B", "R", "Z", "J", "Q", "C", "D"}, 7: {"S", "B", "D", "J", "M", "F", "T", "R"}, 8: {"L", "H", "R", "B", "T", "V", "M"},
		9: {"Q", "P", "D", "S", "V"}}

	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		x := strings.Fields(fileScanner.Text())
		quant, _ := strconv.Atoi(x[1])
		from, _ := strconv.Atoi(x[3])
		to, _ := strconv.Atoi(x[5])
		var el []string
		if len(m[from]) <= quant {
			el = m[from]
		} else {
			el = m[from][len(m[from])-quant:]
		}
		m[to] = append(m[to], el...)
		m[from] = m[from][:len(m[from])-quant]

	}
	fmt.Println("Part 2:")

	for idx := 1; idx < len(m); idx++ {
		if len(m[idx]) > 0 {
			fmt.Print(m[idx][len(m[idx])-1])
		}
	}

}

func main() {
	part1()
	part2()
}
