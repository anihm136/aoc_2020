package main

import (
	"fmt"
)

func main() {
	var mapline string
	maxwidth := 31
	curIndex := [5]int{1, 3, 5, 7, 1}
	incIndex := [5]int{1, 3, 5, 7, 1}
	trees := [5]int{0, 0, 0, 0, 0}
	for i := 0; i < 323; i++ {
		fmt.Scanln(&mapline)
		if i == 0 {
			continue
		}
		for arridx, index := range curIndex[:4] {
			if mapline[index] == byte('#') {
				trees[arridx] += 1
			}
		}

		for arridx, index := range incIndex[:4] {
			curIndex[arridx] = (curIndex[arridx] + index) % maxwidth
		}
		if i%2 == 0 {
			if mapline[curIndex[4]] == byte('#') {
				trees[4] += 1
			}
			curIndex[4] = (curIndex[4] + incIndex[4]) % maxwidth
		}
	}
	ans := 1
	for _, val := range trees {
		ans *= val
	}
	fmt.Println(ans)
}
