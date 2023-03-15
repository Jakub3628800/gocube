package utils

// return n choose k
func Comb(n int, k int) int {
	if k == 0 {
		return 1
	}
	return (n * Comb(n-1, k-1)) / k
}
