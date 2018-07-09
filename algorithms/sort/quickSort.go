package main

import "fmt"

func quickSort(array []DATA, left int, right int) {
	if left < right {

		// 设置分水岭
		mid := array[left]

		// 设置哨兵
		head, tail := left, right
		for {
			// 从右向左找，找到第一个比分水岭小的数
			for array[tail] >= mid && head < tail {
				tail--
			}

			// 从左向右找，找到第一个比分水岭大的数
			for array[head] <= mid && head < tail {
				head++
			}

			// 如果哨兵相遇，则退出循环
			if head >= tail {
				break
			}

			// 交换左右两侧的值
			array[head], array[tail] = array[tail], array[head]
			fmt.Println(array)
		}

		// 将分水岭移到哨兵相遇点
		array[left] = array[head]
		array[head] = mid

		// 递归，左右两侧分别排序
		quickSort(array, left, head-1)
		quickSort(array, head+1, right)
	}
	fmt.Println(array)

}

func QuickSort(array []DATA) {
	quickSort(array, 0, len(array)-1)
}

//test
//func main() {
//	array := []int{6, 1, 2, 7, 9, 3, 4, 5, 10, 8,}
//	QuickSort(array)
//
//}
