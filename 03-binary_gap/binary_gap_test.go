package main

import "testing"

// go test . -v --bench  --benchmem -timeout 30s main
// go test . -v --bench  --benchmem -timeout 30s -run ^binary_gap$ main
func TestBinary_Gap(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			if actual := Binary_Gap(tc.input); actual != tc.expected {
				t.Fatalf("Binary_Gap(%d) = %d, want: %d", tc.input, actual, tc.expected)
			}
		})
	}
}

func BenchmarkBinary_Gap(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			Binary_Gap(tc.input)
		}
	}
}
