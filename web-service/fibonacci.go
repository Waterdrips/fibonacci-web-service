package main

func fibonacciSequence(n int) []int {
	if n <= 0 {
		return []int{}
	}

	sequence := make([]int, n)
	sequence[0] = 0
	if n > 1 {
		sequence[1] = 1
	}
	for i := 2; i < n; i++ {
		sequence[i] = sequence[i-1] + sequence[i-2]
	}
	return sequence
}
