package typo

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

// sliceToString converts array and slice reflect.Value to string
func sliceToString(value reflect.Value) (string, error) {
	var result string
	first := true
	for i := 0; i < value.Len(); i++ {
		str, err := toString(value.Index(i))
		if nil != err {
			return ``, err
		}
		if `` == str {
			continue
		}
		if first {
			first = false
		} else {
			result += ` `
		}
		result += str
	}

	return result, nil
}

// toString converts reflect.Value to string. If translation fails, returns default value and error message
func toString(value reflect.Value) (string, error) {
	typeValue := value.Type()
	if nil == typeValue {
		return writeNullInString, nil
	}
	switch typeValue.Kind() {
	case reflect.Bool:
		return strconv.FormatBool(value.Interface().(bool)), nil
	case reflect.Int:
		return strconv.FormatInt(int64(int(value.Int())), 10), nil
	case reflect.Int8:
		return strconv.FormatInt(int64(int8(value.Int())), 10), nil
	case reflect.Int16:
		return strconv.FormatInt(int64(int16(value.Int())), 10), nil
	case reflect.Int32:
		return strconv.FormatInt(int64(int32(value.Int())), 10), nil
	case reflect.Int64:
		return strconv.FormatInt(value.Int(), 10), nil
	case reflect.Uint:
		return strconv.FormatUint(value.Uint(), 10), nil
	case reflect.Uint8:
		return strconv.FormatUint(uint64(uint8(value.Uint())), 10), nil
	case reflect.Uint16:
		return strconv.FormatUint(uint64(uint16(value.Uint())), 10), nil
	case reflect.Uint32:
		return strconv.FormatUint(uint64(uint32(value.Uint())), 10), nil
	case reflect.Uint64:
		return strconv.FormatUint(value.Uint(), 10), nil
	case reflect.Uintptr:
		return strconv.FormatUint(uint64(uint(value.Uint())), 10), nil
	case reflect.Float32:
		return strconv.FormatFloat(value.Float(), 'G', 32, 32), nil
	case reflect.Float64:
		return strconv.FormatFloat(value.Float(), 'G', 64, 64), nil
	case reflect.Complex64:
		return strconv.FormatComplex(value.Complex(), 'G', 64, 64), nil
	case reflect.Complex128:
		return strconv.FormatComplex(value.Complex(), 'G', 128, 128), nil
	case reflect.Array:
		return sliceToString(value)
	case reflect.Pointer:
		if value.IsNil() {
			return writeNullInString, nil
		}
		return toString(value.Elem())
	case reflect.Slice:
		if value.IsNil() {
			return ``, nil
		}
		return sliceToString(value)
	case reflect.String:
		return value.Interface().(string), nil
	case reflect.Interface:
		return fmt.Sprint(value.Interface()), nil
	case reflect.Struct, reflect.Map:
		jsn, err := json.Marshal(value.Interface())
		return string(jsn), err
	case reflect.Func:
		return ``, errors.New(`argument cannot be a function call`)
	case reflect.Chan, reflect.UnsafePointer:
		return ``, errors.New(typeValue.String() + ` format cannot be converted to string`)
	}

	return ``, errors.New(`unknown format ` + typeValue.String() + ` and cannot be converted to string`)
}

// ToString converts a value of unspecified type to string. If translation fails, returns default value and error message
func ToString(value any) (result string, err error) {
	err = toTree(reflect.ValueOf(value), reflect.ValueOf(&result).Elem())

	return result, err
}

// ToStringValue converts a value of unspecified type to string. If translation fails, returns default value
func ToStringValue(value any) (result string) {
	_ = toTree(reflect.ValueOf(value), reflect.ValueOf(&result).Elem())

	return result
}
