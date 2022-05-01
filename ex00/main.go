package main

import (
	p "piscine"
)

func checkmate(board []string){
	if p.CheckBoard(board) == false {
		p.PrintStr("Error >> board is invalid.\n")
		return
	}
	if p.ConvertBoard(board) == true {
		p.PrintStr("Success\n")
	} else {
		p.PrintStr("Fail\n")
	}
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