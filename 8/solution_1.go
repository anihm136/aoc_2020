package main

import (
	"fmt"
	"strings"
	"strconv"
	"bufio"
	"os"
)

func main() {
	f, _ := os.Open("test.txt")
	defer f.Close()

	var lines []string
	s := bufio.NewScanner(f)
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	hasRun := make([]bool, len(lines))
	var parts []string
	var amt int
	insIdx := 0
	acc := 0
	for {
		if hasRun[insIdx] {
			break
		}
		hasRun[insIdx] = true
		parts = strings.Split(lines[insIdx], " ")
		switch parts[0] {
		case "nop":
			insIdx += 1
		case "jmp":
			amt, _ = strconv.Atoi(parts[1])
			insIdx += amt
		case "acc":
			amt, _ = strconv.Atoi(parts[1])
			acc += amt
			insIdx += 1
		}
	}
	fmt.Println(acc)
}
