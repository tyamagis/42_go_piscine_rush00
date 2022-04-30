package piscine

// king is 1
// board is square (row == column)

func CheckBoard(board []string) bool {
	row := 0;
	column := 0;
	num_king := 0;
	for i, s := range board {
		// check line length
		if i == 0 {
			row = p.StrLen(s)
		} else if (p.StrLen(s) != row) {
			return false
		}
		// count King piece
		if p.StrRFind(s, 'K') {
			num_king++
			if num_king > 1 {
				return false
			}
		}
		column = i;
	}
	// square or not
	if row != column {
		return false
	}
	return true
}