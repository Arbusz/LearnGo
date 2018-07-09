package main

func mergeSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}
	mid := len(array) / 2
	left := mergeSort(array[:mid])
	right := mergeSort(array[mid:])

	return merge(left, right)

}

func merge(left, right []int) (result []int) {
	lp, rp := 0, 0
	for lp < len(left) && rp < len(right) {
		if left[lp] > right[rp] {
			result = append(result, right[rp])
			rp ++
		} else {
			result = append(result, left[lp])
			lp++
		}
	}
	result = append(result, left[lp:]...)
	result = append(result, right[rp:]...)

}
