package main

import (
	// "fmt"
	"fmt"
	"strconv"
    	"bufio"
    	"os"
)

// import "strings"
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
	start := 50
	sample, err := readLines("sample.txt")
	if err != nil {
    		panic(err)
	}

	count := 0
	// fmt.Println(len(sample))

	for _, val := range(sample) {
		ch := string(val[0])
		value, err := strconv.Atoi(val[1:])

		fmt.Println(start, ch, value)
		if err != nil {
			return
		}
	
		curr := value / 100

		if value > 100 {
			for value > 100 {
				value = value % 100
			}
		}

		count += curr
		// fmt.Println(curr, " curr")
		

		if ch == "L"{
			start = start - value
			if(start < 0) {
				if start != -value {
					count += 1
				}
				start = 100 + start
				if (start == 0) {
					count += 1
				}
			} else if (start == 0) {
				count += 1
			}

		} else if ch == "R" {
			start = start + value
			if start > 100 {
				count += 1
				start = start % 100
				if(start == 0) {
					count += 1
				}
			} else if start == 100 {
				count += 1
				start = 0
			}
		// fmt.Println(start, " R")
		}
		fmt.Println(ch, count)
	}
	fmt.Println(count)

}
