package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Rotate waypoint around ship by given amount
func RotateWaypoint(x *int, y *int, amt int, direction byte) {
	if amt == 180 {
		*x, *y = -*x, -*y
	} else if (direction == 'R' && amt == 90) || (direction == 'L' && amt == 270) { // x = y, y = -x
		*x, *y = *y, -*x
	} else { // x = -y, y = x
		*x, *y = -*y, *x
	}
}

func Move(command string, wp_x *int, wp_y *int, ns *int, ew *int) {
	instruction := command[0]
	amount, _ := strconv.Atoi(command[1:])
	switch instruction {
	case 'R', 'L':
		RotateWaypoint(wp_x, wp_y, amount, instruction)
	case 'F':
		*ns += *wp_y * amount
		*ew += *wp_x * amount
	case 'N':
		*wp_y += amount
	case 'S':
		*wp_y -= amount
	case 'E':
		*wp_x += amount
	case 'W':
		*wp_x -= amount
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

	wp_x := 10
	wp_y := 1
	ns := 0
	ew := 0

	for s.Scan() {
		Move(s.Text(), &wp_x, &wp_y, &ns, &ew)
	}

	fmt.Println("Manhattan distance:", Abs(ns)+Abs(ew))
}
