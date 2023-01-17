package main

import (
	"errors"
	"fmt"
)

type homoNode struct {
	id   int
	name string
	age  int
	next *homoNode
}

type homoStack struct {
	arr []homoNode
	len int
}

func (s *homoStack) push(h homoNode) {
	s.arr = append(s.arr, h)
	s.len++
}

func (s *homoStack) pop() (temp homoNode) {
	if s.len == 0 {
		fmt.Printf("Nothing to pop from stack")
		return
	}
	temp = s.arr[s.len-1]
	s.arr = s.arr[:s.len-1]
	s.len--
	return
}

type homoList struct {
	head *homoNode
	homoStack
}

func makeList() (list *homoList) {
	head := new(homoNode)
	return &homoList{head: head}
}

func (h *homoList) upDate(homo *homoNode) {
	if h.head == nil {
		fmt.Println("empty list")
		return
	}
	temp := h.head
	//找到对应位置
	for {
		if temp.next == nil {
			break
		}
		if temp.id == homo.id {
			temp.name = homo.name
			temp.age = homo.age
			break
		}
		temp = temp.next
	}
}

func (h *homoList) add(homo *homoNode) {
	temp := h.head
	for { //这个循环是用来定位的
		//实现按id从小到大顺序插入
		if temp.next == nil {
			break
		}
		if temp.next.id > homo.id {
			homo.next = temp.next
			temp.next = homo
			return
		} else if temp.next.id == homo.id {
			fmt.Println("添加id编号重复,无法保证正确顺序")
			break
		}
		temp = temp.next
	}
	//homo.next = temp.next
	temp.next = homo
}

func (h *homoList) show() {
	temp := h.head
	if temp.next == nil {
		fmt.Println("empty!")
		return
	}

	for {
		if temp.next == nil {
			break
		}
		temp = temp.next //避过遍历头节点
		fmt.Printf("HomoId %d,HomoName %s,HomoAge %d\n", temp.id, temp.name, temp.age)
	}
}

func (h *homoList) delete(id int) {
	if h.head == nil {
		fmt.Println("empty!")
		return
	}
	temp := h.head
	for {
		if temp.next == nil {
			break
		}
		if temp.next.id == id { //执行删除操作
			temp.next = temp.next.next
			break
		}
		temp = temp.next
	}
}

func (h *homoList) searchByEnd(id int) (dstHomo *homoNode) {
	if id <= 0 || h.head == nil || id > h.getLen() {
		fmt.Println(errors.New("请在有效链表中查询有效id"))
		return
	}
	temp := h.head
	for i := 0; i < h.getLen()-id+1; i++ {
		temp = temp.next
	}
	dstHomo = temp
	return
}

func (h *homoList) getLen() (len int) {
	if h.head == nil {
		fmt.Println("unDefined list")
		return
	}
	temp := h.head
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
		len++
	}
	return
}

func (h *homoList) reverseList() { //单链表反转
	if h.head == nil {
		fmt.Println("链表为空,不可反转")
		return
	} else if h.head.next == nil || h.head.next.next == nil {
		fmt.Println("不需要反转")
		return
	}
	temp := h.head.next
	helper := new(homoNode)
	reverseHead := new(homoNode)
	//创建一个新的表头,然后用temp遍历原表的每一个节点,每遍历一个节点,就插入到reverseHead的下一位,相当于链表插入的操作
	for {

		helper = temp.next
		temp.next = reverseHead.next //这句真的神中神,屏蔽了第一次的差异
		reverseHead.next = temp
		temp = helper
		if temp == nil {
			break
		}
		//需要创建一个替身
		//用值代替(低效)
		//helper := new(homoNode)
		//helper.id = temp.id
		//helper.name = temp.name
		//helper.age = temp.age
		//if reverseHead.next == nil {
		//	reverseHead.next = helper
		//	continue
		//} //首个插入是特殊情况
		//helper.next = reverseHead.next
		//reverseHead.next = helper
	}
	//将head指向reverseHead.next
	h.head.next = reverseHead.next
}

func (h *homoList) printFromEnd() {
	if h.head == nil {
		fmt.Println("无效链表")
		return
	}
	if h.head.next == nil {
		fmt.Println("空链表")
		return
	}
	temp := h.head
	//先将节点压入栈中
	stack := homoStack{make([]homoNode, 0), 0}
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
		stack.push(*temp)
	}
	//开始压出栈内的元素
	length := stack.len
	for i := 0; i < length; i++ {
		fmt.Println(stack.pop())
	}
}

func (h *homoList) merge(list *homoList) {
	//用两层循环搞定
	//假设是list2并入list1
	if h.head == nil || h.head == nil {
		fmt.Println("有无效链表")
		return
	}
	if h.head.next == nil || h.head.next == nil {
		fmt.Println("有空链表,无意义合并")
		return
	}
	temp1 := h.head.next
	temp2 := list.head.next
	//这两个表都是有序的,只要temp1遍历完就结束了
	for {
		if temp2.id <= temp1.next.id { //插入
			temp := temp2.next
			//temp2.next = temp1.next
			temp1.next = temp2
			temp2 = temp
			continue
		}
		temp1 = temp1.next
		if temp1.next == nil {
			temp1.next = temp2
			return
		}

	}

}

func main() {

	homoList1 := makeList()
	homo1 := &homoNode{id: 1, name: "先辈", age: 114514}
	homo2 := &homoNode{id: 3, name: "淳平", age: 414}
	homo3 := &homoNode{id: 5, name: "真寻", age: 41}
	homo4 := &homoNode{id: 7, name: "朴秀", age: 14}
	homoList1.add(homo1)
	homoList1.add(homo3)
	homoList1.add(homo2)
	homoList1.add(homo4)
	//homoList1.show()
	homoList2 := makeList()
	homo10 := &homoNode{id: 2, name: "先辈1", age: 114514}
	homo20 := &homoNode{id: 4, name: "淳平1", age: 414}
	homo30 := &homoNode{id: 6, name: "真寻1", age: 41}
	homo40 := &homoNode{id: 8, name: "朴秀1", age: 14}
	homoList2.add(homo10)
	homoList2.add(homo30)
	homoList2.add(homo20)
	homoList2.add(homo40)
	//homoList2.show()
	homoList1.merge(homoList2)
	homoList1.show()

}
