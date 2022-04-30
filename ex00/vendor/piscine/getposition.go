package piscine

// pieces are...
// P, B, R, Q and K

func isKing(r rune) bool {
	if r == 'K' {
		return true
	}
	return false
}

func isPiece(r rune) bool {
	if r == 'K' || r == 'P' || r == 'B' || r == 'R' || r == 'Q' {
		return true
	}
	return false
}

func pawn(board []string, pi, pj int) int {
	for i, _ := range board {
		rs := []rune(board[i])
		for j, pc := range rs {
			if i == pi-1 && j == pj-1 && isKing(pc) {
				return 1
			} else if i == pi-1 && j == pj+1 && isKing(pc) {
				return 1
			}
		}
		board[i] = string(rs)
	}
	return 0
}

func bishop(board []string, bi, bj int) int {
	for i, _ := range board {
		rs := []rune(board[i])
		for j, pc := range rs {
			if (i+j == bi+bj || i-j == bi-bj) && !isKing(pc) {
				return 1
			}
		}
		board[i] = string(rs)
	}
	return 0
}

func rook(board []string, ri, rj int) int {
	for i, _ := range board {
		rs := []rune(board[i])
		for j, pc := range rs {
			if (i == ri || j == rj) && !isKing(pc) {
				return 1
			}
		}
		board[i] = string(rs)
	}
	return 0
}

func convAscii(board []string) {
	for i, _ := range board {
		rs := []rune(board[i])
		for j, rc := range rs {
			if rc >= 127 || !isPiece(rc) {
				rs[j] = '.'
			}
		}
		board[i] = string(rs)
	}
}

func ConvertBoard(board []string) bool {
	convAscii(board)
	check := 0
	for i, s := range board {
		for j, c := range s {
			if c == 'P' {
				check += pawn(board, i, j)
			} else if c == 'B' {
				check += bishop(board, i, j)
			} else if c == 'R' {
				check += rook(board, i, j)
			} else if c == 'Q' {
				check += bishop(board, i, j)
				check += rook(board, i, j)
			}
			if check != 0 {
				return true
			}
		}
	}
	return false
}
