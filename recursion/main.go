package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//104. 二叉树的最大深度
// 对月每个节点而言， 当前的深度 等于 max(左，右） + 1
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

// 110. 平衡二叉树
func isBalanced(root *TreeNode) bool {
	return height(root) >= 0
}
func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftHeight := height(root.Left)
	rightHeight := height(root.Right)
	// 即使止损 有一个不平衡了 就不用判断了
	if leftHeight == -1 || rightHeight == -1 || abs(leftHeight-rightHeight) > 1 {
		return -1
	}
	return max(leftHeight, rightHeight) + 1
}

// 求最大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// 求绝对值
func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

//青蛙跳台阶问题
func numWays(n int) int {
	if n == 0 {
		return 1
	}
	if n <= 2 {
		return n
	}
	pre, pre2 := 2, 1
	for i := 3; i <= n; i++ {
		pre2, pre = pre, (pre+pre2)%1000000007
	}
	return pre
}

// 学完二叉树章节回顾剩下题目
func main() {
	fmt.Println(numWays(4))
}
