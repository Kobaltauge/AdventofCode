package main

import "fmt"

func monkeyop(m int, old uint64) uint64 {
	switch m {
	case 0:
		return old * 7
	case 1:
		return old + 4
	case 2:
		return old + 2
	case 3:
		return old + 7
	case 4:
		return old * 17
	case 5:
		return old + 8
	case 6:
		return old + 6
	case 7:
		return old * old
	}
	return 0
}

func monkeytest(m int, val uint64) bool {
	switch m {
	case 0:
		if val%17 == 0 {
			return true
		} else {
			return false
		}
	case 1:
		if val%3 == 0 {
			return true
		} else {
			return false
		}
	case 2:
		if val%5 == 0 {
			return true
		} else {
			return false
		}
	case 3:
		if val%7 == 0 {
			return true
		} else {
			return false
		}
	case 4:
		if val%11 == 0 {
			return true
		} else {
			return false
		}
	case 5:
		if val%19 == 0 {
			return true
		} else {
			return false
		}
	case 6:
		if val%2 == 0 {
			return true
		} else {
			return false
		}
	case 7:
		if val%13 == 0 {
			return true
		} else {
			return false
		}
	}
	return false
}

func part2() {
	counter := map[int]uint64{0: 0, 1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0}
	monkeys := map[int][]uint64{
		0: []uint64{54, 89, 94},
		1: []uint64{66, 71},
		2: []uint64{76, 55, 80, 55, 55, 96, 78},
		3: []uint64{93, 69, 76, 66, 89, 54, 59, 94},
		4: []uint64{80, 54, 58, 75, 99},
		5: []uint64{69, 70, 85, 83},
		6: []uint64{89},
		7: []uint64{62, 80, 58, 57, 93, 56},
	}

	monkeytar := map[int][]int{
		0: []int{5, 3},
		1: []int{0, 3},
		2: []int{7, 4},
		3: []int{5, 2},
		4: []int{1, 6},
		5: []int{2, 7},
		6: []int{0, 1},
		7: []int{6, 4},
	}

	for round := 1; round <= 10000; round++ {
		fmt.Println(round)
		for x := 0; x <= 7; x++ {
			items := len(monkeys[x])
			for i := 0; i < len(monkeys[x]); i++ {
				val := monkeys[x][i]
				// fmt.Println(i, val, (monkeys[x]))
				w := monkeyop(x, val)
				// w = (int(w / 3))
				if monkeytest(x, w) {
					xx := monkeytar[x][0]
					monkeys[xx] = append(monkeys[xx], w)
				} else {
					xx := monkeytar[x][1]
					monkeys[xx] = append(monkeys[xx], w)
				}
			}

			counter[x] = counter[x] + uint64(items)
			fmt.Println(x, uint64(items), counter[x])
			monkeys[x] = monkeys[x][0:0]
		}
	}
	var st uint64 = 0
	var nd uint64 = 0
	for i := 0; i <= 7; i++ {
		x := counter[i]
		if x > st {
			nd = st
			st = x
		}
		if (x != st) && (x > nd) {
			nd = x
		}

	}
	fmt.Println("Part 1:", st*nd)
}

func main() {
	// part1()
	part2()
}
