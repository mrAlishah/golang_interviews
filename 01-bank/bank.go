package main

import "fmt"

func Bank(R string, V []int) []int {
	//Arrange
	A, B := 0, 0

	// Validation
	if R[0] == 'A' {
		B = V[0]
	} else {
		A = V[0]
	}

	if len(R) > 1 {
		if R[0] == 'A' {
			A = V[1]
		} else {
			B = V[1]
		}
	}

	//Logic
	balanceA, balanceB := A, B
	minBalance := false
	Print := func(r byte, v int) {
		fmt.Printf("+%d=>%c => A=%d, B=%d\n", v, r, balanceA, balanceB)
	}
	Print('-', 0)
	for i := 0; i < len(R); i++ {
		switch R[i] {
		case 'A':
			{
				balanceA += V[i]
				balanceB -= V[i]
				Print('A', V[i])
				if balanceB < 0 {
					balanceA += balanceB
					balanceB = 0
					minBalance = true
					break
				}
			}
		case 'B':
			{
				balanceB += V[i]
				balanceA -= V[i]
				Print('B', V[i])
				if balanceA < 0 {
					balanceB += balanceA
					balanceA = 0
					minBalance = true
					break
				}

			}
		} //switch
	} //for

	//Result
	Print('-', 0)
	if minBalance {
		return []int{balanceA, balanceB}
	} else {
		return []int{A, B}
	}
}
