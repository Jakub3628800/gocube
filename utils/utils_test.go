package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChoose(t *testing.T) {
	r := Comb(42, 2)
	if r != 861 {
		t.Fatalf("4 over 5 should be 10 but was")
	}
	assert.Equal(t, 861, r, "The two words should be the same.")
}

func TestBaseToDec(t *testing.T) {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7}
	r := BaseToDec(arr, 3)

	assert.Equal(t, 21324, r, "The two words should be the same.")
}
