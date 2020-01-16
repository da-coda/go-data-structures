package main

type Comparable interface {
	CompareTo(comparable Comparable) (int, error)
	GreaterThen(comparable Comparable) (bool, error)
	LesserThen(comparable Comparable) (bool, error)
	EqualTo(comparable Comparable) (bool, error)
}

type Printer interface {
	Print(space int, spaceSize int)
}

type Sizer interface {
	Depth() int
	Width() int
}

type TraversableTreeNode interface {
	HasNext() bool
	HasLeft() bool
	HasRight() bool
}

type TreeNode interface {
	Comparable
	Printer
	Sizer
	TraversableTreeNode
	Add(key int, payload interface{}) error
	Get(key int) interface{}
}