package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	var sortedValues []string
	var keys []int
	for key := range input {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	for _, key := range keys {
		sortedValues = append(sortedValues, input[key])
	}
	return sortedValues
}
