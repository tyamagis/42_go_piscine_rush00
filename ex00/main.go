package main

import (
	p "piscine"
	u "utils"
)

func checkmate(board []string){
	if p.CheckBoard(board) == false {
		u.PrintStr("Error >> board is invalid.\n")
		return
	}
	if p.ConvertBoard(board) == true {
		u.PrintStr("Fail\n")
	} else {
		u.PrintStr("Success\n")
	}
}

func main(){
	board := []string{
		"Q...",
		"..K.",
		".RP.",
		"Q.B.",
	}

	checkmate(board);
}