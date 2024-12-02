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
	f, err := os.Open("input.txt")
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
		dir := "null"
		for i := 0; i < len(report)-1; i++ {
			diff := report[i+1] - report[i]
			if diff < 0 {
				if dir == "null" || dir == "down" {
					dir = "down"
				} else {
					ok = false
					break
				}
				diff = -diff
			} else {
				if dir == "null" || dir == "up" {
					dir = "up"
				} else {
					ok = false
					break
				}
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

func part2() {
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
		dir := "null"
		damper := 1
		for i := 0; i < len(report)-1; i++ {
			diff := report[i+1] - report[i]
			if diff < 0 {
				if dir == "null" || dir == "down" {
					dir = "down"
				} else {
					if damper == 0 {
						ok = false
						break
					} else {
						damper -= 1
					}
				}
				diff = -diff
			} else {
				if dir == "null" || dir == "up" {
					dir = "up"
				} else {
					if damper == 0 {
						ok = false
						break
					} else {
						damper -= 1
					}
				}
			}
			if diff >= 1 && diff <= 3 {
				ok = true
			} else {
				if damper == 0 {
					ok = false
					break
				} else {
					damper -= 1
				}
			}
		}
		if ok && damper == 0 {
			count += 1
		}
	}
	f.Close()

	fmt.Println("Part 2:", count)
	// 658 to low
}

func main() {
	part1()
	part2()
}
