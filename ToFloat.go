package typo

import (
	"errors"
	"reflect"
	"strconv"
)

// toFloat64 converts reflect.Value to float64. If translation fails, returns default value and error message
func toFloat64(value reflect.Value, kindResult reflect.Kind) (float64, error) {
	typeValue := value.Type()
	if nil == typeValue {
		return 0, nil
	}
	switch typeValue.Kind() {
	case reflect.Bool:
		if value.Interface().(bool) {
			return 1.0, nil
		}
		return 0.0, nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(value.Int()), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return float64(value.Uint()), nil
	case reflect.Float32, reflect.Float64:
		return value.Float(), nil
	case reflect.String:
		return strconv.ParseFloat(value.String(), 64)
	case reflect.Pointer:
		if value.IsNil() {
			return 0, nil
		}
		return toFloat64(value.Elem(), kindResult)
	case reflect.Func:
		return 0, errors.New(`argument cannot be a function call`)
	case reflect.Complex64, reflect.Complex128, reflect.Array, reflect.Slice, reflect.Struct, reflect.Map, reflect.Chan, reflect.Interface, reflect.UnsafePointer:
		return 0, errors.New(typeValue.String() + ` format cannot be converted to ` + kindResult.String())
	}

	return 0, errors.New(`unknown format ` + typeValue.String() + ` and cannot be converted to ` + kindResult.String())
}

// ToFloat32 converts a value of unspecified type to float32. If translation fails, returns default value and error message
func ToFloat32(value any) (result float32, err error) {
	err = toTree(reflect.ValueOf(value), reflect.ValueOf(&result).Elem())

	return result, err
}

// ToFloat64 converts a value of unspecified type to float64. If translation fails, returns default value and error message
func ToFloat64(value any) (result float64, err error) {
	err = toTree(reflect.ValueOf(value), reflect.ValueOf(&result).Elem())

	return result, err
}

// ToFloat32Value converts a value of unspecified type to float32. If the translation fails, returns the default value
func ToFloat32Value(value any) (result float32) {
	_ = toTree(reflect.ValueOf(value), reflect.ValueOf(&result).Elem())

	return result
}

// ToFloat64Value converts a value of unspecified type to float64. If the translation fails, returns the default value
func ToFloat64Value(value any) (result float64) {
	_ = toTree(reflect.ValueOf(value), reflect.ValueOf(&result).Elem())

	return result
}
