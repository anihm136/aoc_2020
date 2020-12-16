package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Check if array contains a value (python's 'in' is so much easier)
func Contains(arr []int, val int) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

// Find a valid assignment of indices to keywords
func FindPath(curpath []int, matrix map[string][]bool, mapping []string, index int) ([]int, bool) {
	if len(curpath) == len(mapping) { // All keywords have a unique index
		return curpath, true
	}

	curNode := mapping[index] // the keyword we are trying to find an index for
	var found bool
	for idx, valid := range matrix[curNode] {
		if valid && !Contains(curpath, idx) { // index is a valid assignment for curNode, and does not appear in the path
			curpath, found = FindPath(append(curpath, idx), matrix, mapping, index+1)
			if found {
				return curpath, true
			}
		}
	}
	return curpath, false // should not reach here if a valid solution exists
}

// Convert a slice of integer strings to ints
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

	rules := make(map[string][][]int)
	s := bufio.NewScanner(f)
	for s.Scan() {
		if s.Text() == "" { // break at first blank line
			break
		}
		line := strings.Split(s.Text(), ": ")
		k, v := line[0], line[1]
		/*
		 * Each rule is stored as a map from a string (keyword) to an array of pairs
		 * of start and end values Does a better structure exist for pairs? (like
		 * std::pair)
		 */
		for _, r := range strings.Split(v, " or ") {
			x := strings.Split(r, "-")
			start, _ := strconv.Atoi(x[0])
			end, _ := strconv.Atoi(x[1])
			rules[k] = append(rules[k], []int{start, end})
		}
	}

	s.Scan() // skip unnecessary lines
	s.Scan()
	myticket := ConvertToInt(strings.Split(s.Text(), ",")) // values in own ticket for later use
	s.Scan()
	s.Scan()

	// Store only the valid tickets (as per part 1)
	valid := make([][]int, 0)
	for s.Scan() {
		vals := strings.Split(s.Text(), ",")
		valInt := ConvertToInt(vals)
		okArr := true // denotes whether the entire ticket is valid
		for _, val := range valInt {
			ok := false // denotes whether the current value within a ticket is valid
			for _, r := range rules {
				for _, currange := range r {
					if val >= currange[0] && val <= currange[1] { // value is valid for some rule
						ok = true
						break
					}
				}
				if ok {
					break
				}
			}
			okArr = okArr && ok
		}
		if okArr { // all values in the ticket are valid
			valid = append(valid, valInt)
		}
	}

	/* Find out which indices are valid for each keyword for each ticket
	 * validRanges maps each keyword to a 2d bool slice, where rows represent
	 * tickets and columns represent indices in the ticket A cell is true if the
	 * index(column) is valid for that ticket(row) for that keyword(key of the
	 * map)
	 */
	validRanges := make(map[string][][]bool)
	for k := range rules {
		validRanges[k] = make([][]bool, len(valid))
		for i := range valid {
			validRanges[k][i] = make([]bool, len(valid[0]))
		}
	}

	// 4 loop (geddit?)
	for idx, val := range valid {
		for i, v := range val {
			for k, r := range rules {
				for _, cur := range r {
					if v >= cur[0] && v <= cur[1] {
						validRanges[k][idx][i] = true
						break
					}
				}
			}
		}
	}

	/*
	* Combine the data from all the tickets to get a mapping from keywords to
	* valid indices If an index is valid for a keyword in all the tickets, then
	* it is valid for the keywords
	 */
	validIndices := make(map[string][]bool)
	for k := range rules {
		validIndices[k] = make([]bool, len(valid[0]))
	}

	var ok bool
	for key := range validIndices {
		for i := 0; i < len(valid[0]); i++ {
			ok = true
			for j := 0; j < len(valid); j++ {
				ok = ok && validRanges[key][j][i]
			}
			if ok {
				validIndices[key][i] = true
			}
		}
	}

	/* Create a slice so that order of accessing the keywords is constant (seems
	 * hacky, maybe there's a better way)
	 */
	mapping := make([]string, 0)
	for k := range validIndices {
		mapping = append(mapping, k)
	}

	// Count the number of valid indices for the keyword
	ValidCount := func(s string) int {
		ret := 0
		for _, val := range validIndices[s] {
			if val {
				ret += 1
			}
		}
		return ret
	}

	/* Sort the keywords in ascending order of number of valid positions
	 * Looks like there's exactly one ordering when sorted (may not always hold
	 * true)
	 */
	sort.Slice(mapping, func(i, j int) bool {
		return ValidCount(mapping[i]) < ValidCount(mapping[j])
	})

	// Find a valid assignment of indices to keywords
	assignment, _ := FindPath(make([]int, 0), validIndices, mapping, 0)
	prod := 1
	for i, v := range mapping {
		if strings.HasPrefix(v, "departure") {
			prod *= myticket[assignment[i]]
		}
	}
	fmt.Println("Day 16 part 2:", prod)
}
