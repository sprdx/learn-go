package main

/*
Algorithm:
1. Check if shift number, k, greater than or equal 26
2. If true, modulo k with lenght of alphabet, 26
3. Set empty variable for result
4. Do looping as much as length of string that will be encrypted
5. Check if character of string between in range 65 and 90
6. If true, check if rune of character, e,  plus shift number, k, is greater than 90
	6a. if true, add string of rune (64 + e + k - 90) to result variable
	6b. if false, add string of rune e + k
9. If false, Check if character of string between in range 97 and 122
10. If true, check if rune of character, e,  plus shift number, k, is greater than 122
	10a. if true, add string of rune (96 + e + k - 122) to result variable
	10b. if false, add string of rune e + k
11. If false, add string of rune e
*/

//  Caesar's cipher shifts each letter by a number of letters.
func caesarCipher(s string, k int32) string {
	// Write your code here
	if k >= 26 {
		k = k % 26
	}

	en := ""

	for _, e := range s {
		if e >= 65 && e <= 90 {
			if e+k > 90 {
				en += string(64 + (e + k - 90))
			} else {
				en += string(e + k)
			}

		} else if e >= 97 && e <= 122 {
			if e+k > 122 {
				en += string(96 + (e + k - 122))
			} else {
				en += string(e + k)
			}
		} else {
			en += string(e)
		}
	}
	return en
}

// func main() {
// 	fmt.Println(caesarCipher("abcd", 3)) // return "defg"
// }
