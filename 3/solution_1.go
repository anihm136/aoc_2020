package main

import (
	"fmt"
)

func main() {
	var mapline string
	maxwidth := 31
	index := 3
	trees := 0
	for i := 0; i < 323; i++ {
		fmt.Scanln(&mapline)
		if i==0 {
			continue
		}
		if mapline[index]==byte('#') {
			trees += 1
		}
		fmt.Println(index)
		index = (index+3)%maxwidth
	}
	fmt.Println(trees)
}
