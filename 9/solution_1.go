package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, _ := os.Open("test.txt")
	defer f.Close()

	s := bufio.NewScanner(f)
	var readInt int
	idx := 0
	readOrder := [1000]int{}
	possibleSums := map[int][]int{}
	sumSet := map[int]int{}

	for idx < 25 {
		if !s.Scan() {
			break
		}
		readInt, _ = strconv.Atoi(s.Text())
		readOrder[idx] = readInt
		idx++
		for k, v := range possibleSums {
			possibleSums[k] = append(v, readInt)
			sumSet[k+readInt] += 1
		}
		possibleSums[readInt] = make([]int, 0)
	}

	var firstWrong int
	for s.Scan() {
		readInt, _ = strconv.Atoi(s.Text())
		if _, ok := sumSet[readInt]; !ok {
			firstWrong = readInt
			break
		}
		readOrder[idx] = readInt
		for _, v := range possibleSums[readOrder[idx-25]] {
			sumSet[readOrder[idx-25]+v] -= 1
			if sumSet[readOrder[idx-25]+v] == 0 {
				delete(sumSet, readOrder[idx-25]+v)
			}
		}
		delete(possibleSums, readOrder[idx-25])

		for k, v := range possibleSums {
			possibleSums[k] = append(v, readInt)
			sumSet[k+readInt] += 1
		}
		possibleSums[readInt] = make([]int, 0)
		idx += 1
	}

	fmt.Println("Day 9 part 1:", firstWrong)
}
