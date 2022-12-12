package main

import "fmt"

func monkeyop(m int, old uint64) uint64 {
	switch m {
	case 0:
		return old * 19
	case 1:
		return old + 6
	case 2:
		return old * old
	case 3:
		return old + 3
	}
	return 0
}

func monkeytest(m int, val uint64) bool {
	switch m {
	case 0:
		if val%23 == 0 {
			return true
		} else {
			return false
		}
	case 1:
		if val%19 == 0 {
			return true
		} else {
			return false
		}
	case 2:
		if val%13 == 0 {
			return true
		} else {
			return false
		}
	case 3:
		if val%17 == 0 {
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
		0: []uint64{79, 98},
		1: []uint64{54, 65, 75, 74},
		2: []uint64{79, 60, 97},
		3: []uint64{74},
	}

	monkeytar := map[int][]int{
		0: []int{2, 3},
		1: []int{2, 0},
		2: []int{1, 3},
		3: []int{0, 1},
	}

	for round := 1; round <= 10000; round++ {
		fmt.Println(round)
		for x := 0; x <= 3; x++ {
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
