package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	lines := strings.Split(string(content), "\n")
	line := strings.Split(lines[0], ",")
	var prog []int

	for _, no := range line {
		x, err := strconv.Atoi(no)
		if err != nil {
			fmt.Println(err)
			return
		}
		prog = append(prog, x)
	}
	// prog := [...]int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}
	// prog := [...]int{2, 4, 4, 5, 99, 0}
	// prog := [...]int{1, 1, 1, 4, 99, 5, 6, 0, 99}
	pointer := 0
	prog[1] = 12
	prog[2] = 2
	for {
		inst := prog[pointer]
		if inst == 1 {
			res := prog[prog[pointer+1]] + prog[prog[pointer+2]]
			prog[prog[pointer+3]] = res
		} else if inst == 2 {
			res := prog[prog[pointer+1]] * prog[prog[pointer+2]]
			prog[prog[pointer+3]] = res
		} else if inst == 99 {
			break
		} else {
			fmt.Println("Error")
		}
		pointer += 4
	}
	// fmt.Println(prog)
	fmt.Println(prog[0])
}

func part2() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	lines := strings.Split(string(content), "\n")
	line := strings.Split(lines[0], ",")
	var prog []int

	for _, no := range line {
		x, err := strconv.Atoi(no)
		if err != nil {
			fmt.Println(err)
			return
		}
		prog = append(prog, x)
	}
	// prog := [...]int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}
	// prog := [...]int{2, 4, 4, 5, 99, 0}
	// prog := [...]int{1, 1, 1, 4, 99, 5, 6, 0, 99}
	initprog := make([]int, len(prog))
	copy(initprog, prog)
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb <= 100; verb++ {
			copy(prog, initprog)
			result := calc(prog, noun, verb)
			if result == 19690720 {
				fmt.Println(100*noun + verb)
			}
		}
	}
}

func calc(prog []int, noun int, verb int) int {
	prog[1] = noun
	prog[2] = verb
	pointer := 0
	for {
		inst := prog[pointer]
		if inst == 1 {
			res := prog[prog[pointer+1]] + prog[prog[pointer+2]]
			prog[prog[pointer+3]] = res
		} else if inst == 2 {
			res := prog[prog[pointer+1]] * prog[prog[pointer+2]]
			prog[prog[pointer+3]] = res
		} else if inst == 99 {
			break
		} else {
			fmt.Println("Error")
		}
		pointer += 4
	}
	// fmt.Println(prog)
	return prog[0]
}

func main() {
	part1()
	part2()
}
