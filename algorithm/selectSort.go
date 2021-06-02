package algorithm

import "fmt"

/*
	选择排序
选择第一个元素作为最小数，与其他元素进行比较，如果其他元素小于该数，那么最小值的就是其他元素（注意索引的变化）
*/

func selectionSort(arr []int) {
	length := len(arr)
	for i := 0; i < length; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i] // 将最小值和初始参考值交换位置
		fmt.Printf("第%v次选择：%v\n", i+1, arr)
	}
	fmt.Println(arr)
}

func TestSelectionSort() {
	array := []int{15, 18, 62, 9, 54, 32, 13, 10, 6, 3}
	selectionSort(array)
}
