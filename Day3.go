package main

import (
	"fmt"
	// "strings"
	"os"
    	"bufio"
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

func get_max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func getLargestKDigits(str string, k int) int {
    n := len(str)
    result := ""
    startIndex := 0
    
    for i := 0; i < k; i++ {
        // How many digits we still need to pick (including current)
        remaining := k - i
        // Window end: leave enough digits for future picks
        windowEnd := n - remaining + 1
        
        // Find max digit in the current window
        maxDigit := -1 // Start below '0'
        maxIndex := startIndex
        
        for j := startIndex; j < windowEnd; j++ {
 	    // fmt.Println(int(str[j] - '0'), " sca")
            if int(str[j] - '0') > maxDigit {
                maxDigit = int(str[j] - '0')
                maxIndex = j
            }
        }
        
        // result += string(maxDigit)
	small := strconv.Itoa(maxDigit)
	fmt.Println(small, "small")
	result += small
        startIndex = maxIndex + 1  // Next search starts after this digit
    }
 
	num, err := strconv.Atoi(result)
	if err != nil {
    		fmt.Println("Error:", err)
	}
	
	// fmt.Println(num)
    	return num
}

func main() {
	
	sample, err := readLines("Day3.txt")
	if err != nil {
    		panic(err)
	}

	// keymap := make(map[int][]int)
	var ans uint64

	for _, val := range(sample) {
		fmt.Println(len(val), "val")
		curr := getLargestKDigits(val, 12)
		ans += uint64(curr)

		// f_max := int(val[len(val) - 2] - '0')
		// smax := int(val[len(val) - 1] - '0')

		// for i := len(val) - 3; i >= 0; i -- {
		// 	num := int(val[i] - '0')
		// 	if num >= f_max {
		// 		smax = get_max(smax, f_max)
		// 		f_max = num
		// 	}
		// }

		// fmt.Println(f_max, smax)
		// small := strconv.Itoa(f_max) + strconv.Itoa(smax)
		// // fmt.Println(small, "small")
		// num1, err := strconv.Atoi(small)
		// if err != nil {
		// 	return 
		// }

		// ans += num1
    
		// for i, ch := range val[:len(val) - 1] {
		// 	num := int(ch - '0')
		// 	max = get_max(num, max)
		// 	num1 = int(val[i+1] - '0')
		// 	smax = 
		// 	if(smax > max && i + 1 != len(val) - 1) {
		// 		max = smax
		// 	}
		// 	// val, ok := keymap[num]
		// 	// arr := make([]int, len(val), len(val))
		// 	// if ok {
		// 	// 	arr = keymap[num]
		// 	// }
		// 	// arr = append(arr, i)
		// 	// keymap[num] = arr
		// }
	}
	fmt.Println(ans)
}

