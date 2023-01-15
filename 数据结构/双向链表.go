package main

import "fmt"

type node struct {
	id   int
	name string
	next *node
	pre  *node
}

type doubleList struct {
	head *node
}

func (l *doubleList) add(newNode *node) {
	if l.head == nil {
		fmt.Println("无效链表")
		return
	}
	temp := l.head
	for {
		if temp.next == nil {
			break
		}
		if temp.next.id >= newNode.id { //有序添加
			newNode.next = temp.next.next
			temp.next.pre = newNode
			temp.next = newNode
			newNode.pre = temp
			return
		}
		temp = temp.next
	}
	temp.next = newNode
	newNode.pre = temp
}

func (l *doubleList) delete(id int) {
	if l.head == nil {
		fmt.Println("无效链表")
		return
	}
	temp := l.head
	for {
		if temp.next == nil {
			fmt.Println("Not Found")
			break
		}
		temp = temp.next
		if temp.id == id {
			if temp.next == nil {
				temp.pre.next = temp.next
				break
			} //说明是最后一个元素
			temp.next.pre = temp.pre
			temp.pre.next = temp.next
			fmt.Println("Delete successfully")
			break
		}
	}
}

func (l *doubleList) show() {
	if l.head == nil {
		fmt.Println("无效链表")
		return
	}
	temp := l.head
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
		fmt.Println(temp.id, temp.name)
	}
}

func (l *doubleList) update(newNode *node) {
	if l.head == nil {
		fmt.Println("无效链表")
		return
	}
	temp := l.head
	for {
		if temp.next == nil {
			fmt.Println("Not Found")
			break
		}
		temp = temp.next
		if temp.id == newNode.id { //执行修改
			temp.name = newNode.name
			break
		}
	}

}

func main() {
	n1 := &node{id: 1, name: "a"}
	n2 := &node{id: 20, name: "b"}
	n3 := &node{id: 15, name: "c"}
	n4 := &node{id: 15, name: "aac"}
	doubleList := doubleList{head: new(node)}
	doubleList.add(n1)
	doubleList.add(n3)
	doubleList.add(n2)

	//doubleList.delete(2)
	//doubleList.delete(1)
	doubleList.update(n4)
	doubleList.show()
}
