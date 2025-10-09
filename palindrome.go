package main

import "strings"

func reverse(s string) string {
	var rev strings.Builder
	for i := len(s) - 1; i >= 0; i-- {
		word := string(s[i])
		rev.WriteString(word)
	}
	return rev.String()
}

func palindrome(s string) bool {
	rev := reverse(s)
	return rev == s

}
