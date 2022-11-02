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
	akt_jolt := 0
	idx := 0
	count := 0

	res := totheend(numbers, idx, akt_jolt, count)

	fmt.Println(res)
}

func totheend(numbers []int, idx int, akt_jolt int, count int) int {
	max := numbers[len(numbers)-1]
	if akt_jolt > max {
		return 0
	} else if akt_jolt == max {
		return 1
	}

	if (akt_jolt+numbers[idx]) <= max && (akt_jolt+numbers[idx]) <= akt_jolt+3 {
		count++
		idx++
		akt_jolt = akt_jolt + (numbers[idx] - akt_jolt)
		count += totheend(numbers, idx, akt_jolt, count)
		return count
	}
	if (akt_jolt+numbers[idx+1]) <= max && (akt_jolt+numbers[idx]) <= akt_jolt+3 {
		count++
		idx += 2
		akt_jolt = akt_jolt + (numbers[idx+1] - akt_jolt)
		count += totheend(numbers, idx, akt_jolt, count)
		return count
	}
	if (akt_jolt+numbers[idx+2]) <= max && (akt_jolt+numbers[idx]) <= akt_jolt+3 {
		count++
		idx += 3
		akt_jolt = akt_jolt + (numbers[idx+2] - akt_jolt)
		count += totheend(numbers, idx, akt_jolt, count)
		return count
	}

	return 0
}

func main() {
	part1()
	part2()
}
