package main

import (
	"bufio"
	"fmt"
	"os"
)

func part1() {
	f, err := os.Open("sample.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	var treemap []string
	fileScanner := bufio.NewScanner(f)

	for fileScanner.Scan() {
		treemap = append(treemap, fileScanner.Text())
	}
	f.Close()

	x := 0
	y := 0
	counter := 0

	for y <= len(treemap) {
		x += 3
		if x >= len(treemap[0])-1 {
			x = x - len(treemap[0]) - 1
		}
		y += 1
		if string(treemap[x][y]) == "#" {
			counter += 1
		}

	}
	fmt.Println(counter)
	// x := strings.Fields(fileScanner.Text())
	// subarr := strings.Split(x[0], "-")
	// min, _ := strconv.Atoi(subarr[0])
	// max, _ := strconv.Atoi(subarr[len(subarr)-1])
	// subarr = strings.Split(x[1], ":")
	// lett := subarr[0]
	// occur := strings.Count(x[2], lett)
	// if occur >= min && occur <= max {
	// 	counter++
	// }
	// }
}

func main() {
	part1()
}
