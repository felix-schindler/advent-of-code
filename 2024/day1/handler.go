package day1

import (
	"felix-schindler/aoc-24/helper"
	"log"
	"slices"
	"strconv"
	"strings"
)

func Handler(call2 bool) {
	// Get input
	lines := helper.TextFileByLines("day1/input.txt")

	// Prepare input
	ints1 := make([]int, 0)
	ints2 := make([]int, 0)

	// For each line, convert to int and append to slice
	for index, line := range lines {
		if line != "" {
			// Split line by space
			nums := strings.Split(line, "   ")

			if len(nums) == 2 {
				// Convert to int
				i1, _ := strconv.Atoi(nums[0])
				i2, _ := strconv.Atoi(nums[1])

				// Append to slice
				ints1 = append(ints1, i1)
				ints2 = append(ints2, i2)
			} else {
				log.Fatal(index, "contained")
			}
		}
	}

	// Calculate the result
	var result int
	if call2 {
		log.Println("Part 2")
		result = part2(ints1, ints2)
	} else {
		log.Println("Part 1")
		result = part1(ints1, ints2)
	}

	// Print the results
	log.Println(result)
}

func part1(ints1, ints2 []int) int {
	// Sort both int slices
	slices.Sort(ints1)
	slices.Sort(ints2)

	// Check if lengths are equal for the for-loop
	if len(ints1) != len(ints2) {
		log.Fatal("Lengths of slices are not equal")
	}

	// Calculate the differences
	differences := make([]int, 0)
	for i := 0; i < len(ints1); i++ {
		differences = append(differences, helper.Abs(ints2[i]-ints1[i]))
	}

	// Calculate the sum of the differences
	sum := 0
	for _, diff := range differences {
		sum += diff
	}

	// Return the sum
	return sum
}

func part2(ints1, ints2 []int) int {
	dict := helper.CountOccurence(ints2)

	sum := 0
	for _, i := range ints1 {
		sum += i * dict[i]
	}
	return sum
}
