package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ConvertToInt(arr []string) []int {
	ret := make([]int, 0)
	var temp int
	for _, val := range arr {
		temp, _ = strconv.Atoi(val)
		ret = append(ret, temp)
	}

	return ret
}

func main() {
	f, _ := os.Open("test.txt")
	defer f.Close()

	ranges := make(map[string][][]int)
	s := bufio.NewScanner(f)
	for s.Scan() {
		if s.Text() == "" {
			break
		}
		line := strings.Split(s.Text(), ": ")
		k, v := line[0], line[1]
		for _, r := range strings.Split(v, " or ") {
			x := strings.Split(r, "-")
			start, _ := strconv.Atoi(x[0])
			end, _ := strconv.Atoi(x[1])
			ranges[k] = append(ranges[k], []int{start, end})
		}
	}

	// Skip useless lines
	s.Scan()
	s.Scan()
	s.Scan()
	s.Scan()

	/*
	* Find errors in tickets and sum the errors
	* We start by assuming that that the field is
	* incorrect. Then we iterate through all the
	* rules and check if the field fits any of them.
	* If it does, the field is valid and we break
	* to next iteration. Else, add to running sum
	* and continue
	 */
	err := 0
	for s.Scan() {
		vals := strings.Split(s.Text(), ",")
		valInt := ConvertToInt(vals)
		for _, val := range valInt {
			ok := false
			for _, r := range ranges {
				for _, currange := range r {
					if val >= currange[0] && val <= currange[1] {
						ok = true
						break
					}
				}
				if ok {
					break
				}
			}
			if !ok {
				err += val
			}
		}
	}

	fmt.Println("Day 16 part 1:", err)
}
