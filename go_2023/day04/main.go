package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func stringsToIntegers(lines []string) []int {
	var integers []int
	for _, line := range lines {
		if line != "" {
			//fmt.Println(i, line)
			n, err := strconv.Atoi(line)
			if err != nil {
				// handle error
				fmt.Println(err)
				os.Exit(2)
				continue
			}
			integers = append(integers, n)
		}
	}

	return integers
}

func part1() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	result := 0

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		x := fileScanner.Text()
		rest := strings.Split(x, ":")[1]
		win := stringsToIntegers(strings.Split((strings.Split(rest, "|")[0]), " "))
		own := stringsToIntegers(strings.Split((strings.Split(rest, "|")[1]), " "))
		res := 0
		for _, x := range own {
			if slices.Contains(win, x) {
				if res == 0 {
					res++
				} else {
					res = res * 2
				}
			}
		}
		result = result + res
	}

	f.Close()
	fmt.Println(result)
}

func part2() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	cards := make([]int, 300)
	result := 0
	re := regexp.MustCompile("[0-9]+")
	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		x := fileScanner.Text()
		game, _ := strconv.Atoi(re.FindAllString(strings.Split(x, ":")[0], -1)[0])
		cards[game]++
		rest := strings.Split(x, ":")[1]
		win := stringsToIntegers(strings.Split((strings.Split(rest, "|")[0]), " "))
		own := stringsToIntegers(strings.Split((strings.Split(rest, "|")[1]), " "))
		res := 0
		for _, x := range own {
			if slices.Contains(win, x) {
				res++
			}
		}
		for x := 1; x <= cards[game]; x++ {
			for idx := game + 1; idx <= game+res; idx++ {
				cards[idx]++

			}
		}
	}
	for _, card := range cards {

		result = result + card
	}

	f.Close()
	fmt.Println(result)
}

func main() {
	fmt.Println("Part 1")
	part1()
	fmt.Println("Part 2")
	part2()
}
