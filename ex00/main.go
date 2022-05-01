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
		u.PrintStr("Success\n")
	} else {
		u.PrintStr("Fail\n")
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