package main

import (
	"fmt"
	"sort"
)

// findMedian return median value of array
func findMedian(arr []int32) int32 {
	// Sort array
	sort.SliceStable(arr, func(i, j int) bool { return arr[i] < arr[j] })
	// Check if length of array is odd or even
	if len(arr)%2 != 0 {
		return arr[((len(arr)+1)/2)-1]
	} else {
		return arr[(len(arr)+1)/2]
	}
}

func main() {
	fmt.Println(findMedian([]int32{3, 4, 2, 4, 5}))
}
