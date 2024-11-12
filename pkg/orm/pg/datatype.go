package pg

type DataType[T string | int16 | int32 | int64 | float32 | float64 | bool] struct {
	sqlvalue string
}

type DataTypes interface {
	string | int16 | int32 | int64 | float32 | float64 | bool
}

func SmallInt() DataType[int16] {
	return DataType[int16]{
		sqlvalue: "SMALLINT",
	}
}

func BigInt() DataType[int64] {
	return DataType[int64]{
		sqlvalue: "BIGINT",
	}
}
