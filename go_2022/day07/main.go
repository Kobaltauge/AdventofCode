package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Directory struct {
	Parent string
	Files  map[string]int
	Childs []string
}

func part1() {

	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)
	parent := ""
	actual := ""
	m := make(map[string]Directory)

	for fileScanner.Scan() {
		raw := strings.Fields(fileScanner.Text())
		fmt.Println(raw, parent)
		sw := ""
		if _, err := strconv.Atoi(raw[0]); err == nil {
			sw = "no"
		} else {
			sw = raw[0]
		}

		switch sw {
		case "$":
			switch raw[1] {
			case "cd":
				parent = actual
				actual = raw[2]
				_, ok := m[actual]
				if !ok {
					m[actual] = Directory{Parent: parent}
				}
			case "ls":
				continue
			}
		case "dir":
			m[raw[1]] = Directory{Parent: actual}
			ad := m[actual]
			ad.Childs = append(ad.Childs, raw[1])

		case "no":
			v, _ := strconv.Atoi(raw[0])
			ad := m[actual]
			af := ad.Files
			// https://stackoverflow.com/questions/18042439/go-append-to-slice-in-struct
			fmt.Println(raw, parent)
			af[raw[1]] = v
			fmt.Println(raw, parent)

		}
		// for _, x := range raw {
		// 	fmt.Println(x)
		// 	x = string(x)
		// 	switch x {
		// 	case "$":
		// 		fmt.Println("command")
		// 	case "dir":
		// 		fmt.Println("dir")
		// 	}
		// }
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
