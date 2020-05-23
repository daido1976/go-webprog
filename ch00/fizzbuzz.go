package main

import "fmt"

func main() {
	i := scanNumber()

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

func scanNumber() int {
	var i int

	fmt.Printf("Please input number.\n")
	fmt.Scan(&i)
	return i
}
