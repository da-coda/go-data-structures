package bst

type BinarySearchTree struct {
	root *node
}

// Creates a new Binary Search Tree.
func NewBinarySearchTree() BinarySearchTree {
	return BinarySearchTree{}
}

func (b BinarySearchTree) Print(space int, spaceSize int) {
	b.root.Print(space, spaceSize)
}

func (b BinarySearchTree) Height() int {
	return b.root.Height()
}

func (b BinarySearchTree) Width() int {
	return b.root.Width()
}

func (b *BinarySearchTree) Add(key int, payload interface{}) error {
	return b.root.Add(key, payload)
}

func (b *BinarySearchTree) Remove(key int) error {
	return b.root.Remove(key, b.root)
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
