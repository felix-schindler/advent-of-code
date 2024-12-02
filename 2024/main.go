package main

import (
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	// Read string form text file
	file, _ := os.ReadFile("input.txt")
	input := string(file)

	// Split string by line
	lines := strings.Split(input, "\n")
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
		differences = append(differences, abs(ints2[i]-ints1[i]))
	}

	// Calculate the sum of the differences
	sum := 0
	for _, diff := range differences {
		sum += diff
	}

	// Print the differences
	log.Println(differences)
	log.Println(sum)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
