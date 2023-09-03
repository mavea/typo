package typo

import (
	"errors"
	"reflect"
)

// listToComplex converts a set of values into complex128
func listToComplex(r reflect.Value, i reflect.Value) (result complex128, err error) {
	var floatR, floatI float64
	if err = toTree(r, reflect.ValueOf(&floatR).Elem()); nil != err {
		return result, err
	}
	if err = toTree(i, reflect.ValueOf(&floatI).Elem()); nil != err {
		return result, err
	}

	return complex(floatR, floatI), nil
}

// toComplex converts reflect.Value to complex128. If translation fails, returns default value and error message
func toComplex(value reflect.Value, kindResult reflect.Kind) (result complex128, err error) {
	typeValue := value.Type()
	if nil == typeValue {
		return result, nil
	}
	switch typeValue.Kind() {
	case reflect.Pointer:
		if value.IsNil() {
			return 0, nil
		}
		return toComplex(value.Elem(), kindResult)
	case reflect.Float32, reflect.Float64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Bool:
		return result, errors.New(`few elements to form a complex`)
	case reflect.Array, reflect.Slice:
		if 1 < value.Len() {
			return listToComplex(value.Index(0), value.Index(1))
		} else {
			return result, errors.New(`few elements to form a complex`)
		}
	case reflect.Map:
		keysValue := value.MapKeys()
		if 1 < len(keysValue) {
			return listToComplex(value.MapIndex(keysValue[0]), value.MapIndex(keysValue[1]))
		} else {
			return result, errors.New(`few elements to form a complex`)
		}
	case reflect.Struct:
		if 1 < value.NumField() {
			return listToComplex(value.Field(0), value.Field(1))
		} else {
			return result, errors.New(`few elements to form a complex`)
		}
	case reflect.Complex64, reflect.Complex128:
		return value.Complex(), nil
	case reflect.Func:
		return result, errors.New(`argument cannot be a function call`)
	case reflect.Chan, reflect.Interface, reflect.UnsafePointer, reflect.String:
		return result, errors.New(typeValue.String() + ` format cannot be converted to ` + kindResult.String())
	}

	return result, errors.New(`unknown format ` + typeValue.String() + ` and cannot be converted to ` + kindResult.String())
}

// ToComplex64 converts a value of unspecified type to complex64. If translation fails, returns default value and error message
func ToComplex64(value any) (result complex64, err error) {
	err = toTree(reflect.ValueOf(value), reflect.ValueOf(&result).Elem())

	return result, err
}

// ToComplex128 converts a value of unspecified type to complex128. If translation fails, returns default value and error message
func ToComplex128(value any) (result complex128, err error) {
	err = toTree(reflect.ValueOf(value), reflect.ValueOf(&result).Elem())

	return result, err
}

// ToComplex64Value converts a value of unspecified type to complex64. If the translation fails, returns the default value
func ToComplex64Value(value any) (result complex64) {
	_ = toTree(reflect.ValueOf(value), reflect.ValueOf(&result).Elem())

	return result
}

// ToComplex128Value converts a value of unspecified type to complex128. If the translation fails, returns the default value
func ToComplex128Value(value any) (result complex128) {
	_ = toTree(reflect.ValueOf(value), reflect.ValueOf(&result).Elem())

	return result
}
