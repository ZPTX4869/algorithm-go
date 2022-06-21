package tree

import (
	"strconv"
	"strings"
)

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// BFS
func (this *Codec) serialize(root *TreeNode) string {
	res := make([]string, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr != nil {
			res = append(res, strconv.Itoa(curr.Val))
			queue = append(queue, curr.Left)
			queue = append(queue, curr.Right)
		} else {
			res = append(res, "null")
		}
	}

	return strings.Join(res, ",")
}

// BFS
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "null" {
		return nil
	}

	strs := strings.Split(data, ",")
	val, _ := strconv.Atoi(strs[0])
	root := &TreeNode{Val: val}
	queue := []*TreeNode{root}

	for i := 1; i < len(strs); i += 2 {
		curr := queue[0]
		queue = queue[1:]
		leftVal := strs[i]
		rightVal := strs[i+1]

		if leftVal != "null" {
			val, _ := strconv.Atoi(leftVal)
			curr.Left = &TreeNode{Val: val}
			queue = append(queue, curr.Left)
		}
		if rightVal != "null" {
			val, _ := strconv.Atoi(rightVal)
			curr.Right = &TreeNode{Val: val}
			queue = append(queue, curr.Right)
		}
	}

	return root
}

// DFS
func (this *Codec) serialize2(root *TreeNode) string {
	if root == nil {
		return "null"
	}

	return strconv.Itoa(root.Val) + "," + this.serialize2(root.Left) + "," + this.serialize2(root.Right)
}

// DFS
func (this *Codec) deserialize2(data string) *TreeNode {
	var build func(pStrs *[]string) *TreeNode
	build = func(pStrs *[]string) *TreeNode {
		rootVal := (*pStrs)[0]
		*pStrs = (*pStrs)[1:]
		if rootVal == "null" {
			return nil
		}

		val, _ := strconv.Atoi(rootVal)
		root := &TreeNode{Val: val}

		root.Left = build(pStrs)
		root.Right = build(pStrs)

		return root
	}

	strs := strings.Split(data, ",")
	return build(&strs)
}
