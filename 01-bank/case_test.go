package main

var testCases = []struct {
	description string
	input1      string
	input2      []int
	expected    []int
}{
	{
		description: "R BAABA  => [2,4]",
		input1:      "BAABA",
		input2:      []int{2, 4, 1, 1, 2},
		expected:    []int{2, 4},
	},
	{
		description: "R ABAB  => [0,15]",
		input1:      "ABAB",
		input2:      []int{10, 5, 10, 15},
		expected:    []int{0, 15},
	},
	{
		description: "R B  => [100,0]",
		input1:      "B",
		input2:      []int{100},
		expected:    []int{100, 0},
	},
}
