package main

/*
Algorithm:
1. Set an array with 100 elements
2. Do looping and increment value of element of the array based appeared value
3. Set an empty slice
4. Do looping and copy value of array into slice

*/

// countingSort returns the number of times each value appears as an array of integers
func countingSort(arr []int32) []int32 {

	count := [100]int32{}

	for _, e := range arr {
		count[e]++
	}

	countingArr := []int32{}

	for _, e := range count {
		countingArr = append(countingArr, e)
	}

	return countingArr
}

// func main() {
// 	fmt.Println(countingSort([]int32{1, 2, 3})) // return {0, 1, 1, 1, 0, ..., 0}
// }
