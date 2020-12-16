package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetIndex(cmd string) int {
	index := ""
	for _, val := range cmd {
		if val == ']' {
			break
		}
		if val < '0' || val > '9' {
			continue
		}
		index += string(val)
	}
	val, _ := strconv.Atoi(index)
	return val
}

func main() {
	f, _ := os.Open("test.txt")
	defer f.Close()

	s := bufio.NewScanner(f)
	var mask1, mask2 int64
	mem := map[int]int64{}
	for s.Scan() {
		sp := strings.Split(s.Text(), " = ")
		cmd := sp[0]
		val := sp[1]
		if cmd == "mask" {
			mask1_s := strings.Replace(val, "1", "0", -1)
			mask1_s = strings.Replace(mask1_s, "X", "1", -1)
			mask2_s := strings.Replace(val, "X", "0", -1)
			mask1, _ = strconv.ParseInt(mask1_s, 2, 0)
			mask2, _ = strconv.ParseInt(mask2_s, 2, 0)
		} else {
			var valInt int64
			valInt, _ = strconv.ParseInt(val, 10, 0)
			valInt = valInt & mask1
			valInt = valInt + mask2
			idx := GetIndex(cmd)
			mem[idx] = valInt
		}
	}

	var sum int64
	for _, val := range mem {
		sum += val
	}
	fmt.Println("Day 14 part 1:", sum)
}
