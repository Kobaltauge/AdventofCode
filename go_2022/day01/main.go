package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)
	calories := 0
	var numbers []int

	for fileScanner.Scan() {
		x := fileScanner.Text()
		if x == "" {
			numbers = append(numbers, calories)
			calories = 0
		} else {
			x, err := strconv.Atoi(fileScanner.Text())
			if err != nil {
				fmt.Println(err)
				return
			}
			calories = calories + x
		}
	}
	f.Close()
	println(numbers)
	sort.Sort(sort.Reverse(sort.IntSlice(numbers)))
	println(numbers)
	fmt.Println("Part 1:", numbers[0])
	fmt.Println("Part 2:", numbers[0]+numbers[1]+numbers[2])
}
