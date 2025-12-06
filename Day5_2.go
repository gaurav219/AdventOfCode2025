package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
	"sort"
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

// func main() {
	
// 	sample, err := readLines("Day5.txt")
// 	if err != nil {
//     		panic(err)
// 	}

// 	fmt.Println(sample)
// 	flag := false

// 	fresh_list := make(map[uint64]int)
// 	count := 0
	
// 	for _, val := range(sample) {
// 		if val == "" {
// 			flag = true
// 			continue
// 		}
// 		if flag {
// 			val, err := strconv.Atoi(val)
// 			if err != nil {
//     				panic(err)
// 			}

// 			_, ok := fresh_list[uint64(val)]
// 			if ok {
// 				count += 1
// 			}
// 			continue
// 		}

// 		split := strings.Split(val, "-")
// 		initial := split[0]
// 		end := split[1]

// 		first_num, err := strconv.Atoi(initial)
// 		if err != nil {
// 			return 
// 		}
// 		end_num, err := strconv.Atoi(end)

// 		if err != nil {
// 			return 
// 		}

// 		for i := first_num; i <= end_num; i ++ {
// 			_, ok := fresh_list[uint64(i)]
// 			if !ok {
// 				fresh_list[uint64(i)] = 0
// 			}
// 		}
// 	}
// }

type Interval struct {
    L int
    R int
}

func mergeIntervals(intervals []Interval) []Interval {
    if len(intervals) == 0 {
        return intervals
    }

    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i].L < intervals[j].L
    })

    merged := []Interval{intervals[0]}
    fmt.Println("merged123", merged)

    for _, cur := range intervals[1:] {
        last := &merged[len(merged)-1]
	// fmt.Println("last", last)
        if cur.L <= last.R {
            if cur.R > last.R {
                last.R = cur.R
            }
        } else {
            merged = append(merged, cur)
        }
    }

    return merged
}

func contains(intervals []Interval, x int) bool {
    if len(intervals) == 0 {
        return false
    }

    i := sort.Search(len(intervals), func(i int) bool {
        return intervals[i].L > x
    })

    if i == 0 {
        return false
    }

    iv := intervals[i-1]
    return x >= iv.L && x <= iv.R
}

func main() {
    // intervals := []Interval{
    //     {3, 5},
    //     {10, 14},
    //     {16, 20},
    //     {12, 18},
    // }

	sample, err := readLines("Day5.txt")
	if err != nil {
    		panic(err)
	}

	fmt.Println(sample)
    	intervals := make([]Interval, 0, 0)
	count := 0
	
	for _, val := range(sample) {
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

		intervals = append(intervals, Interval{first_num, end_num})
	}

	merged := mergeIntervals(intervals)

	for _, val := range(merged) {
		curr := val.R - val.L + 1
		count += curr
	}

	fmt.Println("Count is", count)
}
