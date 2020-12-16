package main

import "fmt"

func main() {
	fmt.Println(findMin([]int{10, 3, 5, 7}))
}

func findMin(arr []int) int {
	result := arr[0]
	for _, v := range arr {
		if v < result {
			result = v
		}
	}
	return result
}
