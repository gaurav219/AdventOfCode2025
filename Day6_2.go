package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	// "strings"
	"strconv"
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

	arr := make([][]string, 0, len(sample))
	fields1 := make([]string, 0, len(sample))
	for _, line := range sample { // i = row index
		// curr := ""
		fields1 = strings.Split(line, "")

		// fmt.Println(fields1)
		arr = append(arr, fields1)

	}

	ans := 0
	flag := false

	final_ans := 0
	for i := 0; i < len(arr[0]); i++ {
		if (arr[len(arr)-1])[i] == "" {
			continue
		}
		operand := arr[len(arr)-1][i]
		if operand == "*" {
			flag = true
			ans = 1
		}
		spaces_count := 0

		num1 := ""
		for j := 0; j < len(arr)-1; j++ {
			if arr[j][i] == " " {
				spaces_count += 1
			} else {
				num1 += arr[j][i]
			}
			// fmt.Println(arr[j][i])
		}

		if spaces_count == len(arr)-1 {
			final_ans += ans
			ans = 0
			flag = false
			continue
		}
		if num1 == "" {
			continue
		}
		num2, err := strconv.Atoi(num1)
		if err != nil {
			fmt.Println("Error:", err)
		}

		if flag {
			ans *= num2
		} else {
			ans += num2
		}

		if i == len(arr[0])-1 {
			final_ans += ans
		}

		// final_ans += ans

	}
	fmt.Println(final_ans)

	// ans := 0
	// for i := range arr[0] {
	// ch := arr[len(arr) - 1][i]
	// flag := true
	// small := 0
	// if ch == "*" {
	// flag = false
	// small = 1
	// }
	// // fmt.Println(i, ch)
	// for j:=0; j < len(arr) - 1; j ++ {
	// num1, err := strconv.Atoi(arr[j][i])
	// // small = num1
	// if err != nil {
	// return
	// }
	// if flag  {
	// small += num1
	// } else {
	// small *= num1
	// }

	// fmt.Println(small, "num1")
	//     	// fmt.Println(arr[j][i])
	// }
	// ans += small
	// }

	// fmt.Println(ans)
}
