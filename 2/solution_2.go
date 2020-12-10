package main

import (
	"fmt"
)

func main() {
	validCount := 0
	var temp string
	var password string
	var char byte
	var pos1 int
	var pos2 int
	for i := 0; i < 1000; i++ {
		fmt.Scanf("%d%1s%d%c%c%1s%s", &pos1, &temp, &pos2, &char, &char, &temp, &password)
		var pos1Match byte
		var pos2Match byte
		if password[pos1-1] == char {
			pos1Match = 1
		}
		if password[pos2-1] == char {
			pos2Match = 1
		}
		validCount += int(pos1Match ^ pos2Match)
		// fmt.Printf("%c %c %c %d\n", password[pos1-1], password[pos2-1], char, pos1Match ^ pos2Match)
	}
	fmt.Println(validCount)
}
