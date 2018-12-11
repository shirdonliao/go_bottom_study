//++++++++++++++++++++++++++++++++++++++++
//Fighting for great,share generate value!
//Build the best soft by golang,let's go!
//++++++++++++++++++++++++++++++++++++++++
//Author:ShirDon <http://www.shirdon.com>
//Email:hcbsts@163.com;  823923263@qq.com
//++++++++++++++++++++++++++++++++++++++++

package main

import "fmt"

type Object interface{}

type TreeNode struct {
	Data       Object
	LeftChild  *TreeNode
	RightChild *TreeNode
}

//(完全)二叉树结构
type Tree struct {
	RootNode *TreeNode
}

//追加元素 (广度优先，即按层级遍历后添加)
func (tr *Tree) Add(object Object) {
	node := &TreeNode{Data: object}
	if tr.RootNode == nil {
		tr.RootNode = node
		return
	}
	queue := []*TreeNode{tr.RootNode}
	for len(queue) != 0 {
		cur_node := queue[0]
		queue = queue[1:]

		if cur_node.LeftChild == nil {
			cur_node.LeftChild = node
			return
		} else {
			queue = append(queue, cur_node.LeftChild)
		}
		if cur_node.RightChild == nil {
			cur_node.RightChild = node
			return
		} else {
			queue = append(queue, cur_node.RightChild)
		}
	}
}

//广度遍历
func (tr *Tree) BreadthTravel() {

	if tr.RootNode == nil {
		return
	}
	queue := []*TreeNode{}
	queue = append(queue, tr.RootNode)

	for len(queue) != 0 {
		//fmt.Printf("len(queue):%d\n", len(queue))
		cur_node := queue[0]
		queue = queue[1:]

		fmt.Printf("%v  ", cur_node.Data)

		if cur_node.LeftChild != nil {
			queue = append(queue, cur_node.LeftChild)
		}
		if cur_node.RightChild != nil {
			queue = append(queue, cur_node.RightChild)
		}
	}

}

/*
深度遍历:
1.先序遍历:根->左->右
2.中序遍历:左->中->右
3.后序遍历:左->右->根
*/

//先序遍历
func (tr *Tree) PreOrder(node *TreeNode) {
	if node == nil {
		return
	}

	fmt.Printf("%v  ", node.Data)

	//if node.LeftChild != nil {
	tr.PreOrder(node.LeftChild)
	//}
	//if node.RightChild != nil {
	tr.PreOrder(node.RightChild)
	//}
}


//中序遍历
func (tr *Tree) InOrder(node *TreeNode) {
	if node == nil {
		return
	}
	tr.InOrder(node.LeftChild)
	fmt.Printf("%v  ", node.Data)
	tr.InOrder(node.RightChild)
}

func (tr *Tree) PostOrder(node *TreeNode)  {
	if node == nil {
		return
	}
	tr.PostOrder(node.LeftChild)
	tr.PostOrder(node.RightChild)
	fmt.Printf("%v  ", node.Data)
}

func main()  {
	tree := Tree{}
	tree.Add(0)
	tree.Add(1)
	tree.Add(2)
	tree.Add(3)
	tree.Add(4)
	tree.Add(5)
	tree.Add(6)
	tree.Add(7)
	tree.Add(8)
	tree.Add(9)

	//广度优先遍历
	//tree.BreadthTravel()
	//fmt.Println("")

	//深度优先 先序遍历
	tree.PreOrder(tree.RootNode)
	fmt.Println("")

	//深度优先  中序遍历
	tree.InOrder(tree.RootNode)
	fmt.Println("")

	//深度优先  后序遍历
	tree.PostOrder(tree.RootNode)
	fmt.Println("")
}

//results:
//0  1  3  7  8  4  9  2  5  6
//7  3  8  1  9  4  0  5  2  6
//7  8  3  9  4  1  5  6  2  0