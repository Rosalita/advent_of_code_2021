package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	input := readAll(scanner)

	var majority string
	var minority string

	// there are 12 characters in each binary number, so loop 12 times once for each character.
	for i := 0; i < 12; i++ {

		// initialise some counters
		countOfOnes := 0
		countOfZeros := 0

		// count the number of 0s and 1s for each line of the input
		for _, number := range input {
			value := string(number[i])
			if value == "1" {
				countOfOnes++
			} else {
				countOfZeros++
			}
		}

		// save majority and minority values.
		if countOfOnes > countOfZeros {
			majority += "1"
			minority += "0"
		} else {
			majority += "0"
			minority += "1"
		}

	}

	gamma, err := strconv.ParseInt(majority, 2, 64)
	if err != nil {
		fmt.Println(err)
	}

	epsilon, err := strconv.ParseInt(minority, 2, 64)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("gamma * epsilon is:", gamma*epsilon)
}

func readAll(scanner *bufio.Scanner) []string {
	var binaryNumbers []string
	for scanner.Scan() {
		binaryNumbers = append(binaryNumbers, scanner.Text())
	}
	return binaryNumbers
}
