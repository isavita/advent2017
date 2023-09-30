package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type State struct {
	write [2]int
	move  [2]int
	next  [2]string
}

var (
	tape         = make(map[int]int)
	cursor       int
	state        string
	stateTrans   = make(map[string]State)
	steps, count int
)

func main() {
	file, _ := os.Open("day25/input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan() // Read the initial state line
	state = strings.Split(scanner.Text(), " ")[3][:1]
	scanner.Scan() // Read the steps line
	stepsStr := strings.Split(scanner.Text(), " ")[5]
	steps, _ = strconv.Atoi(stepsStr)
	for scanner.Scan() {
		parseInput(scanner)
	}
	for i := 0; i < steps; i++ {
		execute()
	}
	for _, v := range tape {
		count += v
	}
	fmt.Println(count)
}

func parseInput(scanner *bufio.Scanner) {
	line := scanner.Text()
	if strings.HasPrefix(line, "In state ") {
		stateName := strings.Split(line, " ")[2][:1]
		var trans State
		for i := 0; i < 2; i++ {
			scanner.Scan() // Read the value condition line
			scanner.Scan() // Read the write action line
			trans.write[i], _ = strconv.Atoi(strings.Split(scanner.Text(), " ")[4][:1])
			scanner.Scan() // Read the move action line
			moveDirection := strings.Split(scanner.Text(), " ")[6]
			if moveDirection == "right." {
				trans.move[i] = 1
			} else {
				trans.move[i] = -1
			}
			scanner.Scan() // Read the next state line
			trans.next[i] = strings.Split(scanner.Text(), " ")[4][:1]
		}
		stateTrans[stateName] = trans
	}
}

func execute() {
	trans := stateTrans[state]
	tapeVal := tape[cursor]
	tape[cursor] = trans.write[tapeVal]
	cursor += trans.move[tapeVal]
	state = trans.next[tapeVal]
}
