package main

import "fmt"

type queueNew struct {
	arr     [5]int
	len     int
	front   int
	rear    int
	maxSize int
}

func (q *queueNew) isEmpty() bool {
	return q.len == 0
}

func (q *queueNew) isFill() bool {
	return q.len == q.maxSize
}

func makeQueue() queueNew {
	return queueNew{len: 0, front: 0, rear: 0, maxSize: 5}
}

func (q *queueNew) addQueue(n int) {
	if q.isFill() {
		fmt.Println("full queue!")
		return
	}

	q.arr[q.rear] = n
	q.rear = (q.rear + 1) % q.maxSize
	q.len++
}

func (q *queueNew) outQueue() int {
	if q.isEmpty() {
		fmt.Println("Empty queue!")
		return -1
	}
	temp := q.arr[q.front]
	q.front = (q.front + 1) % q.maxSize
	q.len--
	return temp
} //1 2 3 4

func (q *queueNew) showQueue() {
	if q.isEmpty() {
		fmt.Println("Empty queue!")
		return
	}
	tempI := q.front
	for i := 0; i < q.len; i++ {
		fmt.Printf("arr[%d]=%d", i, q.arr[tempI])
		tempI = (tempI + 1) % q.maxSize
	} //只让循环控制次数,另起索引
}

func main() {
	q := makeQueue()
	q.addQueue(10)
	q.addQueue(11)
	q.addQueue(12)
	q.addQueue(13)
	q.addQueue(14)
	q.outQueue()
	q.addQueue(15)
	q.showQueue()
}
