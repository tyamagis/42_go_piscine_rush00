package piscine

import "fmt"

// pieces are...
// P, B, R, Q and K

func isPiece(r rune) bool {
	if r == 'P' || r == 'B' || r == 'R' || r == 'Q' || r == 'K' {
		return true
	}
	return false
}

func pawn(board []string, i, j int) {
	for pi, ps := range board {
		for pj, pc := range ps {
			if (pi == i - 1 && pj == j - 1 && !isPiece(pc)) {

			} else if (pi == i - 1 && pj == j + 1 && !isPiece(pc)) {

			}
		}
	}
}

func bishop(board []string, i, j int) {

}

func rook(board []string, i, j int) {

}

func queen(board []string, i, j int) {

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
	fmt.Println(board)
}

func ConvertBoard(board []string) {
	convAscii(board)
	for i, s := range board {
		for j, c := range s {
			if c == 'P' {
				pawn(board, i, j)
			} else if c == 'B' {
				bishop(board, i, j)
			} else if c == 'R' {
				rook(board, i, j)
			} else if c == 'Q' {
				queen(board, i, j)
			}
		}
	}
}
