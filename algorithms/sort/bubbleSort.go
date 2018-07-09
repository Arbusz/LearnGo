package main


func bubbleSort(array []DATA) {
	done := true //判断循环是否该终止
	for i := 0; i < len(array) && done; i++ {
		done = false
		for j := 0; j < len(array)-i-1; j++ {
			done = true
			if array[j] > array[j+1] {
				array[j+1],array[j] = array[j],array[j+1]
			}
		}
	}
}
