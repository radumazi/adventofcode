package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"errors"
)

func day8() {
	file, err := os.Open("./day8_input.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var instructions = make([]map[string]int64, 1000)
	var i int = 0
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		value, _ := strconv.ParseInt(parts[1], 10, 0)
		instructions[i] = map[string]int64{parts[0]: value}
		i++
	}
	//part 1
	fmt.Println(detectLoop(instructions))
	//part 2
	fmt.Println(fixLoop(instructions))
}

func detectLoop(instructions []map[string]int64) (int64, error) {
	var acc, i int64
	var ranInstructions = make(map[int64]bool)
	for i = 0; i < int64(len(instructions)); i++ {
		if _, ok := ranInstructions[i]; ok {
			return acc, errors.New("Loop detected")
		}
		ranInstructions[i] = true
		for instruction, value := range instructions[i] {
			switch instruction {
			case "acc":
				acc += value
				break
			case "jmp":
				i += value - 1
				break
			default:
				break
			}
		}
	}

	return acc, nil
}

func fixLoop(instructions []map[string]int64) int64 {
	var i int64
	for i = 0; i < int64(len(instructions)); i++ {
		for instruction, value := range instructions[i] {
			if (instruction == "jmp") {
				instructions[i] = map[string]int64{"nop": value}
				acc, error := detectLoop(instructions)	
				if (error != nil) {
					instructions[i] = map[string]int64{"jmp": value}
				} else {
					return acc
				}
			}
			if (instruction == "nop") {
				instructions[i] = map[string]int64{"jmp": value}
				acc, error := detectLoop(instructions)	
				if (error != nil) {
					instructions[i] = map[string]int64{"nop": value}
				} else {
					return acc
				}
			}
		}
	}

	return 0
}
