package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type knod struct {
	x, y int
}

func follow(head, tail knod) (knod, knod) {
	x := head.x - tail.x
	y := head.y - tail.y
	xg := 1
	yg := 1
	if (x > 1) || (x < -1) {
		yg = 0
	}
	if (y > 1) || (y < -1) {
		xg = 0
	}
	if (x > xg) || (x < (-1 * xg)) {
		if x > 0 {
			tail.x++
		} else {
			tail.x--
		}
	}
	if (y > yg) || (y < (-1 * yg)) {
		if y > 0 {
			tail.y++
		} else {
			tail.y--
		}
	}
	return head, tail
}

func unique(slice []knod) []knod {
	keys := make(map[knod]bool)
	list := []knod{}
	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func part1() {

	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	head := knod{x: 0, y: 0}
	tail := knod{x: 0, y: 0}

	var path []knod

	for fileScanner.Scan() {
		raw := strings.Fields(fileScanner.Text())
		direc := raw[0]
		quant, _ := strconv.Atoi(raw[1])
		for i := 1; i <= quant; i++ {
			switch direc {
			case "R":
				head.x++
			case "L":
				head.x--
			case "U":
				head.y++
			case "D":
				head.y--
			}
			head, tail = follow(head, tail)
			path = append(path, tail)
		}
	}
	path = unique(path)

	fmt.Println("Part 1:", len(path))
}

func main() {
	part1()
}
