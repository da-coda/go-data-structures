package bst

import "math"

func (N node) Height() int {
	var left int
	var right int
	if !N.HasNext() {
		return 1
	}
	if N.HasLeft() {
		left = N.left.Height()
	}
	if N.HasRight() {
		right = N.right.Height()
	}
	if left > right {
		return 1 + left
	} else {
		return 1 + right
	}
}

func (N node) Width() int {
	size := N.Height()
	return int(math.Pow(2.0, float64(size)))
}
