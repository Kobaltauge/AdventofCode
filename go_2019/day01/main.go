package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func part1() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)
	var numbers []int

	for fileScanner.Scan() {
		x, err := strconv.Atoi(fileScanner.Text())
		if err != nil {
			fmt.Println(err)
			return
		}
		numbers = append(numbers, x)
	}
	f.Close()
	// fmt.Println(numbers)
	sum := 0
	for _, no := range numbers {
		res := no / 3
		// fmt.Println(res, int(res))
		res = res - 2
		sum += res
	}
	fmt.Println(sum)
}

func part2() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)
	var numbers []int

	for fileScanner.Scan() {
		x, err := strconv.Atoi(fileScanner.Text())
		if err != nil {
			fmt.Println(err)
			return
		}
		numbers = append(numbers, x)
	}
	f.Close()
	// fmt.Println(numbers)
	sum := 0
	for _, no := range numbers {
		midsum := 0
	Exit:
		for no > 0 {
			res := no / 3
			res = res - 2
			if res <= 0 {
				no = res
				continue Exit
			}
			midsum += res
			no = res
		}
		sum += midsum
	}
	fmt.Println(sum)
}

func main() {
	part1()
	part2()
}
