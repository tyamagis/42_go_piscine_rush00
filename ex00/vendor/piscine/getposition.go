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

func pawn(board []string, pi, pj int) {
	for i, _ := range board {
		rs := []rune(board[i])
		for j, pc := range rs {
			if (i == pi - 1 && j == pj - 1 && !isPiece(pc)) {
				rs[j] = '1'
			} else if (i == pi - 1 && j == pj + 1 && !isPiece(pc)) {
				rs[j] = '1'
			}
		}
		board[i] = string(rs)
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
}

func ConvertBoard(board []string) {
	fmt.Println("ConvertBoard() >>")
	fmt.Println("first :", board)
	convAscii(board)
	fmt.Println("Ascii[.] :", board)
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
	fmt.Println("at last :", board)
}
