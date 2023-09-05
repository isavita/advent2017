package main

import (
	"fmt"
	"math"
)

func main() {
	target := 325489
	
	// Step 1: Find the square
	sideLength := int(math.Ceil(math.Sqrt(float64(target))))
	if sideLength % 2 == 0 {
		sideLength++
	}
	
	// Step 2: Calculate side length (already done)
	
	// Step 3: Find distance to the nearest middle point
	maxValue := sideLength * sideLength
	stepsFromEdge := (sideLength - 1) / 2
	var distanceToMiddle int
	
	for i := 0; i < 4; i++ {
		middlePoint := maxValue - stepsFromEdge - (sideLength-1) * i
		distance := int(math.Abs(float64(target - middlePoint)))
		if distance < distanceToMiddle || i == 0 {
			distanceToMiddle = distance
		}
	}
	
	// Step 4: Calculate Manhattan Distance
	manhattanDistance := stepsFromEdge + distanceToMiddle
	
	fmt.Println("The Manhattan Distance is:", manhattanDistance)
}

