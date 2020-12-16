package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Move(command string, direction *rune, ns *int, ew *int) {
	newDirection := command[0]
	amount, _ := strconv.Atoi(command[1:])
	directions := map[int]rune{0: 'N', 1: 'E', 2: 'S', 3: 'W'}    // Use these to find the new
	revDirections := map[rune]int{'N': 0, 'E': 1, 'S': 2, 'W': 3} // direction of the ship
	switch newDirection {
	case 'N':
		*ns += amount
	case 'S':
		*ns -= amount
	case 'E':
		*ew += amount
	case 'W':
		*ew -= amount
	case 'R':
		curIndex := revDirections[*direction]
		curIndex = (curIndex + (amount / 90)) % 4
		*direction = directions[curIndex]
	case 'L':
		curIndex := revDirections[*direction]
		curIndex = curIndex - (amount / 90)
		if curIndex < 0 {
			curIndex += 4
		}
		*direction = directions[curIndex]
	case 'F':
		switch *direction {
		case 'N':
			*ns += amount
		case 'S':
			*ns -= amount
		case 'E':
			*ew += amount
		case 'W':
			*ew -= amount
		}
	}
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	f, _ := os.Open("test.txt")

	s := bufio.NewScanner(f)

	direction := 'E'
	ns := 0
	ew := 0

	for s.Scan() {
		Move(s.Text(), &direction, &ns, &ew)
	}

	fmt.Println("Day 12 part 1:", Abs(ns)+Abs(ew))
}
