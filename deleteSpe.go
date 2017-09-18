package main

import "fmt"

func main() {
	var balance = []int{1, 4, 5, 6, 4}
	deleteSpe(balance, 4)
}

func deleteSpe(param []int, value int) {
	var i, j, count int
	var size = len(param)
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
	param = param[0 : size-count+1]
	for i = 0; i < size-count+1; i++ {
		fmt.Printf("%d", param[i])
	}

}
