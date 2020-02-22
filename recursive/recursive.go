package recursive

func Fac(n int) int {
	if n <= 1 {
		return 1
	}
	return (n * Fac(n-1))
}
func Pow(x, n int) int {
	if n == 1 {
		return x
	} else {
		return x * Pow(x, n-1)
	}
}
func Pow1(x, n int) int {
	result := 1

	// multiply result by x n times in the loop
	for i := 1; i < n; i++ {
		result *= x
	}

	return result
}
