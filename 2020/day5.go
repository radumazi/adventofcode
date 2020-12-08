package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"unicode/utf8"
)

func day5() {
	file, err := os.Open("./day5_input.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var max, i int
	var seats [815]int
	for scanner.Scan() {
		line := scanner.Text()
		var row = getNumber("B", "F", 0, 127, line[:7])
		var col = getNumber("R", "L", 0, 7, line[7:])
		var seatID = row*8 + col
		if seatID > max {
			max = seatID
		}
		seats[i] = seatID
		i++
	}
	sort.Ints(seats[0:])
	fmt.Println(seats)
	for i = 1; i < len(seats); i++ {
		if seats[i]-seats[i-1] > 1 {
			fmt.Println(seats[i-1])
			fmt.Println(seats[i])
			fmt.Println("next")
		}
	}
	// fmt.Println(getNumber("B", "F", 0, 127, "FBFBBFF"))
	// fmt.Println(getNumber("R", "L", 0, 7, "RLR"))
}

func getNumber(upChar string, downChar string, min int, max int, sequence string) int {
	if min+1 == max {
		if string(sequence[0]) == upChar {
			return max
		}
		return min
	}
	_, i := utf8.DecodeRuneInString(sequence)
	if string(sequence[0]) == upChar {
		return getNumber(upChar, downChar, (min+max)/2+1, max, sequence[i:])
	}
	return getNumber(upChar, downChar, min, (min+max)/2, sequence[i:])
}
