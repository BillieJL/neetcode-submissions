func isValidSudoku(board [][]byte) bool {
	seen := map[byte]bool{}
	for row := 0; row < 9; row++ {
		clear(seen)
		for i := 0; i < 9; i++ {
			if board[row][i] == '.' {
				continue
			} else if seen[board[row][i]] {
				return false
			}
			seen[board[row][i]] = true
		}
	}

	for col := 0; col < 9; col++ {
		clear(seen)
		for i := 0; i < 9; i++ {
			if board[i][col] == '.' {
				continue
			} else if seen[board[i][col]] {
				return false
			}
			seen[board[i][col]] = true
		}
	}

	for square := 0; square < 9; square++ {
		clear(seen)
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				row := (square/3)*3+i
				col := (square%3)*3+j
				if board[row][col] == '.' {
					continue
				} else if seen[board[row][col]] {
					return false
				}
				seen[board[row][col]] = true
			}
		}
	}

	return true
}