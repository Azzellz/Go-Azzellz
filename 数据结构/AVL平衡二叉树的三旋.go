package main

import (
	"fmt"
	"math"
)

type AVLTreeNode struct {
	weight int
	left   *AVLTreeNode
	right  *AVLTreeNode
}

type AVLTree struct {
	root *AVLTreeNode
}

func (avlNode *AVLTreeNode) addNode(node *AVLTreeNode) {
	if node == nil || avlNode == nil {
		return
	}

	if node.weight < avlNode.weight {
		//放在左边
		if avlNode.left == nil {
			avlNode.left = node
		} else {
			avlNode.left.addNode(node)
		}
	} else {
		//放在右边
		if avlNode.right == nil {
			avlNode.right = node
		} else {
			avlNode.right.addNode(node)
		}
	}
	if avlNode.getRightHeight()-avlNode.getLeftHeight() > 1 {
		if avlNode.right != nil && avlNode.right.getLeftHeight() > avlNode.right.getRightHeight() {
			avlNode.right.rightRotate()
			avlNode.leftRotate()
		} else {
			avlNode.leftRotate()
		}
		return
	}
	if avlNode.getLeftHeight()-avlNode.getRightHeight() > 1 {
		if avlNode.left != nil && avlNode.left.getRightHeight() > avlNode.left.getLeftHeight() {
			avlNode.left.leftRotate()
			avlNode.rightRotate()
		} else {
			avlNode.rightRotate()
		}

	}
}
func (avlNode *AVLTreeNode) infixOrder() {
	if avlNode == nil {
		return
	}
	if avlNode.left != nil {
		avlNode.left.infixOrder()
	}
	fmt.Println(avlNode.weight)
	if avlNode.right != nil {
		avlNode.right.infixOrder()
	}
}
func (avlNode *AVLTreeNode) getLeftHeight() int {

	if avlNode.left == nil {
		return 0
	}
	return avlNode.left.getHeight()
}
func (avlNode *AVLTreeNode) getRightHeight() int {
	if avlNode.right == nil {
		return 0
	}
	return avlNode.right.getHeight()
}
func (avlNode *AVLTreeNode) getHeight() int {
	//func() float64 {
	//	if avlNode.left == nil {
	//		return 0
	//	} else {
	//		return float64(avlNode.left.getHeight())
	//	}
	//}()
	//func() float64 {
	//	if avlNode.right == nil {
	//		return 0
	//	} else {
	//		return float64(avlNode.right.getHeight())
	//	}
	//}()
	return int(math.Max(func() float64 {
		if avlNode.left == nil {
			return 0
		} else {
			return float64(avlNode.left.getHeight())
		}
	}(), func() float64 {
		if avlNode.right == nil {
			return 0
		} else {
			return float64(avlNode.right.getHeight())
		}
	}()) + 1)
}
func (avlNode *AVLTreeNode) leftRotate() {
	//1.创建一个新节点,这个节点的权为当前节点的权
	newNode := &AVLTreeNode{weight: avlNode.weight}
	//2.把新节点的左子树设置为当前节点的左子树
	newNode.left = avlNode.left
	//3.把新节点的右子树设置成当前节点的右子树的左子树
	newNode.right = avlNode.right.left
	//4.把当前节点的权替换成右子节点的值
	avlNode.weight = avlNode.right.weight
	//5.把当前节点的右子树设置成当前节点右子树的右子树
	avlNode.right = avlNode.right.right
	//6.把当前节点的左子树设置成新的节点
	avlNode.left = newNode

} //左旋
func (avlNode *AVLTreeNode) rightRotate() {
	//1.创建一个新节点,这个节点的权为当前节点的权
	newNode := &AVLTreeNode{weight: avlNode.weight}
	//2.把新节点的右子树设置为当前节点的右子树
	newNode.right = avlNode.right
	//3.把新节点的左子树设置成当前节点的左子树的右子树
	newNode.left = avlNode.left.right
	//4.把当前节点的权替换成左子节点的值
	avlNode.weight = avlNode.left.weight
	//5.把当前节点的左子树设置成当前节点左子树的左子树
	avlNode.left = avlNode.left.left
	//6.把当前节点的右子树设置成新的节点
	avlNode.right = newNode
} //右旋

func main() {
	avlTree := &AVLTree{}
	arr := []int{10, 11, 7, 6, 8, 9}
	for i, v := range arr {
		if i == 0 {
			avlTree.root = &AVLTreeNode{weight: v}
		} else {
			avlTree.root.addNode(&AVLTreeNode{weight: v})
		}
	}
	avlTree.root.infixOrder()
	fmt.Println(avlTree.root.getLeftHeight(), avlTree.root.getRightHeight(), avlTree.root.getHeight())
}
