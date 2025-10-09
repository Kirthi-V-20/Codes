package main

import "strconv"

func fizzbuzz(n int) []string {
	var fb []string
	for i := 1; i <= n; i++ {
		if i%15 == 0 {
			fb = append(fb, "FizzBuzz")
		} else if i%3 == 0 {
			fb = append(fb, "Fizz")
		} else if i%5 == 0 {
			fb = append(fb, "Buzz")
		} else {
			fb = append(fb, strconv.Itoa(i))
		}
	}
	return fb
}
