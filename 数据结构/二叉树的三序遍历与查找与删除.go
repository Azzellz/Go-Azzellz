package main

import (
	"fmt"
)

type binaryTree struct {
	root       *binaryTreeNode
	searchFlag bool
}

type binaryTreeNode struct {
	id    int
	name  string
	left  *binaryTreeNode
	right *binaryTreeNode
}

func (tree *binaryTree) preOrder(root *binaryTreeNode) {
	if tree.root == nil {
		fmt.Println("无效二叉树,空根节点")
		return
	}
	fmt.Println(root.id) //遍历操作
	if root.left != nil {
		tree.preOrder(root.left)
	}
	if root.right != nil {
		tree.preOrder(root.right)
	}

} //二叉树的前序遍历

func (tree *binaryTree) infixOrder(root *binaryTreeNode) {
	if tree.root == nil {
		fmt.Println("无效二叉树,空根节点")
		return
	}

	if root.left != nil {
		tree.preOrder(root.left)
	}
	fmt.Println(root.id) //遍历操作
	if root.right != nil {
		tree.preOrder(root.right)
	}
}

func (tree *binaryTree) postOrder(root *binaryTreeNode) {
	if tree.root == nil {
		fmt.Println("无效二叉树,空根节点")
		return
	}

	if root.left != nil {
		tree.preOrder(root.left)
	}
	if root.right != nil {
		tree.preOrder(root.right)
	}
	fmt.Println(root.id) //遍历操作
}

func (tree *binaryTree) preOrderSearch(root *binaryTreeNode, id int) bool {
	if root == nil {
		fmt.Println("无效二叉树,空根节点")
		return false
	}

	flag := false
	fmt.Println("searching....")
	if root.id == id {
		fmt.Println("找到了", root.name)
		return true
	}

	if root.left != nil {
		flag = tree.preOrderSearch(root.left, id)
	}
	if flag {
		return true
	}

	if root.right != nil {
		tree.preOrderSearch(root.right, id)
	}
	if flag {
		return true
	}

	return false

}

func (tree *binaryTree) infixOrderSearch(root *binaryTreeNode, id int) bool {
	if root == nil {
		fmt.Println("无效二叉树,空根节点")
		return false
	}

	flag := false
	fmt.Println("searching....")

	if root.left != nil {
		flag = tree.infixOrderSearch(root.left, id)
	}
	if flag {
		return true
	}

	if root.id == id {
		fmt.Println("找到了", root.name)
		return true
	}

	if root.right != nil {
		flag = tree.infixOrderSearch(root.right, id)
	}
	if flag {
		return true
	}

	return false

}

func (tree *binaryTree) postOrderSearch(root *binaryTreeNode, id int) bool {

	if root == nil {
		fmt.Println("无效二叉树,空根节点")
		return false
	}
	flag := false

	fmt.Println("searching....")

	if root.left != nil {
		flag = tree.postOrderSearch(root.left, id)
	}
	if flag {
		return true
	}

	if root.right != nil {
		flag = tree.postOrderSearch(root.right, id)
	}
	if flag {
		return true
	}

	if root.id == id {
		fmt.Println("找到了", root.name)
		//os.Exit(1)
		return true
	} else {
		return false
	}
}

func (tree *binaryTree) deleteNode(root *binaryTreeNode, id int) {
	if root == nil {
		fmt.Println("无效树")
		return
	}
	if tree.root.id == id {
		tree.root = nil
		return
	} //置空整个树

	if root.left != nil && root.left.id == id {
		fmt.Println("成功删除", root.left.name)
		root.left = nil
		return
	}
	if root.right != nil && root.right.id == id {
		fmt.Println("成功删除", root.right.name)
		root.right = nil
		return
	}

	if root.left != nil {
		tree.deleteNode(root.left, id)
	}
	if root.right != nil {
		tree.deleteNode(root.right, id)
	}
}

func main() {

	tree := binaryTree{root: &binaryTreeNode{id: 0, name: "homo0"}}
	n1 := &binaryTreeNode{id: 1, name: "homo1"}
	n2 := &binaryTreeNode{id: 2, name: "homo2"}
	n3 := &binaryTreeNode{id: 3, name: "homo3"}
	tree.root.left = n1
	tree.root.right = n2
	n2.left = n3

	//遍历
	//tree.preOrder(tree.root)
	//tree.infixOrder(tree.root)
	//tree.postOrder(tree.root)

	//查找
	//tree.preOrderSearch(tree.root, 1)
	//tree.infixOrderSearch(tree.root, 2)
	tree.postOrderSearch(tree.root, 3)
	tree.deleteNode(tree.root, 3)
	tree.postOrderSearch(tree.root, 3)

}
