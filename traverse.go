package main

func (N Node) HasNext() bool {
	return N.HasLeft() || N.HasRight()
}

func (N Node) HasLeft() bool {
	return N.left != nil
}

func (N Node) HasRight() bool {
	return N.right != nil
}