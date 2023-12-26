package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("day1/input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	input := strings.TrimSpace(string(data))
	halfway := len(input) / 2
	sum := 0

	for i := 0; i < len(input); i++ {
		next := (i + halfway) % len(input)
		if input[i] == input[next] {
			sum += int(input[i] - '0')
		}
	}

	fmt.Println(sum)
}
