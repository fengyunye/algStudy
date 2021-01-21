package main

import (
	"fmt"
	"sort"
	"strings"
)

// 两个数组求交集
func intersect(nums1 []int, nums2 []int) []int {
	numMap := map[int]int{}
	for _, v := range nums1 {
		numMap[v] += 1
	}
	k := 0
	for _, v := range nums2 {
		if numMap[v] > 0 {
			numMap[v] -= 1
			nums2[k] = v
			k++
		}
	}
	return nums2[0:k]
}

// 两个有序数组求交集  - 350
func intersectSort(nums1 []int, nums2 []int) []int {
	// 节省空间使用 nums1 或者nums2 来节省空间
	i, j, k := 0, 0, 0
	sort.Ints(nums1)
	sort.Ints(nums2)
	// 因为是排好序的，所以当一个比较小的时候，往后移动指针即可
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			i++
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			// 相等的话同时移动
			nums1[k] = nums1[i]
			i++
			j++
			k++
		}
	}
	return nums1[:k]
}

//编写一个函数来查找字符串数组中的最长公共前缀。
//如果不存在公共前缀，返回空字符串 ""。 - 14
// 解法 1 简单暴力循环匹配

func longestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
		return ""
	}
	for j := 0; j < len(strs[0]); j++ {
		currentStr := strs[0][j : j+1]
		for i := 1; i < len(strs); i++ {
			if len(strs[i]) < j+1 || strs[i][j:j+1] != currentStr {
				return strs[0][:j]
			}
		}
	}
	return strs[0]
}

//给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。

//设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。

//注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

func maxProfit(prices []int) int {
	n := len(prices)
	//dp := make([][2]int, n)
	//dp[0][1] = -prices[0]
	//// 0 - 没有股票的收益
	//// 1 - 持有股票的收益
	//
	//// 在最后一天的时候,手上没有股票的收益最大
	//// 从第二天开始计算
	//// 保证每一天的收益都是最大话
	//// 动态规划， 当天没有股票， 要么是前一天也没有，或者是今天卖掉了,这两种选择一种
	//
	//for i := 1; i < n; i++ {
	//	dp[i][0] = max(dp[i-1][0], dp[i - 1][1]+prices[i])   // 当天没有股票， 要么是前一天有，或者是今天卖掉了
	//	dp[i][1] = max(dp[i-1][0]-prices[i], dp[i-1][1]) // 第二天有股票，要么是前一天没有，今天买入， 要么是前一天也有
	//}
	//return dp[n-1][0]
	dp0 := 0
	dp1 := -prices[0]

	for i := 1; i < n; i++ {
		dp0 = max(dp0, dp1+prices[i])
		dp1 = max(dp0-prices[i], dp1)
	}
	return dp0
}

// 给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。
// 根据规律反转
func rotate(nums []int, k int) {
	arrLen := len(nums)
	reverse(nums)                    // 反转整个数组
	reverse(nums[:k%arrLen])         // 反转前半部分
	reverse(nums[k%arrLen : arrLen]) // 反转后半不服
}

// 反转数组
func reverse(arr []int) {
	arrLen := len(arr)
	for i := 0; i < arrLen/2; i++ {
		arr[i], arr[arrLen-1-i] = arr[arrLen-1-i], arr[i]
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//
//给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。

//不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。

//元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。
// - 27
func removeElement(nums []int, val int) int {
	var slow int
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] == val {
			continue
		}
		nums[slow], nums[fast] = nums[fast], nums[slow]
		slow++
	}
	return slow
}

// 26 移除重复湘
// 快慢指针
// 快指针遇到重复项跳过
// 快慢一直交换，可以将重复项放到最后
func removeDuplicates(nums []int) int {
	slow := 0
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow], nums[fast] = nums[fast], nums[slow]
		}
	}
	return slow + 1
}

//给定一个由 整数 组成的 非空 数组所表示的非负整数，在该数的基础上加一。
//
//最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
//
//你可以假设除了整数 0 之外，这个整数不会以零开头。 - 66
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		temp := digits[i] + 1
		digits[i] = temp % 10
		if digits[i] != 0 {
			// 只要不等于0就没有金威，可以直接返回了，
			return digits
		}
	}
	// 如果走到这里，说明一直在金伟
	return append([]int{1}, digits...)
}

// 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 的那 两个 整数，并返回它们的数组下标。
//
//你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		currentValue := target - nums[i]
		for j := 1 + i; j < len(nums); j++ {
			if currentValue == nums[j] {
				return []int{i, j}
			}
		}
	}
	return []int{0, 0}
}

// 使用map优化，节省遍历次数
func twoSumV2(nums []int, target int) []int {
	var result []int
	numMap := make(map[int]int)
	for key, value := range nums {
		if k, ex := numMap[target-value]; ex {
			return []int{k, key}
		}
		numMap[value] = key
	}
	return result
}

//给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。
// 找数第一次想法就是快慢指针
func threeSum(nums []int) [][]int {
	// 排序之后使用快慢指针
	sort.Ints(nums)
	var result [][]int
	// 固定一个数， 然后 左指针 + 右指针寻找
	for i := 0; i < len(nums); i++ {
		current := nums[i]
		left := i + 1
		right := len(nums) - 1
		// 需要处理重复值 选定值重复的话，需要跳过
		if i == 0 || nums[i] != nums[i-1] {
			for left < right {
				res := current + nums[left] + nums[right]
				if res < 0 {
					left++
				} else if res > 0 {
					right--
				} else {
					result = append(result, []int{nums[i], nums[left], nums[right]})
					// 左指针和右指针也需要判断
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					for left < right && nums[right] == nums[right-1] {
						right--
					}
					left++
					right--
				}
			}
		}

	}
	return result
}

// 纯规律型

//LEETCODEISHIRING
//
//3
//
//L   C   I   R
//E T O E S I I G
//E   D   H   N
//
//0
//1
//2   3  4
//1
//0
//1
//2
//1
//LCIRETOESIIGEDHN
//4
//
//L     D    R
//E   O E  I I
//E C   I H  N
//T     S    G
//LDREOEIIECIHNTSG
func convert(s string, numRows int) string {
	length := len(s)
	if length <= numRows {
		return s
	}
	stringByte := []byte(s)
	stringSince := make([]string, numRows)
	// 按周期 2n - 2
	cycle := 2*numRows - 2
	for i := 0; i < length; i++ {
		mod := i % cycle
		var index int
		if mod < numRows {
			// 小于行数 就是mod
			index = mod
		} else {
			// 大于行数 周长 - mod
			index = cycle - mod
		}
		stringSince[index] += string(stringByte[i])
	}
	fmt.Println(stringSince)
	return strings.Join(stringSince, "")
}

// 合并两个有序数组
func merge(nums1 []int, m int, nums2 []int, n int) {
	// 从最后一位开始迁移
	mCurrent, nCurrent, p := m-1, n-1, m+n-1

	for nCurrent >= 0 {
		if mCurrent >= 0 && nums1[mCurrent] > nums2[nCurrent] {
			nums1[p] = nums1[mCurrent]
			mCurrent--
		} else {
			nums1[p] = nums2[nCurrent]
			nCurrent--
		}
		p--
	}
	fmt.Println(nums1)

}

//斐波那契数
func fib(n int) int {
	// 递归解法效率太低
	//if n <= 1 {
	//	return n
	//}
	//return fib(n-1) + fib(n -2)
	if n <= 1 {
		return n
	}
	pre, prePre, current := 1, 0, 0
	for i := 2; i <= n; i++ {
		current = pre + prePre
		prePre = pre
		pre = current
	}
	return current
}
func main() {
	/**********************求交集****************************/
	//array1 := []int{1, 2, 3, 2, 1}
	//array2 := []int{2, 2, 1}
	//fmt.Println(res)
	/**********************走最长公共前缀****************************/
	//res := intersectSort(array1, array2)
	//strs := []string{"flower", "flow", "flight"}
	//res := longestCommonPrefix(strs)
	/**********************股票脉脉最佳实际****************************/
	//prices := []int{7, 1, 5, 3, 6, 4}
	//res := maxProfit(prices)
	//fmt.Println(res)
	/**********************翻转数组****************************/
	//array1 := []int{1, 2, 3, 4, 5, 6}
	//rotate(array1, 2)
	/**********************原地删除和移除重复项****************************/
	//array1 := []int{1, 2, 2, 3, 4, 5, 6, 7, 8}
	////removeElement(array1, 2)
	//removeDuplicates(array1)

	/**********************原地删除和移除重复项****************************/
	//array1 := []int{9, 9, 9, 9, 9}
	//fmt.Println(plusOne(array1))

	/**********************求和****************************/
	//array1 := []int{2, 5, 5, 11}
	//fmt.Println(twoSumV2(array1, 10))
	/**********************3数求和****************************/
	//array1 := []int{-1, 0, 1, 2, -1, -4}
	//fmt.Println(threeSum(array1))
	/**********************z字形变换****************************/
	//res := convert("PAYPALISHIRING", 3)
	//fmt.Println(res)
	/**********************合并两个有序数组****************************/
	//array1 := []int{1, 2, 3, 0, 0, 0}
	//array2 := []int{2, 3, 5}
	//merge(array1, 3, array2, 3)
	/****************************斐波那契数*****************************/
	fmt.Println(fib(2))
}
