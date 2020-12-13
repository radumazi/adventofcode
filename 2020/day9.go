package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func day9() {
	file, err := os.Open("./day9_input.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var numbers = make([]int64, 1600)
	var i int = 0
	for scanner.Scan() {
		line := scanner.Text()
		number, _ := strconv.ParseInt(line, 10, 0)
		numbers[i] = number
		i++
	}
	//part 1
	var invalidNumber = checkInput9(numbers, 25)
	fmt.Println(invalidNumber)
	// var invalidNumber = 10884537
	//part 2
	fmt.Println(detectSequence(numbers, invalidNumber))
}

func checkInput9(numbers []int64, preambleSize int) int64 {
	var preamble = make([]int64, preambleSize)
	for i := 0; i < len(numbers); i++ {
		if i >= preambleSize {
			var validation = checkNumber(preamble, numbers[i])
			if (!validation) {
				return numbers[i]
			}
		}
		preamble[i%preambleSize] = numbers[i]
	}

	return -1
}

func checkNumber(preamble []int64, number int64) bool {
	var valid = false
	for i := 0; i<len(preamble); i++ {
		for j := 0; j<len(preamble); j++ {
			if preamble[i]+preamble[j] == number {
				return true
			}
		}
	}

	return valid
}

func detectSequence(numbers []int64, invalidNumber int64) int64 {
	var sum int64 = 0
	var min, max int64
	for i := 0; i<len(numbers); i++ {
		sum = 0
		for j:= i; j<len(numbers); j++ {
			sum += numbers[j]
			if (sum == invalidNumber) {
				min = numbers[i]
				max = numbers[i]
				for k := i; k<=j; k++ {
					if numbers[k] < min {
						min = numbers[k]
					}
					if numbers[k] > max {
						max = numbers[k]
					}
				}
				return min+max
			}
			if (sum > invalidNumber) {
				break
			}
		}
	}

	return sum
}