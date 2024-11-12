package pg

import (
	"fmt"
	"reflect"
)

type varChar[T string] struct {
	length int
}

func VarChar(length int) varChar[string] {
	return varChar[string]{
		length: length,
	}
}

func (f varChar[T]) sqlValue() string {
	return fmt.Sprintf("VARCHAR(%d)", f.length)
}

func (f varChar[T]) structMemberType() reflect.Type {
	var zero [0]T
	return reflect.TypeOf(zero).Elem()
}
