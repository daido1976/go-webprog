package main

import (
	"fmt"
	"strconv"
)

func main() {
	i := scanNumber()
	outputStr := fizzBuzz(i)

	fmt.Printf(outputStr + "\n")
}

func scanNumber() int {
	var i int

	fmt.Printf("Please input number.\n")
	fmt.Scan(&i)
	return i
}

func fizzBuzz(i int) string {
	if i%15 == 0 {
		return "FizzBuzz"
	} else if i%3 == 0 {
		return "Fizz"
	} else if i%5 == 0 {
		return "Buzz"
	} else {
		return strconv.Itoa(i)
	}
}
