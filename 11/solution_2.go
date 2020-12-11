package main

import (
	"bufio"
	"fmt"
	"os"
)

func FindOccupied(arr [][]byte, row int, col int) int {
	numOccupied := 0
	curRow := row
	curCol := col
	breakCondition := func() bool {
		return arr[curRow][curCol] == 'L'
	}
	incCondition := func() bool {
		return arr[curRow][curCol] == '#'
	}
	for {
		curRow -= 1
		if curRow < 0 || breakCondition() {
			break
		}
		if incCondition(){
			numOccupied += 1
			break
		}
	}
	curRow = row
	for {
		curRow += 1
		if curRow >= len(arr) || breakCondition() {
			break
		}
		if incCondition() {
			numOccupied += 1
			break
		}
	}
	curRow = row
	for {
		curCol -= 1
		if curCol < 0 || breakCondition() {
			break
		}
		if incCondition() {
			numOccupied += 1
			break
		}
	}
	curCol = col
	for {
		curCol += 1
		if curCol >= len(arr[0]) || breakCondition() {
			break
		}
		if incCondition() {
			numOccupied += 1
			break
		}
	}
	curCol = col
	for {
		curRow -= 1
		curCol -= 1
		if curRow < 0 || curCol < 0 || breakCondition() {
			break
		}
		if incCondition() {
			numOccupied += 1
			break
		}
	}
	curRow = row
	curCol = col
	for {
		curRow -= 1
		curCol += 1
		if curRow < 0 || curCol >= len(arr[0]) || breakCondition() {
			break
		}
		if incCondition() {
			numOccupied += 1
			break
		}
	}
	curRow = row
	curCol = col
	for {
		curRow += 1
		curCol -= 1
		if curRow >= len(arr) || curCol < 0 || breakCondition() {
			break
		}
		if incCondition() {
			numOccupied += 1
			break
		}
	}
	curRow = row
	curCol = col
	for {
		curRow += 1
		curCol += 1
		if curRow >= len(arr) || curCol >= len(arr[0]) || breakCondition() {
			break
		}
		if incCondition() {
			numOccupied += 1
			break
		}
	}
	return numOccupied
}

func main() {
	f, _ := os.Open("test.txt")
	defer f.Close()

	s := bufio.NewScanner(f)
	var input [][]byte

	for s.Scan() {
		input = append(input, []byte(s.Text()))
	}

	changed := true
	for changed {
		changed = false
		change := make([][]byte, len(input))
		for i := 0; i < len(input); i++ {
			change[i] = make([]byte, len(input[0]))
		}

		for i := 0; i < len(input); i++ {
			for j := 0; j < len(input[0]); j++ {
				handled := false
				occ := FindOccupied(input, i, j)
				if input[i][j] == 'L' {
					swap := occ == 0
					if swap {
						change[i][j] = '#'
						handled = true
						changed = true
					}
				} else if input[i][j] == '#' {
					if occ >= 5 {
						change[i][j] = 'L'
						handled = true
						changed = true
					}
				}
				if !handled {
					change[i][j] = input[i][j]
				}
			}
		}
		input = change
	}

	numOccupied := 0
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[0]); j++ {
			if input[i][j] == '#' {
				numOccupied += 1
			}
		}
	}
	fmt.Println(numOccupied)
}
