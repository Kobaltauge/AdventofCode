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

	result := 0
	red := 12
	green := 13
	blue := 14

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		y := fileScanner.Text()
		x := strings.Split(y, ";")
		ok := true
		gameid, _ := strconv.Atoi(strings.Split(strings.Split(x[0], ":")[0], " ")[1])
		y = y[strings.Index(y, ":")+2:]
		x = strings.Split(y, ";")
		for _, val := range x {
			allcubes := map[string]int{
				"blue":  0,
				"red":   0,
				"green": 0,
			}
			for _, val2 := range strings.Split(val, ",") {
				amount := 0
				tmp := strings.Split((strings.TrimSpace(val2)), " ")
				amount, _ = strconv.Atoi(tmp[0])
				color := tmp[1]
				allcubes[color] = allcubes[color] + amount
			}
			if allcubes["blue"] > blue || allcubes["red"] > red || allcubes["green"] > green {
				ok = false
			}
		}
		if ok {
			result = result + gameid
		}

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
		y := fileScanner.Text()
		y = y[strings.Index(y, ":")+2:]
		x := strings.Split(y, ";")
		allcubes := map[string]int{
			"blue":  0,
			"red":   0,
			"green": 0,
		}
		for _, val := range x {
			for _, val2 := range strings.Split(val, ",") {
				tmp := strings.Split((strings.TrimSpace(val2)), " ")
				amount, _ := strconv.Atoi(tmp[0])
				color := tmp[1]
				if allcubes[color] < amount {
					allcubes[color] = amount
				}
			}
		}
		result = result + (allcubes["red"] * allcubes["green"] * allcubes["blue"])
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
