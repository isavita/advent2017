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

	passphrases := strings.Split(strings.TrimSpace(string(data)), "\n")
	validCount := 0

	for _, passphrase := range passphrases {
		words := strings.Fields(passphrase)
		wordSet := make(map[string]bool)

		valid := true
		for _, word := range words {
			if wordSet[word] {
				valid = false
				break
			}
			wordSet[word] = true
		}

		if valid {
			validCount++
		}
	}

	fmt.Println("The number of valid passphrases is:", validCount)
}

