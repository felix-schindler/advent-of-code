package helper

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func CountOccurence(slice []int) map[int]int {
	dict := make(map[int]int)

	for _, i := range slice {
		dict[i]++
	}

	return dict
}

func StringToInt(slice []string) []int {
	ints := make([]int, 0)

	for _, ascii := range slice {
		i, _ := strconv.Atoi(ascii)
		ints = append(ints, i)
	}

	return ints
}

func TextFileByLines(path string) []string {
	// Read string form text file
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	input := string(file)

	// Split string by line
	return strings.Split(input, "\n")
}
