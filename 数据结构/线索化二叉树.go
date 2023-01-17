package main

import "fmt"

type threadedBinaryTreeNode struct {
	id        int
	name      string
	left      *threadedBinaryTreeNode
	right     *threadedBinaryTreeNode
	leftType  int //0表示指向左子树,1表示指向前驱节点
	rightType int //0表示指向右子树,1表示指向后继节点
}

type threadedBinaryTree struct {
	root *threadedBinaryTreeNode
}

var preNode *threadedBinaryTreeNode = nil

func (tree *threadedBinaryTree) infixOrderThreading(node *threadedBinaryTreeNode) {
	if tree.root == nil {
		fmt.Println("无效树")
		return
	}
	if node == nil {
		return
	}

	//左递归
	tree.infixOrderThreading(node.left)

	//线索化操作
	//前驱指向
	if node.left == nil {
		node.left = preNode
		node.leftType = 1
	}
	//后继指向
	if preNode != nil && preNode.right == nil {
		preNode.right = node
		preNode.rightType = 1
	}
	preNode = node //状态转移方程

	//右递归
	tree.infixOrderThreading(node.right)

} //中序线索化

func (tree *threadedBinaryTree) infixOrderTraversal(node *threadedBinaryTreeNode) {
	if tree.root == nil {
		fmt.Println("无效树")
		return
	}
	//node := tree.root
	//for node != nil {
	//
	//	for node.leftType == 0 {
	//		node = node.left
	//	}
	//	fmt.Println(node.id)
	//
	//	for node.rightType == 1 {
	//		node = node.right
	//		fmt.Println(node.id)
	//	}
	//
	//	node = node.right
	//
	//}

	if node == nil {
		return
	}

	if node.leftType == 0 {
		tree.infixOrderTraversal(node.left)
	}
	if node.rightType == 1 || node.leftType == 1 {
		fmt.Println(node.id)
		return
	}
	fmt.Println(node.id)
	if node.rightType == 0 {
		tree.infixOrderTraversal(node.right)
	}

}

func main() {

	n1 := &threadedBinaryTreeNode{id: 0}
	n2 := &threadedBinaryTreeNode{id: 1}
	n3 := &threadedBinaryTreeNode{id: 2}
	n4 := &threadedBinaryTreeNode{id: 3}
	n5 := &threadedBinaryTreeNode{id: 4}
	n6 := &threadedBinaryTreeNode{id: 5}
	tree := threadedBinaryTree{root: n1}
	n1.right = n3
	n1.left = n2
	n2.right = n5
	n2.left = n4
	n3.left = n6
	tree.infixOrderThreading(tree.root)
	fmt.Println(n4.right.id, n5.right.id)
	tree.infixOrderTraversal(tree.root)

}
