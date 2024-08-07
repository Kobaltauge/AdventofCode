package main

import (
	"fmt"
)

func splitToDigits(n int) []int {
	var ret []int

	for n != 0 {
		ret = append([]int{n % 10}, ret...)
		n /= 10
	}
	return ret
}

func part1() {
	count := 0
	for i := 240298; i <= 784956; i++ {
		slice := splitToDigits(i)
		incre := true
		double := false

		for idx, c := range slice[:len(slice)-1] {
			if slice[idx+1] < c {
				incre = false
				break
			}
			if slice[idx+1] == c && !double {
				double = true
			}
		}
		if double && incre {
			count++
		}

	}
	fmt.Println(count)
}

func part2() {
	count := 0
	for i := 240298; i <= 784956; i++ {
		// i = 111122
		i = 777889
		slice := splitToDigits(i)
		incre := true
		double := false
		triple := false
		even := false

		for idx, c := range slice[:len(slice)-1] {
			if slice[idx+1] < c {
				incre = false
				break
			}
			if slice[idx+1] == c && !double {
				double = true
			}
			if double && idx+2 < len(slice) {
				if slice[idx+2] == c {
					triple = true
				}
			}
			if !even {
				count := 0
				for t, _ := range slice {
					if c == slice[t] {
						count++
					}
				}
				if count%2 == 0 {
					even = true
					triple = false
				} else {
					even = false
				}
			}
		}

		if double && incre && !triple && even {
			fmt.Println(i)
			count++
		}

	}
	// wrong 470(low) 707 575(low)
	fmt.Println(count)
}

func main() {
	part1()
	part2()
}
