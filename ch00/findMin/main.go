package main

func main() {
	return
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

func findMinWithIdx(arr []int) map[string]int {
	minVal := arr[0]
	result := map[string]int{"idx": 0, "value": 0}
	for idx, v := range arr {
		if v < minVal {
			minVal = v
			result = map[string]int{"idx": idx, "value": v}
		}
	}
	return result
}
