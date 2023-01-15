package main

import (
	"fmt"
)

type circleNode struct {
	id   int
	name string
	next *circleNode
	pre  *circleNode
}

type circleList struct {
	first *circleNode
}

func (l *circleList) add(new *circleNode) {
	if l.first == nil {
		fmt.Println("无效链表")
		return
	}
	temp := l.first
	for {
		if temp.next == l.first || temp.next == nil {
			break
		}
		temp = temp.next
	}

	temp.next = new
	new.pre = temp
	new.next = l.first
	l.first.pre = new
}
func (l *circleList) addSomeNodes(number int) {
	if l.first == nil {
		fmt.Println("无效链表")
		return
	}
	//rand.Seed(time.Now().UnixNano())
	for i := 2; i <= number; i++ {
		temp := &circleNode{id: i}
		l.add(temp)
	}
}

func (l *circleList) show() {
	if l.first == nil {
		fmt.Println("无效链表")
		return
	}
	temp := l.first
	for {
		fmt.Println(temp.id)
		temp = temp.next
		if temp == l.first || temp == nil {
			break
		}
	}
}

func (l *circleList) Joseph(start, count int) {
	if l.first == nil {
		fmt.Println("无效链表")
		return
	}
	temp := l.first
	helper := temp
	//先把helper移动到最后一位
	for {
		if helper.next == l.first {
			break
		}
		helper = helper.next
	}
	//将temp移动到start位置
	for i := 0; i < start-1; i++ {
		temp = temp.next
		helper = helper.next
	}
	//开始
	for {
		if temp == helper {
			break
		} //只剩一个节点
		for i := 0; i < count-1; i++ {
			temp = temp.next
			helper = helper.next
		} //移动待删除位置
		//将temp所指向的节点删除
		fmt.Println(temp.id, "出局")
		temp = temp.next
		helper.next = temp
	}
	fmt.Println(temp.id, "获胜")
}

func main() {
	circlelist := circleList{first: &circleNode{id: 1}}
	circlelist.addSomeNodes(5)
	circlelist.show()
	circlelist.Joseph(1, 3)
}
