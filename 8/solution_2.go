package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var lines []string
var switchIdx int

func FindLoop(insIdx int, hasRun []bool, didSwitch bool, acc int) bool {
	if insIdx == len(lines) {
		fmt.Println("Reached end! Accumulator is", acc)
		return true
	}
	if hasRun[insIdx] {
		return false
	}

	hasRun[insIdx] = true
	parts := strings.Split(lines[insIdx], " ")
	amt, _ := strconv.Atoi(parts[1])
	var b1,b2 bool

	switch parts[0] {
	case "acc":
		b1 = FindLoop(insIdx+1, hasRun, didSwitch, acc+amt)
		if !b1 { // Have taken wrong path, backtrack
			hasRun[insIdx] = false
		}
		return b1
	case "nop", "jmp":
		var incIdxOrig, incIdxChange int
		if parts[0] == "jmp" {
			incIdxOrig = amt
			incIdxChange = 1
		} else {
			incIdxOrig = 1
			incIdxChange = amt
		}
		b1 = FindLoop(insIdx+incIdxOrig, hasRun, didSwitch, acc) // First try keeping instruction the same
		if !b1 && !didSwitch {
			b2 = FindLoop(insIdx+incIdxChange, hasRun, true, acc) // change instruction
			if b2 { // This is the instruction to switch
				switchIdx = insIdx
			} else { // Have taken wrong path, backtrack
				hasRun[insIdx] = false
			}
			return b2
		}
		return b1
	}
	fmt.Println("Something's wrong...")
	return false // should be unreachable
}

func main() {
	f, _ := os.Open("test.txt")
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	hasRun := make([]bool, len(lines))
	_ = FindLoop(0, hasRun, false, 0)
}
