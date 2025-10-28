package day7

import (
	"felix-schindler/aoc-24/helper"
	"log"
	"strconv"
	"strings"
)

func part1(lines []string) int {
	sum := 0

	for _, line := range lines {
		if line == "" {
			continue
		}

		split := strings.Split(line, ": ")

		if len(split) != 2 {
			log.Fatalln("Invalid line format", split)
		}

		result, _ := strconv.Atoi(split[0])
		nums := helper.StringToInt(strings.Split(split[1], " "))

		if dfs(nums, 1, nums[0], result) {
			sum += result
		}
	}

	return sum
}

func dfs(numbers []int, index, current, target int) bool {
	if index == len(numbers) {
		return current == target
	}

	// Try addition
	if dfs(numbers, index+1, current+numbers[index], target) {
		return true
	}

	// Try multiplication
	if dfs(numbers, index+1, current*numbers[index], target) {
		return true
	}

	return false
}
