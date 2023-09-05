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
	checksum := 0

	for _, line := range lines {
		nums := strings.Fields(line)
		minVal, _ := strconv.Atoi(nums[0])
		maxVal, _ := strconv.Atoi(nums[0])

		for _, numStr := range nums {
			num, _ := strconv.Atoi(numStr)
			if num < minVal {
				minVal = num
			}
			if num > maxVal {
				maxVal = num
			}
		}

		checksum += (maxVal - minVal)
	}

	fmt.Println("The checksum is:", checksum)
}

