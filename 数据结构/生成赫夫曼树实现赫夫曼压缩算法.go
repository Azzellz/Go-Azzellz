package main

import (
	"fmt"
	"strconv"
)

type sortInterface interface {
	// Len Len方法返回集合中的元素个数
	Len() int
	// Less Less方法报告索引i的元素是否比索引j的元素小
	Less(i, j int) bool
	// Swap Swap方法交换索引i和j的两个元素
	Swap(i, j int)
}

type huffmanQueue struct {
	arr []huffmanNode
	len int
}

func (q *huffmanQueue) add(node huffmanNode) {
	q.arr = append(q.arr, node)
	q.len++
}
func (q *huffmanQueue) out() (node huffmanNode) {
	q.len--
	node = q.arr[0]
	q.arr = q.arr[1:]
	return
}
func (q *huffmanQueue) huffmanConv() {
	//fmt.Println(q.arr)
	for q.len > 1 {
		//sort.Stable(q)
		for i := 0; i < q.len-1; i++ {
			for j := 0; j < q.len-1-i; j++ {
				if q.arr[j].weight > q.arr[j+1].weight {
					q.arr[j], q.arr[j+1] = q.arr[j+1], q.arr[j]
				}
			}
		}
		//将队列的第一个和第二个元素弹出,对应着两个最小值,然后按两个最小权值之和生成一个新的节点
		leftNode := q.out()
		rightNode := q.out()
		newNode := huffmanNode{weight: leftNode.weight + rightNode.weight, left: &leftNode, right: &rightNode}
		q.add(newNode)
	}
}
func (q *huffmanQueue) Len() int           { return len(q.arr) }
func (q *huffmanQueue) Less(i, j int) bool { return q.arr[i].weight-q.arr[j].weight < 0 }
func (q *huffmanQueue) Swap(i, j int)      { q.arr[i], q.arr[j] = q.arr[j], q.arr[i] }

type huffmanNode struct {
	value  rune
	weight int //权值
	left   *huffmanNode
	right  *huffmanNode
}

type huffmanTree struct {
	root *huffmanNode
}

func (node *huffmanNode) getCodes(code string, str string, formMap map[rune]string) {
	str += code
	if node == nil {
		return
	}
	if node.value == 0 {
		node.left.getCodes("0", str, formMap)
		node.right.getCodes("1", str, formMap)
	} else {
		formMap[node.value] = str
	} //有效数据

} //用来得到对应的赫夫曼编码
func transferStrToHuffmanArr(str string) ([]rune, map[rune]string) {
	//先将字符串转化为rune切片
	runeArr := make([]rune, 0)
	for _, v := range str {
		runeArr = append(runeArr, v)
	}
	//将rune切片的值转移到map中
	countMap := make(map[rune]int)
	for _, v := range runeArr {
		if _, ok := countMap[v]; ok {
			countMap[v]++
		} else {
			countMap[v] = 1
		}
	}

	//遍历map生成node队列
	keySlice := make([]rune, 0)
	for i := range countMap {
		keySlice = append(keySlice, i)
	}
	//fmt.Println(keySlice)
	for i := 0; i < len(keySlice)-1; i++ {
		for j := 0; j < len(keySlice)-i-1; j++ {
			if keySlice[j] > keySlice[j+1] {
				keySlice[j], keySlice[j+1] = keySlice[j+1], keySlice[j]
			}
		}
	}
	//fmt.Println(keySlice)
	nodeQueue := huffmanQueue{arr: make([]huffmanNode, 0), len: 0}
	for _, v := range keySlice {
		nodeQueue.add(huffmanNode{value: v, weight: countMap[v]})
	}

	//将node队列转化成赫夫曼树
	nodeQueue.huffmanConv()
	//nodeQueue.arr[0].preOrder()
	//生成赫夫曼编码对照表
	formMap := make(map[rune]string)
	//要解决map无序遍历的问题,将key值取出来排序,然后按顺序从map中取value

	for _, v := range keySlice {
		formMap[v] = ""
	}

	root := nodeQueue.out()
	fmt.Println(root.left, root.left.left)

	root.getCodes("", "", formMap) //会导致无序
	//按runeArr拼接成一个长字符串
	huffmanStr := ""
	for _, v := range runeArr {
		huffmanStr += formMap[v]
	}

	huffmanRune := make([]rune, 0)
	for len(huffmanStr) > 8 {
		temp, _ := strconv.ParseInt(huffmanStr[:8], 2, 8)
		huffmanRune = append(huffmanRune, rune(temp))
		huffmanStr = huffmanStr[8:]
		//fmt.Println(huffmanStr)
	}
	temp, _ := strconv.ParseInt(huffmanStr, 2, 8)
	huffmanRune = append(huffmanRune, rune(temp))
	//fmt.Println(transferRunesToStr(huffmanRune))
	return huffmanRune, formMap

	//fmt.Println(huffmanByte, len(huffmanByte))
}
func transferRunesToStr(arr []rune) (res string) {
	for _, v := range arr {
		res += strconv.FormatInt(int64(v), 2)
	}
	return
}
func (tree *huffmanTree) buildTree(arr []int) {

	//先生成一组包含赫夫曼节点的队列
	nodeQueue := huffmanQueue{arr: make([]huffmanNode, 0), len: 0}
	for _, v := range arr {
		nodeQueue.add(huffmanNode{weight: v})
	}
	//对生成的队列进行排序

	//for nodeQueue.len > 1 {
	//	sort.Sort(&nodeQueue)
	//	//将队列的第一个和第二个元素弹出,对应着两个最小值,然后按两个最小权值之和生成一个新的节点
	//	leftNode := nodeQueue.out()
	//	rightNode := nodeQueue.out()
	//	newNode := huffmanNode{weight: leftNode.weight + rightNode.weight, left: &leftNode, right: &rightNode}
	//	nodeQueue.add(newNode)
	//}
	nodeQueue.huffmanConv()
	temp := nodeQueue.out()
	tree.root = &temp

} //生成赫夫曼树的主方法
func (tree *huffmanTree) preOrder() {
	if tree.root != nil {
		tree.root.preOrder()
	} else {
		fmt.Println("无效树")
	}
}
func (node *huffmanNode) preOrder() {
	if node == nil {
		return
	}
	fmt.Println(node.weight)
	node.left.preOrder()
	node.right.preOrder()
}
func huffmanDecode(arr []rune, formMap map[rune]string) {
	strKey := transferRunesToStr(arr)
	//反转一下对照表的key和value
	newFormMap := make(map[string]rune)
	for i, v := range formMap {
		newFormMap[v] = i
	}
	//按照新表解析字符串
	resStr := ""
	count := 0
	tempLen := len(strKey)
	for i := 0; i <= tempLen; i++ {
		out, ok := newFormMap[strKey[:count]]
		if ok {
			resStr += string(out)
			strKey = strKey[count:]
			count = 0
		}
		count++
	}
	fmt.Println(resStr)
}

func main() {
	//tree := new(huffmanTree)
	//arr := []int{13, 7, 8, 3, 29, 6, 1}
	//fmt.Println()
	//tree.buildTree(arr)
	//tree.root.preOrder()
	str := "you"
	arr, formMap := transferStrToHuffmanArr(str)
	//fmt.Println(arr)
	huffmanDecode(arr, formMap)

	//temp := "10101000"
	//fore, _ := strconv.Atoi(temp)
	//fmt.Println(byte(strconv.FormatInt(int64(fore), 2)))
}
