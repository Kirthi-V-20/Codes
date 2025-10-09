package main

import "strings"

func Pangram(s string) string {
	s = strings.ToLower(s)
	alpha := "abcdefghijklmnopqrstuvwxyz"

	for i := 0; i < len(alpha); i++ {
		if !strings.Contains(s, string(alpha[i])) {
			return "Not pangram"
		}

	}
	return "Pangram"

}
