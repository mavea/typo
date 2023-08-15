package typo

import (
	"errors"
	"reflect"
)

// arrayFromStruct filling an array with structure elements
func arrayFromStruct(value reflect.Value, result reflect.Value, pointerIntegerI *int) error {
	lenResult := result.Len()
	for i := 0; i < value.NumField(); i++ {
		if lenResult == *pointerIntegerI {
			return errors.New(`the passed variable or its part of type ` + value.Type().String() + ` cannot be placed in type ` + result.Elem().Type().String() + ` because it is greater than it`)
		}
		fieldValue := value.Field(i)
		if reflect.Struct == fieldValue.Kind() {
			if err := arrayFromStruct(fieldValue, result, pointerIntegerI); nil != err {
				return err
			}
		} else {
			if err := toTree(fieldValue, result.Elem().Index(*pointerIntegerI)); nil != err {
				return err
			}
			*pointerIntegerI++
		}
	}

	return nil
}

// sliceFromStruct filling an slice with structure elements
func sliceFromStruct(value reflect.Value, result *reflect.Value) error {
	typeResult := (*result).Type()
	for i := 0; i < value.NumField(); i++ {
		fieldValue := value.Field(i)
		if reflect.Struct == fieldValue.Kind() {
			if err := sliceFromStruct(fieldValue, result); nil != err {
				return err
			}
		} else {
			item := reflect.New(typeResult.Elem())
			if err := toTree(fieldValue, item.Elem()); nil != err {
				return err
			}
			*result = reflect.Append(*result, item.Elem())
		}
	}

	return nil
}

// sliceFromStruct filling a structure with slice elements
func sliceToStruct(value reflect.Value, result reflect.Value, pointerIntegerI *int) error {
	lenResult := result.NumField()
	typeResult := result.Type()
	i := 0
	for ; *pointerIntegerI < value.Len(); *pointerIntegerI++ {
		fieldResult := result.Field(i)
		if reflect.Struct == fieldResult.Kind() && `` == searchNameInStructureTag(typeResult.Field(i).Tag.Get(structureTagContainingNameTheElementJson)) {
			if err := sliceToStruct(value, fieldResult, pointerIntegerI); nil != err {
				return err
			}
		} else {
			if err := toTree(value.Index(*pointerIntegerI), fieldResult); nil != err {
				return err
			}
		}
		i++
		if i == lenResult {
			return nil
		}
	}

	return nil
}

// sliceFromStruct filling a structure with elements of another structure
func structToStruct(value reflect.Value, result reflect.Value) error {
	typeResult := result.Type()
	for i := 0; i < result.NumField(); i++ {
		fieldResult := result.Field(i)
		fieldTypeResult := typeResult.Field(i)
		tagName := searchNameInStructureTag(fieldTypeResult.Tag.Get(structureTagContainingNameTheElementJson))
		if reflect.Struct == fieldResult.Kind() && `` == tagName {
			if err := structToStruct(value, fieldResult); nil != err {
				return err
			}
		} else {
			if `` != tagName {
				vF := value.FieldByName(tagName)
				if !vF.IsValid() {
					vF = findFieldByTag(value, tagName)
				}
				if vF.IsValid() {
					if err := toTree(vF, fieldResult); nil != err {
						return err
					}
					continue
				}
			}
			tagName = fieldTypeResult.Name
			vF := value.FieldByName(tagName)
			if vF.IsValid() {
				if err := toTree(vF, fieldResult); nil != err {
					return err
				}
			}
		}
	}

	return nil
}

// findFieldByTag search for a structure element by its tag
func findFieldByTag(value reflect.Value, tag string) (field reflect.Value) {
	typeValue := value.Type()
	for i := 0; i < value.NumField(); i++ {
		if searchNameInStructureTag(typeValue.Field(i).Tag.Get(structureTagContainingNameTheElementJson)) == tag {
			return value.Field(i)
		}
	}

	return field
}

// mapToStruct filling a structure with map elements
func mapToStruct(value reflect.Value, result reflect.Value) error {
	typeResult := result.Type()
	typeValue := value.Type()
	for i := 0; i < result.NumField(); i++ {
		fieldValue := result.Field(i)
		field := typeResult.Field(i)
		keyTo := searchNameInStructureTag(field.Tag.Get(structureTagContainingNameTheElementJson))
		if reflect.Struct == fieldValue.Kind() && `` == keyTo {
			if err := mapToStruct(value, fieldValue); nil != err {
				return err
			}
		} else {
			if `` == keyTo {
				keyTo = field.Name
			}
			key := reflect.New(typeValue.Key())
			if err := toTree(reflect.ValueOf(keyTo), key.Elem()); nil != err {
				return err
			}
			elemValue := value.MapIndex(key.Elem())
			if elemValue.IsValid() {
				if err := toTree(elemValue, fieldValue); nil != err {
					return err
				}
			}
		}
	}

	return nil
}

// mapToStruct filling a map with structure elements
func mapFromStruct(value reflect.Value, result reflect.Value) error {
	typeResult := result.Type()
	typeValue := value.Type()
	for i := 0; i < value.NumField(); i++ {
		fieldValue := value.Field(i)
		if reflect.Struct == fieldValue.Kind() {
			if err := mapFromStruct(fieldValue, result); nil != err {
				return err
			}
		} else {
			fieldTypeValue := typeValue.Field(i)
			tag := searchNameInStructureTag(fieldTypeValue.Tag.Get(structureTagContainingNameTheElementJson))
			keyVal := reflect.ValueOf(tag)
			if `` == tag {
				keyVal = reflect.ValueOf(fieldTypeValue.Name)
			}
			key := reflect.New(typeResult.Key())
			if err := toTree(keyVal, key.Elem()); nil != err {
				return err
			}
			item := reflect.New(typeResult.Elem())
			if err := toTree(fieldValue, item.Elem()); nil != err {
				return err
			}
			result.SetMapIndex(key.Elem(), item.Elem())
		}
	}

	return nil
}
