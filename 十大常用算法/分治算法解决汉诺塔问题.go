package main

import "fmt"

func Hanoitower(num int, A, B, C string) {
	//将塔分为两个部分,最下面的一层和最底层自最上层这个部分
	if num == 1 {
		fmt.Println("第1个盘从", A, "移动到", C)
	} else {
		//将第一部分从A移动到B，借助C
		Hanoitower(num-1, A, C, B)
		//把第二部分从A移动到C
		fmt.Println("第", num, "个盘从", A, "移动到", C)
		//把B处的盘子借助A移动到C
		Hanoitower(num-1, B, A, C)
	}
}

func main() {
	Hanoitower(3, "A", "B", "C")
}
