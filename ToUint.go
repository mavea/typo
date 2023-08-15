package typo

import (
	"errors"
	"reflect"
	"strconv"
)

// toUint64 converts reflect.Value to uint64. If translation fails, returns default value and error message
func toUint64(value reflect.Value, kindResult reflect.Kind) (uint64, error) {
	typeValue := value.Type()
	if nil == typeValue {
		return 0, nil
	}
	switch typeValue.Kind() {
	case reflect.Bool:
		if value.Interface().(bool) {
			return 1, nil
		}
		return 0, nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		i64 := value.Int()
		if 0 > i64 {
			return 0, errors.New(kindResult.String() + ` cannot be negative`)
		}
		return uint64(i64), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return value.Uint(), nil
	case reflect.Float32, reflect.Float64:
		f64 := value.Float()
		if 0.0 > f64 {
			return 0, errors.New(kindResult.String() + ` cannot be negative`)
		}
		return uint64(f64), nil
	case reflect.Pointer:
		if value.IsNil() {
			return 0, nil
		}
		return toUint64(value.Elem(), kindResult)
	case reflect.String:
		f64, err := strconv.ParseFloat(value.String(), 64)
		if nil != err {
			return 0, err
		}
		if 0.0 > f64 {
			return 0, errors.New(kindResult.String() + ` cannot be negative`)
		}
		return uint64(f64), err
	case reflect.Func:
		return 0, errors.New(`argument cannot be a function call`)
	case reflect.Complex64, reflect.Complex128, reflect.Array, reflect.Slice, reflect.Struct, reflect.Map, reflect.Chan, reflect.Interface, reflect.UnsafePointer:
		return 0, errors.New(typeValue.String() + ` format cannot be converted to ` + kindResult.String())
	}

	return 0, errors.New(`unknown format ` + typeValue.String() + ` and cannot be converted to ` + kindResult.String())
}

// ToUint converts a value of unspecified type to uint. If translation fails, returns default value and error message
func ToUint(value any) (result uint, err error) {
	err = toTree(reflect.ValueOf(value), reflect.ValueOf(&result).Elem())

	return result, err
}

// ToUint8 converts a value of unspecified type to uint8. If translation fails, returns default value and error message
func ToUint8(value any) (result uint8, err error) {
	err = toTree(reflect.ValueOf(value), reflect.ValueOf(&result).Elem())

	return result, err
}

// ToUint16 converts a value of unspecified type to uint16. If translation fails, returns default value and error message
func ToUint16(value any) (result uint16, err error) {
	err = toTree(reflect.ValueOf(value), reflect.ValueOf(&result).Elem())

	return result, err
}

// ToUint32 converts a value of unspecified type to uint32. If translation fails, returns default value and error message
func ToUint32(value any) (result uint32, err error) {
	err = toTree(reflect.ValueOf(value), reflect.ValueOf(&result).Elem())

	return result, err
}

// ToUint64 converts a value of unspecified type to uint64. If translation fails, returns default value and error message
func ToUint64(value any) (result uint64, err error) {
	err = toTree(reflect.ValueOf(value), reflect.ValueOf(&result).Elem())

	return result, err
}

// ToUintptr converts a value of unspecified type to uintptr. If translation fails, returns default value and error message
func ToUintptr(value any) (result uintptr, err error) {
	err = toTree(reflect.ValueOf(value), reflect.ValueOf(&result).Elem())

	return result, err
}

// ToUintValue converts a value of unspecified type to uint. If the translation fails, returns the default value
func ToUintValue(value any) (result uint) {
	_ = toTree(reflect.ValueOf(value), reflect.ValueOf(&result).Elem())

	return result
}

// ToUint8Value converts a value of unspecified type to uint8. If the translation fails, returns the default value
func ToUint8Value(value any) (result uint8) {
	_ = toTree(reflect.ValueOf(value), reflect.ValueOf(&result).Elem())

	return result
}

// ToUint16Value converts a value of unspecified type to uint16. If the translation fails, returns the default value
func ToUint16Value(value any) (result uint16) {
	_ = toTree(reflect.ValueOf(value), reflect.ValueOf(&result).Elem())

	return result
}

// ToUint32Value converts a value of unspecified type to uint32. If the translation fails, returns the default value
func ToUint32Value(value any) (result uint32) {
	_ = toTree(reflect.ValueOf(value), reflect.ValueOf(&result).Elem())

	return result
}

// ToUint64Value converts a value of unspecified type to uint64. If the translation fails, returns the default value
func ToUint64Value(value any) (result uint64) {
	_ = toTree(reflect.ValueOf(value), reflect.ValueOf(&result).Elem())

	return result
}

// ToUintptrValue converts a value of unspecified type to uintptr. If the translation fails, returns the default value
func ToUintptrValue(value any) (result uintptr) {
	_ = toTree(reflect.ValueOf(value), reflect.ValueOf(&result).Elem())

	return result
}
