package cores

import (
	"fmt"
	"reflect"
	"strconv"
)

func EqualsReflection(value, other any) bool {
	return reflect.DeepEqual(value, other)
}

func PassTypeIndirectReflection(t reflect.Type) reflect.Type {
	switch t.Kind() {
	case reflect.Pointer:
		return PassTypeIndirectReflection(t.Elem())
	default:
		return t
	}
}

func TypeEqualsReflection(value reflect.Type, other reflect.Type) bool {
	return PassTypeIndirectReflection(value) == PassTypeIndirectReflection(other)
}

func GetValueReflection(value any) reflect.Value {
	var val reflect.Value
	switch value.(type) {
	case reflect.Type:
		panic("should not pass reflect type")
	case reflect.Value:
		val = value.(reflect.Value)
	default:
		val = reflect.ValueOf(value)
	}
	return val
}

func IsValidReflection(value any) bool {
	val := GetValueReflection(value)
	// must be not zero value
	if val.IsValid() {
		switch val.Kind() {
		case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice:
			// chan, func, interface, map, pointer, slice
			// will be considered as nullable value
			return !val.IsNil()
		default:
			// not chan, func, interface, map, pointer, slice
			// will be considered as notnull value
			return true
		}
	}
	// zero or null value
	return false
}

func PassValueIndirectReflection(value any) reflect.Value {
	val := GetValueReflection(value)

	if !IsValidReflection(val) {
		return val
	}

	switch val.Kind() {
	case reflect.Interface, reflect.Pointer:
		return PassValueIndirectReflection(val.Elem())
	default:
		return val
	}
}

func GetTypeReflection(value any) reflect.Type {
	return PassValueIndirectReflection(value).Type()
}

func GetKindReflection(value any) reflect.Kind {
	return PassValueIndirectReflection(value).Kind()
}

func IsCountableReflection[T any](value T) bool {
	val := PassValueIndirectReflection(value)
	if !IsValidReflection(val) {
		panic("invalid value")
	}
	switch val.Kind() {
	case reflect.Array, reflect.Slice, reflect.String, reflect.Chan:
		return true
	default:
		return false
	}
}

func GetSizeReflection(value any) int {
	val := PassValueIndirectReflection(value)
	if !IsCountableReflection(val) {
		panic("value is not countable")
	}
	return val.Len()
}

func IsZeroOrEmptyReflection(value any) bool {
	if !IsValidReflection(value) {
		return true
	}
	if IsCountableReflection(value) {
		return GetSizeReflection(value) == 0
	}
	return false
}

func IsStringableReflection(value any) bool {
	val := PassValueIndirectReflection(value)
	if !IsValidReflection(val) {
		panic("invalid value")
	}
	switch val.Kind() {
	case reflect.String:
		return true
	default:
		return false
	}
}

func IsExportedFieldReflection(value any) bool {
	val := PassValueIndirectReflection(value)
	if !IsValidReflection(val) {
		return false
	}
	return val.CanInterface()
}

func IsExportedFieldByIndexReflection(value any, index int) bool {
	val := PassValueIndirectReflection(value)
	if !IsValidReflection(val) {
		return false
	}
	return IsExportedFieldReflection(val.Field(index))
}

func IsExportedFieldByNameReflection(value any, name string) bool {
	val := PassValueIndirectReflection(value)
	if !IsValidReflection(val) {
		return false
	}
	return IsExportedFieldReflection(val.FieldByName(name))
}

func ToStringReflection(value any) string {
	if value == nil {
		return "<null>"
	}

	if val, ok := value.(StringableImpl); ok {
		return val.ToString()
	}

	val := PassValueIndirectReflection(value)
	if !IsValidReflection(val) {
		return "<null>"
	}

	switch val.Kind() {
	case reflect.Bool:
		v := val.Bool()
		return strconv.FormatBool(v)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v := val.Int()
		return strconv.FormatInt(v, 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v := val.Uint()
		return strconv.FormatUint(v, 10)
	case reflect.Uintptr:
		v := Unwrap(Cast[uintptr](val.Interface()))
		return ToString(v)
	case reflect.Float32, reflect.Float64:
		v := val.Float()
		return strconv.FormatFloat(v, 'f', -1, 64)
	case reflect.Complex64, reflect.Complex128:
		v := val.Complex()
		return strconv.FormatComplex(v, 'f', -1, 128)
	case reflect.String:
		return val.String()
	case reflect.Struct:
		if IsTimeUtcISO8601(val) {
			return strconv.Quote(ToTimeUtcStringISO8601(val))
		}
		return "<struct>"
	case reflect.Array, reflect.Slice:
		return "<array>"
	case reflect.Map:
		return "<map>"
	default:
		return fmt.Sprint(val.Interface())
	}
}

func PassValueIndirect[T any](value any) T {
	var ok bool
	var temp T
	KeepVoid(ok, temp)
	val := PassAnyValueIndirect(value)

	if val == nil {
		return temp
	}

	if temp, ok = Cast[T](val); ok {
		return temp
	}

	panic(NewThrow("invalid data type", ErrDataTypeInvalid))
}

func PassAnyValueIndirect(value any) any {
	val := PassValueIndirectReflection(value)
	if !IsValidReflection(val) {
		return nil
	}

	return val.Interface()
}
