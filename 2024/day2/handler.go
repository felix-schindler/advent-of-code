package day2

import (
	"felix-schindler/aoc-24/helper"
	"log"
	"os"
	"strconv"
	"strings"
)

func Handler(call2 bool) {
	// Read string form text file
	file, err := os.ReadFile("day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := string(file)

	// Split string by line
	lines := strings.Split(input, "\n")

	// Calculate the result
	var result int
	if call2 {
		log.Println("Part 2")
		result = part2(lines)
	} else {
		log.Println("Part 1")
		result = part1(lines)
	}

	// Print the results
	log.Println(result)
}

func part1(lines []string) int {
	// Count of save records
	count := 0

	for _, line := range lines {
		// Make an int array for each line
		asciiInt := strings.Split(line, " ")
		ints := make([]int, len(asciiInt))

		for _, ascii := range asciiInt {
			i, _ := strconv.Atoi(ascii)
			ints = append(ints, i)
		}

		last := -1
		lastVal := ints[len(ints)-1]
		for _, int := range ints {
			if last != -1 {
				// Check if the difference is between 1 and 3
				diff := helper.Abs(int - last)
				validDiff := diff >= 1 && diff <= 3

				// Check if all numbers are either decreasing or increasing
				ascending := last < int && int < lastVal
				descending := last > int && int > lastVal

				// Break loop if the difference is not valid or the numbers are not increasing or decreasing
				if !validDiff {
					break
				}

				// Increase safe record count
				count++
			}

			last = int
		}
	}

	return count
}

func part2(lines []string) int {
	return -1
}
