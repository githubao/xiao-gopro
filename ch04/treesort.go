// 树结构
// author: baoqiang
// time: 2019/3/4 下午7:11
package ch04

import "fmt"

type tree struct {
	value       int
	left, right *tree
}

func RunTreeSort() {
	vals := []int{2, 1, 3}
	TreeSort(vals)
	fmt.Println(vals)
}

func TreeSort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

// 递归把左边的加入，加入自己，把右边的加入
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

// 小值插在左边，大值插在右边
func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}

	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}

	return t
}
