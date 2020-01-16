package main

import (
	"fmt"
	"github.com/da-coda/go-data-structures/bst"
)

func main() {
	root := bst.NewBinarySearchTree(10, "root")
	_ = root.Add(8, "Hello")
	_ = root.Add(13, "World")
	root.Print(0, root.Depth())
	fmt.Print(root.Get(8))
}
