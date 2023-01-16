package Delegation

import (
	"fmt"
)

// Delegation 本包尝试对委托进行实现

type ConfirmFunc interface {
	Confirm(function interface{}, args interface{}) //函数执行
	//Name() string//记录函数名
} //用于调用函数方对函数进行类型断言

const (
	CarryFuncsRemainly   = iota //执行完保留函数
	CarryFuncsUnremainly        //执行完不保留函数
) //定义宏来区分执行模式

type delegator struct {
	funcs funcQueue
} //Delegator用于维护一个基于接口的函数执行队列

func NewDelegator() *delegator {
	return &delegator{}
} //工厂模式返回一个Delegator指针

type funcQueue struct {
	FuncList []ConfirmFunc //空接口切片,用来装函数
	NameList []string      //记录函数名
	len      int
} //funcQueue用于实现一个基于Go语言切片的队列

// funcQueue的方法:对函数队列的增删改查以及执行函数的多种模式
func (queue *funcQueue) addFunc(voidFunc ConfirmFunc, name string) {
	if queue.FuncList == nil {
		queue.FuncList = make([]ConfirmFunc, 0)
		queue.FuncList = append(queue.FuncList, voidFunc)
		queue.NameList = append(queue.NameList, name)
		queue.len++
		return
	} //如果未初始化queue中的函数切片,则要先帮忙先初始化一下,顺便完成添加第一个函数的操作

	queue.FuncList = append(queue.FuncList, voidFunc)
	queue.NameList = append(queue.NameList, name)
	queue.len++

} //函数入队

func (queue *funcQueue) deleteFunc(name string) bool {
	var index = -1
	for i, v := range queue.NameList {
		if v == name {
			index = i
			break
		}
	}
	if index == -1 {
		return false
	}

	queue.NameList = append(queue.NameList[:index], queue.NameList[index+1:]...)
	queue.FuncList = append(queue.FuncList[:index], queue.FuncList[index+1:]...)
	queue.len--
	return true

} //执行一次O(n)的搜索,删除name匹配的函数,返回bool

func (queue *funcQueue) showFunc() {
	if queue.len == 0 {
		fmt.Println("Delegator内未登记函数")
	}
	fmt.Println("Delegator内登记的函数有:")
	for _, v := range queue.NameList {
		fmt.Printf("%v ", v)
	}
	fmt.Println()
} //遍历NameList的所有函数名

func (queue *funcQueue) doFuncUnremainly(doNum int, args interface{}) {
	if queue.len == 0 {
		fmt.Println("Empty delegator")
		return
	}
	if doNum > queue.len {
		tempLen := queue.len
		for i := 0; i < tempLen; i++ {
			temp := queue.FuncList[0]
			queue.FuncList = queue.FuncList[1:]
			queue.NameList = queue.NameList[1:]
			queue.len--
			temp.Confirm(temp, args)
		}
	} else {
		for i := 0; i < doNum; i++ {
			temp := queue.FuncList[0]
			queue.FuncList = queue.FuncList[1:]
			queue.NameList = queue.NameList[1:]
			queue.len--
			temp.Confirm(temp, args)
		}
	}

} //函数出队

func (queue *funcQueue) doFuncRemainly(doNum int, args interface{}) {
	if queue.len == 0 {
		fmt.Println("Empty delegator")
		return
	}

	if doNum > queue.len { //完全执行
		for _, v := range queue.FuncList {
			v.Confirm(v, args)
		}
	} else {
		for i := 0; i < doNum; i++ {
			queue.FuncList[i].Confirm(queue.FuncList[i], args)
		}
	}

} //遍历执行队列,不出队

//================================================================

//delegator的应用主方法:包含添加函数,执行函数,查看函数列表,搜索函数,删除委托函数
//可拓展内容:保留函数执行和不保留执行,可以使用传入本包内的宏来判定执行模式

func (delegator *delegator) RegisterFunc(addFunc ConfirmFunc, name string) {
	delegator.funcs.addFunc(addFunc, name)
} //为委托登记函数

func (delegator *delegator) CarryFunc(carryNum int, carryMode int, args ...interface{}) {
	//不保留执行,执行完函数被消耗
	if carryMode == CarryFuncsUnremainly {
		delegator.funcs.doFuncUnremainly(carryNum, args)
	}
	//================================================================

	//保留执行,执行完函数不被消耗
	if carryMode == CarryFuncsRemainly {
		delegator.funcs.doFuncRemainly(carryNum, args)
	}
	//================================================================

} //执行n次函数,如果n大于委托内函数的数量,则会完全执行委托内的所有函数
func (delegator *delegator) ShowFunc() {

} //按登记顺序打印所有存在于委托内的函数,输出函数名

func (delegator *delegator) SearchFunc(name string) interface{} {
	for i, v := range delegator.funcs.NameList {
		if v == name {
			return delegator.funcs.FuncList[i]
		}
	}
	return nil
} //该方法执行一次O(n)复杂度的遍历,返回目标函数的空接口对象,需要User自行进行类型断言
func (delegator *delegator) DeleteFunc(name string) bool {
	return delegator.funcs.deleteFunc(name)
} //执行一次O(n)的搜索,删除name匹配的函数,返回bool

//================================================================
