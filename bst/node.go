package bst

import "errors"

type node struct {
	key     int
	payload interface{}
	left    *node
	right   *node
}

func (N *node) Add(key int, payload interface{}) error {
	eq, err := N.EqualTo(node{key: key})

	if err != nil {
		return err
	}

	if eq {
		return nil
	}
	if greater, _ := N.LesserThen(node{key: key}); greater {
		if N.right == nil {
			N.right = &node{key: key, payload: payload}
		} else {
			return N.right.Add(key, payload)
		}
	} else {
		if N.left == nil {
			N.left = &node{key: key, payload: payload}
		} else {

			return N.left.Add(key, payload)
		}
	}

	return nil
}

func (N node) Get(key int) interface{} {
	if eq, _ := N.EqualTo(node{key: key}); eq {
		return N.payload
	}
	if !N.HasNext() {
		return nil
	}
	if greater, _ := N.LesserThen(node{key: key}); greater {
		if N.HasRight() {
			return N.right.Get(key)
		} else {
			return nil
		}
	} else {
		if N.HasLeft() {
			return N.left.Get(key)
		} else {
			return nil
		}
	}
}

func (N *node) Remove(key int, parent TreeNode) error {
	parentNode := parent.(*node)

	if N == nil {
		return errors.New("value to be deleted does not exist in the tree")
	}
	comp, _ := N.CompareTo(node{key: key})
	switch {
	case comp > 0:
		return N.left.Remove(key, N)
	case comp < 0:
		return N.right.Remove(key, N)
	default:
		if !N.HasNext() {
			return N.replaceNode(parentNode, nil)
		}
		if !N.HasLeft() {
			return N.replaceNode(parentNode, N.right)
		}
		if !N.HasRight() {
			return N.replaceNode(parentNode, N.left)
		}

		replacement, replParent := N.left.InorderSuccessor(N)
		N.key = replacement.key
		N.payload = replacement.payload

		return replacement.Remove(replacement.key, replParent)
	}

}

func (N *node) InorderSuccessor(parent *node) (*node, *node) {
	if N == nil {
		return nil, parent
	}
	if !N.HasRight() {
		return N, parent
	} else {
		return N.right.InorderSuccessor(N)
	}
}

func (N *node) replaceNode(parent, replacement *node) error {
	if N == nil {
		return errors.New("replaceNode() not allowed on a nil node")
	}

	if N == parent.left {
		parent.left = replacement
		return nil
	}
	parent.right = replacement
	return nil
}
