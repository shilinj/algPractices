package algorithm

import "fmt"

/*
	冒泡排序
相邻的元素比较，如果a[i] > a[i+1], 那么二者交互位置；
数组有N个元素，则最多需要进行N次冒泡，且第N次冒泡选择出第N大的元素
*/

func bubbleSort(array []int) {
	length := len(array)
	if length < 2 {
		fmt.Println(array)
		return
	}

	for i := 0; i < length; i++ {
		for j := 0; j < length-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
			// fmt.Printf("\t第%v第%v次比较：%v\n", i+1, j+1, array)
		}
		fmt.Printf("第%v轮冒泡：%v\n", i+1, array)
	}

	fmt.Println(array)
}

func TestBubbleSort() {
	x := []int{15, 18, 62, 9, 54, 32, 13, 10, 6, 3}
	bubbleSort(x)
}
