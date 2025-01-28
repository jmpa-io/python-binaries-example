package solution

// Find the highest power of 2 that divides N.
func Solution(N int) int {
	var K int
	for N%2 == 0 {
		N /= 2
		K++
	}
	return K
}
