package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	addOper = 0
	subOper = 0
	mulOper = 1
	divOper = 1
) //定义四个运算符的优先级

type numStack struct {
	arr []int
	len int
} //数栈
func (s *numStack) push(n int) {
	s.arr = append(s.arr, n)
	s.len++
}
func (s *numStack) pop() (temp int) {
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

type oper struct {
	stand    int //象征的符号
	priority int //运算优先级
}
type operStack struct {
	arr []oper
	len int
} //符号栈
type intStack struct {
	arr []int
	len int
}
type stringStack struct {
	arr []string
	len int
}

func (s *stringStack) push(n string) {
	s.arr = append(s.arr, n)
	s.len++
}
func (s *stringStack) pop() (temp string) {
	if s.len == 0 {
		fmt.Printf("Nothing to pop from stack")
		return
	}
	temp = s.arr[s.len-1]
	s.arr = s.arr[:s.len-1]
	s.len--
	return
}
func (s *intStack) push(n int) {
	s.arr = append(s.arr, n)
	s.len++
}
func (s *intStack) pop() (temp int) {
	if s.len == 0 {
		fmt.Printf("Nothing to pop from stack")
		return
	}
	temp = s.arr[s.len-1]
	s.arr = s.arr[:s.len-1]
	s.len--
	return
}
func (s *operStack) push(n oper) {
	s.arr = append(s.arr, n)
	s.len++
}
func (s *operStack) pop() (temp oper) {
	if s.len == 0 {
		fmt.Printf("Nothing to pop from stack")
		return
	}
	temp = s.arr[s.len-1]
	s.arr = s.arr[:s.len-1]
	s.len--
	return
}
func (c *calculator) comparePriority(temp oper) {

	if c.operStack.len == 0 {
		c.operStack.push(temp)
		return
	}                                                                 //空符号栈,直接入
	if temp.priority <= c.operStack.arr[c.operStack.len-1].priority { //从数栈弹出两个数,再从符号栈弹出一个符号,进行运算后将结果压入数栈,再将这个符号压入栈中.
		c.operCal()
		c.operStack.push(temp)
	} else {
		c.operStack.push(temp)
	} //直接将符号压入符号栈中
}

type calculator struct {
	operStack
	numStack
	intStack
	stringStack
}

func (c *calculator) operCal() {
	n1 := c.numStack.pop()
	n2 := c.numStack.pop()
	//注意运算顺序是n2对n1
	opr := c.operStack.pop()
	var result int
	switch opr.stand {
	case '+':
		result = n2 + n1
	case '-':
		result = n2 - n1
	case '*':
		result = n2 * n1
	case '/':
		result = n2 / n1
	}
	c.numStack.push(result)
}

func (c *calculator) calToEnd() {
	for {
		if c.numStack.len == 1 {
			break
		}
		c.operCal()
	}
} //完全运算,使数栈中只剩最后一个结果数字

func (c *calculator) Calculate(s string) int {
	//遍历字符串中的每个字符,判断入哪个栈
	//以下是中缀表达式的实现

	//var temp int
	//for i := 0; i <= len(s)-1; i++ {
	//	if c.judgeIsNumber(int(s[i])) { //最后一位需要考虑
	//		temp *= 10
	//		temp += int(s[i] - '0')
	//		if i == len(s)-1 {
	//			c.numStack.push(temp)
	//		} //最后一个元素直接入栈
	//		continue
	//	} //如果下一个还是数字的话就拼接成二位数
	//	if temp != 0 {
	//		c.numStack.push(temp)
	//		temp = 0 //temp要清空
	//	} //说明是多位数
	//	c.switchOprAndPush(int(s[i]))
	//}
	//c.calToEnd()
	//return c.numStack.pop()

	//逆波兰(后缀表达式)实现
	//为方便区分,导入的字符串是按空格划分的,所以要先调用strings包里面的一个方法来划分空格得到字符串切片,之后遍历字符串切片
	arrStr := strings.Split(s, " ")
	for i := 0; i < len(arrStr); i++ {
		if c.isStrOpr(arrStr[i]) {
			c.calByRPN(arrStr[i][0])
			continue
		}
		temp, err := strconv.Atoi(arrStr[i])
		if err != nil {
			fmt.Println(err)
			break
		}
		c.intStack.push(temp)
	}
	return c.intStack.pop()
} //总方法

func (c *calculator) judgeIsNumber(n int) bool {
	if n <= '9' && n >= '0' {
		return true
	}
	return false
}

func (c *calculator) calByRPN(opr byte) {
	n1 := c.intStack.pop()
	n2 := c.intStack.pop()
	var result int
	switch opr {
	case '+':
		result = n2 + n1
	case '-':
		result = n2 - n1
	case '*':
		result = n2 * n1
	case '/':
		result = n2 / n1
	}
	c.intStack.push(result)
}

func (c *calculator) judgeIsOpr(s int) bool {
	switch s {
	case '+':
		return true
	case '-':
		return true
	case '*':
		return true
	case '/':
		return true
	default:
		return false
	}
}

func (c *calculator) isStrOpr(s string) bool {
	return s == "+" || s == "*" || s == "-" || s == "/" || s == "(" || s == ")"
}

func (c *calculator) isCharOpr(s byte) bool {
	return s == '+' || s == '*' || s == '-' || s == '/' || s == '(' || s == ')'
}

func (c *calculator) switchOprAndPushByRPN(s int) {

	switch {
	case s == '+':
		opr := oper{stand: s, priority: addOper}
		c.operStack.push(opr)
	case s == '-':
		opr := oper{stand: s, priority: subOper}
		c.operStack.push(opr)
	case s == '*':
		opr := oper{stand: s, priority: mulOper}
		c.operStack.push(opr)
	case s == '/':
		opr := oper{stand: s, priority: divOper}
		c.operStack.push(opr)
	default:
		fmt.Println("本计算器还不认得这个符号哦,请换个能识别的吧-")
	}

} //判断入哪个栈

func (c *calculator) switchOprAndPush(s int) {

	switch {
	case s == '+':
		opr := oper{stand: s, priority: addOper}
		c.comparePriority(opr)
	case s == '-':
		opr := oper{stand: s, priority: subOper}
		c.comparePriority(opr)
	case s == '*':
		opr := oper{stand: s, priority: mulOper}
		c.comparePriority(opr)
	case s == '/':
		opr := oper{stand: s, priority: divOper}
		c.comparePriority(opr)
	case s <= '9' && s >= '0':
		c.numStack.push(s - '0')
	default:
		fmt.Println("本计算器还不认得这个符号哦,请换个能识别的吧-")
	}

} //判断入哪个栈

func (c *calculator) switchOprForInfixToSuffix(s string) oper {

	switch {
	case s == "+":
		opr := oper{priority: addOper}
		return opr
	case s == "-":
		opr := oper{priority: subOper}
		return opr
	case s == "*":
		opr := oper{priority: mulOper}
		return opr
	case s == "/":
		opr := oper{priority: divOper}
		return opr
	case s == ")" || s == "(":
		opr := oper{priority: -1}
		return opr
	default:
		fmt.Println("本计算器还不认得这个符号哦,请换个能识别的吧-")
	}
	return *new(oper) //随便做的

} //判断入哪个栈

func (c *calculator) toStringArr(s string) []string {
	//第一个问题,先把传进来的字符串变成字符串切片
	var array []string
	temp := 0
	rightKuoHaoNum := 0 //右括号数

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			array = append(array, string(s[i]))
			continue
		} else if s[i] == ')' {
			if rightKuoHaoNum >= 1 {
				//遇到末尾的右括号
				if i == len(s)-1 {
					array = append(array, strconv.Itoa(temp))
					array = append(array, string(s[i]))
					rightKuoHaoNum++
					temp = 0
					continue
				}
				if temp != 0 {
					array = append(array, strconv.Itoa(temp))
					array = append(array, string(s[i]))
					rightKuoHaoNum++
					temp = 0
					continue
				}
				array = append(array, string(s[i]))
				rightKuoHaoNum++
				temp = 0
				continue
			}
			array = append(array, strconv.Itoa(temp))
			array = append(array, string(s[i]))
			rightKuoHaoNum++
			temp = 0
			continue
		}
		if c.isCharOpr(s[i]) { //字符直接添加
			if temp == 0 { //右括号出来的
				array = append(array, string(s[i]))
				continue
			}
			array = append(array, strconv.Itoa(temp))
			array = append(array, string(s[i]))
			temp = 0
			continue
		}
		//说明是数字
		temp *= 10
		temp += int(s[i]) - '0'
		if i == len(s)-1 {
			array = append(array, strconv.Itoa(temp))
		}
	}
	return array
}

func (c *calculator) infixToSuffix(s string) string {
	//s是传进来的表达式
	arr := c.toStringArr(s)
	duringTimeStack := stringStack{arr: make([]string, 0)} //存放中间结果的栈
	oprStack := stringStack{arr: make([]string, 0)}        //存放操作符的栈
	for i := 0; i < len(arr); i++ {
		if !c.isStrOpr(arr[i]) { //是数字,直接入栈
			duringTimeStack.push(arr[i])
			continue
		}
		//遇到运算符
		if c.isStrOpr(arr[i]) {
			//遇到括号的两种逻辑
			if oprStack.len == 0 || arr[i] == "(" { //是左括号或者符号栈为空则直接入栈
				oprStack.push(arr[i])
				continue
			}
			if arr[i] == ")" { //遇到右括号
				for oprStack.arr[oprStack.len-1] != "(" {
					duringTimeStack.push(oprStack.pop())
				} //如果操作符栈顶是左括号时停止,然后在后面弹出左括号,消除这对括号
				oprStack.pop()
				continue
			}
			//遇到其他的算术运算符,比较一下优先级
			opr1 := c.switchOprForInfixToSuffix(arr[i]) //遍历到当前符号的优先级
			for {
				//先创建一个运算符对象,这是为了比较优先级
				if oprStack.len == 0 {
					oprStack.push(arr[i])
					break
				}
				opr2 := c.switchOprForInfixToSuffix(oprStack.arr[oprStack.len-1]) //栈顶符号的优先级
				if opr1.priority > opr2.priority {                                //直接入栈
					oprStack.push(arr[i])
					break
				} else {
					duringTimeStack.push(oprStack.pop())
				}
			}
		}
	}
	//弹出符号栈的剩余元素入中间栈
	for oprStack.len != 0 {
		duringTimeStack.push(oprStack.pop())
	}
	//反序拼接中间栈
	se := ""
	for i := 0; i < duringTimeStack.len; i++ {
		se += duringTimeStack.arr[i] + " "
	}
	return se[:len(se)-1]
}

func main() {
	//要召唤一个计算器
	//先得献祭两个栈
	numstack := numStack{arr: make([]int, 0)}
	oprstack := operStack{arr: make([]oper, 0)}
	//发生献祭事件
	//2030年计算机降临于Goland
	calculator := calculator{operStack: oprstack, numStack: numstack}
	s := "1+((2+3)*4)-5"
	fmt.Println(calculator.Calculate("3 4 + 5 * 6 -"))
	fmt.Println(calculator.toStringArr(s))
	fmt.Println(calculator.infixToSuffix(s))
	fmt.Println(calculator.Calculate(calculator.infixToSuffix(s)))

	fmt.Println("good!")

}
