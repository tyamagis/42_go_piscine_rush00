package piscine

import (
	"ft"
)

func StrRFind(s string, r rune) bool {
	for _, c := range s {
		if c == r {
			return true
		}
	}
	return false
}

func StrLen(s string) int {
	len := 0
	for range s {
		len++
	}
	return len
}

func PrintStr(s string) {
	for _, c := range s {
		ft.PrintRune(c)
	}
}
