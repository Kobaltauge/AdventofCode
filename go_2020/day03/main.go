package main

import (
	"bufio"
	"fmt"
	"os"
)

func part1() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	var treemap []string
	fileScanner := bufio.NewScanner(f)

	for fileScanner.Scan() {
		treemap = append(treemap, fileScanner.Text())
	}
	f.Close()

	x := 0
	y := 0
	counter := 0

	for {
		if string(treemap[y][x]) == "#" {
			counter += 1
		}
		x += 3
		if x > (len(treemap[0]) - 1) {
			x = x - (len(treemap[0]))
		}
		y += 1
		if y == len(treemap) {
			break
		}
	}
	fmt.Println(counter)
}

func part2() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	var treemap []string
	fileScanner := bufio.NewScanner(f)

	for fileScanner.Scan() {
		treemap = append(treemap, fileScanner.Text())
	}
	f.Close()

	sol := 1
	dir := [10]int{1, 1, 3, 1, 5, 1, 7, 1, 1, 2}

	for i := 0; i < 10; i += 2 {
		x := 0
		y := 0
		counter := 0
		for {
			if string(treemap[y][x]) == "#" {
				counter += 1
			}
			x += dir[i]
			if x > (len(treemap[0]) - 1) {
				x = x - (len(treemap[0]))
			}
			y += dir[i+1]
			if y >= len(treemap) {
				break
			}
		}
		// fmt.Println(y, counter)
		sol = sol * counter
	}
	fmt.Println(sol)
}

func main() {
	part1()
	part2()
}
