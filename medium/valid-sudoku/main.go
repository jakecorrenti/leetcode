package medium

// NOTE: still not quite too sure where the idea of the math for determining which box we are in came from.
func isValidSudoku(board [][]byte) bool {
	// lists of seen digits for respective row, column, or 3x3 box
	var rows, cols, boxes []map[byte]bool

	for i := 0; i < 9; i++ {
		rows = append(rows, make(map[byte]bool))
		cols = append(cols, make(map[byte]bool))
		boxes = append(boxes, make(map[byte]bool))
	}

	for i := range board {
		for j, num := range board[i] {
			// i = row, j = col
			// num = board[i][j]
			currGridBox := (i/3)*3 + (j / 3)
			if num == '.' {
				continue
			}
			if rows[i][num] || cols[j][num] || boxes[currGridBox][num] {
				return false
			}
			rows[i][num] = true
			cols[j][num] = true
			boxes[currGridBox][num] = true
		}
	}
	return true
}
