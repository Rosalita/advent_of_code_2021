package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	count := 0

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	depths := readAll(scanner)

	var previousWindow *[]int

	// range over the depths until the third from last depth
	for i := 0; i < len(depths)-2; i++ {

		// Group the depths into a window of three values
		window := depths[i : i+3]

		// If there is no previous window to compare against....
		if previousWindow == nil {
			// Save the first window and continue to next iteration.
			previousWindow = &window
			continue
		}

		if sum(window) > sum(*previousWindow) {
			count++
		}
		previousWindow = &window
	}

	fmt.Println("count is", count)
}

func sum(input []int) int {
	sum := 0
	for i := range input {
		sum += input[i]
	}
	return sum
}

func readAll(scanner *bufio.Scanner) []int {
	var depths []int
	for scanner.Scan() {
		depth, _ := strconv.Atoi(scanner.Text())
		depths = append(depths, depth)
	}
	return depths
}
