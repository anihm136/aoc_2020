package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	wrong := 258585477
	f, _ := os.Open("test.txt")
	defer f.Close()

	s := bufio.NewScanner(f)
	idx := -1
	readOrder := [1000]int{}
	cumulativeSum := [1001]int{}
	var start, end int
	var done bool

	for s.Scan() {
		idx += 1
		readOrder[idx], _ = strconv.Atoi(s.Text())
		if idx == 0 {
			cumulativeSum[1] = readOrder[0]
			continue
		}
		cumulativeSum[idx+1] = cumulativeSum[idx] + readOrder[idx]
		if idx == 1 {
			continue
		}
		for i := idx - 1; i >= 0; i-- {
			if cumulativeSum[idx+1]-cumulativeSum[i] == wrong {
				start = i
				end = idx
				done = true
				break
			}
		}
		if done {
			break
		}
	}

	var min, max int
	min, max = readOrder[start], readOrder[start]
	for _, val := range readOrder[start : end+1] {
		if val < min {
			min = val
		} else if val > max {
			max = val
		}
	}
	fmt.Println(min + max)

}
