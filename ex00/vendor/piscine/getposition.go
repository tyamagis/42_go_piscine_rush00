package piscine

import "fmt"

// pieces are...
// P, B, R, Q

func isPiece(r rune) bool {
	if r == 'P' || r == 'B' || r == 'R' || r == 'Q' {
		return true
	}
	return false
}

func pawn(board []string, i, j int) {
	// for pi, ps := range board {
	// 	for pj, pc := range ps {
	// if (pi == i - 1 && pj == j - 1 && !isPiece(pc)) {

	// } else (pi == i - 1 && pj == j + 1 && !isPiece(pc)) {

	//}
	// 	}
	// }
}

func bishop(board []string, i, j int) {

}

func rook(board []string, i, j int) {

}

func queen(board []string, i, j int) {

}

func convAscii(board []string) {
	for _, s := range board {
		for _, r := range s {
			fmt.Println(r)
		}
	}
}

func ConvertBoard(board []string, size int) {
	convAscii(board)
	// for i, s := range board {
	// 	for j, c := range s {
	// 		if c == 'P' {
	// 			pawn(board, i, j)
	// 		} else if c == 'B' {
	// 			bishop(board, i, j)
	// 		} else if c == 'R' {
	// 			rook(board, i, j)
	// 		} else if c == 'Q' {
	// 			queen(board, i, j)
	// 		}
	// 	}
	// }
}
