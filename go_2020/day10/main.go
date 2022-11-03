package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	sort.Ints(numbers)
	akt_jolt := 0
	count_1 := 0
	count_3 := 0

	for _, adap := range numbers {
		if adap-1 == akt_jolt {
			count_1++
		} else {
			count_3++
		}
		akt_jolt = akt_jolt + (adap - akt_jolt)
	}
	// adapter in laptop
	count_3++
	fmt.Println(count_1, count_3, count_1*count_3)
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

	numbers = append([]int{0}, numbers...)
	sort.Ints(numbers)

	res := totheend(numbers, 0, make(map[int]int))

	fmt.Println(res)
}

func totheend(numbers []int, idx int, cache map[int]int) int {
	// https://stackoverflow.com/a/2050629/9307482
	// checking the cache if value was calculated and return it (dynamic programming)
	if n, ok := cache[idx]; ok {
		return n
	}

	// end of list, end state recursion
	if idx == len(numbers)-1 {
		return 1
	}

	var sol int
	for j := idx + 1; j < len(numbers); j++ {
		// get the next adapters, if they are more then 3 jolts "away" stop
		if diff := numbers[j] - numbers[idx]; diff > 3 {
			break
		}
		//recursion
		sol += totheend(numbers, j, cache)
	}

	// put solution in cache for speed up (dynamic programming)
	cache[idx] = sol
	return sol

}

func examplecode2() {
	// https://github.com/davidporos92/aoc-2020/blob/master/day-10-adapter-array/main.go
	f, err := os.Open("sample.txt")
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

	sort.Ints(numbers)

	//0 for the start and +3 at the end for the laptop
	numbers = append([]int{0}, numbers...)
	numbers = append(numbers, numbers[len(numbers)-1]+3)
	a := make([]int, len(numbers))
	a[len(numbers)-1] = 1

	//start with the highest value
	for i := len(numbers) - 2; i >= 0; i-- {
		// the next adaptor fits everytime so copy the amount of solution in slot
		a[i] = a[i+1]

		//if the nextnextnext adaptor fit, add the solutions to the slot
		if i+3 < len(numbers) && numbers[i+3] <= numbers[i]+3 {
			a[i] += a[i+3]
		}

		//if the nextnext adaptor fit, add the solutions to the slot
		if i+2 < len(numbers) && numbers[i+2] <= numbers[i]+3 {
			a[i] += a[i+2]
		}
	}
	fmt.Println(a[0])
}

func main() {
	part1()
	part2()
}
