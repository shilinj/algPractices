package algorithm

import "fmt"

/*
	快速排序
1. 从数组中挑出一个元素作为基准
2. 重新排列数列，把所有的比基准小的放在基准前面，反之放在后面（一样大可任意一边）完成后基准处在分区的中间位置
3. 通过递归调用把小于基准元素和大于基准元素的子序列进行排序
*/

func quickSort(array []int, left, right int) {
	if left >= right { //一定是left >= right
		return
	}
	temp := array[left]
	start := left
	stop := right
	for right != left {
		for right > left && array[right] >= temp {
			right--
		}
		for left < right && array[left] <= temp {
			left++
		}
		if right > left {
			array[right], array[left] = array[left], array[right]
		}
		fmt.Printf("left=%v, right=%v, arrar=%v\n", left, right, array)
	}
	array[right], array[start] = temp, array[right]
	quickSort(array, start, left)
	quickSort(array, right+1, stop)
}

func TestQuickSort() {
	array := []int{15, 18, 62, 9, 54, 32, 13, 10, 6, 3}
	quickSort(array, 0, len(array)-1)
	fmt.Println(array)
}
