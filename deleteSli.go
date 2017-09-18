package main

import "fmt"

func deleteSli(param []int, s int, t int) {
	var cou int
	var size = len(param)
	param2 := param[0:s]
	param3 := param[t+1 : len(param)]
	param = append(param2, param3...)
	for cou = 0; cou < (size - t + s - 1); cou++ {
		fmt.Printf("%d", param[cou])
	}
}

func main() {
	var balance = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	deleteSli(balance, 3, 6)

}
