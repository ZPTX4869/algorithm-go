package trees

import (
	"math"
)

var lastVisit int

func init()  {
	lastVisit = math.MaxInt64
}

// 递归中序遍历
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if !isValidBST(root.Left) {
		return false
	}

	if root.Val <= lastVisit {
		return false
	} else {
		lastVisit = root.Val
	}

	if !isValidBST(root.Right) {
		return false
	}

	return true
}

// 非递归中序遍历
//func isValidBST(root *TreeNode) bool {
//	if root == nil {
//		panic("Root can not be empty")
//	}
//
//	var stack []*structure.TreeNode
//	lastValue := -math.MaxInt64
//
//	for len(stack) > 0 || root != nil {
//		for root != nil {
//			stack = append(stack, root)
//			root = root.Left
//		}
//
//		curNode := stack[len(stack)-1]
//		if curNode.Val < lastValue {
//			return false
//		} else {
//			lastValue = curNode.Val
//		}
//		if curNode.Right != nil {
//			root = curNode.Right
//		}
//		stack = stack[0 : len(stack)-1]
//	}
//
//	return true
//}

// 递归上下界判断
//func isValidBST(root *TreeNode) bool {
//	if root == nil {
//		panic("Root can not be empty")
//	}
//
//	return helper(root, -math.MaxInt64, math.MaxInt64)
//}

// 非递归中序遍历
//func helper(root *structure.TreeNode, lower, upper int) bool {
//	if root == nil {
//		return true
//	}
//
//	if root.Val < lower || root.Val > upper {
//		return false
//	}
//
//	return helper(root.Left, lower, root.Val) && helper(root.Right, root.Val, upper)
//}
