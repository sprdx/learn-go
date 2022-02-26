package main

import (
	"fmt"
	"strconv"
)

// timeConversion returns a new string representing the input time in 24 hour format.
func timeConversion(s string) string {
	// Check if input is AM or PM
	if s[len(s)-2:] == "AM" {
		// Check if hour is 12
		if s[:2] == "12" {
			// if yes convert 12 to 00 and remove AM
			return "00" + s[2:len(s)-2]
		} else {
			// if no only remove AM
			return s[:len(s)-2]
		}
	} else {
		// if input is PM, check if hour is 12
		if s[:2] == "12" {
			// if yes, remove PM
			return s[:len(s)-2]
		} else {
			// if no, add 12 to hour and remove PM
			hour, _ := strconv.Atoi(s[:2])
			hour24 := hour + 12
			return strconv.Itoa(hour24) + s[2:len(s)-2]
		}
	}
}

func main() {
	fmt.Println(timeConversion("12:03:04AM")) // 00:03:04
}
