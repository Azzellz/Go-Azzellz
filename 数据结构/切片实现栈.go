package main

import "fmt"

type Stack struct {
	arr []int
	len int
}

func (s *Stack) push(n int) {
	s.arr = append(s.arr, n)
	s.len++
}

func (s *Stack) pop() (temp int) {
	if s.len == 0 {
		fmt.Printf("Nothing to pop from stack")
		temp = -1
		return
	}
	temp = s.arr[s.len-1]
	s.arr = s.arr[:s.len-1]
	s.len--
	return
}

func main() {

	stack := Stack{make([]int, 0), 0}
	stack.push(10)
	stack.push(12)
	fmt.Println(stack.pop())
	fmt.Println(stack.pop())
}
