package main

import (
	p "piscine"
	"os"
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
	ac := 0
	for range os.Args {
		ac++
	}
	board := os.Args[1:ac]

	checkmate(board);
}