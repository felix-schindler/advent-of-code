package day3

import (
	"regexp"
	"strconv"
)

func part1(lines []string) int {
	sum := 0
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	for _, line := range lines {
		matches := re.FindAllStringSubmatch(line, -1)

		// For each match [match, group1, group2]
		for _, match := range matches {
			int1, _ := strconv.Atoi(match[1])
			int2, _ := strconv.Atoi(match[2])

			sum += int1 * int2
		}
	}

	return sum
}
