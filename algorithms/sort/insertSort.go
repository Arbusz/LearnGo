package main

import "fmt"

func insertSort(array []int) {
	for i := 1; i < len(array); i++ {
		for j := i; j > 0; j-- {
			if array[j-1] > array[j] {
				array[j-1], array[j] = array[j], array[j-1]
			} else {
				break
			}
		}
	}
	fmt.Println(array)
}

func main() {
	array := []int{6, 1, 2, 7, 9, 3, 11, 5, 10, 8,}
	InsertSort(array)
	fmt.Println(array)

}

//二分查找法优化选择排序
func binarySearch(array []int, k int) (mid int) {
	end := len(array) - 1
	start := 0
	mid = (end + start) / 2
	for count := 1; count <= len(array); count++ {
		if k == array[mid] {
			return
		} else if k > array[mid] {
			start = mid + 1
		} else {
			end = mid - 1
		}
		mid = (end + start) / 2
	}
	return
}

func InsertSort(array []int) {
	for i := 1; i < len(array); i++ {
		if array[i] < array[i-1] {
			location := binarySearch(array, array[i])
			if array[i] > array[location] {
				location += 1
			} else
			{
				location -= 1
			}
			temp := array[i]
			for j := i; j > location; j-- {
				array[j] = array[j-1]
			}
			array[location] = temp
		}
	}
	return
}
