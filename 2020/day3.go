package main

import (
	"bufio"
	"os"
	"fmt"
)

func day3() {
	file, err := os.Open("./day3_input.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var i,markX,markY,treesCount int
	var slopeX = 1;
	var slopeY = 2;
	markX += slopeX;
	markY += slopeY;
	for scanner.Scan() {
		line := scanner.Text()
		if (i == markY) {
			if (string(line[markX]) == "#") {
				treesCount++
			}
			markY = markY+slopeY
			var lineLength = len(line)
			markX = (markX+slopeX)%lineLength	
		}
		i++
	}
	fmt.Println(treesCount)
}