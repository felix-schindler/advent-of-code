package day6

const (
	guard_up    = '^'
	guard_down  = 'v'
	guard_right = '>'
	guard_left  = '<'
	obstruction = '#'
)

func part1(lines []string) int {
	var x, y int
	var dir byte
	found := false

	for i, row := range lines {
		for j, char := range row {
			if char == guard_up || char == guard_down || char == guard_left || char == guard_right {
				x, y = j, i
				dir = byte(char)
				found = true
				break // break inner loop
			}
		}
		if found {
			break // break outer loop
		}
	}

	// Directions: up, right, down, left
	dx := []int{0, 1, 0, -1}
	dy := []int{-1, 0, 1, 0}
	dirs := []byte{guard_up, guard_right, guard_down, guard_left}

	// Convert starting direction to index
	dirIndex := 0
	for i, d := range dirs {
		if d == dir {
			dirIndex = i
			break
		}
	}

	visited := make(map[[2]int]bool)
	visited[[2]int{x, y}] = true

	height, width := len(lines), len(lines[0])

	for {
		// Check next position in current direction
		nx, ny := x+dx[dirIndex], y+dy[dirIndex]

		// If next position is out of bounds, guard leaves the area
		if nx < 0 || nx >= width || ny < 0 || ny >= height {
			break
		}

		// If there's an obstruction, turn right
		if lines[ny][nx] == obstruction {
			dirIndex = (dirIndex + 1) % 4
		} else {
			// Move forward
			x, y = nx, ny
			visited[[2]int{x, y}] = true
		}
	}

	return len(visited)
}
