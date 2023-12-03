package main

import (
	"bytes"
	"fmt"
	"os"
	"slices"
)

func valcheck(data []byte, idx int, lenght int) bool {
	m := map[string]func(int, int) int{
		"n":  func(idx int, lenght int) int { return idx - lenght },
		"ne": func(idx int, lenght int) int { return idx - lenght + 1 },
		"e":  func(idx int, lenght int) int { return idx + 1 },
		"se": func(idx int, lenght int) int { return idx + lenght + 1 },
		"s":  func(idx int, lenght int) int { return idx + lenght },
		"sw": func(idx int, lenght int) int { return idx + lenght - 1 },
		"w":  func(idx int, lenght int) int { return idx - 1 },
		"nw": func(idx int, lenght int) int { return idx - lenght - 1 },
	}

	ways := []string{"n", "ne", "e", "se", "s", "sw", "w", "nw"}

	if (data[m["e"](idx, lenght)] >= 48) && (data[m["e"](idx, lenght)] <= 57) {
		return valcheck(data, idx+1, lenght)
	} else {
		for _, way := range ways {
			if !(m[way](idx, lenght) < 0) || (m[way](idx, lenght) > lenght) {
				//fmt.Println(data[m[way](idx, lenght)])
				//if !(data[m[way](idx, lenght)] == 46 || data[m[way](idx, lenght)] == 10 || data[m["e"](idx, lenght)] >= 48 || data[m["e"](idx, lenght)] >= 57) {
				numbers := []byte{46, 10, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57}
				if !(slices.Contains(numbers, data[m[way](idx, lenght)])) {
					return true
				}
			}

		}
		return false
	}
	return false
}

func part1() {
	data, _ := os.ReadFile("sample.txt")
	grid := bytes.Fields(data)
	lenght := len(grid) + 1
	fmt.Println(grid)
	res := make([]byte, lenght)
	for idx, _ := range data {
		if !(valcheck(data, idx, lenght)) {
			//res[idx] = byte(".")
			copy(res[idx:], ".")
		} else {
			res[idx] = data[idx]
		}
	}

	// fmt.Println("Part 1:", result)
}

func main() {
	part1()
}
