package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func day1() {
	file, err := os.Open("./day1_input.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var numbers = make([]int64, 0)
	for scanner.Scan() {
		line := scanner.Text()
		var number, err = strconv.ParseInt(line, 10, 0)
		numbers = append(numbers, number)
		check(err)
	}
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			for k := 0; k < len(numbers); k++ {
				if numbers[i]+numbers[j]+numbers[k] == 2020 {
					fmt.Println(numbers[i] * numbers[j] * + numbers[k])
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
