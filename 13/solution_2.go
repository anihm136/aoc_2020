package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Convert a list of strings to a map of divisors and
// the remainders they should result in
func ArrToMap(arr []string) map[int]int {
	new := make(map[int]int)
	for i, val := range arr {
		temp, ok := strconv.Atoi(val)
		if ok == nil {
			new[temp] = temp - i
		}
	}
	temp, _ := strconv.Atoi(arr[0])
	new[temp] = 0
	return new
}

// Given a map of divisors to the remainders they should
// result in, return the dividend that satisfies them
func ChineseRemainder(m map[int]int) int {
	prod := 1
	z := make(map[int]int)

	// Product of all divisors
	for key := range m {
		prod *= key
	}

	// z_i is the modular multiplicative inverse of y_i
	// y_i = prod/key
	for key := range m {
		mult := 1
		for (prod/key*mult)%key != 1 {
			mult += 1
		}
		z[key] = mult
	}

	ans := 0
	// ans is sum of (remainder_i*y_i*z_i)
	for key, val := range m {
		ans += val * (prod / key) * z[key]
	}

	return ans % prod
}

func main() {
	f, _ := os.Open("test.txt")
	defer f.Close()

	s := bufio.NewScanner(f)
	_ = s.Scan()
	_ = s.Scan()
	dep := strings.Split(s.Text(), ",")
	depInt := ArrToMap(dep)

	fmt.Println("Day 13 part 2:", (ChineseRemainder(depInt)))
}
