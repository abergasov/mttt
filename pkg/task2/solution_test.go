package task2_test

import (
	"mttt/pkg/task2"
	"testing"
)

type tCase struct {
	A        []int
	Expected int
}

func TestSolution(t *testing.T) {
	table := []tCase{
		{
			A:        []int{130, 191, 200, 10},
			Expected: 140,
		},
		{
			A:        []int{405, 45, 300, 300},
			Expected: 600,
		},
		{
			A:        []int{50, 222, 49, 52, 25},
			Expected: -1,
		},
		{
			A:        []int{30, 909, 3190, 99, 3990, 9009},
			Expected: 9918,
		},
	}
	for _, tc := range table {
		actual := task2.Solution(tc.A)
		if actual != tc.Expected {
			t.Errorf("Expected %d, got %d", tc.Expected, actual)
		}
	}
}
