package main

/*
Algorithm:
1. However much of tower(n), if height(m) is 1, player2 will win.
2. Then, check if number of towers is even or odd. if odd, then the player1 will win.

*/

// towerBreakers return 1 if player1 wins and return 2 if player2 wins
func towerBreakers(n int32, m int32) int32 {

	if m == 1 {
		return 2
	}
	if n%2 != 0 {
		return 1
	}
	return 2
}

// func main() {
// 	fmt.Println(towerBreakers(11, 10)) // return 1
// }
