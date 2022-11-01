package main

import (
	"bufio"
	"fmt"
	"os"
)

type cube struct {
	x, y, z int
}

type hypercube struct {
	x, y, z, w int
}

func part1() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fileScanner := bufio.NewScanner(f)

	actives := make(map[cube]bool)
	x := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		for i, c := range line {
			if c == '#' {
				actives[cube{x, i, 0}] = true
			}
		}
		x++
	}
	f.Close()

	rounds := 6

	for i := 1; i <= rounds; i++ {
		actives = check_and_spin(actives)
	}
	fmt.Println(len(actives))
}

func check_and_spin(actives map[cube]bool) map[cube]bool {
	adjac_count := make(map[cube]int, 26)
	for p := range actives {
		for x := p.x - 1; x <= p.x+1; x++ {
			for y := p.y - 1; y <= p.y+1; y++ {
				for z := p.z - 1; z <= p.z+1; z++ {
					np := cube{x, y, z}
					if np == p {
						continue
					}
					adjac_count[np]++
				}
			}
		}
	}
	new_actives := make(map[cube]bool)

	for p, c := range adjac_count {
		if actives[p] {
			//check if a found adjacent cube is active
			if c == 2 || c == 3 {
				// and it has enough neibours
				new_actives[p] = true
			}
		} else {
			//if not active but has 3 neibours
			if c == 3 {
				new_actives[p] = true
			}
		}
	}

	return new_actives
}

func part2() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fileScanner := bufio.NewScanner(f)

	actives_hyper := make(map[hypercube]bool)
	x := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		for i, c := range line {
			if c == '#' {
				actives_hyper[hypercube{x, i, 0, 0}] = true
			}
		}
		x++
	}
	f.Close()

	rounds := 6

	for i := 1; i <= rounds; i++ {
		actives_hyper = check_and_spin_hyper(actives_hyper)
	}
	fmt.Println(len(actives_hyper))
}

func check_and_spin_hyper(actives map[hypercube]bool) map[hypercube]bool {
	//80 adjacent cubes
	adjac_count := make(map[hypercube]int, 80)
	for p := range actives {
		for x := p.x - 1; x <= p.x+1; x++ {
			for y := p.y - 1; y <= p.y+1; y++ {
				for z := p.z - 1; z <= p.z+1; z++ {
					for w := p.w - 1; w <= p.w+1; w++ {
						np := hypercube{x, y, z, w}
						if np == p {
							continue
						}
						adjac_count[np]++
					}
				}
			}
		}
	}
	new_actives := make(map[hypercube]bool)

	for p, c := range adjac_count {
		if actives[p] {
			//check if a found adjacent hypercube is active
			if c == 2 || c == 3 {
				// and it has enough neibours
				new_actives[p] = true
			}
		} else {
			//if not active but has 3 neibours
			if c == 3 {
				new_actives[p] = true
			}
		}
	}

	return new_actives
}
func main() {
	part1()
	part2()
}
