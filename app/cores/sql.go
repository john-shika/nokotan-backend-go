package cores

import (
	"database/sql"
	"time"
)

func NewNull[T any](value T) sql.Null[T] {
	return sql.Null[T]{V: value, Valid: true}
}

func NewNullString(value string) sql.NullString {
	return sql.NullString{String: value, Valid: true}
}

func NewNullTime(value time.Time) sql.NullTime {
	return sql.NullTime{Time: value, Valid: true}
}

func NewNullBool(value bool) sql.NullBool {
	return sql.NullBool{Bool: value, Valid: true}
}

func NewNullByte(value byte) sql.NullByte {
	return sql.NullByte{Byte: value, Valid: true}
}

func NewNullInt32(value int32) sql.NullInt32 {
	return sql.NullInt32{Int32: value, Valid: true}
}

func NewNullInt64(value int64) sql.NullInt64 {
	return sql.NullInt64{Int64: value, Valid: true}
}

func NewNullFloat64(value float64) sql.NullFloat64 {
	return sql.NullFloat64{Float64: value, Valid: true}
}
