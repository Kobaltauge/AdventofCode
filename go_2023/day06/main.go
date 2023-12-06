package main

import (
	"fmt"
)

func part1() {
	times := []int{54, 94, 65, 92}
	distance := []int{302, 1476, 1029, 1404}

	result := 1
	for i := 0; i < 4; i++ {
		res := 0
		time := times[i]
		dist := distance[i]
		for t := 1; t <= time-1; t++ {
			if (time-t)*t > dist {
				res++
			}
		}
		result = result * res
	}

	fmt.Println(result)
}

func part2() {
	time := 54946592
	dist := 302147610291404

	res := 0
	for t := 1; t <= time-1; t++ {
		if (time-t)*t > dist {
			res++
		}
	}

	fmt.Println(res)
}

func main() {
	fmt.Println("Part 1")
	part1()
	fmt.Println("Part 2")
	part2()
}
