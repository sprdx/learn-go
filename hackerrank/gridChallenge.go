package main

import "sort"

/*
Algorithm:
1. Make 2 variables that determine length of array and length of element/string
2. Do looping as long as length of array
3. Do looping as long as length of string minus 1
4. Check if char j of index i is greater than char j + 1 of index i
5. If true, make empty slice of string variable
	5a. Do looping as long as length of string to do append char of string to the slice variable
	5b. Do sorting the slice variable
	5c. Reassign index i element as empty string
	5d. Do looping as long as length of the slice variable to append the element of the slice to empty string
	5e. Break looping
6. Do looping as long as length of string
7. Do looping as long as length of array
8. Check if char of index n string is greater than char of index n + 1 string
	8a. If true return "NO"
9. Return "YES"
*/

func gridChallenge(grid []string) string {
	// Write your code here
	lenArray := len(grid)
	lenElement := len(grid[0])

	for i := 0; i < lenArray; i++ {
		for j := 0; j < lenElement-1; j++ {
			if grid[i][j] > grid[i][j+1] {
				charArr := []string{}
				for k := 0; k < lenElement; k++ {
					charArr = append(charArr, string(grid[i][k]))
				}
				sort.Strings(charArr)
				grid[i] = ""
				for l := 0; l < len(charArr); l++ {
					grid[i] += charArr[l]
				}
				break
			}
		}
	}
	for m := 0; m < lenElement; m++ {
		for n := 0; n < lenArray-1; n++ {
			if grid[n][m] > grid[n+1][m] {
				return "NO"
			}
		}
	}
	return "YES"
}
