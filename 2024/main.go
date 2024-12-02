package main

import (
	"felix-schindler/aoc-24/day1"
	"felix-schindler/aoc-24/day2"
	"os"
)

func main() {
	switch os.Args[1] {
	case "1":
		day1.Handler(len(os.Args) > 2)
		break
	case "2":
		day2.Handler(len(os.Args) > 2)
	default:
		panic("Invalid day")
	}
}
