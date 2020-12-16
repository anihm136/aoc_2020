package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("test.txt")
	defer f.Close()

	s := bufio.NewScanner(f)
	_ = s.Scan()
	earliest, _ := strconv.Atoi(s.Text()) // Earliest possible departure time
	_ = s.Scan()
	dep := strings.Split(s.Text(), ",") // All bus departure times

	minDiff := earliest // Smallest wait time
	choice := -1        // Index of bus with smallest wait time
	for _, val := range dep {
		if val != "x" {
			curDep, _ := strconv.Atoi(val)
			// How long do we need to wait for this bus after the threshold?
			diff := curDep - (earliest % curDep)
			if diff < minDiff {
				minDiff = diff
				choice = curDep
			}
		}
	}
	fmt.Println("Day 13 part 1:", choice*minDiff)
}
