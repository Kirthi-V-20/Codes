package main

import "fmt"

func fb(n int) {

	fmt.Print("Enter the number : ")
	fmt.Scanln(&n)
	for m := 1; m <= n; m++ {
		if m%3 == 0 && m%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if m%3 == 0 {
			fmt.Println("Fizz")
		} else if m%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(m)
		}
	}
}

func main() {

	fb(15)

}
