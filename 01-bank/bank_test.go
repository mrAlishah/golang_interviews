package main

import "testing"

// go test . -v --bench  --benchmem -timeout 30s main
// go test . -v --bench  --benchmem -timeout 30s -run ^bank$ main
func TestBank(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			if result := Bank(tc.input1, tc.input2); result[0] != tc.expected[0] || result[1] != tc.expected[1] {
				t.Fatalf("Bank(%s,%v) = %v, want: %v", tc.input1, tc.input2, result, tc.expected)
			}
		})
	}
}

func BenchmarkBank(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			Bank(tc.input1, tc.input2)
		}
	}
}
