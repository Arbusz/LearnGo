package main

func shellSort(array []int) {
	for step := len(array) / 2; step >= 1; step /= 2 { //改变步长
		for i := step; i < len(array); i++ { //改变对调指针的位置
			if i >= step && array[i] < array[i-step] {
				array[i], array[i-step] = array[i-step], array[i]
			}
		}
	}
}
func main() {
	array := []int{6, 1, 2, 7, 9, 3, 4, 5, 10, 8,}
	shellSort(array)

}
