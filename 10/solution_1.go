package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	f, _ := os.Open("test.txt")
	defer f.Close()

	s := bufio.NewScanner(f)
	jolts := make([]int, 0, 100)
	var temp int
	for s.Scan() {
		temp, _ = strconv.Atoi(s.Text())
		jolts = append(jolts, temp)
	}

	sort.Ints(jolts)
	prev := 0
	var j1, j3 int
	for i := 0; i < len(jolts); i++ {
		if jolts[i]-prev == 1 {
			j1++
		} else if jolts[i]-prev == 3 {
			j3++
		}
		prev = jolts[i]
	}
	j3++

	fmt.Println("Day 10 part 1:", j1*j3)
}
