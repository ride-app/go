package pg

import (
	"fmt"
	"reflect"
)

type table struct {
	name       string
	sqlValue   string
	selectType reflect.Type
}

func Table(name string, columns map[string]*column) *table {
	table := table{
		name: name,
	}

	table.sqlValue = fmt.Sprintf("CREATE TABLE %s (", name)

	var fields []reflect.StructField

	for fieldName, column := range columns {
		table.sqlValue = fmt.Sprintf("%s %s,", table.sqlValue, column.sqlValue)
		fields = append(fields, reflect.StructField{Name: fieldName, Type: column.field.structMemberType()})
	}

	table.sqlValue = fmt.Sprintf("%s );", table.sqlValue)

	table.selectType = reflect.StructOf(fields)

	return &table
}

func sdzzsd() {
	table := Table("test", map[string]*column{
		"d": Column("test", VarChar(100)),
		"v": Column("test 1", Int()),
	})

	n := reflect.New(reflect.TypeOf(table.selectType)).Elem().Addr().Interface()

}
