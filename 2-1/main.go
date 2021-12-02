package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	horizontalPosition := 0
	depth := 0

	for scanner.Scan() {

		direction, distance := parse(scanner.Text())

		switch direction {
		case "forward":
			horizontalPosition += distance
			continue
		case "down":
			depth += distance
			continue
		case "up":
			depth -= distance
			continue
		}

	}

	fmt.Println("result is:", horizontalPosition*depth)
}

// parse reads a line of input returning the direction and distance.
func parse(command string) (string, int) {
	result := strings.Split(command, " ")
	direction := result[0]
	distance, _ := strconv.Atoi(result[1])
	return direction, distance
}
