package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var jolts []int
var fact map[int]int

func Factorial(num int) int {
	if num <= 1 {
		return 1
	}
	if _, ok := fact[num]; !ok {
		fact[num] = num * Factorial(num-1)
	}
	return fact[num]
}

func Combination(n int, r int) int {
	return Factorial(n) / Factorial(r) / Factorial(n-r)
}

func main() {
	f, _ := os.Open("test.txt")
	defer f.Close()

	fact = make(map[int]int)
	s := bufio.NewScanner(f)
	var temp int
	jolts = append(jolts, 0)
	for s.Scan() {
		temp, _ = strconv.Atoi(s.Text())
		jolts = append(jolts, temp)
	}

	sort.Ints(jolts)
	jolts = append(jolts, jolts[len(jolts)-1]+3)

	/*
	* Essential jolts are those which have a difference of three on any one
	* side(or both).  These cannot be removed, since it would result in a gap of
	* more than three
	 */
	essential := make([]bool, len(jolts))

	essential[0] = true
	essential[len(essential)-1] = jolts[len(jolts)-1]-jolts[len(jolts)-2] == 3
	for i := 1; i < len(jolts)-1; i++ {
		if jolts[i]-jolts[i-1] == 3 || jolts[i+1]-jolts[i] == 3 {
			essential[i] = true
		}
	}

	total := 1
	lastEssential := 0 // Last index of essential jolt. The first jolt(index 0) is always essential
	for i := 1; i < len(essential); i++ {
		if essential[i] {
			tempTotal := 0
			/*
			* Non-consecutive essential jolts have non-essential jolts between, which
			* may or may not be present in a valid solution
			 */
			if i-lastEssential > 1 {
				/*
				* The minimum number of non-essential jolts between two non-consecutive
				* essential jolts depends on the difference between them. Up to a
				* difference of 3, none are required. Between 4 and 6, one is required
				* and so on
				 */
				min := (jolts[i] - jolts[lastEssential] - 1) / 3
				for j := min; j < i-lastEssential; j++ {
					tempTotal += Combination(i-lastEssential-1, j) // All possible valid combinations of non-essential jolts
				}
				total *= tempTotal // Each set of non-essential jolts is independent of each other, and are multiplied to get the total number of valid solutions
			}
			lastEssential = i
		}
	}

	fmt.Println("Day 10 part 2:", (total))
}
