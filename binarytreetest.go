package main

import (
	"fmt"

	"./binarytree"
)
		
		
func main () {
	var tree binarytree.Treee = binarytree.NewTree()
	tree.Insert(7)
	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(9)
	tree.Insert(4)
	tree.Insert(8)
	tree.Insert(2)

	fmt.Println(tree.String())
}
