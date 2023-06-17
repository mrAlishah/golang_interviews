package main

import (
	"fmt"
	"log"
	"strconv"
)

func Binary_Gap(N int) int {
	binarStr := strconv.FormatInt(int64(N), 2)
	fmt.Println(binarStr)
	distance := 0
	start := 0
	end := 0
	for i, a := range binarStr {
		switch {
		case a == '1':
			{
				div := end - start
				if distance < div {
					distance = div
				}
				start = i
			}
		case a == '0':
			{
				end = i
			}
		default:
			log.Fatalln("a =", a)
		}
	}

	if start == 0 {
		return -1
	}
	return distance
}
