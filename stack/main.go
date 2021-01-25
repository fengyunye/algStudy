package main

import (
	"strconv"
)

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

// 是否闭合
func isValid(s string) bool {
	length := len(s)
	if length%2 == 1 {
		return false
	}
	// 设定一个map
	stringMap := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}
	// 用栈存储
	currentI := 0
	item := make([]byte, length)
	// 字符串转换为byte
	stringByte := []byte(s)

	for i := 0; i < length; i++ {
		// 前半截字符直接入栈
		if stringMap[stringByte[i]] > 0 {
			if currentI == 0 {
				return false
			}

			// 如果当前字符是后半截，切栈顶元素和当前元素不是对应关系
			if stringMap[stringByte[i]] != item[currentI-1] {
				return false
			}
			currentI--

		} else {
			// 入栈
			item[currentI] = stringByte[i]
			currentI++
		}

	}
	return currentI == 0
}

// 比较含退格的字符串
func backspaceCompare(S string, T string) bool {
	return reBuildString(S) == reBuildString(T)
}

func reBuildString(s string) string {
	var SItem []byte
	for i := 0; i < len(s); i++ {
		if s[i] == '#' {
			if len(SItem) == 0 {
				continue
			} else {
				SItem = SItem[:len(SItem)-1]
			}
		} else {
			SItem = append(SItem, s[i])
		}
	}
	return string(SItem)
}

// 棒球比赛
func calPoints(ops []string) int {
	var stack []int
	for i := 0; i < len(ops); i++ {
		if ops[i] == "C" {
			stack = stack[:len(stack)-1]
		} else if ops[i] == "D" {
			current := stack[len(stack)-1]
			stack = append(stack, current*2)
		} else if ops[i] == "+" {
			sum := stack[len(stack)-1] + stack[len(stack)-2]
			stack = append(stack, sum)
		} else {
			intNumber, _ := strconv.Atoi(ops[i])
			stack = append(stack, intNumber)
		}
	}
	sum := 0
	for j := 0; j < len(stack); j++ {
		sum += stack[j]
	}
	return sum
}

// 栈相关
func main() {
	//stack := &Stack{}
	//stack.Create(5)
	//stack.Push(1)
	//stack.Push(2)
	//stack.Push(3)
	//stack.Push(4)
	//stack.Push(5)
	//fmt.Println(stack)
	//number, _ := stack.Pop()
	//fmt.Println(number)
	//fmt.Println(stack)
	//
	//fmt.Println(isValid("{}"))
	//queue := list_by_stack.Constructor()
	//queue.Push(1)
	//queue.Push(2)
	//queue.Push(3)
	//fmt.Println(queue)
	backspaceCompare("ab#c", "ad#c")

}
