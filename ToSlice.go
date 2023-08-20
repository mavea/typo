package typo

import (
	"reflect"
)

// ToSlice converts a value of undefined type to a slice of elements of the specified type. If translation fails, returns default value and error message
func ToSlice[typeResult any](value any) (result []typeResult, err error) {
	if err = toTree(reflect.ValueOf(value), reflect.ValueOf(&result).Elem()); nil != err {
		var r []typeResult
		return r, err
	}

	return result, err
}

// ToSliceValue converts a value of undefined type to a slice of elements of the specified type
//
// The conversion goes up to the first element whose type cast resulted in an error.
func ToSliceValue[typeResult any](value any) (result []typeResult) {
	if err := toTree(reflect.ValueOf(value), reflect.ValueOf(&result).Elem()); nil != err {
		var r []typeResult
		return r
	}

	return result
}
