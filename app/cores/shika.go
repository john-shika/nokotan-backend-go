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
	ShikaObjectDataTypeInt
	ShikaObjectDataTypeUint
	ShikaObjectDataTypeInt8
	ShikaObjectDataTypeUint8
	ShikaObjectDataTypeInt16
	ShikaObjectDataTypeUint16
	ShikaObjectDataTypeInt32
	ShikaObjectDataTypeUint32
	ShikaObjectDataTypeInt64
	ShikaObjectDataTypeUint64
	ShikaObjectDataTypeUintptr
	ShikaObjectDataTypeFloat
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
	case ShikaObjectDataTypeInt:
		return "int"
	case ShikaObjectDataTypeUint:
		return "uint"
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
	case ShikaObjectDataTypeFloat:
		return "float"
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
	GetValue() any
	GetKind() ShikaObjectDataTypeKind
	SetValue(value any)
	IsConfigurable() bool
	IsEnumerable() bool
	IsWritable() bool
	IsValid() bool
}

type ShikaObjectProperty struct {
	Value        any
	Kind         ShikaObjectDataTypeKind
	Get          ShikaHandleGetterFunc
	Set          ShikaHandleSetterFunc
	Configurable bool
	Enumerable   bool
	Writable     bool
}

func NewShikaObjectProperty(value any, kind ShikaObjectDataTypeKind) ShikaObjectPropertyImpl {
	return &ShikaObjectProperty{
		Value:        value,
		Kind:         kind,
		Get:          nil,
		Set:          nil,
		Configurable: true,
		Enumerable:   true,
		Writable:     true,
	}
}

func (shikaObjectProperty *ShikaObjectProperty) GetValue() any {

	// calling getter function
	if shikaObjectProperty.Get != nil {
		return shikaObjectProperty.Get()
	}

	// get value directly
	return shikaObjectProperty.Value
}

func (shikaObjectProperty *ShikaObjectProperty) GetKind() ShikaObjectDataTypeKind {
	return shikaObjectProperty.Kind
}

func (shikaObjectProperty *ShikaObjectProperty) SetValue(value any) {

	// calling setter function
	if shikaObjectProperty.Set != nil {
		shikaObjectProperty.Set(value)
		return
	}

	// set value directly
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

func (shikaObjectProperty *ShikaObjectProperty) IsValid() bool {
	if shikaObjectProperty.Kind == ShikaObjectDataTypeString {
		if v, ok := Cast[string](shikaObjectProperty.Value); ok {
			return v != ""
		}
		return false
	}
	return shikaObjectProperty.Kind != ShikaObjectDataTypeUndefined &&
		shikaObjectProperty.Kind != ShikaObjectDataTypeNull
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

type ShikaVarObjectImpl interface {
	GetName() string
	GetOwnProperty() ShikaObjectPropertyImpl
	GetProperties() []ShikaVarObjectImpl
	SetOwnProperty(property ShikaObjectPropertyImpl)
	SetProperties(properties []ShikaVarObjectImpl)
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

type ShikaVarObject struct {
	Name        string
	OwnProperty ShikaObjectPropertyImpl
	Properties  []ShikaVarObjectImpl
	Attributes  []ShikaObjectAttributeImpl
}

func NewShikaVarObject(name string) ShikaVarObjectImpl {
	return &ShikaVarObject{
		Name:        name,
		OwnProperty: nil,
		Properties:  make([]ShikaVarObjectImpl, 0),
	}
}

func (shikaVarObject *ShikaVarObject) GetName() string {
	return shikaVarObject.Name
}

func (shikaVarObject *ShikaVarObject) GetOwnProperty() ShikaObjectPropertyImpl {
	return shikaVarObject.OwnProperty
}

func (shikaVarObject *ShikaVarObject) GetProperties() []ShikaVarObjectImpl {
	return shikaVarObject.Properties
}

func (shikaVarObject *ShikaVarObject) SetOwnProperty(property ShikaObjectPropertyImpl) {
	shikaVarObject.OwnProperty = property
}

func (shikaVarObject *ShikaVarObject) SetProperties(properties []ShikaVarObjectImpl) {
	shikaVarObject.Properties = properties
}

func (shikaVarObject *ShikaVarObject) PropertiesLength() int {
	return len(shikaVarObject.Properties)
}

func (shikaVarObject *ShikaVarObject) GetPropertyKeys() []string {
	keys := make([]string, 0, len(shikaVarObject.Properties))
	for i, shikaVarObj := range shikaVarObject.Properties {
		KeepVoid(i, shikaVarObj)

		keys = append(keys, shikaVarObj.GetName())
	}
	return keys
}

func (shikaVarObject *ShikaVarObject) GetPropertyValues() []ShikaObjectPropertyImpl {
	values := make([]ShikaObjectPropertyImpl, 0, len(shikaVarObject.Properties))
	for i, shikaVarObj := range shikaVarObject.Properties {
		KeepVoid(i, shikaVarObj)

		values = append(values, shikaVarObj.GetOwnProperty())
	}
	return values
}

func (shikaVarObject *ShikaVarObject) HasPropertyKey(key string) bool {
	for i, shikaVarObj := range shikaVarObject.Properties {
		KeepVoid(i, shikaVarObj)

		if shikaVarObj.GetName() == key {
			return true
		}
	}
	return false
}

func (shikaVarObject *ShikaVarObject) ContainPropertyKeys(keys ...string) bool {
	for i, key := range keys {
		KeepVoid(i, key)

		if !shikaVarObject.HasPropertyKey(key) {
			return false
		}
	}
	return true
}

func (shikaVarObject *ShikaVarObject) GetPropertyByName(name string) ShikaObjectPropertyImpl {
	for i, shikaVarObj := range shikaVarObject.Properties {
		KeepVoid(i, shikaVarObj)

		if shikaVarObj.GetName() == name {
			return shikaVarObj.GetOwnProperty()
		}
	}
	return nil
}

func (shikaVarObject *ShikaVarObject) SetPropertyByName(name string, property ShikaObjectPropertyImpl) {

	// replace existing property
	for i, shikaObj := range shikaVarObject.Properties {
		KeepVoid(i, shikaObj)

		if shikaObj.GetName() == name {
			shikaVarObject.Properties[i].SetOwnProperty(property)
			return
		}
	}

	// create new property
	shikaVarObj := NewShikaVarObject(name)
	shikaVarObj.SetOwnProperty(property)
	shikaVarObject.Properties = append(shikaVarObject.Properties, shikaVarObj)
}

func (shikaVarObject *ShikaVarObject) RemovePropertyByName(name string) {
	for i, shikaVarObj := range shikaVarObject.Properties {
		KeepVoid(i, shikaVarObj)

		if shikaVarObj.GetName() == name {
			j := i + 1
			shikaVarObject.Properties = append(shikaVarObject.Properties[:i], shikaVarObject.Properties[j:]...)
			return
		}
	}
}

func (shikaVarObject *ShikaVarObject) GetAttributesLength() int {
	return len(shikaVarObject.Attributes)
}

func (shikaVarObject *ShikaVarObject) GetAttributes() []ShikaObjectAttributeImpl {
	return shikaVarObject.Attributes
}

func (shikaVarObject *ShikaVarObject) SetAttributes(attributes []ShikaObjectAttributeImpl) {
	shikaVarObject.Attributes = attributes
}

func (shikaVarObject *ShikaVarObject) HasAttributeByName(name string) bool {
	for i, shikaObjAttr := range shikaVarObject.Attributes {
		KeepVoid(i, shikaObjAttr)

		if shikaObjAttr.GetName() == name {
			return true
		}
	}
	return false
}

func (shikaVarObject *ShikaVarObject) ContainAttributeNames(names ...string) bool {
	for i, name := range names {
		KeepVoid(i, name)

		if !shikaVarObject.HasAttributeByName(name) {
			return false
		}
	}
	return true
}

func (shikaVarObject *ShikaVarObject) GetAttributeByName(name string) ShikaObjectAttributeImpl {
	for i, shikaObjAttr := range shikaVarObject.Attributes {
		KeepVoid(i, shikaObjAttr)

		if shikaObjAttr.GetName() == name {
			return shikaObjAttr
		}
	}
	return nil
}

func (shikaVarObject *ShikaVarObject) SetAttributeByName(name string, attribute ShikaObjectAttributeImpl) {

	// replace existing attribute
	for i, shikaObjAttr := range shikaVarObject.Attributes {
		KeepVoid(i, shikaObjAttr)

		if shikaObjAttr.GetName() == name {
			shikaVarObject.Attributes[i] = attribute
			return
		}
	}

	// create new attribute
	shikaVarObject.Attributes = append(shikaVarObject.Attributes, attribute)
}

func (shikaVarObject *ShikaVarObject) RemoveAttributeByName(name string) {
	for i, shikaObjAttr := range shikaVarObject.Attributes {
		KeepVoid(i, shikaObjAttr)

		if shikaObjAttr.GetName() == name {
			j := i + 1
			shikaVarObject.Attributes = append(shikaVarObject.Attributes[:i], shikaVarObject.Attributes[j:]...)
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

func ShikaObjectPropertyConversionPreview(obj any) ShikaObjectPropertyImpl {
	if obj == nil {
		return NewShikaObjectProperty(nil, ShikaObjectDataTypeNull)
	}
	if val, ok := obj.(StringableImpl); ok {
		return NewShikaObjectProperty(val.ToString(), ShikaObjectDataTypeString)
	}
	val := PassValueIndirectReflection(obj)
	if !IsValidReflection(val) {
		return NewShikaObjectProperty(nil, ShikaObjectDataTypeNull)
	}
	switch val.Kind() {
	case reflect.Bool:
		v := val.Bool()
		return NewShikaObjectProperty(v, ShikaObjectDataTypeBool)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v := val.Int()
		return NewShikaObjectProperty(v, ShikaObjectDataTypeInt64)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v := val.Uint()
		return NewShikaObjectProperty(v, ShikaObjectDataTypeUint64)
	case reflect.Uintptr:
		v := Unwrap(Cast[uintptr](val.Interface()))
		return NewShikaObjectProperty(v, ShikaObjectDataTypeUintptr)
	case reflect.Float32, reflect.Float64:
		v := val.Float()
		return NewShikaObjectProperty(v, ShikaObjectDataTypeFloat64)
	case reflect.Complex64, reflect.Complex128:
		v := val.Complex()
		return NewShikaObjectProperty(v, ShikaObjectDataTypeComplex128)
	case reflect.String:
		v := val.String()
		return NewShikaObjectProperty(v, ShikaObjectDataTypeString)
	case reflect.Struct:
		if IsShikaObjectPropertyReflection(val) {
			return GetShikaObjectPropertyReflection(val)
		}

		if IsTimeUtcISO8601(val) {
			return NewShikaObjectProperty(ToTimeUtcStringISO8601(val), ShikaObjectDataTypeTime)
		}

		// scrape any fields
		t := val.Type()
		n := t.NumField()
		temp := make([]ShikaVarObjectImpl, 0)
		for i := 0; i < n; i++ {
			sField := t.Field(i)
			sTag := sField.Tag
			sValue := val.Field(i)
			if !IsExportedFieldReflection(sValue) {
				continue
			}
			pName := ToCamelCase(sField.Name)
			pValue := ShikaObjectPropertyConversionPreview(sValue.Interface())
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
						if !pValue.IsValid() {
							continue
						}
					}
				}
			}

			property := NewShikaVarObject(pName)
			property.SetOwnProperty(pValue)
			temp = append(temp, property)
		}
		return NewShikaObjectProperty(temp, ShikaObjectDataTypeObject)
	case reflect.Array, reflect.Slice:
		size := val.Len()
		values := make([]ShikaObjectPropertyImpl, size)
		for i := 0; i < size; i++ {
			elem := val.Index(i).Interface()
			values[i] = ShikaObjectPropertyConversionPreview(elem)
		}
		return NewShikaObjectProperty(values, ShikaObjectDataTypeArray)
	case reflect.Map:
		size := val.Len()
		iter := val.MapRange()
		values := make([]ShikaVarObjectImpl, size)
		for i := 0; iter.Next(); i++ {
			key := ToStringReflection(iter.Key())
			value := iter.Value().Interface()
			temp := NewShikaVarObject(key)
			temp.SetOwnProperty(ShikaObjectPropertyConversionPreview(value))
			values[i] = temp
		}
		return NewShikaObjectProperty(values, ShikaObjectDataTypeObject)
	default:
		return NewShikaObjectProperty(nil, ShikaObjectDataTypeUndefined)
	}
}
