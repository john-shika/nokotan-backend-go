package cores

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func Equals(value, other any) bool {
	return reflect.DeepEqual(value, other)
}

func GetReflectValue(value any) reflect.Value {
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

func IsReflectValid(value any) bool {
	val := GetReflectValue(value)
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

func ReflectPassValueIndirect(value any) reflect.Value {
	val := GetReflectValue(value)
	if !IsReflectValid(val) {
		panic("invalid value")
	}
	kind := val.Kind()
	if kind == reflect.Interface || kind == reflect.Pointer {
		return ReflectPassValueIndirect(val.Elem())
	}
	return val
}

func GetReflectType(value any) reflect.Type {
	return ReflectPassValueIndirect(value).Type()
}

func GetReflectKind(value any) reflect.Kind {
	return ReflectPassValueIndirect(value).Kind()
}

func IsReflectCountable[T any](value T) bool {
	val := ReflectPassValueIndirect(value)
	if !IsReflectValid(val) {
		panic("invalid value")
	}
	switch val.Kind() {
	case reflect.Array, reflect.Slice, reflect.String, reflect.Chan:
		return true
	default:
		return false
	}
}

func ReflectCount(value any) int {
	val := ReflectPassValueIndirect(value)
	if !IsReflectCountable(val) {
		panic("value is not countable")
	}
	return val.Len()
}

func IsReflectEmpty(value any) bool {
	return ReflectCount(value) == 0
}

func IsReflectStringable(value any) bool {
	val := ReflectPassValueIndirect(value)
	if !IsReflectValid(val) {
		panic("invalid value")
	}
	switch val.Kind() {
	case reflect.String:
		return true
	default:
		return false
	}
}

func ToReflectString(value any) string {
	val := ReflectPassValueIndirect(value)
	if !IsReflectStringable(val) {
		panic("value is not string")
	}
	return val.String()
}

func ReflectRepr(obj any) string {
	if obj == nil {
		return "null"
	}
	if serialize, ok := obj.(JsonableImpl); ok {
		return serialize.ToJson()
	}
	if val, ok := obj.(StringableImpl); ok {
		return strconv.Quote(val.ToString())
	}
	val := ReflectPassValueIndirect(obj)
	if IsReflectValid(val) {
		switch val.Kind() {
		case reflect.Bool:
			return strconv.FormatBool(val.Bool())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return strconv.FormatInt(val.Int(), 10)
		case reflect.Uint, reflect.Uintptr, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return strconv.FormatUint(val.Uint(), 10)
		case reflect.Float32, reflect.Float64:
			return strconv.FormatFloat(val.Float(), 'f', -1, 64)
		case reflect.Complex64, reflect.Complex128:
			return strconv.FormatComplex(val.Complex(), 'f', -1, 128)
		case reflect.String:
			return strconv.Quote(val.String())
		case reflect.Interface:
			return strconv.Quote(fmt.Sprint(val.Interface()))
		case reflect.Struct:
			typ := val.Type()
			numFields := typ.NumField()
			temp := make([]string, 0)
			for i := 0; i < numFields; i++ {
				structField := typ.Field(i)
				structTag := structField.Tag
				fieldKey := strings.ToLower(structField.Name)
				fieldVal := ReflectRepr(val.Field(i).Interface())
				if nameTag, ok := structTag.Lookup("name"); ok {
					fieldKey = nameTag
				}
				if jsonTag, ok := structTag.Lookup("json"); ok {
					KeepVoid(jsonTag, ok)

					// TODO: json tag not implemented yet

					//tokens := strings.Split(jsonTag, ",")
					//size := len(tokens)
					//for j := 0; j < size; j++ {
					//	token := strings.Trim(tokens[j], " ")
					//	switch token {
					//	case "-", "ignore", "ignored":
					//		continue
					//	case "omitempty", "notnull", "required":
					//		if fieldVal == "null" || fieldVal == "\"\"" {
					//			continue
					//		}
					//	}
					//}
				}
				entry := strconv.Quote(fieldKey) + ": " + fieldVal
				temp = append(temp, entry)
			}
			return "{" + strings.Join(temp, ", ") + "}"
		case reflect.Array, reflect.Slice:
			size := val.Len()
			values := make([]string, size)
			for i := 0; i < size; i++ {
				elem := val.Index(i).Interface()
				values[i] = ReflectRepr(elem)
			}
			return "[" + strings.Join(values, ", ") + "]"
		case reflect.Map:
			size := val.Len()
			iter := val.MapRange()
			values := make([]string, size)
			for i := 0; iter.Next(); i++ {
				key := iter.Key().Interface()
				value := iter.Value().Interface()
				values[i] = strconv.Quote(fmt.Sprint(key)) + ": " + ReflectRepr(value)
			}
			return "{" + strings.Join(values, ", ") + "}"
		default:
			panic("unknown type")
		}
	}
	return "undefined"
}
