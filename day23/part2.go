package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func partOne() {
	// Read instructions from input file
	file, err := os.Open("day23/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var instructions []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		instructions = append(instructions, scanner.Text())
	}

	// Initialize registers and mul counter
	registers := make(map[string]int)
	mulCounter := 0

	// Execute the instructions
	for i := 0; i < len(instructions); {
		parts := strings.Fields(instructions[i])
		cmd, x, y := parts[0], parts[1], parts[2]
		valY, _ := strconv.Atoi(y)

		// Resolve value of y if it's a register
		if y >= "a" && y <= "z" {
			valY = registers[y]
		}

		switch cmd {
		case "set":
			registers[x] = valY
		case "sub":
			registers[x] -= valY
		case "mul":
			registers[x] *= valY
			mulCounter++
		case "jnz":
			valX, _ := strconv.Atoi(x)
			if x >= "a" && x <= "z" {
				valX = registers[x]
			}
			if valX != 0 {
				i += valY
				continue
			}
		}
		i++
	}

	fmt.Println("Part One: Number of times mul is invoked:", mulCounter)
}

func partTwo() {
	b := 57*100 + 100000 // Initial value computed for register b
	c := b + 17000       // Initial value computed for register c
	h := 0               // Initialize register h

	for x := b; x <= c; x += 17 { // Simulate the loop from b to c with step 17
		if !isPrime(x) {
			h++
		}
	}

	fmt.Println("Part Two: Value left in register h:", h)
}

func main() {
	partOne()
	partTwo()
}
