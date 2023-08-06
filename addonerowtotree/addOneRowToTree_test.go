package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddOneRow(t *testing.T) {
	bt4 := TreeNode{Val: 4}
	bt2 := TreeNode{Val: 2}
	bt6 := TreeNode{Val: 6}
	bt3 := TreeNode{Val: 3}
	bt1 := TreeNode{Val: 1}
	bt5 := TreeNode{Val: 5}

	bt4.Left = &bt2
	bt4.Right = &bt6
	bt2.Left = &bt3
	bt2.Right = &bt1
	bt6.Left = &bt5

	root := AddOneRowToTree(&bt4, 1, 2)

	res := make([]int, 0)
	PrintBinaryTree(root, &res)
	expected := []int{4, 1, 2, 3, 1, 1, 6, 5}
	assert.Equal(t, expected, res)

}
