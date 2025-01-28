package solution

// NOTES:

// A is a non-empty array consisting of N integers.

// The unique number occurs exactly ONCE in array A.

// A may contain multiple unique numbers; it must return the first one
// (i.e in the lowest position)

// The function should return -1 if there are no unique numbers in A.

func Solution(A []int) int {

	// create a map that counts the number of unique values.
	counter := make(map[int]int)
	for _, i := range A {
		counter[i]++
	}

	// determine first unique value in the counter map.
	for _, i := range A {
		if counter[i] == 1 {
			return i
		}
	}

	// return -1 on no unique numbers.
	return -1
}

func SolutionV2(A []int) int {

	// count all positive integers.
	counter := make(map[int]bool)
	for _, v := range A {
		if v > 0 {
			counter[v] = true
		}
	}

	// find the missing number.
	for i := 1; ; i++ {
		if !counter[i] {
			return i
		}
	}
}
