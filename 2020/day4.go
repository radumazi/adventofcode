package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func day4() {
	file, err := os.Open("./day4_input.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var passportText = ""
	var validPassports int;
	for scanner.Scan() {
		line := scanner.Text()
		passportText += " "+line
		if line == "" {
			// fmt.Println(passportText)
			if(isValidPassport(passportText)) {
				validPassports++
			}
			passportText = ""
		}
	}
	fmt.Println(validPassports)
}

func isValidPassport(passportText string) bool {
	var match = false;
	match, _ = regexp.MatchString("byr:[^ ]+", passportText)
	if (!match) {
		return false;
	}
	match, _ = regexp.MatchString("iyr:[^ ]+", passportText)
	if (!match) {
		return false;
	}
	match, _ = regexp.MatchString("eyr:[^ ]+", passportText)
	if (!match) {
		return false;
	}
	match, _ = regexp.MatchString("hgt:[^ ]+", passportText)
	if (!match) {
		return false;
	}
	match, _ = regexp.MatchString("hcl:[^ ]+", passportText)
	if (!match) {
		return false;
	}
	match, _ = regexp.MatchString("ecl:[^ ]+", passportText)
	if (!match) {
		return false;
	}
	match, _ = regexp.MatchString("pid:[^ ]+", passportText)
	if (!match) {
		return false;
	}
	return true
}