package bst

import "errors"

/*
	Compares calling object key to given object key. Returned key shows if greater, equal or smaller
	and also the difference between both values

	returns positive int if calling object is greater
	returns 0 if calling object is equal
	returns negative if calling object is smaller
*/
func (N node) CompareTo(comparable Comparable) (int, error) {
	node, ok := comparable.(node)
	if !ok {
		return 0, errors.New("can't compare node type to non-node type")
	}
	return N.key - node.key, nil
}

func (N node) GreaterThen(comparable Comparable) (bool, error) {
	node, ok := comparable.(node)
	if !ok {
		return false, errors.New("can't compare node type to non-node type")
	}
	return N.key > node.key, nil
}

func (N node) LesserThen(comparable Comparable) (bool, error) {
	node, ok := comparable.(node)
	if !ok {
		return false, errors.New("can't compare node type to non-node type")
	}
	return N.key < node.key, nil
}

func (N node) EqualTo(comparable Comparable) (bool, error) {
	node, ok := comparable.(node)
	if !ok {
		return false, errors.New("can't compare node type to non-node type")
	}
	return N.key == node.key, nil
}
