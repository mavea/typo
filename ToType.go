package typo

import (
	"errors"
	"reflect"
)

// toTree converts the value passed in Reflect.Value to the specified type passed as Reflect.Value. If translation fails, returns default value and error message
func toTree(value reflect.Value, result reflect.Value) error {
	if !value.IsValid() {
		return nil
	}
	typeResult := result.Type()
	if nil == typeResult {
		return nil
	}
	if reflect.Interface == value.Kind() {
		value = reflect.ValueOf(value.Interface())
	}
	switch result.Kind() {
	case reflect.Bool:
		bo, err := toBool(value)
		if nil != err {
			return err
		}
		result.Set(reflect.ValueOf(bo))

		return nil
	case reflect.Int:
		i64, err := toInt64(value, reflect.Int)
		if nil != err {
			return err
		}
		if -2147483648 > i64 || 2147483647 < i64 {
			return errors.New(`the passed value is not in the range of the int data type`)
		}
		i := int(i64)
		result.Set(reflect.ValueOf(i))

		return nil
	case reflect.Int8:
		i64, err := toInt64(value, reflect.Int8)
		if nil != err {
			return err
		}
		if -128 > i64 || 127 < i64 {
			return errors.New(`the passed value is not in the range of the int8 data type`)
		}
		i8 := int8(i64)
		result.Set(reflect.ValueOf(i8))

		return nil
	case reflect.Int16:
		i64, err := toInt64(value, reflect.Int16)
		if nil != err {
			return err
		}
		if -32768 > i64 || 32767 < i64 {
			return errors.New(`the passed value is not in the range of the int16 data type`)
		}
		i16 := int16(i64)
		result.Set(reflect.ValueOf(i16))

		return nil
	case reflect.Int32:
		i64, err := toInt64(value, reflect.Int32)
		if nil != err {
			return err
		}
		if -2147483648 > i64 || 2147483647 < i64 {
			return errors.New(`the passed value is not in the range of the int32 data type`)
		}
		i32 := int32(i64)
		result.Set(reflect.ValueOf(i32))

		return nil
	case reflect.Int64:
		i64, err := toInt64(value, reflect.Int64)
		if nil != err {
			return err
		}
		if -9223372036854775808 > i64 || 9223372036854775807 < i64 {
			return errors.New(`the passed value is not in the range of the int64 data type`)
		}
		result.Set(reflect.ValueOf(i64))

		return nil
	case reflect.Uint:
		ui64, err := toUint64(value, reflect.Uint)
		if nil != err {
			return err
		}
		if 4294967295 < ui64 {
			return errors.New(`the passed value is not in the range of the uint data type`)
		}
		ui := uint(ui64)
		result.Set(reflect.ValueOf(ui))

		return nil
	case reflect.Uint8:
		ui64, err := toUint64(value, reflect.Uint8)
		if nil != err {
			return err
		}
		if 255 < ui64 {
			return errors.New(`the passed value is not in the range of the uint8 data type`)
		}
		ui8 := uint8(ui64)
		result.Set(reflect.ValueOf(ui8))

		return nil
	case reflect.Uint16:
		ui64, err := toUint64(value, reflect.Uint16)
		if nil != err {
			return err
		}
		if 65535 < ui64 {
			return errors.New(`the passed value is not in the range of the uint16 data type`)
		}
		ui16 := uint16(ui64)
		result.Set(reflect.ValueOf(ui16))

		return nil
	case reflect.Uint32:
		ui64, err := toUint64(value, reflect.Uint32)
		if nil != err {
			return err
		}
		if 4294967295 < ui64 {
			return errors.New(`the passed value is not in the range of the uint32 data type`)
		}
		ui32 := uint32(ui64)
		result.Set(reflect.ValueOf(ui32))

		return nil
	case reflect.Uint64:
		ui64, err := toUint64(value, reflect.Uint64)
		if nil != err {
			return err
		}
		if 18446744073709551615 < ui64 {
			return errors.New(`the passed value is not in the range of the uint64 data type`)
		}
		result.Set(reflect.ValueOf(ui64))

		return nil
	case reflect.Uintptr:
		ui64, err := toUint64(value, reflect.Uintptr)
		if nil != err {
			return err
		}
		uiptr := uintptr(ui64)
		result.Set(reflect.ValueOf(uiptr))

		return nil
	case reflect.Float32:
		fl64, err := toFloat64(value, reflect.Float32)
		if nil != err {
			return err
		}
		fl32 := float32(fl64)
		result.Set(reflect.ValueOf(fl32))

		return nil
	case reflect.Float64:
		fl64, err := toFloat64(value, reflect.Float64)
		if nil != err {
			return err
		}
		result.Set(reflect.ValueOf(fl64))

		return nil
	case reflect.Complex64:
		co128, err := toComplex(value, reflect.Complex64)
		if nil != err {
			return err
		}
		co64 := complex64(co128)
		result.Set(reflect.ValueOf(co64))

		return nil
	case reflect.Complex128:
		co128, err := toComplex(value, reflect.Complex128)
		if nil != err {
			return err
		}
		result.Set(reflect.ValueOf(co128))

		return nil
	case reflect.String:
		st, err := toString(value)
		if nil != err {
			return err
		}
		result.Set(reflect.ValueOf(st))

		return nil
	case reflect.Pointer:
		item := reflect.New(result.Type().Elem())
		if err := toTree(value, item.Elem()); nil != err {
			return err
		}
		if !result.IsZero() {
			result.Elem().Set(item.Elem())
		} else {
			result.Set(item)
		}

		return nil
	case reflect.Array:
		if nil == value.Type() {
			result.SetZero()
			return nil
		}
		elemResult := reflect.New(typeResult)
		kindValue := value.Kind()
		switch kindValue {
		case reflect.Map:
			if value.IsNil() {
				result.SetZero()
				return nil
			}
			lenElemResult := elemResult.Len()
			i := 0
			for _, elem := range value.MapKeys() {
				if err := toTree(value.MapIndex(elem), elemResult.Elem().Index(i)); nil != err {
					return err
				}
				i++
				if i == lenElemResult {
					break
				}
			}
			break
		case reflect.Slice, reflect.Array:
			if reflect.Slice == kindValue {
				if value.IsNil() {
					result.SetZero()
					return nil
				}
			}
			for i := 0; i < value.Len() && i < elemResult.Elem().Len(); i++ {
				if err := toTree(value.Index(i), elemResult.Elem().Index(i)); nil != err {
					return err
				}
			}
			break
		case reflect.Struct:
			i := 0
			if err := arrayFromStruct(value, elemResult, &i); nil != err {
				return err
			}
			break
		default:
			if err := toTree(value, elemResult.Elem().Index(0)); nil != err {
				return err
			}
			break
		}
		if result.CanAddr() {
			result.Set(elemResult.Elem())
		} else {
			return errors.New(`set variable is not can addr`)
		}
		return nil
	case reflect.Slice:
		if nil == value.Type() {
			result.SetZero()
			return nil
		}
		elemResult := reflect.MakeSlice(typeResult, 0, 0)
		kindValue := value.Kind()
		switch kindValue {
		case reflect.Map:
			if value.IsNil() {
				result.SetZero()
				return nil
			}
			for _, elem := range value.MapKeys() {
				item := reflect.New(typeResult.Elem())
				if err := toTree(value.MapIndex(elem), item.Elem()); nil != err {
					return err
				}
				elemResult = reflect.Append(elemResult, item.Elem())
			}
			break
		case reflect.Slice, reflect.Array:
			if reflect.Slice == kindValue {
				if value.IsNil() {
					result.SetZero()
					return nil
				}
			}
			for i := 0; i < value.Len(); i++ {
				item := reflect.New(typeResult.Elem())
				if err := toTree(value.Index(i), item.Elem()); nil != err {
					return err
				}
				elemResult = reflect.Append(elemResult, item.Elem())
			}
			break
		case reflect.Struct:
			if err := sliceFromStruct(value, &elemResult); nil != err {
				return err
			}
			break
		default:
			item := reflect.New(typeResult.Elem())
			if err := toTree(value, item.Elem()); nil != err {
				return err
			}
			elemResult = reflect.Append(elemResult, item.Elem())
			break
		}
		if result.CanAddr() {
			result.Set(elemResult)
		} else {
			return errors.New(`set variable is not can addr`)
		}
		return nil
	case reflect.Map:
		typeValue := value.Type()
		if nil == typeValue {
			result.SetZero()
			return nil
		}
		elemResult := reflect.MakeMap(typeResult)
		kindValue := value.Kind()
		switch kindValue {
		case reflect.Struct:
			if err := mapFromStruct(value, elemResult); nil != err {
				return err
			}
			break
		case reflect.Map:
			if value.IsNil() {
				result.SetZero()
				return nil
			}
			keys := value.MapKeys()
			for _, itemKey := range keys {
				convertItemKey := itemKey.Convert(value.Type().Key())
				key := reflect.New(typeResult.Key())
				if err := toTree(convertItemKey, key.Elem()); nil != err {
					return err
				}
				temp := value.MapIndex(convertItemKey)
				item := reflect.New(typeResult.Elem())
				if err := toTree(temp, item.Elem()); nil != err {
					return err
				}
				elemResult.SetMapIndex(key.Elem(), item.Elem())
			}
			break
		case reflect.Slice, reflect.Array:
			if reflect.Slice == kindValue {
				if value.IsNil() {
					result.SetZero()
					return nil
				}
			}

			for i := 0; i < value.Len(); i++ {
				key := reflect.New(typeResult.Key())
				if err := toTree(reflect.ValueOf(i), key.Elem()); nil != err {
					return err
				}
				item := reflect.New(typeResult.Elem())
				if err := toTree(value.Index(i), item.Elem()); nil != err {
					return err
				}

				elemResult.SetMapIndex(key.Elem(), item.Elem())
			}
			break
		default:
			return errors.New(`format ` + typeValue.String() + ` and cannot be converted to ` + typeResult.String())
		}
		if result.CanAddr() {
			result.Set(elemResult)
		} else {
			return errors.New(`set variable is not can addr`)
		}
		return nil
	case reflect.Struct:
		if !value.IsValid() {
			result.SetZero()
			return nil
		}
		elemResult := reflect.New(typeResult)
		switch value.Kind() {
		case reflect.Struct:
			if err := structToStruct(value, elemResult.Elem()); nil != err {
				return err
			}
			break
		case reflect.Array:
			i := 0
			if err := sliceToStruct(value, elemResult.Elem(), &i); nil != err {
				return err
			}
			break
		case reflect.Slice:
			i := 0
			if err := sliceToStruct(value, elemResult.Elem(), &i); nil != err {
				return err
			}
			break
		case reflect.Map:
			if err := mapToStruct(value, elemResult.Elem()); nil != err {
				return err
			}
			break
		default:
			return errors.New(`format ` + value.Type().String() + ` and cannot be converted to ` + typeResult.String())
		}
		if result.CanAddr() {
			result.Set(elemResult.Elem())
		} else {
			return errors.New(`set variable is not can addr`)
		}
		return nil
	case reflect.Func:
		return errors.New(`argument cannot be a function call`)
	case reflect.Chan, reflect.Interface, reflect.UnsafePointer:
		return errors.New(value.Type().String() + ` format cannot be converted to ` + typeResult.String())
	}

	return errors.New(`unknown format ` + value.Type().String() + ` and cannot be converted to ` + typeResult.String())
}

// ToType converts the passed value of undefined type to the specified type. If translation fails, returns default value and error message
//
// Is a synonym for the function To
func ToType[typeResult any](value any) (result typeResult, err error) {
	if err = toTree(reflect.ValueOf(value), reflect.ValueOf(&result).Elem()); nil != err {
		var r typeResult
		return r, err
	}

	return result, err
}

// To converts the passed value of undefined type to the specified type. If translation fails, returns default value and error message
func To[typeResult any](value any) (result typeResult, err error) {
	if err = toTree(reflect.ValueOf(value), reflect.ValueOf(&result).Elem()); nil != err {
		var r typeResult
		return r, err
	}

	return result, err
}

// ToTypeValue converts the passed value of undefined type to the specified type. If translation fails, returns default value and error message
//
// Is a synonym for the function ToValue
// The conversion goes up to the first element whose type cast resulted in an error.
func ToTypeValue[typeResult any](value any) (result typeResult) {
	if err := toTree(reflect.ValueOf(value), reflect.ValueOf(&result).Elem()); nil != err {
		var r typeResult
		return r
	}

	return result
}

// ToValue converts the passed value of undefined type to the specified type. If translation fails, returns default value and error message
//
// The conversion goes up to the first element whose type cast resulted in an error.
func ToValue[typeResult any](value any) (result typeResult) {
	if err := toTree(reflect.ValueOf(value), reflect.ValueOf(&result).Elem()); nil != err {
		var r typeResult
		return r
	}

	return result
}
