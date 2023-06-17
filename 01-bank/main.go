package main

import "fmt"

func main() {

	fmt.Println(Bank("BAABA", []int{2, 4, 1, 1, 2})) //[2,4]
	fmt.Println(Bank("ABAB", []int{10, 5, 10, 15}))  //[0,15]
	fmt.Println(Bank("B", []int{100}))               //[100,0]
}
