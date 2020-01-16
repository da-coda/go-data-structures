package bst

func (N node) HasNext() bool {
	return N.HasLeft() || N.HasRight()
}

func (N node) HasLeft() bool {
	return N.left != nil
}

func (N node) HasRight() bool {
	return N.right != nil
}
