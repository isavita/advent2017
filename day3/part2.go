package main

import (
	"fmt"
)

func main() {
	target := 325489
	grid := make(map[[2]int]int)
	grid[[2]int{0, 0}] = 1
	
	x, y := 0, 0
	dx, dy := 0, -1
	
	for {
		// Change direction when reaching a corner
		if x == y || (x < 0 && x == -y) || (x > 0 && x == 1-y) {
			dx, dy = -dy, dx
		}
		
		// Move to the next square
		x += dx
		y += dy
		
		// Calculate value for the current square
		value := 0
		for dx := -1; dx <= 1; dx++ {
			for dy := -1; dy <= 1; dy++ {
				value += grid[[2]int{x + dx, y + dy}]
			}
		}
		grid[[2]int{x, y}] = value
		
		// Check if value is greater than the target
		if value > target {
			fmt.Println("The first value written larger than the puzzle input is:", value)
			break
		}
	}
}
