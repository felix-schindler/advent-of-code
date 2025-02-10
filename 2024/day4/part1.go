package day4

func part1(lines []string) int {
	// Make a 2D array (grid) with each letter of the input
	grid := make([][]rune, 0)
	for _, line := range lines {
		if line != "" {
			grid = append(grid, []rune(line))
		}
	}

	// Call the count function
	return countOccurrences(grid, "XMAS")
}

var directions = [][2]int{
	{0, 1},   // Right
	{0, -1},  // Left
	{1, 0},   // Down
	{-1, 0},  // Up
	{1, 1},   // Diagonal Down-Right
	{-1, -1}, // Diagonal Up-Left
	{1, -1},  // Diagonal Down-Left
	{-1, 1},  // Diagonal Up-Right
}

// Helper function to check if a position is valid in the grid
func isValid(x int, y int, rows int, cols int) bool {
	return x >= 0 && x < rows && y >= 0 && y < cols
}

// Function to search for word starting from (x, y) in a specific direction
func searchWord(grid [][]rune, x int, y int, direction [2]int, word string) bool {
	rows, cols := len(grid), len(grid[0])
	for i := 0; i < len(word); i++ {
		nx, ny := x+direction[0]*i, y+direction[1]*i
		if !isValid(nx, ny, rows, cols) || grid[nx][ny] != rune(word[i]) {
			return false
		}
	}
	return true
}

// Function to count all occurrences of a given word
func countOccurrences(grid [][]rune, word string) int {
	rows, cols := len(grid), len(grid[0])
	count := 0

	for x := 0; x < rows; x++ {
		for y := 0; y < cols; y++ {
			for _, direction := range directions {
				if searchWord(grid, x, y, direction, word) {
					count++
				}
			}
		}
	}

	return count
}
