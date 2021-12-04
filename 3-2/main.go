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

	oxygenBinary := processOxygen(input)
	scrubberBinary := processScrubber(input)

	oxygen, err := strconv.ParseInt(oxygenBinary[0], 2, 64)
	if err != nil {
		fmt.Println(err)
	}

	scrubber, err := strconv.ParseInt(scrubberBinary[0], 2, 64)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("oxygen rating is:", oxygen)
	fmt.Println("scrubber rating is:", scrubber)

	fmt.Println("oxygen * scrubber rating is:", oxygen*scrubber)
}

func processOxygen(input []string) []string {

	// next data will be the input for the next call to process
	var nextData []string

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

		// find the most common value for this position
		var mostCommonValue string

		if countOfOnes >= countOfZeros {
			mostCommonValue = "1"
			fmt.Println("most common is 1")
		} else {
			mostCommonValue = "0"
			fmt.Println("most common is 0")
		}

		for _, number := range input {
			value := string(number[i])
			if value == mostCommonValue {
				nextData = append(nextData, number)
			}
		}

		input = nextData
		nextData = []string{}
		if len(input) == 1 {
			return input
		}
	}

	// this line will never be reached and exists only to please the compiler
	return []string{"foo"}
}

func processScrubber(input []string) []string {

	// next data will be the input for the next call to process
	var nextData []string

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

		// find the least common value for this position
		var leastCommonValue string

		if countOfOnes >= countOfZeros {
			leastCommonValue = "0"
		} else {
			leastCommonValue = "1"
		}

		for _, number := range input {
			value := string(number[i])
			if value == leastCommonValue {
				nextData = append(nextData, number)
			}
		}

		input = nextData
		nextData = []string{}
		if len(input) == 1 {
			return input
		}
	}

	// this line will never be reached and exists only to please the compiler
	return []string{"foo"}
}

func readAll(scanner *bufio.Scanner) []string {
	var binaryNumbers []string
	for scanner.Scan() {
		binaryNumbers = append(binaryNumbers, scanner.Text())
	}
	return binaryNumbers
}
