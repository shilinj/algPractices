package algorithm

import "fmt"

/*
	选择排序
1. 在一个长度为 N 的无序数组中，第一次遍历 n-1 个数找到最小的和第一个数交换
2. 第二次从下一个数开始遍历 n-2 个数，找到最小的数和第二个数交换
3. 重复以上操作直到第 n-1 次遍历最小的数和第 n-1 个数交换，排序完成
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
