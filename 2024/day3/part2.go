package day3

import (
	"regexp"
	"strconv"
	"strings"
)

func part2(lines []string) int {
	sum := 0
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)
	enabled := true

	for _, line := range lines {
		matches := re.FindAllStringSubmatch(line, -1)

		// For each match [match, group1, group2]
		for _, match := range matches {
			action := match[0]
			if action == "do()" {
				enabled = true
			} else if action == "don't()" {
				enabled = false
			} else if enabled && strings.HasPrefix(action, "mul") {
				int1, _ := strconv.Atoi(match[1])
				int2, _ := strconv.Atoi(match[2])

				sum += int1 * int2
			}
		}
	}

	return sum
}
