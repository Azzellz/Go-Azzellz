package main

import "fmt"

type Queue struct {
	arr []int
	len int
}

func (q *Queue) addQueue(n ...int) {
	for i := 0; i < len(n); i++ {
		q.arr = append(q.arr, n[i])
		q.len++
	}
}

func (q *Queue) removeQueue() int {
	temp := q.arr[0]
	q.arr = q.arr[1:]
	q.len--
	return temp
}

func (q *Queue) showMembers() {
	if q.len == 0 {
		fmt.Println("nothing..")
		return
	}
	for _, v := range q.arr {
		fmt.Printf("%d\t", v)
	}
}

func main() {
	queue := Queue{arr: make([]int, 0), len: 0}
	queue.addQueue(10)
	queue.removeQueue()
	queue.addQueue(1, 2, 3, 4, 5)
	queue.removeQueue()
	queue.showMembers()

}
