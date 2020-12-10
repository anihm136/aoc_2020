package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	testInput, _ := os.Open("test.txt")
	s := bufio.NewScanner(testInput)
	set := make(map[rune]struct{})
	var exists = struct{}{}
	sum := 0

	for s.Scan() {
		if s.Text() == "" {
			sum += len(set)
			set = make(map[rune]struct{})
			continue
		}
		for _, val := range s.Text() {
			set[val] = exists
		}
	}
	sum += len(set)
	fmt.Println(sum)
}
