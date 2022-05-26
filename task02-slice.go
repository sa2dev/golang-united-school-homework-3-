package homework

func reverse(input []int64) (result []int64) {
	result = make([]int64, len(input), cap(input))
	for i, v := range input {
		result[len(input)-i-1] = v
	}
	return result
}
