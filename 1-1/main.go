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
	var previous *int

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		depth := nextDepth(scanner)

		if previous == nil {
			previous = &depth
		}

		if depth > *previous {
			count++
		}

		previous = &depth
	}

	fmt.Println("count is", count)
}

func nextDepth(scanner *bufio.Scanner) int {
	depth, _ := strconv.Atoi(scanner.Text())
	return depth
}
