package array_queue

type ArrayQueue struct {
	Items  []string
	Length int
	Head   int
	Tail   int
}

// 入队
func (a *ArrayQueue) Enqueue(s string) bool {
	if a.Tail == a.Length {
		if a.Head == 0 {
			return false
		}
		// 需要搬运
		for i := a.Head; i < a.Length; i++ {
			a.Items[i-a.Head] = a.Items[i]
		}
		a.Tail = a.Tail - a.Head
		a.Head = 0
	}
	a.Items[a.Tail] = s
	a.Tail++
	return true
}

// 出队
func (a *ArrayQueue) Dequeue() string {
	if a.Head == a.Tail {
		return ""
	}
	s := a.Items[a.Head]
	a.Head++
	return s
}
