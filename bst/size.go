package bst

import "math"

func (N Node) Depth() int {
	var left int
	var right int
	if !N.HasNext() {
		return 1
	}
	if N.HasLeft() {
		left = N.left.Depth()
	}
	if N.HasRight() {
		right = N.right.Depth()
	}
	if left > right {
		return 1 + left
	} else {
		return 1 + right
	}
}

func (N Node) Width() int {
	size := N.Depth()
	return int(math.Pow(2.0, float64(size)))
}
