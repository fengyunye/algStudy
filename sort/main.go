package main

import (
	"fmt"
)

// 基于比较 从小到大
// 冒泡排序 空间复杂度 O(1) 时间复杂度
func BubbleSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		// 经过一轮反复比较， 一个数已排在最末尾位置, 每次从0开始， 需要冒泡的次数 n - 1 - 已排好的数量
		// 可以优化，如果一轮中一次交换都没有，说明已经是有序数列了
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

// 插入排序
// 插入排序 空间复杂度O(1)
func InsertSort(arr []int) []int {
	// 分为有序空间和无序空间
	// 每次从无序空间中取数与有序空间进行比较
	if len(arr) < 1 {
		return arr
	}
	// 从第一个数
	for i := 1; i < len(arr); i++ {
		// 与有序数列依次比较
		current := arr[i]
		j := i
		for ; j >= 0; j-- {
			if current < arr[j] {
				// 这些数比要插入数的大，往后移动一位
				arr[j+1] = arr[j]
			} else {
				// 因为一直是有序数列，如果中间有一个数是比当前数小，后面的数就都比这个数小了, 而此时的j 就是需要插入的位置
				break
			}
		}
		arr[j] = current
	}
	return arr
}

// 选择排序
// 和插入排序方法类似， 分为两个序列，一个有序，一个无序，每次从无序中找最小的
//  每次都叫唤位置  5  8 5 2    5 -1  5 -2
//  2 8 5 5   5 - 1 就跑到 5 - 2  后面去了   就不稳定了
func SelectSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		min := i
		// 剩下的数中找一个最小的
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = arr[j]
			}
		}
		// 进行交换
		arr[min], arr[i] = arr[i], arr[min]
	}
	return arr
}

// 归并排序
// 分治是思想
// 递归是编程技巧
func MergeSort(arr []int) []int {
	// 确定终止条件
	length := len(arr)
	if length < 2 {
		return arr
	}
	// 分成两块
	middle := length / 2
	left := arr[0:middle]
	right := arr[middle:]
	return Merge(MergeSort(left), MergeSort(right))
}

// 将两个有序的数组合并成一个有序的数组
func Merge(left []int, right []int) []int {
	var result []int
	for len(left) != 0 && len(right) != 0 {
		if left[0] > right[0] {
			result = append(result, right[0])
			right = right[1:]
		} else {
			result = append(result, left[0])
			left = left[1:]
		}
	}
	if len(left) > 0 {
		result = append(result, left...)
	}

	if len(right) > 0 {
		result = append(result, right...)
	}
	return result
}

// 快速排序，思路与快速排序一直
func QuickSort(arr []int) []int {
	return arr
}

func main() {
	fmt.Println("今天来搞搞排序")
	sortArr := []int{2, 9, 3, 4, 8, 3}
	fmt.Println(BubbleSort(sortArr))
	fmt.Println(InsertSort(sortArr))
	fmt.Println(SelectSort(sortArr))
	fmt.Println(MergeSort(sortArr))
}
