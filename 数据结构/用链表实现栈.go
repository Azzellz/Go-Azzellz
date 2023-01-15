package main

import "fmt"

type linkStack struct {
	len int
	stackList
}

type stackList struct {
	head *stackNode
}

type stackNode struct {
	id   int
	next *stackNode
}

func (l *linkStack) push(new *stackNode) {
	if l.head == nil {
		fmt.Println("未初始化栈")
		return
	}
	//if l.head.next == nil{
	//	fmt.Println("空-")
	//	return
	//}
	temp := l.head
	for { //先遍历到表的尾部
		if temp.next == nil {
			break
		}
		temp = temp.next
	}
	temp.next = new
	l.len++
}

func (l *linkStack) pop() *stackNode {
	if l.head == nil {
		fmt.Println("未初始化栈")
		return nil
	}
	if l.head.next == nil {
		fmt.Println("空-")
		return nil
	}
	temp := l.head
	for {
		if temp.next.next == nil {
			break
		}
		temp = temp.next
	}
	new := temp.next
	temp.next = nil
	return new

}

func main() {

	stack := new(linkStack)
	stack.stackList.head = new(stackNode)
	n1 := &stackNode{id: 0}
	n2 := &stackNode{id: 1}
	stack.push(n1)
	stack.push(n2)
	fmt.Println(stack.pop())
	fmt.Println(stack.pop())
	fmt.Println(stack.pop())

}
