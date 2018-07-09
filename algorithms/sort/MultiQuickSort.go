package main

import (
	"fmt"
	"sync"
)

//使用goroutine并发排序

type DATA int


var lock sync.WaitGroup

func MultiQuickSort(array []DATA) {
	lock.Add(1)
	Quicksort(array, 0, len(array)-1)
	lock.Wait()
}

func Quicksort(array []DATA, left int, right int) {
	defer lock.Done()

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

		lock.Add(2)

		// 递归，左右两侧分别排序
		go Quicksort(array, left, head-1)
		go Quicksort(array, head+1, right)
	}
	fmt.Println(array)

}

//test
func main() {
	array := []DATA{6, 1, 2, 7, 9, 3, 4, 5, 10, 8,}
	MultiQuickSort(array)

}
