package main

import "fmt"

func main() {
	var i int
	fmt.Scan(&i)

	if i%15 == 0 {
		fmt.Printf("FizzBuzz\n")
	} else if i%3 == 0 {
		fmt.Printf("Fizz\n")
	} else if i%5 == 0 {
		fmt.Printf("Buzz\n")
	} else {
		fmt.Printf("%d\n", i)
	}
}
