package main

import (
	p "piscine"
)


func checkmate(board []string){
	if p.CheckBoard(board) == false {
		// print error
	}
	// if GetPosition(board) == false {
	// 	// print error
	// }
	// if Judge(board) {
	// 	ft.PrintStr("Success\n")
	// } else {
	// 	ft.PrintStr("Fail\n")
	// }
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