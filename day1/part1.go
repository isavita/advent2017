package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	
	input := strings.TrimSpace(string(data))
	sum := 0

	for i := 0; i < len(input); i++ {
		next := (i + 1) % len(input)
		if input[i] == input[next] {
			sum += int(input[i] - '0')
		}
	}

	fmt.Println("The captcha solution is:", sum)
}
