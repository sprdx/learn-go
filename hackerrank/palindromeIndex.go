package main

/*
Algorithm:
1. Determine value of last index and median of string
2. Do looping until median of string
3. Check if index i element isn't equal to index (lastIndex - i) element
4. if True, check if index i isn't equal to index (lastIndex - i -1) OR index (lastIndex -i - 2) isn't equal to index (i +1)
	4a. if True, check if i is median AND length of string is odd
		4a1. if True, return -1
		4a2. if False, return i
	4b. if False, check if i is median AND length of string is odd
		4b1. if True, return -1
		4b2. if False, return lastIndex - i
5. if False, return -1
*/

func palindromeIndex(s string) int32 {
	var lastIndex int = len(s) - 1
	var medIndex int
	if len(s)%2 != 0 {
		medIndex = len(s) / 2
	} else {
		medIndex = len(s)/2 - 1
	}

	for i := 0; i <= medIndex; i++ {
		if s[i] != s[lastIndex-i] {
			if s[i] != s[lastIndex-i-1] || s[lastIndex-i-2] != s[i+1] {
				if i == medIndex && len(s)%2 != 0 {
					return -1
				} else {
					return int32(i)
				}
			} else {
				if i == medIndex && len(s)%2 != 0 {
					return -1
				}
				return int32(lastIndex - i)
			}
		}
	}
	return -1
}
