package day2

import (
	"felix-schindler/aoc-24/helper"
	"log"
	"slices"
	"strings"
)

func part1(lines []string) int {
	// Count of save records
	count := 0

	for index, line := range lines {
		// Get line num and int array
		lineNum := index + 1
		nums := helper.StringToInt(strings.Split(line, " "))

		// Check if all values are unique
		if !isUnique(nums) {
			log.Println("Numbers of", lineNum, "are not unique")
			continue
		}

		// Check if values are sorted
		if !isSorted(nums) {
			log.Println("Numbers of", lineNum, "are not sorted")
			continue
		}

		// Check difference between all nums
		if !validDiff(nums) {
			log.Println("Numbers of", lineNum, "have invalid difference")
			continue
		}

		count++
	}

	return count
}

func isUnique(nums []int) bool {
	isUnique := true
	occurences := helper.CountOccurence(nums)

	for _, numCount := range occurences {
		if numCount != 1 {
			isUnique = false
			break
		}
	}

	return isUnique
}

func isAscending(nums []int) bool {
	// Check ascending
	return slices.IsSorted(nums)
}

func isDescending(nums []int) bool {
	// Check descending
	descending := nums
	slices.Reverse(descending)
	isDescending := slices.IsSorted(descending)
	return isDescending
}

func isSorted(nums []int) bool {
	if isAscending(nums) {
		return true
	}

	return isDescending(nums)
}

func validDiff(nums []int) bool {
	last := -1
	diffValid := false

	for _, num := range nums {
		if last != -1 {
			diff := helper.Abs(num - last)
			diffValid = diff >= 1 && diff <= 3

			if !diffValid {
				break
			}
		}

		last = num
	}

	return diffValid
}
