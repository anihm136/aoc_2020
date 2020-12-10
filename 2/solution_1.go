package main

import (
	"fmt"
)

func main() {
	validCount := 0
	var temp string
	var password string
	var char byte
	var min int
	var max int
	for i := 0; i < 1000; i++ {
		fmt.Scanf("%d%1s%d%c%c%1s%s", &min, &temp, &max, &char, &char, &temp, &password)
		charCount := 0
		for _, e := range password {
			if e == rune(char) {
				charCount += 1
			}
		}
		if charCount >= min && charCount <= max {
			validCount += 1
		}
		// fmt.Printf("%d %d %c %s\n", min, max, char, password)
	}
	fmt.Println(validCount)
}
