package main

import "fmt"

func main() {
	var balance = []int{2, 9, 3, 6, 8}
	var min int
	min = deleteMin(balance)

	fmt.Printf("最小值是 %d", min)
}

func deleteMin(param []int) int {
	var i, k, c int
	var size = len(param)
	var j = param[0]
	if size > 1 {
		for i = 0; i < size; i++ {
			if j > param[i] {
				j = param[i]
				c = i
			}
		}
		k = param[c]
		param[c] = param[size-1]
		param = param[0 : size-1]
		/*for i = 0; i < len(param); i++ {
			fmt.Printf("%d", param[i])
		}*/
		fmt.Println(param)
		return k
	} else {
		k = j
		param[0] = 0
		return k
	}
}
