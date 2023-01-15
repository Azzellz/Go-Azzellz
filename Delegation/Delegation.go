package Delegation

import "fmt"

// Delegate 本包尝试对委托进行实现

const (
	CarryWithRemain   = iota //执行完保留函数
	CarryWithUnremain        //执行完不保留函数
) //定义宏来区分执行模式

type VoidFunc func() //代表一类空返回值的函数类型

type Delegator struct {
	voidFunc voidFuncQueue
} //Delegator用于维护多种函数执行队列

type voidFuncQueue struct {
	voidFuncList []VoidFunc
	len          int
} //voidFuncQueue用于实现一个基于Go语言切片的队列

//voidFuncQueue的实现方法

func (queue *voidFuncQueue) addFunc(voidFunc VoidFunc) {
	if queue.voidFuncList == nil {
		queue.voidFuncList = make([]VoidFunc, 0)
		queue.voidFuncList = append(queue.voidFuncList, voidFunc)
		queue.len++
		return
	} //如果未初始化queue中的函数切片,则要先帮忙先初始化一下,顺便完成添加第一个函数的操作

	queue.voidFuncList = append(queue.voidFuncList, voidFunc)
	queue.len++

} //函数入队

func (queue *voidFuncQueue) doFuncUnremainly(doNum int) {
	if queue.len == 0 {
		fmt.Println("Empty delegator")
	}
	if doNum > queue.len {
		tempLen := queue.len
		for i := 0; i < tempLen; i++ {
			temp := queue.voidFuncList[0]
			queue.voidFuncList = queue.voidFuncList[1:]
			queue.len--
			temp()
		}
	} else {
		for i := 0; i < doNum; i++ {
			temp := queue.voidFuncList[0]
			queue.voidFuncList = queue.voidFuncList[1:]
			queue.len--
			temp()
		}
	}

} //函数出队

func (queue *voidFuncQueue) doFuncRemainly(doNum int) {
	if queue.len == 0 {
		fmt.Println("Empty delegator")
	}

	if doNum > queue.len { //完全执行
		for _, v := range queue.voidFuncList {
			v()
		}
	} else {
		for i := 0; i < doNum; i++ {
			queue.voidFuncList[i]()
		}
	}

} //遍历执行队列,不出队

//================================================================

//Delegator的应用主方法:包含添加函数,执行函数,查看函数列表
//可拓展内容:保留函数执行和不保留执行,可以使用传入本包内的宏来判定执行模式

func (delegator *Delegator) AddFunc(addFunc VoidFunc) {
	delegator.voidFunc.addFunc(addFunc)
}

func (delegator *Delegator) CarryFunc(carryNum int, carryMode int) {
	//不保留执行,执行完函数被消耗
	if carryMode == CarryWithUnremain {
		delegator.voidFunc.doFuncUnremainly(carryNum)
	}
	//================================================================

	//保留执行,执行完函数不被消耗
	if carryMode == CarryWithRemain {
		delegator.voidFunc.doFuncRemainly(carryNum)
	}
	//================================================================

} //执行n次函数,如果n大于委托内函数的数量,则会完全执行委托内的所有函数

//================================================================
