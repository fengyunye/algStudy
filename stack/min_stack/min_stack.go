// 最小栈 155  用两个栈存储

package min_stack

import "math"

type MinStack struct {
	Stack []int
	Min   []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		Stack: []int{},
		Min:   []int{math.MaxInt64},
	}
}

func (this *MinStack) Push(x int) {
	this.Stack = append(this.Stack, x)
	minNum := this.Min[len(this.Min)-1]
	this.Min = append(this.Min, min(x, minNum))
}

func (this *MinStack) Pop() {
	this.Stack = this.Stack[:len(this.Stack)-1]
	this.Min = this.Min[:len(this.Min)-1]
}

func (this *MinStack) Top() int {
	return this.Stack[len(this.Stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.Min[len(this.Min)-1]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
