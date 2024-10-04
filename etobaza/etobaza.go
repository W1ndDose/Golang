package main

import (
	"fmt"
)

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

type BinarySearchTree struct {
	root *TreeNode
}

func (bst *BinarySearchTree) Insert(key int) {
	if bst.root == nil {
		bst.root = &TreeNode{val: key}
	} else {
		bst.insertRecursive(bst.root, key)
	}
}

func (bst *BinarySearchTree) insertRecursive(node *TreeNode, key int) {
	if key < node.val {
		if node.left == nil {
			node.left = &TreeNode{val: key}
		} else {
			bst.insertRecursive(node.left, key)
		}
	} else {
		if node.right == nil {
			node.right = &TreeNode{val: key}
		} else {
			bst.insertRecursive(node.right, key)
		}
	}
}

func (bst *BinarySearchTree) inOrderTraversal(node *TreeNode, result *[]int) {
	if node != nil {
		bst.inOrderTraversal(node.left, result)
		*result = append(*result, node.val)
		bst.inOrderTraversal(node.right, result)
	}
}

func (bst *BinarySearchTree) GetIthNode(i int) (int, bool) {
	result := []int{}
	bst.inOrderTraversal(bst.root, &result)
	if i >= 0 && i < len(result) {
		return result[i], true
	}
	return 0, false
}

func main() {

	bst := &BinarySearchTree{}
	nodes := []int{17, 6, 5, 20, 19, 18, 11, 14, 12, 13, 2, 4, 10}

	for _, node := range nodes {
		bst.Insert(node)
	}

	i := 5
	if ithNode, ok := bst.GetIthNode(i); ok {
		fmt.Printf("%d-й узел по порядку: %d\n", i, ithNode)
	} else {
		fmt.Printf("%d-й узел по порядку не существует\n", i)
	}
}
