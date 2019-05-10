package main

import (
	"fmt"
)

/**
 * @brief      { reads and integer ans checks for errors, an rune input is an error }
 *
 * @param      name  The name of the variable, for error log
 *
 * @return     { the readed value }
 */
func readcheck(name string) int {
	var aux int
	_, err := fmt.Scanf("%d", &aux)
	if err != nil {
		fmt.Println("There is a problem to read the variable ", name)
		return -1
	}
	return aux
}

func main() {
	// path := make([][]int, 192)
	var landmarks int
	var friends int

	landmarks = readcheck("landmarks")
	friends = readcheck("friends")
	fmt.Println(landmarks)
	fmt.Println(friends)
}
