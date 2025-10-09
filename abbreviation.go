package main

import (
	"strings"
	"unicode"
)

func Abbreviate(s string) string {
	words := strings.Split(s, " ")
	var abbreviate strings.Builder
	for _, word := range words {
		if len(word) > 0 && unicode.IsUpper(rune(word[0])) {
			abbreviate.WriteByte(word[0])
		}
	}
	return abbreviate.String()
}
