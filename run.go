package main

func main() {
	root := Node{key: 10}
	_ = root.Add(8, "Hello")
	_ = root.Add(13, "World")
	root.Print(0, root.Depth())
}
