package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func stringsToIntegers(lines []string) []int {
	var integers []int
	for _, line := range lines {
		if line != "" {
			//fmt.Println(i, line)
			n, err := strconv.Atoi(line)
			if err != nil {
				// handle error
				fmt.Println(err)
				os.Exit(2)
				continue
			}
			integers = append(integers, n)
		}
	}

	return integers
}

func main() {

	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	result := 0
	result2 := 0

	for fileScanner.Scan() {
		x := fileScanner.Text()
		rest := strings.Split(x, " ")
		numb := stringsToIntegers(rest)

		rows := make([][]int, 1000)
		rows[0] = numb
		for idx := 0; idx >= 0; idx++ {
			for i := 0; i < len(rows[idx])-1; i++ {
				rows[idx+1] = append(rows[idx+1], rows[idx][i+1]-rows[idx][i])
			}
			zero := false
			for _, c := range rows[idx] {
				if c == 0 {
					zero = true
				} else {
					zero = false
					break
				}
			}
			if zero {
				rows = rows[:idx+1]
				break
			}
		}
		for idx := len(rows) - 1; idx >= 1; idx-- {
			if idx == len(rows)-1 {
				var ele int = 0
				rows[idx] = append([]int{ele}, rows[idx]...)
				rows[idx] = append([]int{ele}, rows[idx]...)
			}
			val := rows[idx][len(rows[idx])-1] + rows[idx-1][len(rows[idx-1])-1]
			var val2 int = rows[idx-1][0] - rows[idx][0]
			rows[idx-1] = append(rows[idx-1], val)
			rows[idx-1] = append([]int{val2}, rows[idx-1]...)
		}
		result = result + rows[0][len(rows[0])-1]
		result2 = result2 + rows[0][0]
		// for _, value := range rows {
		// 	fmt.Printf("%d - %d\n", len(value), value)
		// }
		// fmt.Println("\n")
	}
	fmt.Println("Part 1:", result)
	fmt.Println("Part 2:", result2)
}
