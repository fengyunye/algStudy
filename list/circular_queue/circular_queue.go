package circular_queue

type MyCircularQueue struct {
	Items  []int
	Length int
	Head   int
	Tail   int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		Items:  make([]int, k+1),
		Length: k + 1,
		Head:   0,
		Tail:   0,
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if (this.Tail+1)%this.Length == this.Head {
		return false
	}
	this.Items[this.Tail] = value
	this.Tail = (this.Tail + 1) % this.Length
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.Head == this.Tail {
		return false
	}
	this.Head = (this.Head + 1) % this.Length
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.Head == this.Tail {
		return -1
	}
	return this.Items[this.Head]
}

func (this *MyCircularQueue) Rear() int {
	if this.Head == this.Tail {
		return -1
	}
	n := this.Tail - 1
	if n < 0 {
		n = this.Length - 1
	}
	return this.Items[n]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.Head == this.Tail
}

func (this *MyCircularQueue) IsFull() bool {
	return (this.Tail+1)%this.Length == this.Head
}
