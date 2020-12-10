package main

import (
	"fmt"
)

func main() {
	set1 := make(map[int]struct{}) // Store values
	set2 := make(map[int]struct{}) // Store potential values for completion
	var exists = struct{}{}
	var temp int
	for i := 0; i < 200; i++ {
		fmt.Scanln(&temp)
		for e := range set1 {
			if _, ok := set2[e+temp]; ok {
				fmt.Println(temp * e * (2020 - temp - e))
				break
			}
		}
		set1[temp] = exists
		set2[2020-temp] = exists
	}
}
