package main

import "fmt"

func Abbreviate(s string) {
	abbr := ""

	if len(s) > 0 && s[0] >= 'A' && s[0] <= 'Z' {
		abbr += string(s[0])
	}

	for i := 1; i < len(s); i++ {
		if s[i-1] == ' ' && s[i] >= 'A' && s[i] <= 'Z' {
			abbr += string(s[i])
		}
	}

	fmt.Println(abbr)
}

func main() {
	Abbreviate("Indian Institute of Management")
}
