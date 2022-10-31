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
	var numbers2 []int

	for fileScanner.Scan() {
		x, err := strconv.Atoi(fileScanner.Text())
		if err != nil {
			fmt.Println(err)
			return
		}
		numbers = append(numbers, x)
		numbers2 = append(numbers2, x)
	}
	f.Close()
	// fmt.Println(numbers)
	// fmt.Println(numbers2)
	for _, v1 := range numbers {
		for _, v2 := range numbers2 {
			if v1+v2 == 2020 {
				fmt.Println(v1 * v2)
			}
		}
	}
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
	var numbers2 []int
	var numbers3 []int

	for fileScanner.Scan() {
		x, err := strconv.Atoi(fileScanner.Text())
		if err != nil {
			fmt.Println(err)
			return
		}
		numbers = append(numbers, x)
		numbers2 = append(numbers2, x)
		numbers3 = append(numbers3, x)
	}
	f.Close()
	// fmt.Println(numbers)
	// fmt.Println(numbers2)
	for _, v1 := range numbers {
		for _, v2 := range numbers2 {
			for _, v3 := range numbers3 {
				if v1+v2+v3 == 2020 {
					fmt.Println(v1 * v2 * v3)
				}
			}
		}
	}
}

func main() {
	part1()
	part2()
}
