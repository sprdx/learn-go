package main

import (
	"fmt"
	"sort"
)

// algorithm
// 1. sort array of integer with length array is odd
// 2. assume array has been sorted
// 3. if lonely integer is last element, then function return the integer
// 4. do looping for i less than length of array minus 1
// 5. if element i-st/nd/th is not match the next, then the element is lonely integer
// 6. if false, do plus 2 to i, and do looping again

func lonelyinteger(a []int32) int32 {

	sort.SliceStable(a, func(i, j int) bool { return a[i] < a[j] })

	for i := 0; i < len(a)-1; i += 2 {
		if a[i] != a[i+1] {
			return a[i]
		}
	}
	return a[len(a)-1]
}

func main() {
	fmt.Println(lonelyinteger([]int32{1, 3, 2, 1, 2})) // result 3
}
