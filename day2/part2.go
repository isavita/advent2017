package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	sum := 0

	for _, line := range lines {
		numsStr := strings.Fields(line)
		var nums []int
		for _, numStr := range numsStr {
			num, _ := strconv.Atoi(numStr)
			nums = append(nums, num)
		}

		for i, num1 := range nums {
			for j, num2 := range nums {
				if i != j && num1 % num2 == 0 {
					sum += num1 / num2
				}
			}
		}
	}

	fmt.Println("The sum of each row's result is:", sum)
}

