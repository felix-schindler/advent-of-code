package day5

import (
	"log"
	"slices"
	"strconv"
	"strings"
)

func part1(lines []string) int {
	raw_rules, raw_updates := splitLines(lines)
	rules := map[string][]string{}
	updates := []string{}

	for _, rule := range raw_rules {
		xy := strings.Split(rule, "|")

		if len(xy) != 2 {
			log.Fatalln("Invalid rule", rule)
		}

		x := xy[0]
		y := xy[1]

		rules[x] = append(rules[x], y)
	}

	for _, update := range raw_updates {
		include := true
		nums := strings.Split(update, ",")

		for i, num := range nums {
			rule, exists := rules[num]

			if exists {
				pre := nums[:i]

				for _, y := range rule {
					if slices.Contains(pre, y) {
						include = false
						break // break outer loop
					}
				}

				if !include {
					break // break outer loop
				}
			}
		}

		if include {
			updates = append(updates, update)
		}
	}

	sum := 0
	for _, valid_update := range updates {
		split := strings.Split(valid_update, ",")
		middleIndex := len(split) / 2
		int, _ := strconv.Atoi(split[middleIndex])
		sum += int
	}

	return sum
}

func splitLines(lines []string) ([]string, []string) {
	splitIndex := -1
	for i, line := range lines {
		if strings.TrimSpace(line) == "" {
			splitIndex = i
			break
		}
	}

	if splitIndex == -1 {
		log.Fatalln("Failed to split into rules and updates")
	}

	return lines[:splitIndex], lines[splitIndex+1:]
}
