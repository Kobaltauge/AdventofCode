package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(i, x int) int {
	res := 0
	var c []int
	c = append(c, 20)
	c = append(c, 60)
	c = append(c, 100)
	c = append(c, 140)
	c = append(c, 180)
	c = append(c, 220)
	if contains(c, i) {
		res = i * x
	}
	return res
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func part1() {

	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	result := 0
	value := 1
	for i := 1; i < 225; {
		fileScanner.Scan()
		raw := strings.Fields(fileScanner.Text())
		fmt.Println(raw)
		cmd := raw[0]
		switch cmd {
		case "noop":
			result = result + check(i, value)
			i++
		case "addx":
			quant, _ := strconv.Atoi(raw[1])
			result = result + check(i, value)
			i++
			result = result + check(i, value)
			i++
			value = value + quant
		}
	}

	fmt.Println("Part 1:", result)
}

func main() {
	part1()
}
