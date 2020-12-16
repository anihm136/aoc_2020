package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Power(base, exponent int) int {
	if exponent != 0 {
		return (base * Power(base, exponent-1))
	} else {
		return 1
	}
}

func GenerateIndices(bitmask string, indexInt int64) []string {
	index := fmt.Sprintf("%36b", indexInt)
	index = strings.ReplaceAll(index, " ", "0")

	numX := strings.Count(bitmask, "X")
	if numX == 0 {
		return []string{index}
	}

	ret := make([]string, 0, Power(2, numX))
	xIndices := make([]int, numX+1)
	xIndices[0] = -1
	for i := 0; i < numX; i++ {
		xIndices[i+1] = strings.Index(bitmask[xIndices[i]+1:], "X") + xIndices[i] + 1
	}
	xIndices = xIndices[1:]

	for i := range bitmask {
		if bitmask[i] == '1' {
			index = index[:i] + "1" + index[i+1:]
		}
	}
	for _, val := range xIndices {
		index = index[:val] + "X" + index[val+1:]
	}

	var rec func(mask string, perm string)
	rec = func(mask string, perm string) {
		if len(perm) == len(xIndices) {
			for i, v := range xIndices {
				mask = mask[:v] + string(perm[i]) + mask[v+1:]
			}
			ret = append(ret, mask)
			return
		}
		perm += string('0')
		rec(mask, perm)
		perm = perm[:len(perm)-1] + string('1')
		rec(mask, perm)
	}

	rec(index, "")
	return ret
}

func GetIndex(cmd string) int64 {
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
	val, _ := strconv.ParseInt(index, 10, 0)
	return val
}

func ConvertToInt(maskStr []string) []int64 {
	ret := []int64{}
	for _, val := range maskStr {
		val, _ := strconv.ParseInt(val, 2, 0)
		ret = append(ret, val)
	}
	return ret
}

func main() {
	f, _ := os.Open("test.txt")
	defer f.Close()

	s := bufio.NewScanner(f)
	var mask string
	mem := map[int64]int{}
	for s.Scan() {
		sp := strings.Split(s.Text(), " = ")
		cmd := sp[0]
		val := sp[1]
		if cmd == "mask" {
			mask = val
		} else {
			var valInt int
			valInt, _ = strconv.Atoi(val)
			idx := GetIndex(cmd)
			for _, i := range ConvertToInt(GenerateIndices(mask, idx)) {
				mem[i] = valInt
			}
		}
	}

	var sum int
	for _, val := range mem {
		sum += val
	}
	fmt.Println("Day 14 part 2:", sum)
}
