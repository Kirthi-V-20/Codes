/*Panagram
  - Write a function that accepts a string s as input and returns
    true if s is a panagram. false otherwise. A panagram is a
    sentence which contains all letters of the alphabet (e.g. "the
    quick brown fox jumps over the lazy dog")*/

package main

import (
	"fmt"
	"strings"
)

func Panagram(a string) {

	a = strings.ToLower(a)
	for char := 'a'; char <= 'z'; char++ {
		if !strings.Contains(a, string(char)) {
			fmt.Println("False")
			return
		} else {
			fmt.Println("True")
			return
		}

	}
}

func main() {
	Panagram("The quick brown fox jumps over the lazy dog")
}
