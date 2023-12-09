package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 3}
	list := inorderTraversal(root)
	fmt.Println(list)
}

func inorderTraversal(root *TreeNode) []int {
	list := make([]int, 0)
	if root == nil {
		return list
	}
	return getTraversalOrder(root, list)
}

func getTraversalOrder(root *TreeNode, list []int) []int {
	if root.Left == nil && root.Right == nil {
		list = append(list, root.Val)
		return list
	} else if root.Left == nil {
		list = append(list, root.Val)
		return getTraversalOrder(root.Right, list)
	} else if root.Right == nil {
		l1 := getTraversalOrder(root.Left, list)
		l1 = append(l1, root.Val)
		return l1
	} else {
		l1 := getTraversalOrder(root.Left, list)
		l1 = append(l1, root.Val)
		l2 := getTraversalOrder(root.Right, l1)
		return l2
	}
}
