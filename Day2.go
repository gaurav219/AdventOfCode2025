package main

import (
	"fmt"
	"strings"
	"os"
    "bufio"
	"strconv"
)

func getLength(str string) int {

	count := 0

	num, err := strconv.Atoi(str)

	if err != nil {
		return -1
	}

	for num > 0 {
		count += 1
		num = num / 10
	}

	return count
}

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
	
	sample, err := readLines("Day2.txt")
	if err != nil {
    		panic(err)
	}
	input := strings.Split(strings.Join(sample, " "), ",")

	var ans uint64
	
	for _, val := range(input) {
		split := strings.Split(val, "-")
		initial := split[0]
		end := split[1]
		first_num, err := strconv.Atoi(initial)
		if err != nil {
			return 
		}
		end_num, err := strconv.Atoi(end)

		if err != nil {
			return 
		}

		initial_length := getLength(initial)
		end_length := getLength(end)

		if(initial_length % 2 != 0) {
			if (end_length % 2 != 0) {
				continue
			}
		} else if(initial_length % 2 == 0) {
			first_mid := val[0:initial_length / 2]
			end_mid := end[0:end_length / 2]

			// fmt.Println(mid)
			first_mid_num, err := strconv.Atoi(first_mid)
			if err != nil {
				return 
			}

			end__mid_num, err := strconv.Atoi(end_mid)
			if err != nil {
				return 
			}

			small := end__mid_num - first_mid_num - 1

			ans += uint64(small)

			manual_left, err := strconv.Atoi(first_mid + first_mid)
			if err != nil {
				return 
			}

			manual_right, err := strconv.Atoi(end_mid + end_mid)
			if err != nil {
				return 
			}

			if(manual_left >= first_num) {
				ans += 1
			}

			if(manual_right <= end_num) {
				ans += 1
			}
		}
	}
}
