package utils

import (
	"testing"
)

func TestComb(t *testing.T) {
	testCases := []struct {
		n      int
		k      int
		result int
	}{
		{5, 0, 1},
		{5, 1, 5},
		{5, 2, 10},
		{5, 5, 1},
		{10, 3, 120},
	}

	for _, tc := range testCases {
		got := Comb(tc.n, tc.k)
		if got != tc.result {
			t.Errorf("Comb(%d, %d) = %d; want %d", tc.n, tc.k, got, tc.result)
		}
	}
}
