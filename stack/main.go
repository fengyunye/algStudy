package main

import "fmt"

type Stack struct {
	Items   []int
	Current int
	Length  int
}

// 创建一个栈
func (s *Stack) Create(n int) {
	s.Length = n
	s.Current = 0
	s.Items = make([]int, n)
}

// 入栈
func (s *Stack) Push(number int) bool {
	if s.Current >= s.Length {
		return false
	}
	s.Items[s.Current] = number
	s.Current++
	return true
}

// 出站
func (s *Stack) Pop() (int, bool) {
	if s.Current == 0 {
		return 0, false
	}
	res := s.Items[s.Current-1]
	s.Current--
	return res, true
}

// 栈相关
func main() {
	stack := &Stack{}
	stack.Create(5)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	fmt.Println(stack)
	number, _ := stack.Pop()
	fmt.Println(number)
	fmt.Println(stack)

}
