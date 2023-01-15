package main

import "fmt"

type binarySortTreeNode struct {
	weight int
	left   *binarySortTreeNode
	right  *binarySortTreeNode
}

type binarySortTree struct {
	root *binarySortTreeNode
}

func (bstNode *binarySortTreeNode) addNode(node *binarySortTreeNode) {
	if node == nil || bstNode == nil {
		return
	}

	if node.weight < bstNode.weight {
		//放在左边
		if bstNode.left == nil {
			bstNode.left = node
		} else {
			bstNode.left.addNode(node)
		}
	} else {
		//放在右边
		if bstNode.right == nil {
			bstNode.right = node
		} else {
			bstNode.right.addNode(node)
		}
	}
}

// searchNode 从该节点开始查找,返回的一个目标节点,没有找到则会返回nil
func (bstNode *binarySortTreeNode) searchNode(weight int) *binarySortTreeNode {
	if bstNode == nil {
		fmt.Println("无效根节点")
		return nil
	}
	if bstNode.weight == weight {
		return bstNode
	} else if weight < bstNode.weight {
		//向左递归查找
		if bstNode.left == nil {
			return nil
		}
		return bstNode.left.searchNode(weight)
	} else {
		//向右递归查找
		if bstNode.right == nil {
			return nil
		}
		return bstNode.right.searchNode(weight)
	}
}

// searchParentNode 从该节点开始查找,返回的一个目标节点的父节点,没有找到则会返回nil
func (bstNode *binarySortTreeNode) searchParentNode(weight int) *binarySortTreeNode {
	if (bstNode.left != nil && bstNode.left.weight == weight) || (bstNode.right != nil && bstNode.right.weight == weight) {
		return bstNode
	}
	if weight < bstNode.weight && bstNode.left != nil {
		//向左边递归查找
		return bstNode.left.searchParentNode(weight)
	} else if weight >= bstNode.weight && bstNode.right != nil {
		//向右边找
		return bstNode.right.searchParentNode(weight)
	} else {
		//没有父节点
		return nil
	}

}

// deleteNode 从该节点开始找,删除一个目标节点
func (bstNode *binarySortTreeNode) deleteNode(weight int) int {
	if bstNode == nil {
		fmt.Println("无效树")
		return -1
	}
	if bstNode.weight == weight { //删除的是根节点
		bstNode = nil
		return -1
	}

	if bstNode.searchParentNode(weight) == nil { //发现二叉树只有一个节点
		bstNode = nil
		return -1
	}
	targetNode := bstNode.searchNode(weight)
	if targetNode == nil {
		fmt.Println("未找到目标节点")
		return 0
	}
	parentNode := bstNode.searchParentNode(weight)
	if targetNode.left == nil && targetNode.right == nil { //目标节点是叶子节点
		if parentNode.left == targetNode {
			parentNode.left = nil
		} else if parentNode.right == targetNode {
			parentNode.right = nil
		}
	} else if targetNode.left == nil || targetNode.right == nil { //目标节点是子树节点
		if parentNode.left == targetNode {
			if targetNode.left == nil {
				parentNode.left = targetNode.right
			} else if targetNode.right == nil {
				parentNode.left = targetNode.left
			}

		} else if parentNode.right == targetNode {
			if targetNode.left == nil {
				parentNode.right = targetNode.right
			} else if targetNode.right == nil {
				parentNode.right = targetNode.left
			}
		}
	} else if targetNode.left != nil && targetNode.right != nil {
		val := bstNode.deleteMinNode(targetNode)
		targetNode.weight = val
	}
	return 1

}
func (bstNode *binarySortTreeNode) deleteMinNode(node *binarySortTreeNode) int {
	temp := node
	for temp.left != nil {
		temp = temp.left
	} //到树的最左端就是最小值
	//此时temp指向最小的节点
	bstNode.deleteNode(temp.weight)
	return temp.weight
}
func (bstNode *binarySortTreeNode) infixOrder() {
	if bstNode == nil {
		return
	}
	if bstNode.left != nil {
		bstNode.left.infixOrder()
	}
	fmt.Println(bstNode.weight)
	if bstNode.right != nil {
		bstNode.right.infixOrder()
	}

}

func (tree *binarySortTree) deleteNode(weight int) {
	code := tree.root.deleteNode(weight)
	if code == -1 {
		tree.root = nil
	} //删除根节点
}

func main() {
	tree := binarySortTree{}
	arr := []int{8, 3, 4, 6, 7, 1, 10, 14}
	for i, v := range arr {
		if i == 0 {
			tree.root = &binarySortTreeNode{weight: v}
		} else {
			tree.root.addNode(&binarySortTreeNode{weight: v})
		}
	}
	tree.root.infixOrder()
	tree.deleteNode(1)
	fmt.Println(tree.root.searchNode(14))
}
