package main

import "fmt"

func palindrome(s string) {

	rev := ""
	for i := len(s) - 1; i >= 0; i-- {
		rev += string(s[i])
	}
	if s == rev {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}

func main() {
	palindrome("level")
}
