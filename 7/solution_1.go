package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var exists struct{}
var contained map[string][]string
var contains map[string]struct{}

func RecSearch(search string) {
	for _, val := range contained[search] {
		contains[val] = exists
		RecSearch(val)
	}
}

func main() {
	contained = make(map[string][]string)
	contains = make(map[string]struct{})
	exists = struct{}{}

	testInput, _ := os.Open("test.txt")
	s := bufio.NewScanner(testInput)
	var a, b, temp []string

	re := regexp.MustCompile(`(?:\d+ )?(\w+ \w+).?`)
	for s.Scan() {
		a = strings.Split(s.Text(), " bags contain ")
		b = strings.Split(a[1], ", ")
		for _, val := range b {
			temp = re.FindStringSubmatch(val)
			if temp[1] != "no other" {
				if _, ok := contained[temp[1]]; !ok {
					contained[temp[1]] = make([]string, 0, 5)
				}
				contained[temp[1]] = append(contained[temp[1]], a[0])
			}
		}
	}
	RecSearch("shiny gold")
	fmt.Println("Day 7 part 1:", len(contains))
}
