package main

import "fmt"

type Treenode struct {
	value       int
	Left, Right *Treenode
}

// func (root *Treenode) traverse() {
// 	if root == nil {
// 		return
// 	}
// 	root.Left.traverse()
// 	fmt.Print(root.value)
// 	root.Right.traverse()
// }
func creatTreeNode(value int) *Treenode {
	return &Treenode{value: value}
}
func main() {
	// 创建struct
	var root2 Treenode
	root2 = Treenode{value: 3}
	root2.Left = &Treenode{}
	root2.Right = &Treenode{5, nil, nil}
	root2.Right.Left = new(Treenode)

	// nodes := []Treenode{
	// 	{value: 3},
	// 	{},
	// 	{6, nil, &root2},
	// }
	// fmt.Println(nodes)
	// 工厂函数
	root2.Left.Right = creatTreeNode(2)
	fmt.Println(root2.Left)
}
