package main

import (
	"fmt"
	"sort"
)

/*
	Algorithm:
	1. Do sorting array with odd length number
	2. Set variable that will be contained zigzag pattern
	3. Do looping to transfer elements from index zero to index median -1 into zigzag var
	4. Do reversed looping to transfer elements from last index to median index into zigzag var
	5. Return zigzag var
*/

func findZigZagSequence(arr []int32) []int32 {
	sort.SliceStable(arr, func(i, j int) bool { return arr[i] < arr[j] })

	zigZagArr := []int32{}
	median := (len(arr) / 2)

	for i := 0; i < median; i++ {
		zigZagArr = append(zigZagArr, arr[i])
	}

	for i := len(arr) - 1; i >= median; i-- {
		zigZagArr = append(zigZagArr, arr[i])
	}

	return zigZagArr
}

func main() {
	fmt.Println(findZigZagSequence([]int32{2, 6, 7, 9, 3})) // return []{2, 3, 9, 7, 6}
}
