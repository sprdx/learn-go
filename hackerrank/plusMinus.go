package main

import "fmt"

// plusMinus calculate the ratios of its elements that are positive, negative, and zero. Print the decimal value of each fraction on a new line with 6 places after the decimal
func plusMinus(arr []int32) {

	// Algorithm :
	// Set variables for each group elements that count occurrence type based number whether positive, negative or zero
	var countPlus float32
	var countNegative float32
	var countZero float32
	var lenghtArray = float32(len(arr))

	// Do looping for the entire array element, check if element is positive, negative or zero, and do increment value of count
	for _, e := range arr {
		if e > 0 {
			countPlus++
		} else if e < 0 {
			countNegative++
		} else {
			countZero++
		}
	}

	// Print ratio of each element groups with 6 numbers after the decimal
	fmt.Printf("%.6f\n", countPlus/lenghtArray)
	fmt.Printf("%.6f\n", countNegative/lenghtArray)
	fmt.Printf("%.6f\n", countZero/lenghtArray)
}

func main() {
	plusMinus([]int32{1, 1, 0, -1, -1}) // result 0.400000, 0.400000, 0.200000
}
