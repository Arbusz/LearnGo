package main

import "fmt"

func main() {
	var balance = []int{2, 9, 3, 6, 8}
	var min int
	min = deleteMin(balance, 5)

	fmt.Printf("最小值是 %d", min)
}

func deleteMin(param []int, size int) int {
	var i, k, c int
	var j = param[0]
	if size > 1 {
		for i = 0; i < size; i++ {
			if j > param[i] {
				j = param[i]
				c = i
			}
		}
		k = param[c]
		for j = i; j < size-1; j++ {
			param[j] = param[j+1]
		}
		return k
	} else {
		k = j
		param[0] = 0
		return k
	}
}
