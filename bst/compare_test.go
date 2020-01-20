package bst

import "testing"

func Test_node_CompareTo(t *testing.T) {
	type fields struct {
		key     int
		payload interface{}
		left    *node
		right   *node
	}
	type args struct {
		comparable Comparable
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{
			"greater by 5",
			fields{5, nil, nil, nil},
			args{node{0, nil, nil, nil}},
			5,
			false,
		}, {
			"equal",
			fields{5, nil, nil, nil},
			args{node{5, nil, nil, nil}},
			0,
			false,
		}, {
			"lesser by 5",
			fields{0, nil, nil, nil},
			args{node{5, nil, nil, nil}},
			-5,
			false,
		}, {
			"non-node comparable",
			fields{0, nil, nil, nil},
			args{mockComparable{}},
			0,
			true,
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
			got, err := N.CompareTo(tt.args.comparable)
			if (err != nil) != tt.wantErr {
				t.Errorf("CompareTo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CompareTo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_node_EqualTo(t *testing.T) {
	type fields struct {
		key     int
		payload interface{}
		left    *node
		right   *node
	}
	type args struct {
		comparable Comparable
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			"equal",
			fields{5, nil, nil, nil},
			args{node{5, nil, nil, nil}},
			true,
			false,
		}, {
			"not equal",
			fields{5, nil, nil, nil},
			args{node{0, nil, nil, nil}},
			false,
			false,
		}, {
			"non-node comparable",
			fields{0, nil, nil, nil},
			args{mockComparable{}},
			false,
			true,
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
			got, err := N.EqualTo(tt.args.comparable)
			if (err != nil) != tt.wantErr {
				t.Errorf("EqualTo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("EqualTo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_node_GreaterThen(t *testing.T) {
	type fields struct {
		key     int
		payload interface{}
		left    *node
		right   *node
	}
	type args struct {
		comparable Comparable
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			"greater",
			fields{5, nil, nil, nil},
			args{node{0, nil, nil, nil}},
			true,
			false,
		}, {
			"lesser",
			fields{-5, nil, nil, nil},
			args{node{0, nil, nil, nil}},
			false,
			false,
		}, {
			"non-node comparable",
			fields{0, nil, nil, nil},
			args{mockComparable{}},
			false,
			true,
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
			got, err := N.GreaterThen(tt.args.comparable)
			if (err != nil) != tt.wantErr {
				t.Errorf("GreaterThen() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GreaterThen() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_node_LesserThen(t *testing.T) {
	type fields struct {
		key     int
		payload interface{}
		left    *node
		right   *node
	}
	type args struct {
		comparable Comparable
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			"lesser",
			fields{5, nil, nil, nil},
			args{node{10, nil, nil, nil}},
			true,
			false,
		}, {
			"greater",
			fields{5, nil, nil, nil},
			args{node{0, nil, nil, nil}},
			false,
			false,
		}, {
			"non-node comparable",
			fields{0, nil, nil, nil},
			args{mockComparable{}},
			false,
			true,
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
			got, err := N.LesserThen(tt.args.comparable)
			if (err != nil) != tt.wantErr {
				t.Errorf("LesserThen() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("LesserThen() got = %v, want %v", got, tt.want)
			}
		})
	}
}

type mockComparable struct {
}

func (m mockComparable) CompareTo(comparable Comparable) (int, error) {
	panic("implement me")
}

func (m mockComparable) GreaterThen(comparable Comparable) (bool, error) {
	panic("implement me")
}

func (m mockComparable) LesserThen(comparable Comparable) (bool, error) {
	panic("implement me")
}

func (m mockComparable) EqualTo(comparable Comparable) (bool, error) {
	panic("implement me")
}
