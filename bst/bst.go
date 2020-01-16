package bst

type BinarySearchTree struct {
	root *node
}

// Creates a new Binary Search Tree.
// A BST can only exists when it has at least a root node
func NewBinarySearchTree(key int, payload interface{}) BinarySearchTree {
	return BinarySearchTree{&node{key: key, payload: payload}}
}

func (b BinarySearchTree) Print(space int, spaceSize int) {
	b.root.Print(space, spaceSize)
}

func (b BinarySearchTree) Depth() int {
	return b.root.Depth()
}

func (b BinarySearchTree) Width() int {
	return b.root.Width()
}

func (b *BinarySearchTree) Add(key int, payload interface{}) error {
	return b.root.Add(key, payload)
}

func (b BinarySearchTree) Get(key int) interface{} {
	return b.root.Get(key)
}

func (b BinarySearchTree) HasNext() bool {
	return b.root.HasNext()
}

func (b BinarySearchTree) HasLeft() bool {
	return b.root.HasLeft()
}

func (b BinarySearchTree) HasRight() bool {
	return b.root.HasRight()
}
