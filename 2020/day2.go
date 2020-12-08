package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func day2() {
	file, err := os.Open("./day2_input.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	r, _ := regexp.Compile("([0-9]+)-([0-9]+) ([a-z]): ([a-z]+)")
	var counter = 0
	for scanner.Scan() {
		line := scanner.Text()
		var match = r.FindStringSubmatch(line)
		var min, _ = strconv.ParseInt(match[1], 10, 0)
		var max, _ = strconv.ParseInt(match[2], 10, 0)
		var char = match[3]
		var password = match[4]
		if isValidPassword(min, max, char, password) {
			counter++
		}
	}

	fmt.Println(counter)

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

// func isValidPassword(min int64, max int64, char string, password string) bool {
// 	var counter int64 = 0
// 	for _, letter := range password {
// 		if string(letter) == char {
// 			counter++
// 		}
// 	}
// 	return counter >= min && counter <= max
// }

func isValidPassword(min int64, max int64, char string, password string) bool {
	var minOcc = string(password[min-1]) == char
	var maxOcc = string(password[max-1]) == char
	return (minOcc || maxOcc) && !(minOcc && maxOcc)
}