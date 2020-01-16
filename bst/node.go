package bst

type Node struct {
	key     int
	payload interface{}
	left    *Node
	right   *Node
}

func NewRootNode(key int, payload interface{}) Node {
	return Node{key: key, payload: payload}
}

func (N *Node) Add(key int, payload interface{}) error {
	eq, err := N.EqualTo(Node{key: key})

	if err != nil {
		return err
	}

	if eq {
		return nil
	}
	greater, _ := N.LesserThen(Node{key: key})
	if greater {
		if N.right == nil {
			N.right = &Node{key: key, payload: payload}
		} else {
			return N.right.Add(key, payload)
		}
	} else {
		if N.left == nil {
			N.left = &Node{key: key, payload: payload}
		} else {

			return N.left.Add(key, payload)
		}
	}

	return nil
}
