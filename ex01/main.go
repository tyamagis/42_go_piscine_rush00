package main

import (
	p "piscine"
	u "utils"
	"os"
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
	ac := 0
	for range os.Args {
		ac++
	}
	board := os.Args[1:ac]

	checkmate(board);
}