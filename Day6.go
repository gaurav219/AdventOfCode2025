package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func main() {

	sample, err := readLines("Day6.txt")
	if err != nil {
		panic(err)
	}
	// fmt.Println(sample)

	arr := make([][]string, len(sample))
	// fmt.Println(len(arr))

	// curr := strings.Fields(sample[0])
	// fmt.Println("curr", curr)
	// for k := range(arr) {
	// 	arr[k] = make([]int, len(curr[0]))
	// }

	// fmt.Println(len(arr[0]))
	for i, line := range sample { // i = row index
		fields := strings.Fields(line) // splits on any whitespace
		// fmt.Println(fields)

		// Allocate the row with correct number of columns
		arr[i] = make([]string, len(fields))

		for j, f := range fields { // j = column index
			// num, err := strconv.Atoi(f)
			// if err != nil {
			//     panic(err)
			// }
			arr[i][j] = f
		}
	}

	ans := 0
	fmt.Println(len(arr[0]))
	for i := range arr[0] {
		ch := arr[len(arr)-1][i]
		flag := true
		small := 0
		if ch == "*" {
			flag = false
			small = 1
		}
		// fmt.Println(len(arr[i]))
		for j := 0; j < len(arr)-1; j++ {
			num1, err := strconv.Atoi(arr[j][i])
			// small = num1
			if err != nil {
				return
			}
			if flag {
				small += num1
			} else {
				small *= num1
			}

			fmt.Println(small, "num1")
			// fmt.Println(arr[j][i])
		}
		ans += small
	}

	fmt.Println(ans)
	// for _, val := range(sample) {
	// 	curr := strings.Fields(val)
	// 	// fmt.Println(curr[i], "curr")
	// 	// len(curr[[0]])
	// 	for _, val1 := range(curr) {
	// 		num1, err := strconv.Atoi(val1)
	// 		if err != nil {
	// 			return
	// 		}
	// 		fmt.Println(num1, "num1")
	// 		 arr[i][j] = num1

	// 	}
	// }
}
