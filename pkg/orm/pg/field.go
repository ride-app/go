package pg

import (
	"reflect"
)

type Field interface {
	structMemberType() reflect.Type
	sqlValue() string
}
