package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("first: ", i)
	}

	// add break
	fmt.Println("----------")
	for i := 0; i < 5; i++ {
		if i > 1 {
			break
		}
		fmt.Println("second: ", i)
	}

	// add continue
	fmt.Println("----------")
	for i := 0; i < 5; i++ {
		if i == 1 {
			continue
		}
		fmt.Println("third: ", i)
	}

	// like while loops in another languages
	fmt.Println("----------")
	i := 0
	for i < 5 {
		i++
		if i < 2 {
			continue
		}
		fmt.Println("fourth: ", i)

	}
}
