package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var monkeys = [int][]int{ 	0: [54, 89, 94],
							1: [66, 71],
							2: [76, 55, 80, 55, 55, 96, 78],
							3: [93, 69, 76, 66, 89, 54, 59, 94],
							4: [80, 54, 58, 75, 99],
							5: [69, 70, 85, 83],
							6: [89],
							7: [62, 80, 58, 57, 93, 56]}

var monkeytar := [int][]int{0: [5,3],
	1: [0,3],
	2: [7,4],
	3: [5,2],
	4: [1,6],
	5: [2,7],
	6: [0,1],
	7: [6,4]

}

func monkeyop(m, old int) int {
	switch m {
		0:
		return old*7
		1:
		return old +4
		2:
		return old +2
		3:
		return old+7
		4:
		return old *17
		5:
		return old+8
		6:
		return old+6
		7:
		return old*old
	}
}

func monkeytest(m, val int) bool {
	switch m {
		0:
		if val%17 == 0 {return true} else {false}
		1:
		if val%3 == 0 {return true} else {false}
		2:
		if val%5 == 0 {return true} else {false}
		3:
		if val%7 == 0 {return true} else {false}
		4:
		if val%11 == 0 {return true} else {false}
		5:
		if val%19 == 0 {return true} else {false}
		6:
		if val%2 == 0 {return true} else {false}
		7:
		if val%13 == 0 {return true} else {false}
	}
}

							
func part1() {
	for round := 0; round >0;{
	for _,x:= range(monkeys) {
		for _,i:= range(monkeys[x]){
			w := monkeyop(x, i)
			w = (int(w/300))
			if (monkeytest(x, w)) {
				monkeys[x] = append(monkeytar[x][0], w)
			} else {
				monkeys[x] = append(monkeytar[x][1], w)
			}

		}
	}
}
	fmt.Println("Part 1:", result)
}

func main() {
	part1()
}
