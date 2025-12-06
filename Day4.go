package main

import (
	"fmt"
	"os"
	"bufio"
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
	
	sample, err := readLines("Day4.txt")
	if err != nil {
    		panic(err)
	}
	
	// n := len(sample)
	arr := make([][]rune, len(sample))

	for i, row := range sample {
    		arr[i] = []rune(row)  // split each string into runes (characters)
	}

	sample(arr)
	// fmt.Println(string(arr[0][2]))

}

func sample(i int, j int, arr [][]int, count int) {
	
	if(i < 0 || j < 0 || i >= len(arr) || j >= len(arr)) {
		return
	}

	if(rune(arr[i][j]) == '@') {
		curr := getCount(i, j , arr, 0)
		if curr < 3 {
			count += 1
		}
	}

	sample(i, j + 1, arr, count)
	sample(i + 1, j, arr, count)
	
}

func getCount(i int, j int, arr [][]int, count int) int {

	if(i < 0 || j < 0 || i >= len(arr) || j >= len(arr)) {
		return 0
	}

	if(rune(arr[i][j]) == '.') {
		return 0
	}

	ans1 := getCount(i - 1, j - 1, arr, count + 1)
	ans2 := getCount(i - 1, j, arr, count + 1)
	ans3 := getCount(i - 1, j + 1, arr, count + 1)
	ans4 := getCount(i, j - 1, arr, count + 1)
	ans5 := getCount(i, j + 1, arr, count + 1)
	ans6 := getCount(i + 1, j - 1, arr, count + 1) 
	ans7 := getCount(i + 1, j, arr, count + 1)
	ans8 := getCount(i + 1, j + 1, arr, count + 1)

	return ans1 + ans2 + ans3 + ans4 + ans5 + ans6 + ans7 + ans8
}
