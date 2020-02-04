package bst

import "testing"

func Test_node_Height(t *testing.T) {
	type fields struct {
		key     int
		payload interface{}
		left    *node
		right   *node
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			"Height 3",
			fields{0, nil, &node{left: &node{}}, nil},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			N := node{
				key:     tt.fields.key,
				payload: tt.fields.payload,
				left:    tt.fields.left,
				right:   tt.fields.right,
			}
			if got := N.Height(); got != tt.want {
				t.Errorf("Height() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_node_Width(t *testing.T) {
	type fields struct {
		key     int
		payload interface{}
		left    *node
		right   *node
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			N := node{
				key:     tt.fields.key,
				payload: tt.fields.payload,
				left:    tt.fields.left,
				right:   tt.fields.right,
			}
			if got := N.Width(); got != tt.want {
				t.Errorf("Width() = %v, want %v", got, tt.want)
			}
		})
	}
}
