package main

import "fmt"

type student struct {
	name string
	age  int
	id   int
	next *student
}

type studentList struct {
	head *student
}

type studentHashTable struct {
	buckets []studentList
}

func makeStudentHashTable(len int) *studentHashTable {
	temp := new(studentHashTable)
	temp.buckets = make([]studentList, len)
	return temp
}

func (s *student) showSelf() {
	fmt.Printf("id: %d,name: %s,age: %d\n", s.id, s.name, s.age)
}

func (sh *studentHashTable) hashFunc(id int) int {
	return id % len(sh.buckets)
}

func (sh *studentHashTable) addStudent(stu *student) {
	//先根据学生的id进行散列
	sh.buckets[sh.hashFunc(stu.id)].addStudent(stu)
}

func (sh *studentHashTable) findStudent(id int) {
	if sh.buckets[sh.hashFunc(id)].head == nil || sh.buckets[sh.hashFunc(id)].head.next == nil {
		fmt.Println("未能找到")
		return
	}
	sh.buckets[sh.hashFunc(id)].findStudent(id)
}

func (sh *studentHashTable) removeStudent(id int) {
	sh.buckets[sh.hashFunc(id)].removeStudent(id)
}

func (sh *studentHashTable) updateStudent(id int, stu *student) {
	sh.buckets[sh.hashFunc(id)].updateStudent(id, stu)
}

func (s *studentList) addStudent(stu *student) {
	if s.head == nil {
		head := &student{}
		s.head = head
		//如果链表未初始化则先要帮它初始化一下
	}
	temp := s.head
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = stu
	fmt.Println("添加成功")
	stu.showSelf()
}

func (s *studentList) removeStudent(idToRemove int) {
	if s.head == nil || s.head.next == nil {
		fmt.Println("空List")
		return
	}
	temp := s.head
	for temp.next != nil {
		if temp.next.id == idToRemove {
			temp.next = temp.next.next
			fmt.Println("成功删除")
			return
		}
		temp = temp.next
	}
	fmt.Println("删除失败")

}

func (s *studentList) findStudent(idToFind int) {
	if s.head == nil || s.head.next == nil {
		fmt.Println("空List")
		return
	}
	temp := s.head
	for temp.next != nil {
		if temp.next.id == idToFind {
			fmt.Println("Find successfully")
			temp.next.showSelf()
			return
		}
		temp = temp.next
	}
	fmt.Println("Not Find")
}

func (s *studentList) updateStudent(idToUpdate int, stu *student) {
	if s.head == nil || s.head.next == nil {
		fmt.Println("空List")
		return
	}
	temp := s.head
	for temp.next != nil {
		if temp.next.id == idToUpdate {
			fmt.Println("Update successfully")
			temp.next.id = stu.id
			temp.next.name = stu.name
			temp.next.age = stu.age
			return
		}
		temp = temp.next
	}
	fmt.Println("Not Find")
}

func main() {
	//list := studentList{head: new(student)}
	s1 := &student{id: 0, name: "homo0", age: 114514}
	s2 := &student{id: 1, name: "homo1", age: 11451419}
	//list.addStudent(s1)
	//list.findStudent(0)
	//list.removeStudent(0)
	//list.findStudent(0)
	hashList := makeStudentHashTable(10)
	hashList.addStudent(s1)
	hashList.addStudent(s2)
	hashList.findStudent(1)
	hashList.removeStudent(1)
	hashList.findStudent(1)

}
