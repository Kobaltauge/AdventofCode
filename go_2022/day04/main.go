package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	occur := 0

	for fileScanner.Scan() {
		x := strings.Fields(fileScanner.Text())
		subarr := strings.Split(x[0], ",")
		min1, _ := strconv.Atoi(strings.Split(subarr[0], "-")[0])
		max1, _ := strconv.Atoi(strings.Split(subarr[0], "-")[1])
		min2, _ := strconv.Atoi(strings.Split(subarr[1], "-")[0])
		max2, _ := strconv.Atoi(strings.Split(subarr[1], "-")[1])

		if min1 >= min2 && max1 <= max2 || min1 <= min2 && max1 >= max2 {
			occur++
		}
	}
	f.Close()
	fmt.Println("Part 1:", occur)
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func contains(ints []int, num int) bool {
	for _, v := range ints {
		if v == num {
			return true
		}
	}
	return false
}

func part2() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	occur := 0

	for fileScanner.Scan() {
		x := strings.Fields(fileScanner.Text())
		subarr := strings.Split(x[0], ",")
		min1, _ := strconv.Atoi(strings.Split(subarr[0], "-")[0])
		max1, _ := strconv.Atoi(strings.Split(subarr[0], "-")[1])
		min2, _ := strconv.Atoi(strings.Split(subarr[1], "-")[0])
		max2, _ := strconv.Atoi(strings.Split(subarr[1], "-")[1])

		a := makeRange(min1, max1)
		b := makeRange(min2, max2)

		for _, x := range a {
			if contains(b, x) {
				occur++
				break
			}
		}

	}
	f.Close()
	fmt.Println("Part 2:", occur)
}

func main() {
	part1()
	part2()
}
