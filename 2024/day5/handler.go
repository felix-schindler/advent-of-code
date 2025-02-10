package day5

import (
	"felix-schindler/aoc-24/helper"
	"log"
)

func Handler(call2 bool) {
	// Get input
	lines := helper.TextFileByLines("day5/input.txt")

	if call2 {
		log.Println("Part 2", part2(lines))
	} else {
		log.Println("Part 1", part1(lines))
	}
}
