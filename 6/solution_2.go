package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	testInput, _ := os.Open("test.txt")
	s := bufio.NewScanner(testInput)
	m := make(map[rune]int)
	sum := 0
	lines := 0

	for s.Scan() {
		if s.Text() == "" {
			for _, val := range m {
				if val == lines {
					sum += 1
				}
			}
			m = make(map[rune]int)
			lines = 0
			continue
		}
		for _, val := range s.Text() {
			m[val] += 1
		}
		lines += 1
	}
	for _, val := range m {
		if val == lines {
			sum += 1
		}
	}
	fmt.Println("Day 6 part 2:", sum)
}
