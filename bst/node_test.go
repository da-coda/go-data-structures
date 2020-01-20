package bst

import (
	"reflect"
	"testing"
)

func Test_node_Add(t *testing.T) {
	type fields struct {
		key     int
		payload interface{}
		left    *node
		right   *node
	}
	type args struct {
		key     int
		payload interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		left    bool
	}{
		{
			"add left",
			fields{2, nil, nil, nil},
			args{1, nil},
			false,
			true,
		}, {
			"add right",
			fields{2, nil, nil, nil},
			args{3, nil},
			false,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			N := &node{
				key:     tt.fields.key,
				payload: tt.fields.payload,
				left:    tt.fields.left,
				right:   tt.fields.right,
			}
			if err := N.Add(tt.args.key, tt.args.payload); (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.left && !N.HasLeft() {
				t.Errorf("Add() didn't add left node")
			} else if !tt.left && !N.HasRight() {
				t.Errorf("Add() didn't add right node")

			}
		})
	}
}

func Test_node_Get(t *testing.T) {
	type fields struct {
		key     int
		payload interface{}
		left    *node
		right   *node
	}
	type args struct {
		key int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
	}{
		{
			"nil",
			fields{5, nil, nil, nil},
			args{1},
			nil,
		},
		{
			"get Node",
			fields{5, nil, &node{3, "Hello World", nil, nil}, nil},
			args{3},
			"Hello World",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			N := &node{
				key:     tt.fields.key,
				payload: tt.fields.payload,
				left:    tt.fields.left,
				right:   tt.fields.right,
			}
			if got := N.Get(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
