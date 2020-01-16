package main

import "errors"

/*
	Compares calling object key to given object key. Returned key shows if greater, equal or smaller
	and also the difference between both values

	returns positive int if calling object is greater
	returns 0 if calling object is equal
	returns negative if calling object is smaller
*/
func (N Node) CompareTo(comparable Comparable) (int, error) {
	node, ok := comparable.(Node)
	if !ok {
		return 0, errors.New("can't compare Node type to non-Node type")
	}
	return N.key - node.key, nil
}

func (N Node) GreaterThen(comparable Comparable) (bool, error) {
	node, ok := comparable.(Node)
	if !ok {
		return false, errors.New("can't compare Node type to non-Node type")
	}
	return N.key > node.key, nil
}

func (N Node) LesserThen(comparable Comparable) (bool, error) {
	node, ok := comparable.(Node)
	if !ok {
		return false, errors.New("can't compare Node type to non-Node type")
	}
	return N.key < node.key, nil
}

func (N Node) EqualTo(comparable Comparable) (bool, error) {
	node, ok := comparable.(Node)
	if !ok {
		return false, errors.New("can't compare Node type to non-Node type")
	}
	return N.key == node.key, nil
}