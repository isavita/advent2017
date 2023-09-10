package main

import (
	"fmt"
)

func main() {
	steps := 343
	currentPos := 0
	valueAfterZero := 0

	for i := 1; i <= 50000000; i++ {
		// Calculate the position where the new element will be inserted
		currentPos = (currentPos + steps) % i
		// Check if the new element will be inserted immediately after 0
		if currentPos == 0 {
			valueAfterZero = i
		}
		// Update the current position for the next iteration
		currentPos++
	}

	fmt.Printf("The value after 0 is: %d\n", valueAfterZero)
}
