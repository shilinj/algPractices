package algorithm

import "fmt"

func selectionSort(arr []int) {
	length := len(arr)
	for i := 0; i < length; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
		// fmt.Printf("第%v次选择：%v\n", i+1, arr)
	}
	fmt.Println(arr)
}

func TestSelectionSort() {
	array := []int{15, 18, 62, 9, 54, 32, 13, 10, 6, 3}
	selectionSort(array)
}
