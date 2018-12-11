//++++++++++++++++++++++++++++++++++++++++
//Fighting for great,share generate value!
//Build the best soft by golang,let's go!
//++++++++++++++++++++++++++++++++++++++++
//Author:ShirDon <http://www.shirdon.com>
//Email:hcbsts@163.com;  823923263@qq.com
//++++++++++++++++++++++++++++++++++++++++

package main

import (
	"fmt"
)

type tree struct {
	left  *tree
	right *tree
	data  int
}

func create(index int, value []int) (T *tree) {
	T = &tree{}
	T.data = value[index-1]
	fmt.Printf("index %v \n", index)
	if index < len(value)-1 && 2*index <= len(value) && 2*index+1 <= len(value) {
		T.left = create(2*index, value)
		T.right = create(2*index+1, value)
	}
	return T

}

func show(treeNode *tree) {
	if treeNode != nil {
		fmt.Printf("%v ", treeNode.data)
		if treeNode.left != nil {
			show(treeNode.left)
		}
		if treeNode.right != nil {
			show(treeNode.right)
		}
	} else {
		return
	}
}

func main() {
	value := []int{1, 2, 3, 4, 5}
	TreeRoot := create(1, value)
	show(TreeRoot)

}
