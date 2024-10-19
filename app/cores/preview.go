package cores

import (
	"strconv"
	"strings"
	"time"
)

func ShikaJsonEncodeIndentPermutatePreview(shikaObjectProperty ShikaObjectPropertyImpl, indent int, start int) string {
	var err error
	var t time.Time
	KeepVoid(err, t)

	end := start + indent
	whiteSpaceStart := strings.Repeat(" ", start)
	whiteSpaceEnd := strings.Repeat(" ", end)
	KeepVoid(end, whiteSpaceStart, whiteSpaceEnd)

	if shikaObjectProperty == nil {
		return "undefined"
	}

	switch shikaObjectProperty.GetKind() {
	case ShikaObjectDataTypeUndefined:
		return "undefined"
	case ShikaObjectDataTypeNull:
		return "null"
	case ShikaObjectDataTypeBool:
		v := Unwrap(Cast[bool](shikaObjectProperty.GetValue()))
		return strconv.FormatBool(v)
	case ShikaObjectDataTypeInt8:
		v := Unwrap(Cast[int8](shikaObjectProperty.GetValue()))
		return strconv.FormatInt(int64(v), 10)
	case ShikaObjectDataTypeUint8:
		v := Unwrap(Cast[uint8](shikaObjectProperty.GetValue()))
		return strconv.FormatUint(uint64(v), 10)
	case ShikaObjectDataTypeInt16:
		v := Unwrap(Cast[int16](shikaObjectProperty.GetValue()))
		return strconv.FormatInt(int64(v), 10)
	case ShikaObjectDataTypeUint16:
		v := Unwrap(Cast[uint16](shikaObjectProperty.GetValue()))
		return strconv.FormatUint(uint64(v), 10)
	case ShikaObjectDataTypeInt32:
		v := Unwrap(Cast[int32](shikaObjectProperty.GetValue()))
		return strconv.FormatInt(int64(v), 10)
	case ShikaObjectDataTypeUint32:
		v := Unwrap(Cast[uint32](shikaObjectProperty.GetValue()))
		return strconv.FormatUint(uint64(v), 10)
	case ShikaObjectDataTypeInt64:
		v := Unwrap(Cast[int64](shikaObjectProperty.GetValue()))
		return strconv.FormatInt(v, 10)
	case ShikaObjectDataTypeUint64:
		v := Unwrap(Cast[uint64](shikaObjectProperty.GetValue()))
		return strconv.FormatUint(v, 10)
	case ShikaObjectDataTypeUintptr:
		v := Unwrap(Cast[uintptr](shikaObjectProperty.GetValue()))
		return ToStringReflection(v)
	case ShikaObjectDataTypeFloat32:
		v := Unwrap(Cast[float32](shikaObjectProperty.GetValue()))
		return strconv.FormatFloat(float64(v), 'f', -1, 32)
	case ShikaObjectDataTypeFloat64:
		v := Unwrap(Cast[float64](shikaObjectProperty.GetValue()))
		return strconv.FormatFloat(v, 'f', -1, 64)
	case ShikaObjectDataTypeComplex64:
		v := Unwrap(Cast[complex64](shikaObjectProperty.GetValue()))
		return strconv.FormatComplex(complex128(v), 'f', -1, 64)
	case ShikaObjectDataTypeComplex128:
		v := Unwrap(Cast[complex128](shikaObjectProperty.GetValue()))
		return strconv.FormatComplex(v, 'f', -1, 128)
	case ShikaObjectDataTypeString:
		v := Unwrap(Cast[string](shikaObjectProperty.GetValue()))
		return strconv.Quote(v)
	case ShikaObjectDataTypeArray:
		shikaObjectProperties := Unwrap(Cast[[]ShikaObjectPropertyImpl](shikaObjectProperty.GetValue()))
		size := len(shikaObjectProperties)
		values := make([]string, size)
		for i := 0; i < size; i++ {
			elem := shikaObjectProperties[i]
			v := ShikaJsonEncodeIndentPermutatePreview(elem, indent, end)
			if len(v) > 0 {
				values[i] = whiteSpaceEnd + v
				continue
			}
			values[i] = whiteSpaceEnd + "undefined"
		}
		if len(values) > 0 {
			return "[\n" + strings.Join(values, ",\n") + "\n" + whiteSpaceStart + "]"
		}
		return "[]"
	case ShikaObjectDataTypeObject:
		shikaVarObjects := Unwrap(Cast[[]ShikaVarObjectImpl](shikaObjectProperty.GetValue()))
		size := len(shikaVarObjects)
		values := make([]string, size)
		for i := 0; i < size; i++ {
			elem := shikaVarObjects[i]
			k := strconv.Quote(ToCamelCase(elem.GetName()))
			v := elem.GetOwnProperty()
			values[i] = whiteSpaceEnd + k + ": " + ShikaJsonEncodeIndentPermutatePreview(v, indent, end)
		}
		if len(values) > 0 {
			return "{\n" + strings.Join(values, ",\n") + "\n" + whiteSpaceStart + "}"
		}
		return "{}"
	case ShikaObjectDataTypeTime:
		if t, err = GetTimeUtcISO8601(shikaObjectProperty.GetValue()); err != nil {
			return "undefined"
		}
		v := ToTimeUtcStringISO8601(t)
		return strconv.Quote(v)
	default:
		return "undefined"
	}
}

func ShikaJsonEncodeIndentPreview(obj any, indent int) string {
	shikaObjectProperty := ShikaObjectPropertyConversionPreview(obj)
	return ShikaJsonEncodeIndentPermutatePreview(shikaObjectProperty, indent, 0)
}

func ShikaJsonEncodePreview(obj any) string {
	return ShikaJsonEncodeIndentPreview(obj, 4)
}

func ShikaYamlEncodeIndentPermutatePreview(shikaObjectProperty ShikaObjectPropertyImpl, indent int, start int) string {
	var err error
	var t time.Time
	KeepVoid(err, t)

	end := start + indent
	whiteSpaceStart := strings.Repeat(" ", start)
	whiteSpaceEnd := strings.Repeat(" ", end)
	KeepVoid(end, whiteSpaceStart, whiteSpaceEnd)

	if shikaObjectProperty == nil {
		return "undefined"
	}

	switch shikaObjectProperty.GetKind() {
	case ShikaObjectDataTypeUndefined:
		return "undefined"
	case ShikaObjectDataTypeNull:
		return "null"
	case ShikaObjectDataTypeBool:
		v := Unwrap(Cast[bool](shikaObjectProperty.GetValue()))
		return strconv.FormatBool(v)
	case ShikaObjectDataTypeInt8:
		v := Unwrap(Cast[int8](shikaObjectProperty.GetValue()))
		return strconv.FormatInt(int64(v), 10)
	case ShikaObjectDataTypeUint8:
		v := Unwrap(Cast[uint8](shikaObjectProperty.GetValue()))
		return strconv.FormatUint(uint64(v), 10)
	case ShikaObjectDataTypeInt16:
		v := Unwrap(Cast[int16](shikaObjectProperty.GetValue()))
		return strconv.FormatInt(int64(v), 10)
	case ShikaObjectDataTypeUint16:
		v := Unwrap(Cast[uint16](shikaObjectProperty.GetValue()))
		return strconv.FormatUint(uint64(v), 10)
	case ShikaObjectDataTypeInt32:
		v := Unwrap(Cast[int32](shikaObjectProperty.GetValue()))
		return strconv.FormatInt(int64(v), 10)
	case ShikaObjectDataTypeUint32:
		v := Unwrap(Cast[uint32](shikaObjectProperty.GetValue()))
		return strconv.FormatUint(uint64(v), 10)
	case ShikaObjectDataTypeInt64:
		v := Unwrap(Cast[int64](shikaObjectProperty.GetValue()))
		return strconv.FormatInt(v, 10)
	case ShikaObjectDataTypeUint64:
		v := Unwrap(Cast[uint64](shikaObjectProperty.GetValue()))
		return strconv.FormatUint(v, 10)
	case ShikaObjectDataTypeUintptr:
		v := Unwrap(Cast[uintptr](shikaObjectProperty.GetValue()))
		return ToStringReflection(v)
	case ShikaObjectDataTypeFloat32:
		v := Unwrap(Cast[float32](shikaObjectProperty.GetValue()))
		return strconv.FormatFloat(float64(v), 'f', -1, 32)
	case ShikaObjectDataTypeFloat64:
		v := Unwrap(Cast[float64](shikaObjectProperty.GetValue()))
		return strconv.FormatFloat(v, 'f', -1, 64)
	case ShikaObjectDataTypeComplex64:
		v := Unwrap(Cast[complex64](shikaObjectProperty.GetValue()))
		return strconv.FormatComplex(complex128(v), 'f', -1, 64)
	case ShikaObjectDataTypeComplex128:
		v := Unwrap(Cast[complex128](shikaObjectProperty.GetValue()))
		return strconv.FormatComplex(v, 'f', -1, 128)
	case ShikaObjectDataTypeString:
		v := Unwrap(Cast[string](shikaObjectProperty.GetValue()))
		return strconv.Quote(v)
	case ShikaObjectDataTypeArray:

		// modified white space for indentation string
		start = start - 2             // truncate the hyphen and whitespace
		start = Unwrap(Max(start, 0)) // truncate safety number
		end = start + indent
		end = end - 2             // truncate the hyphen and whitespace
		end = Unwrap(Max(end, 0)) // truncate safety number
		whiteSpaceStart = strings.Repeat(" ", start)
		whiteSpaceEnd = strings.Repeat(" ", end)
		KeepVoid(end, whiteSpaceStart, whiteSpaceEnd)

		shikaObjectProperties := Unwrap(Cast[[]ShikaObjectPropertyImpl](shikaObjectProperty.GetValue()))
		size := len(shikaObjectProperties)
		values := make([]string, size)
		for i := 0; i < size; i++ {
			elem := shikaObjectProperties[i]
			v := ShikaYamlEncodeIndentPermutatePreview(elem, indent, end)
			if len(v) > 0 {
				if v[0] == '\n' {
					s := len(whiteSpaceStart) + 3 // truncate the newline, hyphen and whitespace
					values[i] = whiteSpaceStart + "- " + v[s:]
					continue
				}
				values[i] = whiteSpaceStart + "- " + v
				continue
			}
			values[i] = whiteSpaceStart + "- " + "undefined"
		}
		if len(values) > 0 {
			return "\n" + strings.Join(values, "\n")
		}
		return "[]"
	case ShikaObjectDataTypeObject:
		shikaVarObjects := Unwrap(Cast[[]ShikaVarObjectImpl](shikaObjectProperty.GetValue()))
		size := len(shikaVarObjects)
		values := make([]string, size)
		for i := 0; i < size; i++ {
			elem := shikaVarObjects[i]
			k := ToSnakeCase(elem.GetName())
			v := elem.GetOwnProperty()
			values[i] = whiteSpaceStart + k + ": " + ShikaYamlEncodeIndentPermutatePreview(v, indent, end)
		}
		if len(values) > 0 {
			return "\n" + strings.Join(values, "\n")
		}
		return "{}"
	case ShikaObjectDataTypeTime:
		if t, err = GetTimeUtcISO8601(shikaObjectProperty.GetValue()); err != nil {
			return "undefined"
		}
		v := ToTimeUtcStringISO8601(t)
		return strconv.Quote(v)
	default:
		return "undefined"
	}
}

func ShikaYamlEncodeIndentPreview(obj any, indent int) string {
	shikaObjectProperty := ShikaObjectPropertyConversionPreview(obj)
	return ShikaYamlEncodeIndentPermutatePreview(shikaObjectProperty, indent, 0)
}

func ShikaYamlEncodePreview(obj any) string {
	return ShikaYamlEncodeIndentPreview(obj, 4)
}
