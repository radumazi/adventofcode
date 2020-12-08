package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func day7() {
	file, err := os.Open("./day7_input.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	q, _ := regexp.Compile("([a-z ]+) bags")
	r, _ := regexp.Compile("([0-9]+) ([a-z ]+) bag[s]*")
	var bags = make(map[string]map[string]int64)
	for scanner.Scan() {
		line := scanner.Text()
		var bag = q.FindStringSubmatch(line)
		var matches = r.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			if _, ok := bags[bag[1]]; !ok {
				bags[bag[1]] = make(map[string]int64)
			}
			bags[bag[1]][match[2]], _ = strconv.ParseInt(match[1], 10, 0)
		}
	}
	//part 1
	fmt.Println(len(searchColors([]string{"shiny gold"}, bags)) - 1)
	//part 2
	fmt.Println(countBagsInside("shiny gold", bags))
}

func searchColors(colors []string, bags map[string]map[string]int64) []string {
	var foundSomething = false
	for colorKey := range bags {
		for _, color := range colors {
			if _, ok := bags[colorKey][color]; ok {
				var alreadyThere = false
				for _, value := range colors {
					if colorKey == value {
						alreadyThere = true
					}
				}
				if !alreadyThere {
					colors = append(colors, colorKey)
					foundSomething = true
				}
			}
		}
	}
	if foundSomething {
		return searchColors(colors, bags)
	}

	return colors
}

func countBagsInside(colorKey string, bags map[string]map[string]int64) int64 {
	if _, ok := bags[colorKey]; !ok {
		return 0
	}
	var total int64
	for color, value := range bags[colorKey] {
		total += value + value * countBagsInside(color, bags)
	}

	return total
}