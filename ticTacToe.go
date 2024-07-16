package ticTacToe

const X byte = 'X'
const O byte = 'O'

func checkVertical(board []string, target byte) bool {
	for col, val := range board[0] {
		if val == rune(target) {
			if board[1][col] == target &&
				board[2][col] == target &&
				board[3][col] == target {
				return true
			}
		}
	}
	return false
}

func checkHorizontal(board []string, target byte) bool {
	for _, row := range board {
		if row[0] == target && row[1] == target && row[2] == target && row[3] == target {
			return true
		}
	}
	return false
}

func checkDiagonal(board []string, target byte) bool {
	if board[0][0] == target && board[1][1] == target && board[2][2] == target && board[3][3] == target {
		return true
	} else if board[0][3] == target && board[1][2] == target && board[2][1] == target && board[3][0] == target {
		return true
	}
	return false
}

func checkCorner(board []string, target byte) bool {
	// board is of a constant size, for corners just check directly
	return board[0][0] == target &&
		board[0][3] == target &&
		board[3][0] == target &&
		board[3][3] == target

}

func checkSquare(board []string, target byte) bool {
	for row_idx, row := range board {
		for col_idx := range row {
			if row_idx < 3 && col_idx < 3 {
				if board[row_idx][col_idx] == target && board[row_idx+1][col_idx] == target && board[row_idx][col_idx+1] == target && board[row_idx+1][col_idx+1] == target {
					return true

				}
			}
		}
	}
	return false
}

// given a string representation of a board, return
// either "X" if X won the game, "O" if O won the game, or "" if neither won
func CheckWinner(board []string) string {
	target := X
	if checkCorner(board, target) || checkDiagonal(board, target) ||
		checkHorizontal(board, target) || checkSquare(board, target) || checkVertical(board, target) {
		return string(target)
	}
	target = O
	if checkCorner(board, target) || checkDiagonal(board, target) ||
		checkHorizontal(board, target) || checkSquare(board, target) || checkVertical(board, target) {
		return string(target)
	}
	return ""
}

func AnyMovesLeft(board []string) bool {
	for row := range board {
		for col := range board[row] {
			if board[row][col] != X && board[row][col] != O {
				return true
			}
		}
	}
	return false
}

func IsGameOver(board []string) bool {
	winner := CheckWinner(board)
	if winner == "" {
		return !AnyMovesLeft(board)
	}
	return true
}
