package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func firstDigit(s string) int {
	for _, r := range s {
		if unicode.IsDigit(r) {
			return int(r - '0')
		}
	}
	return 0
}

func lastDigit(s string) int {
	runes := []rune(s)
	for i := len(runes) - 1; i >= 0; i-- {
		if unicode.IsDigit(runes[i]) {
			return int(runes[i] - '0')
		}
	}
	return 0
}

func DigitPrefix(s string) string {
	for i, r := range s {
		if unicode.IsDigit(r) {
			return s[:i]
		}
	}
	return s
}

func DigitSuffix(s string) string {
	runes := []rune(s)
	for i := len(runes) - 1; i >= 0; i-- {
		if unicode.IsDigit(runes[i]) {
			return s[i+1:]
		}
	}
	return s
}

func FrontString(s string) string {
	things := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for i, r := range s {
		if unicode.IsDigit(r) {
			return s
		}
		for idx, _ := range things {
			if strings.Contains(s[:i+1], things[idx]) {
				// 1 so only the first finding is replaced, because it conflicts when searching later from the end
				s = strings.Replace(s, things[idx], strconv.Itoa(idx+1), 1)
				return s
			}
		}
	}
	return s
}

func BackString(s string) string {
	things := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	runes := []rune(s)
	for i := len(runes) - 1; i >= 0; i-- {
		if unicode.IsDigit(runes[i]) {
			return s
		}
		for idx, _ := range things {
			if strings.Contains(s[i:], things[idx]) {
				// -1 to replace all. this doesn't matter when searching from the end
				s = strings.Replace(s, things[idx], strconv.Itoa(idx+1), -1)
				return s
			}
		}
	}
	return s
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
		first := firstDigit(x)
		last := lastDigit(x)
		res := first*10 + last
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
	result := 0

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		x := fileScanner.Text()
		fmt.Println(x)
		x = FrontString(x)
		fmt.Println(x)
		x = BackString(x)
		fmt.Println(x)
		first := firstDigit(x)
		last := lastDigit(x)

		res := first*10 + last
		fmt.Println(res)
		fmt.Println("")
		result = result + res

	}
	f.Close()
	fmt.Println(result)
}

func main() {
	part1()
	part2()
}
