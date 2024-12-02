package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func stringsToIntegers(lines []string) ([]int, error) {
	integers := make([]int, 0, len(lines))
	for _, line := range lines {
		n, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		integers = append(integers, n)
	}
	return integers, nil
}

func part1() {
	f, err := os.Open("sample.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)
	count := 0
	for fileScanner.Scan() {
		x := fileScanner.Text()
		x1 := strings.Fields(x)
		report, _ := stringsToIntegers(x1)
		ok := false
		for i := 0; i < len(report)-1; i++ {
			diff := report[i+1] - report[i]
			if diff < 0 {
				inc := false
				diff = -diff
			} else {
				inc := true
			}
			if diff >= 1 && diff <= 3 {
				ok = true
			} else {
				ok = false
				break
			}
		}
		if ok {
			count += 1
		}
	}
	f.Close()

	fmt.Println("Part 1:", count)
}

func main() {
	part1()
}
