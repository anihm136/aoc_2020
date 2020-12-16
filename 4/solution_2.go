package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func validate(key string, value string) bool {
	// byr (Birth Year) - four digits; at least 1920 and at most 2002.
	// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
	// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
	// hgt (Height) - a number followed by either cm or in:
	//     If cm, the number must be at least 150 and at most 193.
	//     If in, the number must be at least 59 and at most 76.
	// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
	// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
	// pid (Passport ID) - a nine-digit number, including leading zeroes.
	// cid (Country ID) - ignored, missing or not.
	switch key {
	case "byr", "iyr", "eyr":
		value, ok := strconv.Atoi(value)
		if ok != nil {
			return false
		}
		switch key {
		case "byr":
			if value < 1920 || value > 2002 {
				return false
			}
		case "iyr":
			if value < 2010 || value > 2020 {
				return false
			}
		case "eyr":
			if value < 2020 || value > 2030 {
				return false
			}
		}
	case "hgt":
		match, err := regexp.Match(`^\d{2,3}(cm)|(in)$`, []byte(value))
		if err != nil || !match {
			return false
		}
		suffix := value[len(value)-2:]
		intValue, _ := strconv.Atoi(value[:len(value)-2])
		if (suffix == "cm" && (intValue < 150 || intValue > 193)) || (suffix == "in" && (intValue < 59 || intValue > 76)) {
			return false
		}
	case "ecl":
		validVals := [...]string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
		for _, v := range validVals {
			if value == v {
				return true
			}
		}
		return false
	case "pid":
		match, err := regexp.Match(`^\d{9}$`, []byte(value))
		if err != nil || !match {
			return false
		}
	case "hcl":
		match, err := regexp.Match(`^#[0-9a-f]{6}$`, []byte(value))
		if err != nil || !match {
			return false
		}
	}
	return true
}

func main() {
	testInput, _ := os.Open("test.txt")
	s := bufio.NewScanner(testInput)
	re := regexp.MustCompile(`([^:]+):(\S+) ?`)
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
			if validate(val[1], val[2]) {
				chk[fieldmap[val[1]]] = true
			}
		}
	}
	for _, val := range chk {
		allFields = allFields && val
	}
	if allFields {
		valid += 1
	}
	fmt.Println("Day 4 part 2:", valid)
}
