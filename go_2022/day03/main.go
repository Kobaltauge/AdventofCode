package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func SliceIndex(limit int, predicate func(i int) bool) int {
	for i := 0; i < limit; i++ {
		if predicate(i) {
			return i
		}
	}
	return -1
}

func part1() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	var alph = [...]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	sum := 0

	for fileScanner.Scan() {
		x := fileScanner.Text()
		left, right := x[0:len(x)/2], x[len(x)/2:len(x)]
		// partsum := 0
		for idx := 0; idx <= len(x)/2; idx++ {
			if strings.Contains(left, string(right[idx])) {
				j := SliceIndex(len(alph), func(i int) bool { return alph[i] == string(right[idx]) }) + 1
				sum = sum + j
				break
			}

		}
	}

	fmt.Println("Part 1:", sum)
	f.Close()
}

func part2() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	var alph = [...]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	sum := 0

	for fileScanner.Scan() {
		a := fileScanner.Text()
		fileScanner.Scan()
		b := fileScanner.Text()
		fileScanner.Scan()
		c := fileScanner.Text()
		x := a

		if len(a) > len(b) {
			if len(a) > len(c) {
				x = a
			} else {
				x = c
			}
		} else {
			if len(b) > len(c) {
				x = b
			} else {
				x = c
			}
		}
		for idx := 0; idx <= len(x); idx++ {
			if strings.Contains(a, string(x[idx])) && strings.Contains(b, string(x[idx])) && strings.Contains(c, string(x[idx])) {
				j := SliceIndex(len(alph), func(i int) bool { return alph[i] == string(x[idx]) }) + 1
				sum = sum + j
				break
			}

		}
	}

	fmt.Println("Part 2:", sum)
	f.Close()
}

func main() {
	part1()
	part2()
}
