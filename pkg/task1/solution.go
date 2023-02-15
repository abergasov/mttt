package task1

func Solution(T []int, A []int) int {
	// n := len(T)
	learned := make(map[int]bool)
	toLearn := make([]int, 0)
	for _, skill := range A {
		if learned[skill] {
			continue
		}
		toLearn = append(toLearn, skill)
		for len(toLearn) > 0 {
			curr := toLearn[len(toLearn)-1]
			toLearn = toLearn[:len(toLearn)-1]
			if learned[curr] {
				continue
			}
			learned[curr] = true
			toLearn = append(toLearn, T[curr])
		}
	}
	return len(learned)
}
