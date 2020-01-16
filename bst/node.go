package bst

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

func (N *node) Get(key int) interface{} {
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
