package list_by_stack

// 用栈实现队列  232
type MyQueue struct {
	Current  []int
	Exchange []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		Current:  []int{},
		Exchange: []int{},
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	var currentLen, exchangeLen int
	// 先将当前的数据全部入到交换栈
	for len(this.Current) != 0 {
		currentLen = len(this.Current)
		this.Exchange = append(this.Exchange, this.Current[currentLen-1])
		this.Current = this.Current[:currentLen-1]
	}

	this.Current = append(this.Current, x)
	// 再将交换机的数据放至当前栈
	for len(this.Exchange) != 0 {
		exchangeLen = len(this.Exchange)
		this.Current = append(this.Current, this.Exchange[exchangeLen-1])
		this.Exchange = this.Exchange[:exchangeLen-1]
	}
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	current := this.Current[len(this.Current)-1]
	this.Current = this.Current[:len(this.Current)-1]
	return current

}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	return this.Current[len(this.Current)-1]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.Current) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
