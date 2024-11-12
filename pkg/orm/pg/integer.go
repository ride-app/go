package pg

import (
	"reflect"
)

type integer[T int] struct {
}

func Int() integer[int] {
	return integer[int]{}
}

func (f integer[T]) sqlValue() string {
	return "INTEGER"
}

func (f integer[T]) structMemberType() reflect.Type {
	var zero [0]T
	return reflect.TypeOf(zero).Elem()
}
