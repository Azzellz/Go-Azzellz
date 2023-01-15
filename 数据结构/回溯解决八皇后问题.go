package main

import "fmt"

type queens struct {
	arr [8]int
	//数组的下标表示第几个皇后,值表示皇后的位置
	count int
}

func (q *queens) print() {
	for i := 0; i < 8; i++ {
		fmt.Print(q.arr[i], " ")
	}
	fmt.Println()
	q.count++
} //打印解

func (q *queens) judge(n int) bool {
	for i := 0; i < n; i++ {
		if q.arr[i] == q.arr[n] || abs(n, i) == abs(q.arr[i], q.arr[n]) {
			return false
		} //判断是否在同一列或在同一斜线
	}
	return true
} //用来判断第n个皇后是否满足规则

func (q *queens) check(n int) {
	if n == 8 {
		q.print()
		return
	} //基准条件,表示8个皇后都已经放完
	for i := 0; i < 8; i++ {
		q.arr[n] = i //从0号位开始放起
		if q.judge(n) {
			q.check(n + 1)
		}
	}
} //放置皇后(从0开始

func abs(a, b int) int {
	if a-b < 0 {
		return -(a - b)
	} else {
		return a - b
	}
} //求绝对值

func main() {
	q := queens{}
	q.check(0)
	fmt.Println(q.count)
}
