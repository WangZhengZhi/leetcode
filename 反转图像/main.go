package main

import (
	"fmt"
)

func main() {
	A := [][]int{
		{1, 1, 0},
		{1, 0, 1},
		{0, 0, 0},
	}

	fmt.Println(flipAndInvertImage(A))
}
func flipAndInvertImage(A [][]int) [][]int {
	n := len(A)
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			A[i][j], A[i][n-1-j] = A[i][n-1-j], A[i][j]

		}

	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if A[i][j] == 0 {
				A[i][j] = 1

			} else {
				A[i][j] = 0
			}

		}

	}
	return A

}
