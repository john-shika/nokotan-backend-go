package cores

import (
	"fmt"
	"reflect"
	"strconv"
)

func EqualsReflect(value, other any) bool {
	return reflect.DeepEqual(value, other)
}

func PassTypeIndirectReflect(value any) reflect.Type {
	typ := GetTypeReflect(value)
	switch typ.Kind() {
	case reflect.Pointer:
		return PassTypeIndirectReflect(typ.Elem())
	default:
		return typ
	}
}

func TypeEqualsReflect(value, other any) bool {
	return PassTypeIndirectReflect(value) == PassTypeIndirectReflect(other)
}

func GetTypeReflect(value any) reflect.Type {
	var val reflect.Value
	switch value.(type) {
	case reflect.Type:
		return value.(reflect.Type)
	case reflect.Value:
		val = value.(reflect.Value)
	default:
		val = reflect.ValueOf(value)
	}
	return val.Type()
}

func GetValueReflect(value any) reflect.Value {
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

func GetKindReflect(value any) reflect.Kind {
	return PassValueIndirectReflect(value).Kind()
}

func IsValidReflect(value any) bool {
	val := GetValueReflect(value)
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

func PassValueIndirectReflect(value any) reflect.Value {
	val := GetValueReflect(value)

	if !IsValidReflect(val) {
		return val
	}

	switch val.Kind() {
	case reflect.Interface, reflect.Pointer:
		return PassValueIndirectReflect(val.Elem())
	default:
		return val
	}
}

func IsCountableReflect[T any](value T) bool {
	val := PassValueIndirectReflect(value)
	if !IsValidReflect(val) {
		panic("invalid value")
	}
	switch val.Kind() {
	case reflect.Array, reflect.Slice, reflect.String, reflect.Chan:
		return true
	default:
		return false
	}
}

func GetSizeReflect(value any) int {
	val := PassValueIndirectReflect(value)
	if !IsCountableReflect(val) {
		panic("value is not countable")
	}
	return val.Len()
}

func IsNoneOrEmptyReflect(value any) bool {
	if !IsValidReflect(value) {
		return true
	}
	if IsCountableReflect(value) {
		return GetSizeReflect(value) == 0
	}
	return false
}

func IsStringableReflect(value any) bool {
	val := PassValueIndirectReflect(value)
	if !IsValidReflect(val) {
		panic("invalid value")
	}
	switch val.Kind() {
	case reflect.String:
		return true
	default:
		return false
	}
}

func IsExportedFieldReflect(value any) bool {
	val := PassValueIndirectReflect(value)
	if !IsValidReflect(val) {
		return false
	}
	return val.CanInterface()
}

func IsExportedFieldAtReflect(value any, index int) bool {
	val := PassValueIndirectReflect(value)
	if !IsValidReflect(val) {
		return false
	}
	return IsExportedFieldReflect(val.Field(index))
}

func IsExportedFieldByNameReflect(value any, name string) bool {
	val := PassValueIndirectReflect(value)
	if !IsValidReflect(val) {
		return false
	}
	return IsExportedFieldReflect(val.FieldByName(name))
}

func ToStringReflect(value any) string {
	if value == nil {
		return "<null>"
	}

	if val, ok := value.(StringableImpl); ok {
		return val.ToString()
	}

	val := PassValueIndirectReflect(value)
	if !IsValidReflect(val) {
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
	val := PassValueIndirectReflect(value)
	if !IsValidReflect(val) {
		return nil
	}

	return val.Interface()
}

func GetNameTypeReflect(value any) string {
	val := PassValueIndirectReflect(value)
	if !IsValidReflect(val) {
		panic("invalid value")
	}
	method := val.MethodByName("GetNameType")
	if method.IsValid() {
		results := method.Call(nil)
		if len(results) != 1 || !results[0].IsValid() {
			panic("invalid results")
		}
		name := Unwrap(CastString(results[0].Interface()))
		return name
	}

	return val.Type().Name()
}
