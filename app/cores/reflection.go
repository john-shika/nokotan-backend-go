package cores

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"reflect"
	"strconv"
	"strings"
	"time"
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

var TimeTypeReflection = reflect.TypeOf(new(time.Time))
var NumericDateTypeReflection = reflect.TypeOf(new(jwt.NumericDate))

func IsDateTimeStringISO8601Reflection(value any) bool {
	val := PassValueIndirectReflection(value)
	if !IsValidReflection(val) {
		return false
	}
	if TypeEqualsReflection(val.Type(), TimeTypeReflection) {
		return true
	} else if TypeEqualsReflection(val.Type(), NumericDateTypeReflection) {
		return true
	}
	return false
}

const (
	TimeFormatISO8601 = "2006-01-02T15:04:05.000Z07:00"
)

func GetDateTimeStringISO8601Reflection(value any) string {
	val := PassValueIndirectReflection(value)
	if !IsValidReflection(val) {
		return ""
	}

	if TypeEqualsReflection(val.Type(), TimeTypeReflection) {
		t := val.Interface().(time.Time)
		return t.UTC().Format(TimeFormatISO8601)
	} else if TypeEqualsReflection(val.Type(), NumericDateTypeReflection) {
		// why not use a pointer for jwt.NumericDate, because
		// has already pass value indirect reflection
		t := val.Interface().(jwt.NumericDate).Time
		return t.UTC().Format(TimeFormatISO8601)
	}
	return ""
}

func ToStringReflection(value any) string {
	val := PassValueIndirectReflection(value)
	if !IsValidReflection(val) {
		return "<undefined>"
	}
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
	case reflect.Struct:
		if IsDateTimeStringISO8601Reflection(val) {
			return strconv.Quote(GetDateTimeStringISO8601Reflection(val))
		}
		return "<struct>"
	case reflect.Array, reflect.Slice:
		return "<array>"
	case reflect.Map:
		return "<map>"
	default:
		return "<undefined>"
	}
}

func PassBackValueIndirectByInterfaceReflection[T any](value any) T {
	val := PassValueIndirectReflection(value)
	if !IsValidReflection(val) {
		return Default[T]()
	}

	var t T
	KeepVoid(t)

	if val.Type() == reflect.TypeOf(t) {
		return Unwrap(Cast[T](val.Interface()))
	}

	panic(NewThrow("invalid data type", ErrDataTypeInvalid))
}

func PassBackAnyValueIndirectByInterfaceReflection(value any) any {
	return PassBackValueIndirectByInterfaceReflection[any](value)
}

func JsonPreviewPermuteIndentReflection(obj any, indent int, start int) string {
	end := start + indent
	whiteSpaceStart := strings.Repeat(" ", start)
	whiteSpaceEnd := strings.Repeat(" ", end)
	if obj == nil {
		return "null"
	}
	if serialize, ok := obj.(JsonableImpl); ok {
		return serialize.ToJson()
	}
	if val, ok := obj.(StringableImpl); ok {
		return strconv.Quote(val.ToString())
	}
	val := PassValueIndirectReflection(obj)
	if !IsValidReflection(val) {
		return "undefined"
	}
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
	case reflect.Struct:
		if IsDateTimeStringISO8601Reflection(val) {
			return strconv.Quote(GetDateTimeStringISO8601Reflection(val))
		}

		// scrape any fields
		t := val.Type()
		n := t.NumField()
		temp := make([]string, 0)
		for i := 0; i < n; i++ {
			structField := t.Field(i)
			structTag := structField.Tag
			field := val.Field(i)
			fieldKey := ToCamelCase(structField.Name)
			fieldValue := "undefined"
			if IsExportedFieldReflection(field) {
				fieldValue = JsonPreviewPermuteIndentReflection(field.Interface(), indent, end)
			}
			if nameTag, ok := structTag.Lookup("name"); ok {
				fieldKey = nameTag
			}
			if jsonTag, ok := structTag.Lookup("json"); ok {
				KeepVoid(jsonTag, ok)

				name := ""
				tokens := strings.Split(jsonTag, ",")
				size := len(tokens)
				if size > 0 {
					name = strings.Trim(tokens[0], " ")
				}

				if len(name) > 0 && name != "-" {
					fieldKey = name
				}

				for j := 1; j < size; j++ {
					token := strings.Trim(tokens[j], " ")
					switch token {
					case "-", "ignore", "ignored":
						continue
					case "omitempty", "notnull", "required":
						if fieldValue == "undefined" ||
							fieldValue == "null" ||
							fieldValue == strconv.Quote("") {
							continue
						}
					}
				}
			}

			// added
			entry := whiteSpaceEnd + strconv.Quote(fieldKey) + ": " + fieldValue
			temp = append(temp, entry)
		}
		return "{\n" + strings.Join(temp, ",\n") + "\n" + whiteSpaceStart + "}"
	case reflect.Array, reflect.Slice:
		size := val.Len()
		values := make([]string, size)
		for i := 0; i < size; i++ {
			elem := val.Index(i).Interface()
			values[i] = whiteSpaceEnd + JsonPreviewPermuteIndentReflection(elem, indent, end)
		}
		if len(values) > 0 {
			return "[\n" + strings.Join(values, ",\n") + "\n" + whiteSpaceStart + "]"
		}
		return "[]"
	case reflect.Map:
		size := val.Len()
		iter := val.MapRange()
		values := make([]string, size)
		for i := 0; iter.Next(); i++ {
			key := iter.Key().Interface()
			value := iter.Value().Interface()
			values[i] = whiteSpaceEnd + strconv.Quote(fmt.Sprint(key)) + ": " + JsonPreviewPermuteIndentReflection(value, indent, end)
		}
		if len(values) > 0 {
			return "{\n" + strings.Join(values, ",\n") + "\n" + whiteSpaceStart + "}"
		}
		return "{}"
	default:
		return "undefined"
	}
}

func JsonPreviewIndentReflection(obj any, indent int) string {
	return JsonPreviewPermuteIndentReflection(obj, indent, 0)
}

func JsonPreviewReflection(obj any) string {
	return JsonPreviewPermuteIndentReflection(obj, 4, 0)
}
