package main

import "fmt"



func selectSort(array []int) {
	for i := 0; i < len(array); i++ {
		min := i
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[min] {
				min = j
			}
		}
		array[min],array[i]= array[i],array[min]
		fmt.Println(array)

	}
}
//test
func main() {
	array := []int{6, 1, 2, 7, 9, 3, 4, 5, 10, 8,}
	selectSort(array)

}