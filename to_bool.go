package typo

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
)

// toBool converts reflect.Value to bool. If translation fails, returns default value and error message
func toBool(value reflect.Value) (bool, error) {
	typeValue := value.Type()
	if nil == typeValue {
		return false, nil
	}
	switch typeValue.Kind() {
	case reflect.Bool:
		return value.Interface().(bool), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int() != 0, nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return value.Uint() != 0, nil
	case reflect.Float32, reflect.Float64:
		return value.Float() != 0.0, nil
	case reflect.Pointer:
		if value.IsNil() {
			return false, nil
		}
		return toBool(value.Elem())
	case reflect.String:
		stringValue := value.String()
		switch strings.ToLower(stringValue) {
		case `false`, `0`, `null`, `nil`, `-`, ``:
			return false, nil
		case `true`, `1`, `+`:
			return true, nil
		}
		f64, err := strconv.ParseFloat(stringValue, 64)
		if nil != err {
			return true, nil
		}
		if 0.0 != f64 {
			return true, nil
		}
		return false, nil
	case reflect.Map:
		if !value.IsNil() {
			return 0 != len(value.MapKeys()), nil
		}
		return false, nil
	case reflect.Slice:
		if !value.IsNil() {
			return 0 != value.Len(), nil
		}
		return false, nil
	case reflect.Complex64, reflect.Complex128:
		if complex(0, 0) == value.Complex() {
			return false, nil
		}
		return true, nil
	case reflect.Func:
		return false, errors.New(`argument cannot be a function call`)
	case reflect.Array, reflect.Struct, reflect.Chan, reflect.Interface, reflect.UnsafePointer:
		return false, errors.New(typeValue.String() + ` format cannot be converted to bool`)
	}

	return false, errors.New(`unknown format ` + typeValue.String() + ` and cannot be converted to bool`)
}

// ToBool converts a value of unspecified type to bool. If translation fails, returns default value and error message
func ToBool(value any) (result bool, err error) {
	err = toTree(reflect.ValueOf(value), reflect.ValueOf(&result).Elem())

	return result, err
}

// ToBoolValue converts a value of unspecified type to bool. If the translation fails, returns the default value
func ToBoolValue(value any) (result bool) {
	_ = toTree(reflect.ValueOf(value), reflect.ValueOf(&result).Elem())

	return result
}
