package utils

// return n choose k
func Comb(n int, k int) int {
	if k == 0 {
		return 1
	}
	return (n * Comb(n-1, k-1)) / k
}

// convert array of numbers in specific base to decimal
func BaseToDec(arr []int, base int) int {
	result := 0
	multiplier := 1

	for i := 0; i < len(arr); i++ {
		result += arr[i] * multiplier
		multiplier = multiplier * base
	}
	return result
}

// convert array of numbers in specific base to decimal
func BaseToTer(arr []int, base int) int {
	result := 0
	multiplier := 1

	for i := 0; i < len(arr); i++ {
		result += arr[i] * multiplier
		multiplier = multiplier * base
	}
	return result
}
