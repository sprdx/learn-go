package main

import "fmt"

/*
Algorithm:
1. Set variable to save result of finding possible max operation
2. Do looping as much as an half of length of matrix
3. Do checking for each cell of submatrix upper-left quadran
4. When checking for highest cell value, compare init cell value to three others possible cell if assumme we do reverse rows/columns
5. Add obtained value to result variable
6. Return possibleMaxSum variable
*/

// flippingMatrix returns possible maximum sum of submatrix n x n quadran from 2n x 2n matrix
func flippingMatrix(matrix [][]int32) int32 {
	// Write your code here
	var possibleMaxSum int32
	for i := 0; i < len(matrix)/2; i++ {
		for j := 0; j < len(matrix)/2; j++ {
			highest := matrix[i][j]
			if highest < matrix[i][len(matrix[i])-1-j] {
				highest = matrix[i][len(matrix[i])-1-j]
			}
			if highest < matrix[len(matrix[i])-1-i][j] {
				highest = matrix[len(matrix[i])-1-i][j]
			}
			if highest < matrix[len(matrix[i])-1-i][len(matrix[i])-1-j] {
				highest = matrix[len(matrix[i])-1-i][len(matrix[i])-1-j]
			}
			possibleMaxSum += highest
		}
	}
	return possibleMaxSum
}

func main() {
	fmt.Println(flippingMatrix([][]int32{{1, 2}, {3, 4}})) // return 4
	fmt.Println(flippingMatrix([][]int32{
		{1, 2, 3, 4},
		{3, 4, 5, 6},
		{5, 6, 7, 8},
		{7, 8, 9, 10}})) // return 10 + 9 + 8 + 7 = 34
}
