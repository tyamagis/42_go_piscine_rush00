package main

import (
	p "piscine"
)


func checkmate(board []string){
	if p.CheckBoard(board) == false {
		// print error
	}
	p.ConvertBoard(board)
	// if Judge(board) {
	// 	ft.PrintStr("Success\n")
	// } else {
	// 	ft.PrintStr("Fail\n")
	// }
}

func main(){
	board := []string{
		"aBcdefg",
		"hijKlmn",
		"oPQRstu",
		"vwxyzAb",
		"CDEFGHI",
		"JkLMNOp",
		"rSTUVWX",
	}

	checkmate(board);
}