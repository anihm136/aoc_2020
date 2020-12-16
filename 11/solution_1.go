package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("test.txt")
	defer f.Close()

	s := bufio.NewScanner(f)
	var input [][]byte

	for s.Scan() {
		input = append(input, []byte(s.Text()))
	}

	changed := true
	chk := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	for changed {
		changed = false
		change := make([][]byte, len(input))
		for i := 0; i < len(input); i++ {
			change[i] = make([]byte, len(input[0]))
		}

		for i := 0; i < len(input); i++ {
			for j := 0; j < len(input[0]); j++ {
				handled := false
				if input[i][j] == 'L' {
					swap := true
					for _, val := range chk {
						if i+val[0] >= 0 && j+val[1] >= 0 && i+val[0] < len(input) && j+val[1] < len(input[0]) {
							if input[i+val[0]][j+val[1]] == '#' {
								swap = false
								break
							}

						}
					}
					if swap {
						change[i][j] = '#'
						handled = true
						changed = true
					}
				} else if input[i][j] == '#' {
					numOccupied := 0
					for _, val := range chk {
						if i+val[0] >= 0 && j+val[1] >= 0 && i+val[0] < len(input) && j+val[1] < len(input[0]) {
							if input[i+val[0]][j+val[1]] == '#' {
								numOccupied += 1
							}

						}
					}
					if numOccupied >= 4 {
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
	fmt.Println("Day 11 part 1:", numOccupied)
}
