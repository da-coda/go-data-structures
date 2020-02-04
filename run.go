package main

import (
	"github.com/da-coda/go-data-structures/bst"
)

func main() {
	root := bst.NewBinarySearchTree(10, "root")
	_ = root.Add(8, "Hello")
	_ = root.Add(13, "World")
	_ = root.Add(3, "World")
	_ = root.Add(10, "World")
	_ = root.Add(20, "World")
	_ = root.Add(17, "World")
	_ = root.Add(9, "World")
	root.Print(0, root.Height())
	_ = root.Remove(8)
	root.Print(0, root.Height())

}
