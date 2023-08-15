package typo

import (
	"reflect"
)

// ToSlice converts a value of undefined type to a slice of elements of the specified type. If translation fails, returns default value and error message
func ToSlice[T any](value any) (result []T, err error) {
	err = toTree(reflect.ValueOf(value), reflect.ValueOf(&result))

	return result, err
}

// ToSliceValue converts a value of undefined type to a slice of elements of the specified type
//
// The conversion goes up to the first element whose type cast resulted in an error.
func ToSliceValue[T any](value any) (result []T) {
	_ = toTree(reflect.ValueOf(value), reflect.ValueOf(&result))

	return result
}
