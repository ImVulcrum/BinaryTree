package main

import (
	"fmt"

	"./binarytree"
)

func main() {
	var tree binarytree.Tree = binarytree.NewTree(true)

	//clstree.InsertList([]int{7, 5, 3, 9, 4, 8, 2})
	tree.InsertList([]int{67, 99, 20, 52, 55, 53, 80, 28, 14, 10, 51, 71, 16, 32, 73, 66, 98, 82, 21, 57, 34, 79, 40, 63, 94, 45, 25, 22, 84, 83})
	//tree.InsertList([]int{62, 80, 50, 40, 70, 90, 60, 70, 30, 80, 59, 61, 89, 91})

	fmt.Println(tree)
}
