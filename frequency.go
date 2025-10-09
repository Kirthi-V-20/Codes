package main

import "fmt"

func Frequency(s string) map[string]int {
	freq := map[string]int{}

	for i := 0; i < len(s); i++ {
		char := string(s[i])
		freq[char]++
	}

	return freq
}

func main() {
	fmt.Println(Frequency("apple"))

}
