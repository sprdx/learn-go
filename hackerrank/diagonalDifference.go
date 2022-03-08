package main

/*
Algorithm:
1. Do looping throughout the lenght of array
2. Each looping add diagonal element value to variable which primary diagonal have pattern arr[i][i]
   and secondary diagonal have pattern arr[i][len(arr)-1-i]
3. After looping finished, to obtain absolute value of difference between diagonal, check the diagonal that have higher value

*/
func diagonalDifference(arr [][]int32) int32 {
	// Write your code here
	var diff int32
	var diagonal1 int32
	var diagonal2 int32

	for i := 0; i < len(arr); i++ {
		diagonal1 += arr[i][i]
		diagonal2 += arr[i][len(arr)-1-i]
	}

	if diagonal1 > diagonal2 {
		diff = diagonal1 - diagonal2
	} else {
		diff = diagonal2 - diagonal1
	}

	return diff
}

// func main() {
// 	diff := diagonalDifference([][]int32{{1, 2, 3}, {3, 4, 5}, {5, 6, 7}})
// 	fmt.Println(diff) // result 12 - 12 = 0
// }
