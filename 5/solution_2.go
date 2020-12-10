package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	f, _ := os.Open("test.txt")
	s := bufio.NewScanner(f)
	var minRow, maxRow, minCol, maxCol, curSeat int
	seats := make([]int, 0)
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
		seats = append(seats, curSeat)
	}
	sort.Ints(seats)
	for idx, val := range seats {
		if seats[idx+1] == val+2 {
			fmt.Println(val + 1)
			break
		}
	}

}
