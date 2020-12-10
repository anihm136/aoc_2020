package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	testInput, _ := os.Open("test.txt")
	s := bufio.NewScanner(testInput)
	re := regexp.MustCompile(`(?P<key>[^:]+):\S+ ?`)
	var fieldmap = map[string]int{
		"byr": 0,
		"iyr": 1,
		"eyr": 2,
		"hgt": 3,
		"hcl": 4,
		"ecl": 5,
		"pid": 6,
		"cid": 7,
	}
	valid := 0
	chk := [...]bool{false, false, false, false, false, false, false, true}
	var allFields bool = true
	for s.Scan() {
		if s.Text() == "" {
			for _, val := range chk {
				allFields = allFields && val
			}
			if allFields {
				valid += 1
			}
			allFields = true
			chk = [...]bool{false, false, false, false, false, false, false, true}
			continue
		}
		for _, val := range re.FindAllStringSubmatch(s.Text(), -1) {
			chk[fieldmap[val[1]]] = true
		}
	}
	for _, val := range chk {
		allFields = allFields && val
	}
	if allFields {
		valid += 1
	}
	fmt.Println(valid)
}
