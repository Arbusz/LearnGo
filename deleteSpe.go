package main

import "fmt"

func main() {
	var balance = []int{1, 4, 5, 6, 4}
	deleteSpe(balance, 5, 4)
}

func deleteSpe(param []int, size int, value int) {
	var i, j, count int
	for i = 0; i < size; i++ {
		if param[i] == value {
			count++
			param[i] = 0
			for j = i; j < size-count; j++ {
				param[j] = param[j+1]
			}
			j = 0
		}
	}
	for i = 0; i < size; i++ {
		fmt.Printf("%d", param[i])
	}

}
