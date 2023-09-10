package main

import (
	"fmt"
)

func main() {
	steps := 343
	buffer := []int{0}
	currentPos := 0

	for i := 1; i <= 2017; i++ {
		currentPos = (currentPos + steps) % len(buffer)
		// Insert new element after the current position
		buffer = append(buffer[:currentPos+1], buffer[currentPos:]...)
		buffer[currentPos+1] = i
		// Update the current position
		currentPos++
	}

	// Find the value immediately after the last inserted value (2017)
	for i, val := range buffer {
		if val == 2017 {
			fmt.Printf("The value after 2017 is: %d\n", buffer[(i+1)%len(buffer)])
			break
		}
	}
}
