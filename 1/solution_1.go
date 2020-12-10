package main

import (
	"fmt"
)

func main() {
	set := make(map[int]struct{}) // Use map with keys as set elements and values as empty struct
	var exists = struct{}{} // Empty struct for map values
	var temp int
	for i := 0; i < 200; i++ {
		fmt.Scanln(&temp)
		if _,ok := set[2020-temp]; ok {
			fmt.Println(temp*(2020-temp))
			break
		}
		set[temp] = exists
	}
}
