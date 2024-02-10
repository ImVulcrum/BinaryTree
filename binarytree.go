package main

import (
	"fmt"

	"./binarytree"
)

func main() {
	var tree binarytree.Tree = binarytree.NewTree()

	//tree.InsertList([]int{7, 5, 3, 9, 4, 8, 2})
	tree.InsertList([]int{26, 73, 91, 45, 60, 42, 32, 81, 84, 37, 43, 70, 94, 36, 51, 66, 19, 24, 99, 54, 75, 67, 88, 46, 56})
	//tree.InsertList([]int{62, 80, 50, 40, 70, 90, 60, 70, 30, 80, 59, 61, 89, 91})

	fmt.Println(tree)
}
