package main

import (
	"./binarytree"
)

func main() {
	var tree binarytree.Tree = binarytree.NewTree()

	// tree.InsertList([]int{7, 5, 3, 999999, 4, 8, 2})
	// tree.InsertList([]int{62, 80, 50, 40, 70, 90, 60, 70, 30, 50, 80, 59, 61, 89, 91})
	tree.InsertList([]int{62, 80, 50, 40, 70, 90, 60, 70, 30, 80, 59, 61, 89, 91})

	tree.Print(100)
}
