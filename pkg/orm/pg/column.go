package pg

import "fmt"

type column struct {
	name     string
	field    Field
	sqlValue string
}

func Column(name string, dataType Field) *column {
	return &column{
		name:     name,
		field:    dataType,
		sqlValue: fmt.Sprintf("%s %s", name, dataType.sqlValue()),
	}
}
