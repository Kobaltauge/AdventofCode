package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func stringsToIntegers(lines []string) ([]int, error) {
	integers := make([]int, 0, len(lines))
	for _, line := range lines {
		n, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		integers = append(integers, n)
	}
	return integers, nil
}

func part1() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	var subarr1 []string
	var subarr2 []string
	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		x := fileScanner.Text()
		subarr1 = append(subarr1, strings.Fields(x)[0])
		subarr2 = append(subarr2, strings.Fields(x)[1])

	}
	f.Close()
	subarr3, _ := stringsToIntegers(subarr1)
	subarr4, _ := stringsToIntegers(subarr2)
	slices.Sort(subarr3)
	slices.Sort(subarr4)
	sum := 0
	for i, _ := range subarr3 {
		x := subarr3[i] - subarr4[i]
		fmt.Println(x)
		if x < 0 {
			x = -x
		}
		sum = x + sum
	}
	fmt.Println("Part 1:", sum)
}

func part2() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	var subarr1 []string
	var subarr2 []string
	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		x := fileScanner.Text()
		subarr1 = append(subarr1, strings.Fields(x)[0])
		subarr2 = append(subarr2, strings.Fields(x)[1])

	}
	f.Close()
	subarr3, _ := stringsToIntegers(subarr1)
	subarr4, _ := stringsToIntegers(subarr2)
	sum := 0
	for _, i := range subarr3 {
		count := 0
		for _, j := range subarr4 {
			if j == i {
				count = count + 1
			}
		}
		sum = i*count + sum
	}

	fmt.Println("Part 2:", sum)
}

func main() {
	part1()
	part2()
}
