package day2

import (
	"felix-schindler/aoc-24/helper"
	"strings"
)

// Same as 1 but when there's only one issue, we skip it
func part2(lines []string) int {
	// Count of save records
	count := 0

	for _, line := range lines {
		// Convert line to int array
		nums := helper.StringToInt(strings.Split(line, " "))

		// Sum of all issues
		sum := countUniqueIssues(nums) + countDiffIssues(nums) + countSortIssues(nums)

		// Check difference between all nums
		if sum > 1 {
			continue
		}

		count++
	}

	return count
}

func countUniqueIssues(nums []int) int {
	issues := 0
	occurences := helper.CountOccurence(nums)

	for _, numCount := range occurences {
		if numCount > 1 {
			issues++
		}
	}

	return issues
}

// Count how many numbers are NOT sorted in ascending or descending order
func countSortIssues(nums []int) int {
	issues := 0

	// Check for ascending order deviations
	if isAscending(nums) {
		for i := 0; i < len(nums)-1; i++ {
			if nums[i] > nums[i+1] {
				issues++
			}
		}
	}

	// Check for descending order deviations
	if isDescending(nums) {
		for i := 0; i < len(nums)-1; i++ {
			if nums[i] < nums[i+1] {
				issues++
			}
		}
	}

	return issues
}

func countDiffIssues(nums []int) int {
	last := -1
	issues := 0

	for _, num := range nums {
		if last != -1 {
			diff := helper.Abs(num - last)
			diffValid := diff >= 1 && diff <= 3
			if !diffValid {
				issues++
			}
		}

		last = num
	}

	return issues
}
