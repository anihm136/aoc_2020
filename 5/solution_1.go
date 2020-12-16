package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("test.txt")
	s := bufio.NewScanner(f)
	var minRow, maxRow, minCol, maxCol, curSeat int
	maxSeat := 0
	for s.Scan() {
		minRow = 0
		maxRow = 128
		minCol = 0
		maxCol = 8
		for _, e := range []byte(s.Text()) {
			switch e {
			case 'F':
				maxRow -= (maxRow - minRow) / 2
			case 'B':
				minRow += (maxRow - minRow) / 2
			case 'L':
				maxCol -= (maxCol - minCol) / 2
			case 'R':
				minCol += (maxCol - minCol) / 2
			}
		}
		curSeat = minRow*8 + minCol
		if curSeat > maxSeat {
			maxSeat = curSeat
		}
	}
	fmt.Println("Day 5 part 1:", maxSeat)
}
