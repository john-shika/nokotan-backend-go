package cores

import (
	"strings"
)

// var ErrDataTypeInvalid = errors.New("exception: invalid data type")

var EmptyString string

func KeepVoid(_ ...any) {
	// do nothing...
}

func Default[T any]() T {
	var temp T
	return temp
}

func NoErr(err error) {
	if err != nil {
		panic(err)
	}
}

type ErrOrOkImpl interface {
	~bool | any
}

func IsOk[T ErrOrOkImpl](eOk T) bool {
	var ok bool
	var bOk bool
	var temp any
	KeepVoid(ok, bOk, temp)

	if temp, ok = CastAny(eOk); !ok {
		return false
	}

	if bOk, ok = CastBool(temp); !ok {
		return false
	}

	return bOk
}

func IsErr[T ErrOrOkImpl](eOk T) bool {
	var ok bool
	var err error
	var temp any
	KeepVoid(ok, err, temp)

	if temp, ok = CastAny(eOk); !ok {
		return false
	}

	if err, ok = CastErr(temp); !ok {
		return false
	}

	return err != nil
}

func CastErr[E ErrOrOkImpl](eOk E) (error, bool) {
	var ok bool
	var err error
	var temp any
	KeepVoid(ok, err, temp)

	if temp, ok = CastAny(eOk); !ok {
		return nil, false
	}

	if err, ok = temp.(error); !ok {
		return nil, false
	}

	return err, true
}

func Unwrap[T any, E ErrOrOkImpl](result T, eOk E) T {
	var ok bool
	var err error
	KeepVoid(ok, err)

	if err, ok = CastErr(eOk); !ok {
		if !IsOk(eOk) {
			panic("invalid data type")
		} else {
			return result
		}
	}
	
	NoErr(err)
	return result
}

type StackCollectionImpl[T any] interface {
	*T | []T
}

type MapCollectionImpl[T any] interface {
	StackCollectionImpl[T] | map[string]T
}

type Map[T any] map[string]T
type MapAny = Map[any]

func (m Map[T]) HasKey(key string) bool {
	var ok bool
	var value T
	KeepVoid(ok, value)
	if value, ok = m[key]; !ok {
		return false
	}
	return true
}

func (m Map[T]) GetValue(key string) T {
	var ok bool
	var value T
	KeepVoid(ok, value)
	if value, ok = m[key]; !ok {
		return Default[T]()
	}
	return value
}

func (m Map[T]) SetValue(key string, value T) bool {
	var ok bool
	var temp T
	KeepVoid(ok, temp, value)
	if temp, ok = m[key]; !ok {
		return false
	}
	m[key] = value
	return true
}

func Cast[T any](value any) (T, bool) {
	temp, ok := value.(T)
	return temp, ok
}

func CastAny(value any) (any, bool) {
	return Cast[any](value)
}

func CastBool(value any) (bool, bool) {
	return Cast[bool](value)
}

func CastInt8(value any) (int8, bool) {
	return Cast[int8](value)
}

func CastUint8(value any) (uint8, bool) {
	return Cast[uint8](value)
}

func CastInt16(value any) (int16, bool) {
	return Cast[int16](value)
}

func CastUint16(value any) (uint16, bool) {
	return Cast[uint16](value)
}

func CastInt(value any) (int, bool) {
	return Cast[int](value)
}

func CastUint(value any) (uint, bool) {
	return Cast[uint](value)
}

func CastInt32(value any) (int32, bool) {
	return Cast[int32](value)
}

func CastUint32(value any) (uint32, bool) {
	return Cast[uint32](value)
}

func CastInt64(value any) (int64, bool) {
	return Cast[int64](value)
}

func CastUint64(value any) (uint64, bool) {
	return Cast[uint64](value)
}

func CastFloat32(value any) (float32, bool) {
	return Cast[float32](value)
}

func CastFloat64(value any) (float64, bool) {
	return Cast[float64](value)
}

func CastString(value any) (string, bool) {
	return Cast[string](value)
}

func CastPtr[T any](value any) (*T, bool) {
	temp, ok := value.(*T)
	return temp, ok
}

func IsNoneOrEmpty(value string) bool {
	return value == "" || value == "\x00"
}

func IsNoneOrEmptyWhiteSpace(value string) bool {
	if IsNoneOrEmpty(value) {
		return true
	}
	temp := strings.TrimSpace(value)
	return temp == "" ||
		temp == "\x00" ||
		temp == "\xA0" ||
		temp == "\t" ||
		temp == "\r" ||
		temp == "\n" ||
		temp == "\r\n"
}
