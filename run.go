package main

import "datastructures/bst"

func main() {
	root := bst.NewRootNode(10, "root")
	_ = root.Add(8, "Hello")
	_ = root.Add(13, "World")
	root.Print(0, root.Depth())
}
