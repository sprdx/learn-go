package main

import (
	"fmt"
)

// miniMaxSum finds the minimum and maximum values that can be calculated by summing exactly four of the five integers. Then print the respective minimum and maximum values as a single line of two space-separated long integers.
func miniMaxSum(arr []int32) {
	// Set first element of array as lowest value
	var low int32 = arr[0]
	var high int32
	var Sum int64

	// Find the lowest and highest number of array and add each elements to Sum
	for _, e := range arr {
		Sum += int64(e)
		if e > high {
			high = e
		} else if e < low {
			low = e
		}
	}

	// Minimum sum of four elements is sum substracted by highest number
	// Maximum sum of four elements is sum substracted by lowest number
	var miniSumOfFourElements int64 = Sum - int64(high)
	var maxSumOfFourElements int64 = Sum - int64(low)

	fmt.Println(miniSumOfFourElements, maxSumOfFourElements)
}

// func main() {
// 	miniMaxSum([]int32{1, 2, 3, 4, 5}) // result 10 14
// }
