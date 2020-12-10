package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"strconv"
)

var contains map[string]map[string]int

func RecSearch(search string) int {
	ans := 0
	checkNext := contains[search]
	for k, v := range checkNext {
		ans += v
		if k != "no other" {
			ans += v * RecSearch(k)
		}
	}
	return ans
}

func main() {
	contains = make(map[string]map[string]int)

	testInput, _ := os.Open("test.txt")
	s := bufio.NewScanner(testInput)
	var a, b, temp []string

	re := regexp.MustCompile(`(?P<Num>\d+ )?(?P<Type>\w+ \w+).?`)
	for s.Scan() {
		a = strings.Split(s.Text(), " bags contain ")
		b = strings.Split(a[1], ", ")
		if _,ok := contains[a[0]]; !ok {
			contains[a[0]] = make(map[string]int)
		}
		for _, val := range b {
			temp = re.FindStringSubmatch(val)
			if temp[1] != "no other" {
				intval, _ := strconv.Atoi(strings.TrimRight(temp[1], " "))
				contains[a[0]][temp[2]] = intval
			} else {
				contains[a[0]][temp[1]] = 0
			}
		}
	}
	fmt.Println(RecSearch("shiny gold"))
}
