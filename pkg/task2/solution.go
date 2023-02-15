package task2

import "math"

func Solution(A []int) int {
	m := make(map[int][]int)
	for i := 0; i < len(A); i++ {
		firstDigit := A[i] / int(math.Pow10(int(math.Log10(float64(A[i])))))
		lastDigit := A[i] % 10
		key := firstDigit*10 + lastDigit
		m[key] = append(m[key], A[i]) // key to idetify next possible candidats
	}

	maxSum := -1
	for _, values := range m {
		n := len(values)
		for i := 0; i < n-1; i++ {
			for j := i + 1; j < n; j++ {
				sum := values[i] + values[j]
				if sum > maxSum {
					maxSum = sum
				}
			}
		}
	}
	return maxSum
}
