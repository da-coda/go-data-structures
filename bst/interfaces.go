package bst

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

type Traversable interface {
	HasNext() bool
	HasLeft() bool
	HasRight() bool
}

type Manipulable interface {
	Add(key int, payload interface{}) error
}

type Accessible interface {
	Get(key int) interface{}
}

type Tree interface {
	Printer
	Sizer
	Manipulable
	Accessible
	Traversable
}

type TreeNode interface {
	Comparable
	Tree
}
