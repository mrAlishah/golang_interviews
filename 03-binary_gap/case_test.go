package main

var testCases = []struct {
	description string
	input       int
	expected    int
}{
	{
		description: "checking 9",
		input:       9,
		expected:    2,
	},
	{
		description: "checking 529",
		input:       529,
		expected:    4,
	},
	{
		description: "checking 4",
		input:       4,
		expected:    -1,
	},
	{
		description: "checking 20",
		input:       20,
		expected:    1,
	},
	{
		description: "checking 15",
		input:       15,
		expected:    0,
	},
	{
		description: "checking 32",
		input:       32,
		expected:    -1,
	},
}
