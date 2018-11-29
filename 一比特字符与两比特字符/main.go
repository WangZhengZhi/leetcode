package main

import (
	"fmt"
)

func main() {
	bits := []int{1, 0, 0}
	fmt.Println(isOneBitCharacter(bits))
}
func isOneBitCharacter(bits []int) bool {
	var i int
	for i < len(bits)-1 {
		i += bits[i] + 1
		// if bits[i]==0 i++,else i+=2

	}
	return i == len(bits)-1
}
