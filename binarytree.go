package main

import (
	"./binarytree"
)

func main() {
	var tree binarytree.Tree = binarytree.NewTree()

	tree.InsertList([]int{7, 5, 3, 9, 4, 8, 2})

	tree.Print(100)
}
