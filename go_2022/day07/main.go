package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func part1() {

	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	var path = map[string]string{}

	for fileScanner.Scan() {
		raw := strings.Fields(fileScanner.Text())
		fmt.Println(raw)
		switch raw[0] {
		case "$":
			switch raw[0] {

			}

		}
		for _, x := range raw {
			fmt.Println(x)
			x = string(x)
			switch x {
			case "$":
				fmt.Println("command")
			case "dir":
				fmt.Println("dir")
			}
		}
	}
	// quant, _ := strconv.Atoi(x[1])
	// from, _ := strconv.Atoi(x[3])
	// to, _ := strconv.Atoi(x[5])

	// }

	// for idx := 1; idx <= quant; idx++ {
	// 	el := m[from][len(m[from])-1]
	// 	m[to] = append(m[to], el)
	// 	m[from] = m[from][:len(m[from])-1]
	// }
	// }
	// fmt.Println("Part 1:")
	// for idx := 1; idx < len(m); idx++ {
	// fmt.Print(m[idx][len(m[idx])-1])
	// }
	// fmt.Println("")
}

func main() {
	part1()
}
