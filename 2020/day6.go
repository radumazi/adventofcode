package main

import (
	"bufio"
	"fmt"
	"os"
)

func day6() {
	file, err := os.Open("./day6_input.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var passportText = ""
	var answers = make(map[string]int)
	var total = 0
	var noLines = 0
	for scanner.Scan() {
		noLines++
		line := scanner.Text()
		passportText += line
		if line == "" {
			for _, letter := range passportText {
				if string(letter) != "" {
					answers[string(letter)]++
				}
			}
			for _, value := range answers {
				if value == noLines-1 {
					total++
				}
			}
			answers = make(map[string]int)
			passportText = ""
			noLines = 0
		}
	}
	fmt.Println(total)
}
