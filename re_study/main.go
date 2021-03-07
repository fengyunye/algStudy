package main

import "fmt"

// 冒泡排序 空间复杂度 O(1) 时间复杂度
func bobbleSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length-1; i++ {
		flag := false
		// 经过一轮反复比较， 一个数已排在最末尾位置, 每次从0开始， 需要冒泡的次数 n - 1 - 已排好的数量
		// 可以优化，如果一轮中一次交换都没有，说明已经是有序数列了
		for j := 0; j < length-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = true
			}
		}

		if flag == false {
			break
		}
	}
	return arr
}

// 插入排序
// 有序空间和无序空间，
// 每次从无序空间选择一个数和有序空间进行比较，确定插入位置
func insertSort(arr []int) []int {
	length := len(arr)
	for i := 1; i < length; i++ {
		currentValue := arr[i]
		j := i - 1
		for ; j >= 0; j-- {
			if currentValue > arr[j] {
				break
			} else {
				arr[j+1] = arr[j]
			}
		}
		arr[j+1] = currentValue
	}
	return arr
}

// 有序空间和无序空间，
// 插入是 选择一个数，确定插入位置，然后插入， 重点是选择插入的位置
// 选择是 选择最小的数，和前一位兑换， 终点是选择最小的数
func selectSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		arr[min], arr[i] = arr[i], arr[min]
	}
	return arr
}

// 归并排序
// 把数组分为左，右两个数组， 保证左右是有序的，然后全部组合起来依然有序
func mergeSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}
	middle := length / 2
	left := arr[:middle]
	right := arr[middle:]
	return mergeArr(mergeSort(left), mergeSort(right))
}
func mergeArr(leftArr, rightArr []int) []int {
	var result []int
	for len(leftArr) != 0 && len(rightArr) != 0 {
		if leftArr[0] > rightArr[0] {
			result = append(result, rightArr[0])
			rightArr = rightArr[1:]
		} else {
			result = append(result, leftArr[0])
			leftArr = leftArr[1:]
		}
	}
	if len(leftArr) > 0 {
		result = append(result, leftArr...)
	}

	if len(rightArr) > 0 {
		result = append(result, rightArr...)
	}

	return result
}

// 快速排序
// 与归并排序的不同在于，避免一个新的数组空间
// 将数据分为 l  p  r
func quickSort(arr []int) []int {
	return _quickSort(arr, 0, len(arr)-1)
}
func _quickSort(arr []int, left, right int) []int {
	for right > left {
		partIndex := part(arr, left, right)
		_quickSort(arr, left, partIndex-1)
		_quickSort(arr, partIndex+1, right)
	}
	return arr
}

// 选出基准点的同时交换数据
// 要保证不使用新的数组接受亦能交换
// 双指针
// i, j
func part(arr []int, left, right int) int {
	// 使用选择
	pivot := right
	i := left
	// 双指针
	// j 代表数组当前指针
	//  0 ~ i - 1 达标已处理空间
	//
	for j := i; j < right; i++ {
		if arr[pivot] > arr[j] {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	// i 代表基准点真的位置
	arr[i], arr[pivot] = arr[pivot], arr[i]
	return i
}

// 交换
func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func main() {
	arrayInt := []int{456, 526, 4, 2, 54}
	fmt.Println(mergeSort(arrayInt))
}
