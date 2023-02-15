package task1_test

import (
	"mttt/pkg/task1"
	"testing"
)

type tCase struct {
	T        []int
	A        []int
	Expected int
}

func TestSolution(t *testing.T) {
	table := []tCase{
		{
			T:        []int{0, 0, 1, 1},
			A:        []int{2},
			Expected: 3,
		},
		{
			T:        []int{0, 0, 0, 0, 2, 3, 3},
			A:        []int{2, 5, 6},
			Expected: 5,
		},
		{
			T:        []int{0, 0, 1, 2},
			A:        []int{1, 2},
			Expected: 3,
		},
		{
			T:        []int{0, 3, 0, 0, 5, 0, 5},
			A:        []int{4, 2, 6, 1, 0},
			Expected: 7,
		},
	}
	for _, tc := range table {
		actual := task1.Solution(tc.T, tc.A)
		if actual != tc.Expected {
			t.Errorf("Expected %d, got %d", tc.Expected, actual)
		}
	}
}
