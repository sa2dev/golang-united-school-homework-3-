package homework

import (
	"sort"
)

func sortMapValues(input map[int]string) (result []string) {
	l := len(input)
	result = make([]string, 0, l)

	keys := make([]int, 0, l)
	for k := range input {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, d := range keys {
		result = append(result, input[d])
	}
	return result
}
