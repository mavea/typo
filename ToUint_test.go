package typo

import (
	"errors"
	"math"
	"testing"
)

func getElementForTestUInt(num int) (testStructure, bool) {
	switch num {
	case 0:
		return newTest(num, `simple integer conversion test`, toTest(ToUint(1)), ToUintValue(1), uint(1), nil)
	case 1:
		var i int = 2
		return newTest(num, `simple integer conversion test`, toTest(ToUint(i)), ToUintValue(i), uint(i), nil)
	case 2:
		var i int = 2
		return newTest(num, `point integer conversion test`, toTest(ToUint(&i)), ToUintValue(&i), uint(i), nil)
	case 3:
		var i int8 = 8
		return newTest(num, `simple integer8 conversion test`, toTest(ToUint(i)), ToUintValue(i), uint(i), nil)
	case 4:
		var i int8 = 8
		return newTest(num, `point integer8 conversion test`, toTest(ToUint(&i)), ToUintValue(&i), uint(i), nil)
	case 5:
		var i int16 = 16
		return newTest(num, `simple integer16 conversion test`, toTest(ToUint(i)), ToUintValue(i), uint(i), nil)
	case 6:
		var i int16 = 16
		return newTest(num, `point integer16 conversion test`, toTest(ToUint(&i)), ToUintValue(&i), uint(i), nil)
	case 7:
		var i int32 = 32
		return newTest(num, `simple integer32 conversion test`, toTest(ToUint(i)), ToUintValue(i), uint(i), nil)
	case 8:
		var i int32 = 32
		return newTest(num, `point integer32 conversion test`, toTest(ToUint(&i)), ToUintValue(&i), uint(i), nil)
	case 9:
		var i int64 = 64
		return newTest(num, `simple integer64 conversion test`, toTest(ToUint(i)), ToUintValue(i), uint(i), nil)
	case 10:
		var i int64 = 64
		return newTest(num, `point integer64 conversion test`, toTest(ToUint(&i)), ToUintValue(&i), uint(i), nil)
	case 11:
		var ui uint = 51
		return newTest(num, `simple uint conversion test`, toTest(ToUint(ui)), ToUintValue(ui), uint(ui), nil)
	case 12:
		var ui uint = 51
		return newTest(num, `point uint conversion test`, toTest(ToUint(&ui)), ToUintValue(&ui), uint(ui), nil)
	case 13:
		var ui uint8 = 58
		return newTest(num, `simple uint8 conversion test`, toTest(ToUint(ui)), ToUintValue(ui), uint(ui), nil)
	case 14:
		var ui uint8 = 58
		return newTest(num, `point uint8 conversion test`, toTest(ToUint(&ui)), ToUintValue(&ui), uint(ui), nil)
	case 15:
		var ui uint16 = 116
		return newTest(num, `simple uint16 conversion test`, toTest(ToUint(ui)), ToUintValue(ui), uint(ui), nil)
	case 16:
		var ui uint16 = 116
		return newTest(num, `point uint16 conversion test`, toTest(ToUint(&ui)), ToUintValue(&ui), uint(ui), nil)
	case 17:
		var ui uint32 = 102
		return newTest(num, `simple uint32 conversion test`, toTest(ToUint(ui)), ToUintValue(ui), uint(ui), nil)
	case 18:
		var ui uint32 = 102
		return newTest(num, `point uint32 conversion test`, toTest(ToUint(&ui)), ToUintValue(&ui), uint(ui), nil)
	case 19:
		var ui uint64 = 104
		return newTest(num, `simple uint64 conversion test`, toTest(ToUint(ui)), ToUintValue(ui), uint(ui), nil)
	case 20:
		var ui uint64 = 104
		return newTest(num, `point uint64 conversion test`, toTest(ToUint(&ui)), ToUintValue(&ui), uint(ui), nil)
	case 21:
		var ui uintptr = 123
		return newTest(num, `simple uintptr conversion test`, toTest(ToUint(ui)), ToUintValue(ui), uint(ui), nil)
	case 22:
		var ui uintptr = 123
		return newTest(num, `point uintptr conversion test`, toTest(ToUint(&ui)), ToUintValue(&ui), uint(ui), nil)
	case 23:
		var f float32 = math.Pi
		return newTest(num, `simple float32 conversion test`, toTest(ToUint(f)), ToUintValue(f), uint(f), nil)
	case 24:
		var f float32 = math.Pi
		return newTest(num, `pointer float32 conversion test`, toTest(ToUint(&f)), ToUintValue(&f), uint(f), nil)
	case 25:
		var f float64 = math.Pi
		return newTest(num, `simple float64 conversion test`, toTest(ToUint(f)), ToUintValue(f), uint(f), nil)
	case 26:
		var f float64 = math.Pi
		return newTest(num, `pointer float64 conversion test`, toTest(ToUint(&f)), ToUintValue(&f), uint(f), nil)
	case 27:
		var elemNull *uint
		return newTest(num, `pointer pointer uint = nil conversion test`, toTest(ToUint(&elemNull)), ToUintValue(&elemNull), uint(0), nil)
	case 28:
		b := true
		return newTest(num, `simple bool = true conversion test`, toTest(ToUint(b)), ToUintValue(b), uint(1), nil)
	case 29:
		b := true
		return newTest(num, `pointer bool = true conversion test`, toTest(ToUint(&b)), ToUintValue(&b), uint(1), nil)
	case 30:
		b := false
		return newTest(num, `simple bool = false conversion test`, toTest(ToUint(b)), ToUintValue(b), uint(0), nil)
	case 31:
		b := false
		return newTest(num, `pointer bool = false conversion test`, toTest(ToUint(&b)), ToUintValue(&b), uint(0), nil)
	case 32:
		return newTest(num, `pointer nil conversion test`, toTest(ToUint(nil)), ToUintValue(nil), uint(0), nil)

	}

	return testStructure{}, false
}

func getElementForTestUintDefault(num int) (testStructure, bool) {
	switch num {
	case 0:
		return newTest(num, `simple 2147483647 conversion test`, toTest(ToUint(2147483647)), ToUintValue(2147483647), uint(2147483647), nil)
	case 1:
		return newTest(num, `simple -2147483648 conversion test`, toTest(ToUint(-2147483648)), ToUintValue(-2147483648), uint(0), errors.New(`uint cannot be negative`))
	case 2:
		slice := []int{4, 3}
		return newTest(num, `simple []int{4, 3} conversion test`, toTest(ToUint(slice)), ToUintValue(slice), uint(0), errors.New(`[]int format cannot be converted to uint`))
	case 3:
		slice := []int{4, 3}
		return newTest(num, `pointer []int{4, 3} conversion test`, toTest(ToUint(slice)), ToUintValue(slice), uint(0), errors.New(`[]int format cannot be converted to uint`))

	}

	return testStructure{}, false
}

func TestToUint(t *testing.T) {
	testByLists(t, []testListStructure{
		{
			Name:           `General tests`,
			FuncionGetList: getElementForTestUInt,
		},
		{
			Name:           `Personal tests`,
			FuncionGetList: getElementForTestUintDefault,
		},
	})
}

func getElementForTestUint8(num int) (testStructure, bool) {
	switch num {
	case 0:
		return newTest(num, `simple 127 conversion test`, toTest(ToUint8(127)), ToUint8Value(127), uint8(127), nil)
	case 1:
		return newTest(num, `simple -128 conversion test`, toTest(ToUint8(-128)), ToUint8Value(-128), uint8(0), errors.New(`uint8 cannot be negative`))
	case 2:
		arrayBool := [2]bool{true, false}
		return newTest(num, `simple []int{4, 3} conversion test`, toTest(ToUint8(arrayBool)), ToUint8Value(arrayBool), uint8(0), errors.New(`[2]bool format cannot be converted to uint8`))
	case 3:
		arrayBool := [2]bool{true, false}
		return newTest(num, `pointer []int{4, 3} conversion test`, toTest(ToUint8(arrayBool)), ToUint8Value(arrayBool), uint8(0), errors.New(`[2]bool format cannot be converted to uint8`))
	case 4:
		var f float64 = math.Pi + float64(987654318.0)
		return newTest(num, `simple pi+ conversion test`, toTest(ToUint8(f)), ToUint8Value(f), uint8(0), errors.New(`the passed value is not in the range of the uint8 data type`))

	}

	return testStructure{}, false
}

func TestToUint8(t *testing.T) {
	testByLists(t, []testListStructure{
		{
			Name:           `Personal tests`,
			FuncionGetList: getElementForTestUint8,
		},
	})
}

func getElementForTestUint16(num int) (testStructure, bool) {
	switch num {
	case 0:
		return newTest(num, `simple 32767 conversion test`, toTest(ToUint16(32767)), ToUint16Value(32767), uint16(32767), nil)
	case 1:
		return newTest(num, `simple -32768 conversion test`, toTest(ToUint16(-32768)), ToUint16Value(-32768), uint16(0), errors.New(`uint16 cannot be negative`))
	case 2:
		maper := map[int]int{1: 12, 3: 14}
		return newTest(num, `simple map[int]int conversion test`, toTest(ToUint16(maper)), ToUint16Value(maper), uint16(0), errors.New(`map[int]int format cannot be converted to uint16`))
	case 3:
		maper := map[int]int{1: 12, 3: 14}
		return newTest(num, `pointer map[int]int conversion test`, toTest(ToUint16(&maper)), ToUint16Value(&maper), uint16(0), errors.New(`map[int]int format cannot be converted to uint16`))

	}

	return testStructure{}, false
}

func TestToUint16(t *testing.T) {
	testByLists(t, []testListStructure{
		{
			Name:           `Personal tests`,
			FuncionGetList: getElementForTestUint16,
		},
	})
}

func getElementForTestUint32(num int) (testStructure, bool) {
	switch num {
	case 0:
		return newTest(num, `simple 2147483647 conversion test`, toTest(ToUint32(2147483647)), ToUint32Value(2147483647), uint32(2147483647), nil)
	case 1:
		return newTest(num, `simple -2147483648 conversion test`, toTest(ToUint32(-2147483648)), ToUint32Value(-2147483648), uint32(0), errors.New(`uint32 cannot be negative`))
	case 4:
		var f float64 = math.Pi + float64(987654318.0)
		return newTest(num, `simple pi+ conversion test`, toTest(ToUint32(f)), ToUint32Value(f), uint32(f), nil)
	case 5:
		return newTest(num, `simple []int{} conversion test`, toTest(ToUint32([]int{})), ToUint32Value([]int{}), uint32(0), errors.New(`[]int format cannot be converted to uint32`))
	case 6:
		arrayBool := [2]string{"true", "false"}
		return newTest(num, `simple [2]string{} conversion test`, toTest(ToUint32(arrayBool)), ToUint32Value(arrayBool), uint32(0), errors.New(`[2]string format cannot be converted to uint32`))
	case 7:
		arrayBool := [2]string{"true", "false"}
		return newTest(num, `pointer [2]string{} conversion test`, toTest(ToUint32(arrayBool)), ToUint32Value(arrayBool), uint32(0), errors.New(`[2]string format cannot be converted to uint32`))
	}

	return testStructure{}, false
}

func TestToUint32(t *testing.T) {
	testByLists(t, []testListStructure{
		{
			Name:           `Personal tests`,
			FuncionGetList: getElementForTestUint32,
		},
	})
}

func getElementForTestUint64(num int) (testStructure, bool) {
	switch num {
	case 0:
		return newTest(num, `simple 9223372036854775807 conversion test`, toTest(ToUint64(9223372036854775807)), ToUint64Value(9223372036854775807), uint64(9223372036854775807), nil)
	case 1:
		return newTest(num, `simple -9223372036854775808 conversion test`, toTest(ToUint64(-9223372036854775808)), ToUint64Value(-9223372036854775808), uint64(0), errors.New(`uint64 cannot be negative`))
	case 2:
		return newTest(num, `simple float32(-2.544) conversion test`, toTest(ToUint64(float32(-2.544))), ToUint64Value(float32(-2.544)), uint64(0), errors.New(`uint64 cannot be negative`))
	case 3:
		return newTest(num, `simple string -2.544 conversion test`, toTest(ToUint64(`-2.544`)), ToUint64Value(`-2.544`), uint64(0), errors.New(`uint64 cannot be negative`))
	case 4:
		return newTest(num, `simple string -2 conversion test`, toTest(ToUint64(`-2`)), ToUint64Value(`-2`), uint64(0), errors.New(`uint64 cannot be negative`))
	case 5:
		var f float64 = math.Pi + float64(987654318.0)
		return newTest(num, `simple pi+ conversion test`, toTest(ToUint64(f)), ToUint64Value(f), uint64(f), nil)
	case 6, 7:
		type structure struct {
			Val int `json:"val"`
		}
		structureItem := structure{Val: 43}
		if 6 == num {
			return newTest(num, `simple structure conversion test`, toTest(ToUint64(structureItem)), ToUint64Value(structureItem), uint64(0), errors.New(`typo.structure format cannot be converted to uint64`))
		} else {
			return newTest(num, `pointer structure conversion test`, toTest(ToUint64(&structureItem)), ToUint64Value(&structureItem), uint64(0), errors.New(`typo.structure format cannot be converted to uint64`))
		}
	case 8:
		return newTest(num, `simple []string{} conversion test`, toTest(ToUint64([]string{})), ToUint64Value([]string{}), uint64(0), errors.New(`[]string format cannot be converted to uint64`))
	case 9:
		return newTest(num, `simple []interface{}{1} conversion test`, toTest(ToUint64([]interface{}{1})), ToUint64Value([]interface{}{1}), uint64(0), errors.New(`[]interface {} format cannot be converted to uint64`))
	}

	return testStructure{}, false
}

func TestToUint64(t *testing.T) {
	testByLists(t, []testListStructure{
		{
			Name:           `Personal tests`,
			FuncionGetList: getElementForTestUint64,
		},
	})
}

func getElementForTestUintptr(num int) (testStructure, bool) {
	switch num {
	case 0:
		return newTest(num, `simple 9223372036854775807 conversion test`, toTest(ToUintptr(9223372036854775807)), ToUintptrValue(9223372036854775807), uintptr(9223372036854775807), nil)
	case 1:
		return newTest(num, `simple -9223372036854775808 conversion test`, toTest(ToUintptr(-9223372036854775808)), ToUintptrValue(-9223372036854775808), uintptr(0), errors.New(`uintptr cannot be negative`))
	case 2:
		return newTest(num, `simple float32(-2.544) conversion test`, toTest(ToUintptr(float32(-2.544))), ToUintptrValue(float32(-2.544)), uintptr(0), errors.New(`uintptr cannot be negative`))
	case 3:
		return newTest(num, `simple string -2.544 conversion test`, toTest(ToUintptr(`-2.544`)), ToUintptrValue(`-2.544`), uintptr(0), errors.New(`uintptr cannot be negative`))
	case 4:
		return newTest(num, `simple string -2 conversion test`, toTest(ToUintptr(`-2`)), ToUintptrValue(`-2`), uintptr(0), errors.New(`uintptr cannot be negative`))
	case 5:
		var f float64 = math.Pi + float64(987654318.0)
		return newTest(num, `simple pi+ conversion test`, toTest(ToUintptr(f)), ToUintptrValue(f), uintptr(f), nil)
	case 6, 7:
		type structure struct {
			Val int `json:"val"`
		}
		structureItem := structure{Val: 43}
		if 6 == num {
			return newTest(num, `simple structure conversion test`, toTest(ToUintptr(structureItem)), ToUintptrValue(structureItem), uintptr(0), errors.New(`typo.structure format cannot be converted to uintptr`))
		} else {
			return newTest(num, `pointer structure conversion test`, toTest(ToUintptr(&structureItem)), ToUintptrValue(&structureItem), uintptr(0), errors.New(`typo.structure format cannot be converted to uintptr`))
		}
	case 8:
		return newTest(num, `simple []string{} conversion test`, toTest(ToUintptr([]string{})), ToUintptrValue([]string{}), uintptr(0), errors.New(`[]string format cannot be converted to uintptr`))
	case 9:
		return newTest(num, `simple []interface{}{1} conversion test`, toTest(ToUintptr([]interface{}{1})), ToUintptrValue([]interface{}{1}), uintptr(0), errors.New(`[]interface {} format cannot be converted to uintptr`))
	case 10:
		return newTest(num, `simple float32(2.544) conversion test`, toTest(ToUintptr(float32(2.544))), ToUintptrValue(float32(2.544)), uintptr(2), nil)
	case 11:
		return newTest(num, `simple float64(2.544) conversion test`, toTest(ToUintptr(float64(2.544))), ToUintptrValue(float64(2.544)), uintptr(2), nil)
	case 12:
		return newTest(num, `simple string 2.544 conversion test`, toTest(ToUintptr(`2.544`)), ToUintptrValue(`2.544`), uintptr(2), nil)
	case 13:
		return newTest(num, `simple string 2 conversion test`, toTest(ToUintptr(`2`)), ToUintptrValue(`2`), uintptr(2), nil)
	case 14:
		return newTest(num, `simple string - conversion test`, toTest(ToUintptr(`-`)), ToUintptrValue(`-`), uintptr(0), errors.New(`strconv.ParseFloat: parsing "-": invalid syntax`))
	case 15:
		return newTest(num, `simple string + conversion test`, toTest(ToUintptr(`+`)), ToUintptrValue(`+`), uintptr(0), errors.New(`strconv.ParseFloat: parsing "+": invalid syntax`))
	case 16:
		return newTest(num, `simple string / conversion test`, toTest(ToUintptr(`/`)), ToUintptrValue(`/`), uintptr(0), errors.New(`strconv.ParseFloat: parsing "/": invalid syntax`))
	}

	return testStructure{}, false
}

func TestToUintptr(t *testing.T) {
	testByLists(t, []testListStructure{
		{
			Name:           `Personal tests`,
			FuncionGetList: getElementForTestUintptr,
		},
	})
}
