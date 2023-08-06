package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PrintBinaryTree(root *TreeNode, arr *[]int) {

	*arr = append(*arr, root.Val)
	if root.Left != nil {
		PrintBinaryTree(root.Left, arr)
	}
	if root.Right != nil {
		PrintBinaryTree(root.Right, arr)
	}
}

type TreeDepth struct {
	Node  *TreeNode
	Depth int
}

func getAllParents(root *TreeNode, depth int) []*TreeNode {
	parents := make([]*TreeNode, 0)
	stack := make([]TreeDepth, 0)
	tt := TreeDepth{root, 1}
	stack = append(stack, tt)

	for len(stack) > 0 {
		d := stack[0].Depth
		node := stack[0].Node

		if d == depth-1 {
			parents = append(parents, node)
		}

		if node.Left != nil {
			stack = append(stack, TreeDepth{node.Left, d + 1})
		}
		if node.Right != nil {
			stack = append(stack, TreeDepth{node.Right, d + 1})
		}

		stack = stack[1:]
	}
	return parents
}

//	    4                4
//	  2    6    =>     1   1
//	3  1  5          2       6
//	               3  1     5
//
// value 1 depth 2
// just find the parents node(s)
// insert my new nodes
func AddOneRowToTree(root *TreeNode, value int, depth int) *TreeNode {

	if depth == 1 {
		newNode := &TreeNode{Val: value}
		newNode.Left = root
		root = newNode
		return root
	}
	parentNode := getAllParents(root, depth)
	for _, e := range parentNode {

		child := e.Left
		newNode := &TreeNode{Val: value}
		e.Left = newNode
		newNode.Left = child

		childr := e.Right
		newNodeR := &TreeNode{Val: value}
		e.Right = newNodeR
		newNodeR.Right = childr

	}
	return root
}

func main() {
	fmt.Println("Doing nothing")
}
