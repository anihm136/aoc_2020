package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ConvertToInt(arr []string) []int {
	var temp int
	ret := make([]int, 0, len(arr))
	for _, val := range arr {
		temp, _ = strconv.Atoi(val)
		ret = append(ret, temp)
	}
	return ret
}

type history struct {
	prev1 int
	prev2 int
}

func main() {
	f, _ := os.Open("test.txt")
	s := bufio.NewScanner(f)

	s.Scan()
	in := ConvertToInt(strings.Split(s.Text(), ","))
	called := map[int]history{}
	max := len(in)
	idx := 0

	lastSpoken := -1
	for {
		if _, ok := called[in[idx%max]]; !ok {
			lastSpoken = in[idx%max]
			called[lastSpoken] = history{prev1: idx, prev2: -1}
		} else {
			hist := called[lastSpoken]
			if hist.prev2 == -1 {
				lastSpoken = 0
				hist, ok = called[lastSpoken]
			} else {
				lastSpoken = hist.prev1 - hist.prev2
				hist, ok = called[lastSpoken]
			}
			if ok {
				called[lastSpoken] = history{prev1: idx, prev2: hist.prev1}
			} else {
				called[lastSpoken] = history{prev1: idx, prev2: -1}
			}
		}
		idx += 1
		if idx == 2020 {
			fmt.Println("Day 15 part 1:", lastSpoken)
		}
		if idx == 30000000 {
			fmt.Println("Day 15 part 2:", lastSpoken)
			break
		}
	}
}
