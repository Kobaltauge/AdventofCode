package main

import (
	"bufio"
	"fmt"
	"os"
)

func part1() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	score := 0

	for fileScanner.Scan() {
		game := fileScanner.Text()
		opp := game[0]
		me := game[2]
		midscore := 0
		switch opp {
		case 65:
			switch me {
			case 88:
				midscore = 3 + 1
			case 89:
				midscore = 6 + 2
			case 90:
				midscore = 0 + 3
			}
		case 66:
			switch me {
			case 88:
				midscore = 0 + 1
			case 89:
				midscore = 3 + 2
			case 90:
				midscore = 6 + 3
			}
		case 67:
			switch me {
			case 88:
				midscore = 6 + 1
			case 89:
				midscore = 0 + 2
			case 90:
				midscore = 3 + 3
			}
		}
		score = score + midscore
	}
	fmt.Println("Part 1:", score)
	f.Close()
}

func part2() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	score := 0

	for fileScanner.Scan() {
		game := fileScanner.Text()
		opp := game[0]
		me := game[2]
		midscore := 0
		switch opp {
		case 65:
			switch me {
			case 88:
				midscore = 0 + 3
			case 89:
				midscore = 3 + 1
			case 90:
				midscore = 6 + 2
			}
		case 66:
			switch me {
			case 88:
				midscore = 0 + 1
			case 89:
				midscore = 3 + 2
			case 90:
				midscore = 6 + 3
			}
		case 67:
			switch me {
			case 88:
				midscore = 0 + 2
			case 89:
				midscore = 3 + 3
			case 90:
				midscore = 6 + 1
			}
		}
		score = score + midscore
	}
	fmt.Println("Part 2:", score)
	f.Close()
}

func main() {
	part1()
	part2()
}
