package typo

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

// toInt64 converts reflect.Value to int64. If translation fails, returns default value and error message
func toInt64(value reflect.Value, kindResult reflect.Kind) (int64, error) {
	typeValue := value.Type()
	if nil == typeValue {
		return 0, nil
	}
	kindValue := typeValue.Kind()
	switch kindValue {
	case reflect.Bool:
		if value.Interface().(bool) {
			return 1, nil
		}
		return 0, nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int(), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return int64(value.Uint()), nil
	case reflect.Float32, reflect.Float64:
		return int64(value.Float()), nil
	case reflect.Pointer:
		if value.IsNil() {
			return 0, nil
		}
		return toInt64(value.Elem(), kindResult)
	case reflect.String, reflect.Interface:
		if reflect.Interface == kindValue {
			value = reflect.ValueOf(fmt.Sprint(value.Interface()))
		}
		f64, err := strconv.ParseFloat(value.String(), 64)
		if nil != err {
			return 0, err
		}
		return int64(f64), err
	case reflect.Func:
		return 0, errors.New(`argument cannot be a function call`)
	case reflect.Complex64, reflect.Complex128, reflect.Array, reflect.Slice, reflect.Struct, reflect.Map, reflect.Chan, reflect.UnsafePointer:
		return 0, errors.New(typeValue.String() + ` format cannot be converted to ` + kindResult.String())
	}

	return 0, errors.New(`unknown format ` + typeValue.String() + ` and cannot be converted to ` + kindResult.String())
}

// ToInt converts a value of unspecified type to int. If translation fails, returns default value and error message
func ToInt(value any) (result int, err error) {
	err = toTree(reflect.ValueOf(value), reflect.ValueOf(&result).Elem())

	return result, err
}

// ToInt8 converts a value of unspecified type to int8. If translation fails, returns default value and error message
func ToInt8(value any) (result int8, err error) {
	err = toTree(reflect.ValueOf(value), reflect.ValueOf(&result).Elem())

	return result, err
}

// ToInt16 converts a value of unspecified type to int16. If translation fails, returns default value and error message
func ToInt16(value any) (result int16, err error) {
	err = toTree(reflect.ValueOf(value), reflect.ValueOf(&result).Elem())

	return result, err
}

// ToInt32 converts a value of unspecified type to int32. If translation fails, returns default value and error message
func ToInt32(value any) (result int32, err error) {
	err = toTree(reflect.ValueOf(value), reflect.ValueOf(&result).Elem())

	return result, err
}

// ToInt64 converts a value of unspecified type to int64. If translation fails, returns default value and error message
func ToInt64(value any) (result int64, err error) {
	err = toTree(reflect.ValueOf(value), reflect.ValueOf(&result).Elem())

	return result, err
}

// ToIntValue converts a value of unspecified type to int. If the translation fails, returns the default value
func ToIntValue(value any) (result int) {
	_ = toTree(reflect.ValueOf(value), reflect.ValueOf(&result).Elem())

	return result
}

// ToInt8Value converts a value of unspecified type to int8. If the translation fails, returns the default value
func ToInt8Value(value any) (result int8) {
	_ = toTree(reflect.ValueOf(value), reflect.ValueOf(&result).Elem())

	return result
}

// ToInt16Value converts a value of unspecified type to int16. If the translation fails, returns the default value
func ToInt16Value(value any) (result int16) {
	_ = toTree(reflect.ValueOf(value), reflect.ValueOf(&result).Elem())

	return result
}

// ToInt32Value converts a value of unspecified type to int32. If the translation fails, returns the default value
func ToInt32Value(value any) (result int32) {
	_ = toTree(reflect.ValueOf(value), reflect.ValueOf(&result).Elem())

	return result
}

// ToInt64Value converts a value of unspecified type to int64. If the translation fails, returns the default value
func ToInt64Value(value any) (result int64) {
	_ = toTree(reflect.ValueOf(value), reflect.ValueOf(&result).Elem())

	return result
}
