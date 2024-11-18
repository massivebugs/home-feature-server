package util

import (
	"reflect"
	"testing"
)

func TestSet_Add(t *testing.T) {
	type args struct {
		s Set[any]
		v any
	}

	tests := []struct {
		name string
		args args
		want Set[any]
	}{
		{
			name: "success value is added",
			args: args{
				s: NewSet[any](),
				v: "test_value",
			},
			want: Set[any]{
				items: map[any]bool{"test_value": true},
			},
		},
		{
			name: "success duplicate value is added",
			args: args{
				s: Set[any]{
					items: map[any]bool{"test_value": true},
				},
				v: "test_value",
			},
			want: Set[any]{
				items: map[any]bool{"test_value": true},
			},
		},
		{
			name: "success struct value is added",
			args: args{
				s: NewSet[any](),
				v: struct{ Test string }{Test: "test_value"},
			},
			want: Set[any]{
				items: map[any]bool{struct{ Test string }{Test: "test_value"}: true},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.s.Add(tt.args.v)

			if !reflect.DeepEqual(tt.args.s, tt.want) {
				t.Errorf("Set.Add() = %v, want %v", tt.args.s, tt.want)
			}
		})
	}
}

func TestSet_Remove(t *testing.T) {
	sp := &struct{ Test string }{Test: "test_value"}

	type args struct {
		s Set[any]
		v any
	}

	tests := []struct {
		name string
		args args
		want Set[any]
	}{
		{
			name: "success value is removed",
			args: args{
				s: Set[any]{
					items: map[any]bool{"test_value": true},
				},
				v: "test_value",
			},
			want: Set[any]{
				items: map[any]bool{},
			},
		},
		{
			name: "success value doesn't exist",
			args: args{
				s: Set[any]{
					items: map[any]bool{"test_value": true},
				},
				v: "test_value2",
			},
			want: Set[any]{
				items: map[any]bool{"test_value": true},
			},
		},
		{
			name: "success struct value is removed",
			args: args{
				s: Set[any]{
					items: map[any]bool{struct{ Test string }{Test: "test_value"}: true},
				},
				v: struct{ Test string }{Test: "test_value"},
			},
			want: Set[any]{
				items: map[any]bool{},
			},
		},
		{
			name: "success struct pointer is removed",
			args: args{
				s: Set[any]{
					items: map[any]bool{sp: true},
				},
				v: sp,
			},
			want: Set[any]{
				items: map[any]bool{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.s.Remove(tt.args.v)

			if !reflect.DeepEqual(tt.args.s, tt.want) {
				t.Errorf("Set.Remove() = %v, want %v", tt.args.s, tt.want)
			}
		})
	}
}

func TestSet_Exists(t *testing.T) {
	sp := &struct{ Test string }{Test: "test_value"}

	type args struct {
		s Set[any]
		v any
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success value exists",
			args: args{
				s: Set[any]{
					items: map[any]bool{"test_value": true},
				},
				v: "test_value",
			},
			want: true,
		},
		{
			name: "success value doesn't exist",
			args: args{
				s: Set[any]{
					items: map[any]bool{"test_value": true},
				},
				v: "test_value2",
			},
			want: false,
		},
		{
			name: "success struct value exists",
			args: args{
				s: Set[any]{
					items: map[any]bool{struct{ Test string }{Test: "test_value"}: true},
				},
				v: struct{ Test string }{Test: "test_value"},
			},
			want: true,
		},
		{
			name: "success struct pointer exists",
			args: args{
				s: Set[any]{
					items: map[any]bool{sp: true},
				},
				v: sp,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.s.Exists(tt.args.v)

			if got != tt.want {
				t.Errorf("Set.Exists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Size(t *testing.T) {
	type args struct {
		s Set[any]
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success Set has values",
			args: args{
				s: Set[any]{
					items: map[any]bool{
						"test_value":  true,
						"test_value2": true,
					},
				},
			},
			want: 2,
		},
		{
			name: "success Set is empty",
			args: args{
				s: Set[any]{},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.s.Size()

			if got != tt.want {
				t.Errorf("Set.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_ToSlice(t *testing.T) {

	type args struct {
		s Set[any]
	}

	tests := []struct {
		name string
		args args
		want []any
	}{
		{
			name: "success Set has values",
			args: args{
				s: Set[any]{
					items: map[any]bool{
						"test_value": true,
					},
				},
			},
			want: []any{
				"test_value",
			},
		},
		{
			name: "success Set has no values",
			args: args{
				s: Set[any]{
					items: map[any]bool{},
				},
			},
			want: []any{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.s.ToSlice()

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Set.ToSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
