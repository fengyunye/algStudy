package main

import (
	"fmt"
)

//有一个头结点，没有值域，只有连域，专门存放第一个结点的地址。
//有一个尾结点，有值域，也有链域，链域值始终为NULL。
//所以，在单链表中为找第i个结点或数据元素，必须先找到第i - 1 结点或数据元素，而且必须知道头结点，否者整个链表无法访问。

// 结点
type Node struct {
	Data interface{}
	Next *Node
}

type List struct {
	HeadNode *Node // 头结点
}

type Method interface {
	IsEmpty()
	Length()
	Insert(i int, v interface{}) // 增
	Delete(i int)                // 删
	GetLength() int              // 获取长度
	Search(v interface{}) int    // 查
	isNull() bool                // 判断是否为空
}

// 是否为空
func (l *List) IsEmpty() bool {
	if l.HeadNode == nil {
		return true
	} else {
		return false
	}
}

// 获取长度
func (l *List) Length() int {
	currentNode := l.HeadNode
	currentNum := 0

	for currentNode != nil {
		currentNode = currentNode.Next
		currentNum++
	}
	return currentNum
}

// 头部添加元素
func (l *List) Add(data interface{}) *Node {
	newNode := &Node{
		data,
		nil,
	}
	newNode.Next = l.HeadNode.Next
	l.HeadNode = newNode
	return newNode
}

// 尾部添加元素
func (l *List) Append(data interface{}) {
	newNode := &Node{
		data,
		nil,
	}
	if l.HeadNode == nil {
		// 链表为空 直接赋值
		l.HeadNode = newNode
	} else {
		currentNode := l.HeadNode
		for currentNode.Next != nil {
			currentNode = currentNode.Next
		}
		currentNode.Next = newNode
	}
}

// 指定位置插入结点

func (l *List) Insert(i int, data interface{}) {
	if i < 0 {
		l.Add(data) // 头部插入
	} else if i >= l.Length() {
		l.Append(data) // 尾部插入
	} else {
		newNode := &Node{
			Data: data,
		}
		currentNode := l.HeadNode
		currentI := 0
		for currentI < i {
			currentNode = currentNode.Next
			currentI++
		}
		newNode.Next = currentNode.Next // 新结点的地址 指向原节点的下一个  例如 a ->c    插入b   先让 b->c  再将a 指向b 即可完成交换
		currentNode.Next = newNode      // 原节点
	}
}

// 删除指定值的元素
func (l *List) Remove(data interface{}) {
	currentNode := l.HeadNode
	if currentNode.Data == data {
		// 删除头结点
		l.HeadNode = currentNode.Next
	} else {
		for currentNode.Next != nil {
			if currentNode.Next.Data == data {
				// 查看下一项
				currentNode.Next = currentNode.Next.Next
			} else {
				// 到下一个结点
				currentNode = currentNode.Next
			}
		}
	}
}

// 删除指定位置的元素

func (l *List) RemoveByIndex(i int) {
	currentNode := l.HeadNode
	if i <= 0 {
		//删除头结点
		l.HeadNode = currentNode.Next
	} else if i > l.Length() {
		fmt.Printf("超出链表长度,当前列表长度%d", l.Length())
		return
	} else {
		currentI := 0
		for currentNode.Next != nil {
			// 循环遍历链表，
			if currentI+1 == i {
				break
			}
			currentI++
			currentNode = currentNode.Next
		}
		currentNode.Next = currentNode.Next.Next
	}
}

// 打印
func (l *List) ShowList() {
	if !l.IsEmpty() {
		currentNode := l.HeadNode
		for {
			fmt.Printf("\t%v", currentNode.Data)
			//fmt.Println(currentNode)
			if currentNode.Next != nil {
				currentNode = currentNode.Next
			} else {
				break
			}
		}
	}
}

// 是否包含
func (l *List) Contain(data interface{}) bool {
	currentNode := l.HeadNode
	for currentNode != nil {
		if currentNode.Data == data {
			return true
		}
		currentNode = currentNode.Next
	}
	return false
}

// 将数组转换成链表

func ArrayToList(array []int) *List {
	list := &List{}
	for _, value := range array {
		list.Append(value)
	}
	return list
}

// 单链表反转 - 递归
func ReverseList(n *Node) *Node {
	if n == nil || n.Next == nil {
		return n
	}
	res := ReverseList(n.Next)
	n.Next.Next = n
	n.Next = nil
	return res
}

// 单链表反转 - 迭代
// 直接反转 需要预先存储 pre 和 next
func ReverseListByFor(n *Node) *Node {
	var pre *Node
	for n != nil {
		next := n.Next
		n.Next = pre
		pre = n
		n = next
	}
	return pre
}

// 简单暴力法，map存储，有相同结点，则有换
func HasCycle(head *Node) bool {
	headMap := map[*Node]struct{}{}
	for head != nil {
		if _, ok := headMap[head]; ok {
			return true
		}
		headMap[head] = struct{}{}
		head = head.Next
	}
	return false
}

// 定义快慢指针 如果链表里有环, 跑得慢的和跑的快的一定会相遇，
func HasCycleV2(head *Node) bool {
	// 快慢指针
	if head == nil || head.Next == nil {
		return false
	}
	fast := head.Next.Next
	for fast != nil && head != nil && fast.Next != nil {
		if fast == head {
			return true
		}
		head = head.Next
		fast = fast.Next.Next
	}
	return false
}

// 删除倒数第n个结点，设立哨兵结点result
func removeNthFromEnd(head *Node, n int) *Node {
	// 设定哨兵结点存储结点
	result := &Node{}
	result.Next = head
	// 设定哨兵
	var pre *Node
	cur := result
	currentI := 1
	for head != nil {
		// 因为是删除倒数， 所以移动是有规律
		// head遍历至 尾结点，  想要删除的是倒数 4个结点，  两者相差n-1个结点， 所以当head移动了n-1后， cur也开始移动， 因为cur是从哨兵结点开始的，所以实际还加了一个结点
		if currentI >= n {
			// 需要存储pre结点, 如果是删除的第一个结点，没有哨兵结点，其实是特别麻烦的。
			pre = cur
			cur = cur.Next
		}
		head = head.Next
		currentI++
	}
	pre.Next = pre.Next.Next
	return result.Next
}

// 有序链表合并
func MergeList(node1 *Node, node2 *Node) *Node {
	// 设定哨兵结点，
	result := &Node{}
	current := result
	for node1 != nil && node2 != nil {
		if node1.Data.(int) < node2.Data.(int) {
			current.Next = node1
			node1 = node1.Next
		} else {
			current.Next = node2
			node2 = node2.Next
		}
		current = current.Next
	}
	if node1 != nil {
		current.Next = node1
	}
	if node2 != nil {
		current.Next = node2
	}
	return result.Next
}

// 方法1 - 循环遍历，放入数组中，使用数组，占用空间
// 龟兔法， 快指针到了，满指针才到一半
// 偶数项   6
func middleNode(head *Node) *Node {
	// 快慢指针
	slow, fast := head, head

	for fast != nil {
		if fast.Next != nil && fast.Next.Next != nil {
			fast = fast.Next.Next
			slow = slow.Next
		} else if fast.Next != nil && fast.Next.Next == nil {
			fast = fast.Next
			slow = slow.Next
		} else {
			fast = fast.Next
		}
	}
	return slow
}

//给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。

//请你将两个数相加，并以相同形式返回一个表示和的链表。

//你可以假设除了数字 0 之外，这两个数都不会以 0 开头。

// - 2 链表两数想加
func addTwoNumbers(l1 *Node, l2 *Node) *Node {
	// 新链表存储结果，哨兵返回值
	list := &Node{
		Data: 0,
		Next: nil,
	}
	result := list
	temp := 0 // 代表进位，第一次不进位
	for l1 != nil || l2 != nil {
		if l1 != nil {
			temp += l1.Data.(int)
			l1 = l1.Next
		}
		if l2 != nil {
			temp += l2.Data.(int)
			l2 = l2.Next
		}
		list.Next = &Node{
			Data: temp % 10,
			Next: nil,
		}
		temp = temp / 10
		list = list.Next
	}
	return result.Next
}

func Test() {
	result := &Node{
		Data: 5,
	}
	current := result
	current = &Node{
		Data: 6,
	}
	fmt.Println(result)
	fmt.Println(current)
}

// 83. 删除排序链表中的重复元素
func deleteDuplicates(head *Node) *Node {
	// 定义哨兵
	result := &Node{}
	result.Next = head
	// 因为是排好序的，所以只需要比较相邻
	for head != nil && head.Next != nil {
		if head.Data == head.Next.Data {
			// 删除下一项
			head.Next = head.Next.Next
		} else {
			// 指向下一个节点
			head = head.Next
		}
	}
	return result.Next
}

// 移除
func removeElements(head *Node, val int) *Node {
	// 定义哨兵节点
	result := &Node{}
	result.Next = head
	current := result
	for current != nil && current.Next != nil {
		if current.Next.Data == val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
	return result.Next
}

// a + b = b + a 如果没有相交，一定会同时 =nil 并结束
// 相交链表
func getIntersectionNode(headA, headB *Node) *Node {
	if headA == nil || headB == nil {
		return nil
	}
	l1, l2 := headA, headB
	for l1 != l2 {
		if l1 != nil {
			l1 = l1.Next
		} else {
			l1 = l2
		}

		if l2 != nil {
			l2 = l2.Next
		} else {
			l2 = l1
		}
	}
	return l1
}

func main() {
	//arr1 := [5]int{1, 2, 3, 4, 5}
	//arr2 := [5]int{2, 3, 4, 5, 6}
	//list1 := ArrayToList(arr1[:])
	//list2 := ArrayToList(arr2[:])
	////list := &List{}
	////list.Append(1)
	////list.Append(2)
	////list.Append(3)
	////list.Append(4)
	////list.RemoveByIndex(1)
	////list.HeadNode = ReverseList(list.HeadNode)
	////list1.ShowList()
	////list2.ShowList()
	//mergeList := MergeList(list1.HeadNode, list2.HeadNode)
	//currentList := &List{}
	//currentList.HeadNode = mergeList
	//currentList.ShowList()
	//j := middleNode(currentList.HeadNode)
	//fmt.Println(j)
	//Test()
	//list.Insert(2, 1)
	//list.PrintList()
	//list := CreateNode(1)
	//node2 := CreateNode(2)
	//list.next = node2
	////list.Insert(1, 5)
	//fmt.Println(*list.next)
	//var b = []int{1, 2, 3, 5}
	//ArrayToNode(b)

	arr3 := [3]int{2, 4, 3}
	arr4 := [3]int{5, 6, 4}
	list3 := ArrayToList(arr3[:])
	list4 := ArrayToList(arr4[:])
	newNode := addTwoNumbers(list3.HeadNode, list4.HeadNode)
	showList := &List{
		newNode,
	}
	showList.ShowList()

	arr5 := [3]int{1, 1, 1}

	list5 := ArrayToList(arr5[:])
	deleteDuplicates(list5.HeadNode)
	removeElements(list5.HeadNode, 1)

}
