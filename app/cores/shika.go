package cores

import (
	"reflect"
	"strings"
)

type ShikaObjectDataTypeKind int

const (
	ShikaObjectDataTypeUndefined ShikaObjectDataTypeKind = iota
	ShikaObjectDataTypeNull
	ShikaObjectDataTypeBool
	ShikaObjectDataTypeInt8
	ShikaObjectDataTypeUint8
	ShikaObjectDataTypeInt16
	ShikaObjectDataTypeUint16
	ShikaObjectDataTypeInt32
	ShikaObjectDataTypeUint32
	ShikaObjectDataTypeInt64
	ShikaObjectDataTypeUint64
	ShikaObjectDataTypeFloat32
	ShikaObjectDataTypeFloat64
	ShikaObjectDataTypeComplex64
	ShikaObjectDataTypeComplex128
	ShikaObjectDataTypeString
	ShikaObjectDataTypeArray
	ShikaObjectDataTypeObject
	ShikaObjectDataTypeNamespace
	ShikaObjectDataTypeInterface
	ShikaObjectDataTypeStruct
	ShikaObjectDataTypeClass
	ShikaObjectDataTypeFunction
	ShikaObjectDataTypeAttribute
	ShikaObjectDataTypeTime
)

func (shikaObjectDataTypeKind ShikaObjectDataTypeKind) ToString() string {
	switch shikaObjectDataTypeKind {
	case ShikaObjectDataTypeUndefined:
		return "undefined"
	case ShikaObjectDataTypeNull:
		return "null"
	case ShikaObjectDataTypeBool:
		return "bool"
	case ShikaObjectDataTypeInt8:
		return "int8"
	case ShikaObjectDataTypeUint8:
		return "uint8"
	case ShikaObjectDataTypeInt16:
		return "int16"
	case ShikaObjectDataTypeUint16:
		return "uint16"
	case ShikaObjectDataTypeInt32:
		return "int32"
	case ShikaObjectDataTypeUint32:
		return "uint32"
	case ShikaObjectDataTypeInt64:
		return "int64"
	case ShikaObjectDataTypeUint64:
		return "uint64"
	case ShikaObjectDataTypeFloat32:
		return "float32"
	case ShikaObjectDataTypeFloat64:
		return "float64"
	case ShikaObjectDataTypeComplex64:
		return "complex64"
	case ShikaObjectDataTypeComplex128:
		return "complex128"
	case ShikaObjectDataTypeString:
		return "string"
	case ShikaObjectDataTypeArray:
		return "array"
	case ShikaObjectDataTypeObject:
		return "object"
	case ShikaObjectDataTypeNamespace:
		return "namespace"
	case ShikaObjectDataTypeInterface:
		return "interface"
	case ShikaObjectDataTypeStruct:
		return "struct"
	case ShikaObjectDataTypeClass:
		return "class"
	case ShikaObjectDataTypeFunction:
		return "function"
	case ShikaObjectDataTypeAttribute:
		return "attribute"
	case ShikaObjectDataTypeTime:
		return "time"
	default:
		return "unknown"
	}
}

type ShikaHandleGetterFunc func() any
type ShikaHandleSetterFunc func(value any)

type ShikaObjectPropertyImpl interface {
	GetKind() ShikaObjectDataTypeKind
	GetValue() any
	SetValue(value any)
	IsConfigurable() bool
	IsEnumerable() bool
	IsWritable() bool
}

type ShikaObjectProperty struct {
	Kind         ShikaObjectDataTypeKind
	Value        any
	Get          ShikaHandleGetterFunc
	Set          ShikaHandleSetterFunc
	Configurable bool
	Enumerable   bool
	Writable     bool
}

func (shikaObjectProperty *ShikaObjectProperty) GetKind() ShikaObjectDataTypeKind {
	return shikaObjectProperty.Kind
}

func (shikaObjectProperty *ShikaObjectProperty) GetValue() any {
	if shikaObjectProperty.Get != nil {
		return shikaObjectProperty.Get()
	}
	return shikaObjectProperty.Value
}

func (shikaObjectProperty *ShikaObjectProperty) SetValue(value any) {
	if shikaObjectProperty.Set != nil {
		shikaObjectProperty.Set(value)
		return
	}
	shikaObjectProperty.Value = value
}

func (shikaObjectProperty *ShikaObjectProperty) IsConfigurable() bool {
	return shikaObjectProperty.Configurable
}

func (shikaObjectProperty *ShikaObjectProperty) IsEnumerable() bool {
	return shikaObjectProperty.Enumerable
}

func (shikaObjectProperty *ShikaObjectProperty) IsWritable() bool {
	return shikaObjectProperty.Writable
}

func NewShikaObjectProperty(value any, t ShikaObjectDataTypeKind) ShikaObjectPropertyImpl {
	return &ShikaObjectProperty{
		Kind:         t,
		Value:        value,
		Get:          nil,
		Set:          nil,
		Configurable: true,
		Enumerable:   true,
		Writable:     true,
	}
}

type ShikaObjectAttributeImpl interface {
	GetName() string
	GetParametersLength() int
	GetParameters() []any
}

type ShikaObjectAttribute struct {
	Name       string
	Parameters []any
}

func NewShikaObjectAttribute(name string, parameters ...any) ShikaObjectAttributeImpl {
	return &ShikaObjectAttribute{
		Name:       name,
		Parameters: parameters,
	}
}

func (shikaObjectAttribute *ShikaObjectAttribute) GetName() string {
	return shikaObjectAttribute.Name
}

func (shikaObjectAttribute *ShikaObjectAttribute) GetParametersLength() int {
	return len(shikaObjectAttribute.Parameters)
}

func (shikaObjectAttribute *ShikaObjectAttribute) GetParameters() []any {
	return shikaObjectAttribute.Parameters
}

type ShikaObjectImpl interface {
	GetName() string
	GetOwnProperty() ShikaObjectPropertyImpl
	GetProperties() []ShikaObject
	SetOwnProperty(property ShikaObjectPropertyImpl)
	SetProperties(properties []ShikaObject)
	PropertiesLength() int
	GetPropertyKeys() []string
	GetPropertyValues() []ShikaObjectPropertyImpl
	HasPropertyKey(key string) bool
	ContainPropertyKeys(keys ...string) bool
	GetPropertyByName(name string) ShikaObjectPropertyImpl
	SetPropertyByName(name string, property ShikaObjectPropertyImpl)
	RemovePropertyByName(name string)
	GetAttributesLength() int
	GetAttributes() []ShikaObjectAttributeImpl
	SetAttributes(attributes []ShikaObjectAttributeImpl)
	HasAttributeByName(name string) bool
	ContainAttributeNames(names ...string) bool
	GetAttributeByName(name string) ShikaObjectAttributeImpl
	SetAttributeByName(name string, attribute ShikaObjectAttributeImpl)
	RemoveAttributeByName(name string)
}

type ShikaObject struct {
	Name        string
	OwnProperty ShikaObjectPropertyImpl
	Properties  []ShikaObject
	Attributes  []ShikaObjectAttributeImpl
}

func NewShikaObject(name string) ShikaObjectImpl {
	return &ShikaObject{
		Name:        name,
		OwnProperty: nil,
		Properties:  make([]ShikaObject, 0),
	}
}

func (shikaObject *ShikaObject) GetName() string {
	return shikaObject.Name
}

func (shikaObject *ShikaObject) GetOwnProperty() ShikaObjectPropertyImpl {
	return shikaObject.OwnProperty
}

func (shikaObject *ShikaObject) GetProperties() []ShikaObject {
	return shikaObject.Properties
}

func (shikaObject *ShikaObject) SetOwnProperty(property ShikaObjectPropertyImpl) {
	shikaObject.OwnProperty = property
}

func (shikaObject *ShikaObject) SetProperties(properties []ShikaObject) {
	shikaObject.Properties = properties
}

func (shikaObject *ShikaObject) PropertiesLength() int {
	return len(shikaObject.Properties)
}

func (shikaObject *ShikaObject) GetPropertyKeys() []string {
	keys := make([]string, 0, len(shikaObject.Properties))
	for i, shikaObj := range shikaObject.Properties {
		KeepVoid(i, shikaObj)

		keys = append(keys, shikaObj.Name)
	}
	return keys
}

func (shikaObject *ShikaObject) GetPropertyValues() []ShikaObjectPropertyImpl {
	values := make([]ShikaObjectPropertyImpl, 0, len(shikaObject.Properties))
	for i, shikaObj := range shikaObject.Properties {
		KeepVoid(i, shikaObj)

		values = append(values, shikaObj.OwnProperty)
	}
	return values
}

func (shikaObject *ShikaObject) HasPropertyKey(key string) bool {
	for i, shikaObj := range shikaObject.Properties {
		KeepVoid(i, shikaObj)

		if shikaObj.Name == key {
			return true
		}
	}
	return false
}

func (shikaObject *ShikaObject) ContainPropertyKeys(keys ...string) bool {
	for i, key := range keys {
		KeepVoid(i, key)

		if !shikaObject.HasPropertyKey(key) {
			return false
		}
	}
	return true
}

func (shikaObject *ShikaObject) GetPropertyByName(name string) ShikaObjectPropertyImpl {
	for i, shikaObj := range shikaObject.Properties {
		KeepVoid(i, shikaObj)

		if shikaObj.Name == name {
			return shikaObj.OwnProperty
		}
	}
	return nil
}

func (shikaObject *ShikaObject) SetPropertyByName(name string, property ShikaObjectPropertyImpl) {
	for i, shikaObj := range shikaObject.Properties {
		KeepVoid(i, shikaObj)

		if shikaObj.Name == name {
			shikaObject.Properties[i].OwnProperty = property
			return
		}
	}
}

func (shikaObject *ShikaObject) RemovePropertyByName(name string) {
	for i, shikaObj := range shikaObject.Properties {
		KeepVoid(i, shikaObj)

		if shikaObj.Name == name {
			j := i + 1
			shikaObject.Properties = append(shikaObject.Properties[:i], shikaObject.Properties[j:]...)
			return
		}
	}
}

func (shikaObject *ShikaObject) GetAttributesLength() int {
	return len(shikaObject.Attributes)
}

func (shikaObject *ShikaObject) GetAttributes() []ShikaObjectAttributeImpl {
	return shikaObject.Attributes
}

func (shikaObject *ShikaObject) SetAttributes(attributes []ShikaObjectAttributeImpl) {
	shikaObject.Attributes = attributes
}

func (shikaObject *ShikaObject) HasAttributeByName(name string) bool {
	for i, shikaObjAttr := range shikaObject.Attributes {
		KeepVoid(i, shikaObjAttr)

		if shikaObjAttr.GetName() == name {
			return true
		}
	}
	return false
}

func (shikaObject *ShikaObject) ContainAttributeNames(names ...string) bool {
	for i, name := range names {
		KeepVoid(i, name)

		if !shikaObject.HasAttributeByName(name) {
			return false
		}
	}
	return true
}

func (shikaObject *ShikaObject) GetAttributeByName(name string) ShikaObjectAttributeImpl {
	for i, shikaObjAttr := range shikaObject.Attributes {
		KeepVoid(i, shikaObjAttr)

		if shikaObjAttr.GetName() == name {
			return shikaObjAttr
		}
	}
	return nil
}

func (shikaObject *ShikaObject) SetAttributeByName(name string, attribute ShikaObjectAttributeImpl) {
	for i, shikaObjAttr := range shikaObject.Attributes {
		KeepVoid(i, shikaObjAttr)

		if shikaObjAttr.GetName() == name {
			shikaObject.Attributes[i] = attribute
			return
		}
	}
}

func (shikaObject *ShikaObject) RemoveAttributeByName(name string) {
	for i, shikaObjAttr := range shikaObject.Attributes {
		KeepVoid(i, shikaObjAttr)

		if shikaObjAttr.GetName() == name {
			j := i + 1
			shikaObject.Attributes = append(shikaObject.Attributes[:i], shikaObject.Attributes[j:]...)
			return
		}
	}
}

var ShikaObjectPropertyType = reflect.TypeOf(new(ShikaObjectProperty))

func IsShikaObjectPropertyReflection(value any) bool {
	val := PassValueIndirectReflection(value)
	if !IsValidReflection(val) {
		return false
	}
	if TypeEqualsReflection(val.Type(), ShikaObjectPropertyType) {
		return true
	}
	return false
}

func GetShikaObjectPropertyReflection(value any) ShikaObjectPropertyImpl {
	val := PassValueIndirectReflection(value)
	if !IsValidReflection(val) {
		return nil
	}
	if TypeEqualsReflection(val.Type(), ShikaObjectPropertyType) {
		return val.Interface().(ShikaObjectPropertyImpl)
	}
	return nil
}

func ShikaObjectConversionPreview(obj any) ShikaObjectPropertyImpl {
	if obj == nil {
		return NewShikaObjectProperty(nil, ShikaObjectDataTypeNull)
	}
	if val, ok := obj.(StringableImpl); ok {
		return NewShikaObjectProperty(val.ToString(), ShikaObjectDataTypeString)
	}
	val := PassValueIndirectReflection(obj)
	if !IsValidReflection(val) {
		return NewShikaObjectProperty(nil, ShikaObjectDataTypeUndefined)
	}
	switch val.Kind() {
	case reflect.Bool:
		return NewShikaObjectProperty(val.Bool(), ShikaObjectDataTypeBool)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return NewShikaObjectProperty(val.Int(), ShikaObjectDataTypeInt64)
	case reflect.Uint, reflect.Uintptr, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return NewShikaObjectProperty(val.Uint(), ShikaObjectDataTypeUint64)
	case reflect.Float32, reflect.Float64:
		return NewShikaObjectProperty(val.Float(), ShikaObjectDataTypeFloat64)
	case reflect.Complex64, reflect.Complex128:
		return NewShikaObjectProperty(val.Complex(), ShikaObjectDataTypeComplex128)
	case reflect.String:
		return NewShikaObjectProperty(val.String(), ShikaObjectDataTypeString)
	case reflect.Struct:
		if IsShikaObjectPropertyReflection(val) {
			return GetShikaObjectPropertyReflection(val)
		}

		if IsDateTimeStringISO8601Reflection(val) {
			return NewShikaObjectProperty(GetDateTimeStringISO8601Reflection(val), ShikaObjectDataTypeTime)
		}

		// scrape any fields
		t := val.Type()
		n := t.NumField()
		temp := make([]ShikaObjectImpl, 0)
		for i := 0; i < n; i++ {
			sField := t.Field(i)
			sTag := sField.Tag
			sValue := val.Field(i)
			if !IsExportedFieldReflection(sValue) {
				continue
			}
			pName := ToCamelCase(sField.Name)
			pValue := ShikaObjectConversionPreview(sValue.Interface())
			if nameTag, ok := sTag.Lookup("name"); ok {
				pName = nameTag
			}
			if jsonTag, ok := sTag.Lookup("json"); ok {
				KeepVoid(jsonTag, ok)

				name := ""
				tokens := strings.Split(jsonTag, ",")
				size := len(tokens)
				if size > 0 {
					name = strings.Trim(tokens[0], " ")
				}

				if len(name) > 0 && name != "-" {
					pName = name
				}

				for j := 1; j < size; j++ {
					token := strings.Trim(tokens[j], " ")
					switch token {
					case "-", "ignore", "ignored":
						continue
					case "omitempty", "notnull", "required":
						if pValue.GetKind() == ShikaObjectDataTypeUndefined ||
							pValue.GetKind() == ShikaObjectDataTypeNull ||
							pValue.GetValue() == nil {
							continue
						}
					}
				}
			}

			property := NewShikaObject(pName)
			property.SetOwnProperty(pValue)
			temp = append(temp, property)
		}
		return NewShikaObjectProperty(temp, ShikaObjectDataTypeObject)
	case reflect.Array, reflect.Slice:
		size := val.Len()
		values := make([]ShikaObjectPropertyImpl, size)
		for i := 0; i < size; i++ {
			elem := val.Index(i).Interface()
			values[i] = ShikaObjectConversionPreview(elem)
		}
		return NewShikaObjectProperty(values, ShikaObjectDataTypeArray)
	case reflect.Map:
		size := val.Len()
		iter := val.MapRange()
		values := make([]ShikaObjectImpl, size)
		for i := 0; iter.Next(); i++ {
			key := ToStringReflection(iter.Key())
			value := iter.Value().Interface()
			temp := NewShikaObject(key)
			temp.SetOwnProperty(ShikaObjectConversionPreview(value))
			values[i] = temp
		}
		return NewShikaObjectProperty(values, ShikaObjectDataTypeObject)
	default:
		return NewShikaObjectProperty(nil, ShikaObjectDataTypeUndefined)
	}
}
