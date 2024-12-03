package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

	sum := 0
	for fileScanner.Scan() {
		x := fileScanner.Text()
		//x := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
		re := regexp.MustCompile(`mul\(\d*,\d*\)`)
		findings := re.FindAllString(x, -1)
		for _, find := range findings {
			re2 := regexp.MustCompile(`mul\((\d*),(\d*)\)`)
			all := re2.FindAllStringSubmatch(find, -1)
			a, _ := strconv.Atoi(all[0][1])
			b, _ := strconv.Atoi(all[0][2])
			sum = sum + (a * b)
		}
	}
	fmt.Println("Part 1:", sum)
	// 170807108
}

func part2() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	sum := 0
	compute := true
	for fileScanner.Scan() {
		x := fileScanner.Text()
		//x := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
		re := regexp.MustCompile(`do\(\)|mul\(\d*,\d*\)|don't\(\)`)
		findings := re.FindAllString(x, -1)
		for _, find := range findings {
			if find == "do()" {
				compute = true
				continue
			} else if find == "don't()" {
				compute = false
				continue
			}
			if compute {
				re2 := regexp.MustCompile(`mul\((\d*),(\d*)\)`)
				all := re2.FindAllStringSubmatch(find, -1)
				a, _ := strconv.Atoi(all[0][1])
				b, _ := strconv.Atoi(all[0][2])
				sum = sum + (a * b)
			}
		}
	}
	fmt.Println("Part 2:", sum)
	// 74838033
}

func main() {
	part1()
	part2()
}
