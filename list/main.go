package main

import (
	"alg/list/circular_queue"
	"fmt"
)

func main() {
	obj := circular_queue.Constructor(3)
	obj.EnQueue(1)
	obj.EnQueue(2)
	obj.EnQueue(3)
	obj.EnQueue(4)
	obj.DeQueue()
	r := obj.EnQueue(4)
	fmt.Println(r)
	//queue1.Enqueue("5")
	//queue1.Enqueue("6")
	//queue1.Enqueue("7")
	//queue1.Enqueue("8")
	//queue1.Enqueue("9")
	//queue1.Enqueue("10")
	fmt.Println(obj.Rear())
}
