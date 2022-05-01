package piscine

import u "utils"

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
	for i, s := range board {
		for j, pc := range s {
			if i == pi-1 && j == pj-1 && isKing(pc) {
				return 1
			} else if i == pi-1 && j == pj+1 && isKing(pc) {
				return 1
			}
		}
	}
	return 0
}

func bishop(board []string, bi, bj int) int {
	size := u.StrLen(board[0])
	// left top
	for i, k := bi-1, 0; i >= 0; i-- {
		j := bj - k
		if j < 0 {
			break
		}
		if isKing(rune(board[i][j])) {
			return 1
		}
		if isPiece(rune(board[i][j])) {
			break
		}
		k++
	}
	// right top
	for i, k := bi-1, 0; i >= 0; i-- {
		j := bj + k
		if j >= size {
			break
		}
		if isKing(rune(board[i][j])) {
			return 1
		}
		if isPiece(rune(board[i][j])) {
			break
		}
		k++
	}
	// left bottom
	for i, k := bi+1, 0; i < size; i++ {
		j := bj - k
		if j < 0 {
			break
		}
		if isKing(rune(board[i][j])) {
			return 1
		}
		if isPiece(rune(board[i][j])) {
			break
		}
		k++
	}
	// right bottom
	for i, k := bi+1, 0; i < size; i++ {
		j := bj + k
		if j >= size {
			break
		}
		if isKing(rune(board[i][j])) {
			return 1
		}
		if isPiece(rune(board[i][j])) {
			break
		}
	}
	return 0
}

func rook(board []string, ri, rj int) int {
	size := u.StrLen(board[0])
	// top
	for i := ri - 1; i >= 0; i-- {
		if isKing(rune(board[i][rj])) {
			return 1
		}
		if isPiece(rune(board[i][rj])) {
			break
		}
	}
	// left
	for j := rj - 1; j >= 0; j-- {
		if isKing(rune(board[ri][j])) {
			return 1
		}
		if isPiece(rune(board[ri][j])) {
			break
		}
	}
	// right
	for j := rj + 1; j < size; j++ {
		if isKing(rune(board[ri][j])) {
			return 1
		}
		if isPiece(rune(board[ri][j])) {
			break
		}
	}
	// bottom
	for i := ri + 1; i < size; i++ {
		if isKing(rune(board[i][rj])) {
			return 1
		}
		if isPiece(rune(board[i][rj])) {
			break
		}
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
