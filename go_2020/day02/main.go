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
	counter := 0

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		x := strings.Fields(fileScanner.Text())
		subarr := strings.Split(x[0], "-")
		min, _ := strconv.Atoi(subarr[0])
		max, _ := strconv.Atoi(subarr[len(subarr)-1])
		subarr = strings.Split(x[1], ":")
		lett := subarr[0]
		occur := strings.Count(x[2], lett)
		if occur >= min && occur <= max {
			counter++
		}
	}
	f.Close()
	fmt.Println(counter)
}

func part2() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	counter := 0

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		x := strings.Fields(fileScanner.Text())
		subarr := strings.Split(x[0], "-")
		min, _ := strconv.Atoi(subarr[0])
		max, _ := strconv.Atoi(subarr[len(subarr)-1])
		subarr = strings.Split(x[1], ":")
		lett := subarr[0]
		// Go has no XOR therefor !=
		if (string(x[2][min-1]) == lett) != (string(x[2][max-1]) == lett) {
			counter++
		}
	}
	f.Close()
	fmt.Println(counter)
}

func main() {
	part1()
	part2()
}
