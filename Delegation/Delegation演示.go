package main

import (
	"NewGo/Delegation"
	"fmt"
)

type func1 func(int)
type func2 func(int, int)
type func3 func()

func (func3) Confirm(function interface{}, args interface{}) {
	temp, ok := function.(func3)
	if ok {
		temp()
	} else {
		fmt.Println("error")
	}
}

// 利用func1和func2来体现差异
// 传参的时候只需要考虑传参的位置,由使用者实现接口时正确传入参数,这样就解决了不同函数参数传递的问题(这么tm的这么麻烦)
// 以下两个实现接口的例子作为如何使用委托的教程
// Fucking golang why it don't have 泛型
func (func1) Confirm(function interface{}, args interface{}) {
	temp, ok := function.(func1)
	if ok {
		ar := args.([]interface{})
		i := ar[0].(int) //这里不嫌麻烦的话其实可以再写一层类型断言的ok判断,由User自行决定
		temp(i)
	} else {
		fmt.Println("Delegator调用有误")
	}
}
func (func2) Confirm(function interface{}, args interface{}) {
	temp, ok := function.(func2)
	if ok {
		ar := args.([]interface{})
		i := ar[0].(int)
		j := ar[1].(int)
		temp(i, j)
	} else {
		fmt.Println("Delegator调用有误")
	}
}

func test1(a int) {
	fmt.Println("test1")
	fmt.Println(a)
}
func test2(a int, b int) {
	fmt.Println("test2")
	fmt.Println(a, b)
}

type testor struct{}

func (testor) test3() {
	fmt.Println("test3")
}

func main() {
	//自定义函数类型实现Confirm接口使用委托:
	var f1 func1 = test1
	var f2 func2 = test2
	delegator1 := Delegation.NewDelegator()
	delegator1.RegisterFunc(f1, "f1")
	delegator1.RegisterFunc(f2, "f2")
	delegator1.ShowFunc()
	delegator1.CarryFunc(3, Delegation.CarryFuncsUnremainly, 3, 4)
	//结构体绑定的方法实现Confirm接口使用委托:
	testor1 := testor{}
	//要麻烦一点,如下
	var f3 func3 = testor1.test3
	delegator2 := Delegation.NewDelegator()
	delegator2.RegisterFunc(f3, "f3")
	//演示委托的搜索功能
	f4 := delegator2.SearchFunc("f3")
	f5 := f4.(func3)
	f5()
	delegator2.CarryFunc(1, Delegation.CarryFuncsRemainly)
	delegator2.DeleteFunc("f3")
	delegator2.CarryFunc(1, Delegation.CarryFuncsRemainly)

}
