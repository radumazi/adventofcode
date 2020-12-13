package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	// "math"
)

func day10() {
	file, err := os.Open("./day10_input.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var numbers []int
	for scanner.Scan() {
		line := scanner.Text()
		number, _ := strconv.Atoi(line)
		numbers = append(numbers, number)
	}
	//part 1
	sort.Ints(numbers[0:])
	// fmt.Println(checkInput10(numbers))
	//part 2
	// checkInput10_2(numbers)
	numbers = append([]int{0}, numbers...)
	numbers = append(numbers, numbers[len(numbers)-1]+3)
	fmt.Println(countPossibleAdapters(0, numbers, make(map[int]int)))
}

func checkInput10(numbers []int) int {
	var oneJolts = 0
	var threeJolts = 0
	var nextNumber = 0
	for i := 0; i < len(numbers); i++ {
		if i+1 < len(numbers) {
			nextNumber = numbers[i+1]
		} else {
			nextNumber = numbers[i] + 3
		}
		switch nextNumber - numbers[i] {
		case 1:
			oneJolts++
			break
		case 3:
			threeJolts++
			break
		}
	}

	return oneJolts * threeJolts
}

// func checkInput10_2(numbers []int) {
// 	var removable []int
// 	for i := 0; i < len(numbers); i++ {
// 		if i+2 < len(numbers) {
// 			if numbers[i+2]-numbers[i] <= 3 {
// 				removable = append(removable, numbers[i+1])
// 			}
// 		}
// 	}

// 	fmt.Println(removable)
// 	removableCount := float64(len(removable))
// 	fmt.Println(math.Pow(2, removableCount))
// }

func countPossibleAdapters(fromIndex int, nums []int, visited map[int]int) int {
	if fromIndex >= len(nums)-3 {
		return 1
	}

	num := nums[fromIndex]
	if res, ok := visited[num]; ok {
		return res
	}

	var count int
	for i := fromIndex + 1; i < fromIndex+4; i++ {
		n := nums[i]
		if areCompatible(num, n) {
			count += countPossibleAdapters(i, nums, visited)
		}
	}

	visited[num] = count // store the result
	return count
}

func areCompatible(low, high int) bool {
	return low+1 == high || low+2 == high || low+3 == high
}