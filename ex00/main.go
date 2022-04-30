package main

import (
	p "piscine"
	"ft"
)

func checkmate(board []string){
	if CheckBoard(board) == false {
		// print error
	}
	if GetPosition(board) == false {
		// print error
	}
	if Judge(board) {
		ft.PrintStr("Success\n")
	} else {
		ft.PrintStr("Fail\n")
	}
}

func main(){
	board := []string{
		"R...",
		".K..",
		"..P.",
		"....",
	}

	checkmate(board);
}