package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func day4_2() {
	file, err := os.Open("./day4_input.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var passportText = ""
	var validPassports int
	for scanner.Scan() {
		line := scanner.Text()
		passportText += " " + line
		if line == "" {
			// fmt.Println(passportText)
			if isValidPassport2(passportText) {
				validPassports++
			}
			passportText = ""
		}
	}
	fmt.Println(validPassports)
}

func isValidPassport2(passportText string) bool {
	r, _ := regexp.Compile("byr:([^ ]+)")
	match := r.FindStringSubmatch(passportText)
	if match == nil || match[1] == "" {
		return false
	} else {
		var value, _ = strconv.ParseInt(match[1], 10, 0)
		if value < 1920 || value > 2002 {
			return false
		}
	}
	r, _ = regexp.Compile("iyr:([^ ]+)")
	match = r.FindStringSubmatch(passportText)
	if match == nil || match[1] == "" {
		return false
	} else {
		var value, _ = strconv.ParseInt(match[1], 10, 0)
		if value < 2010 || value > 2020 {
			return false
		}
	}
	r, _ = regexp.Compile("eyr:([^ ]+)")
	match = r.FindStringSubmatch(passportText)
	if match == nil || match[1] == "" {
		return false
	} else {
		var value, _ = strconv.ParseInt(match[1], 10, 0)
		if value < 2020 || value > 2030 {
			return false
		}
	}
	r, _ = regexp.Compile("hgt:([0-9]+)(cm|in)")
	match = r.FindStringSubmatch(passportText)
	if match == nil || match[1] == "" || match[2] == "" {
		return false
	} else {
		var value, _ = strconv.ParseInt(match[1], 10, 0)
		if match[2] == "cm" && (value < 150 || value > 193) {
			return false
		}
		if match[2] == "in" && (value < 59 || value > 76) {
			return false
		}
	}
	r, _ = regexp.Compile("hcl:#([0-9a-f]+)")
	match = r.FindStringSubmatch(passportText)
	if match == nil || match[1] == "" || len(match[1]) != 6 {
		return false
	}
	r, _ = regexp.Compile("ecl:(amb|blu|brn|gry|grn|hzl|oth)")
	match = r.FindStringSubmatch(passportText)
	if match == nil || match[1] == "" {
		return false
	}
	r, _ = regexp.Compile("pid:([0-9]+)")
	match = r.FindStringSubmatch(passportText)
	if match == nil || match[1] == "" || len(match[1]) != 9 {
		return false
	}

	return true
}
