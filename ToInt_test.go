package typo

import (
	"errors"
	"math"
	"testing"
)

func getElementForTestInt(num int) (testStructure, bool) {
	switch num {
	case 0:
		return newTest(num, `simple integer conversion test`, toTest(ToInt(1)), ToIntValue(1), int(1), nil)
	case 1:
		var i int = 2
		return newTest(num, `simple integer conversion test`, toTest(ToInt(i)), ToIntValue(i), int(i), nil)
	case 2:
		var i int = 2
		return newTest(num, `point integer conversion test`, toTest(ToInt(&i)), ToIntValue(&i), int(i), nil)
	case 3:
		var i int8 = 8
		return newTest(num, `simple integer8 conversion test`, toTest(ToInt(i)), ToIntValue(i), int(i), nil)
	case 4:
		var i int8 = 8
		return newTest(num, `point integer8 conversion test`, toTest(ToInt(&i)), ToIntValue(&i), int(i), nil)
	case 5:
		var i int16 = 16
		return newTest(num, `simple integer16 conversion test`, toTest(ToInt(i)), ToIntValue(i), int(i), nil)
	case 6:
		var i int16 = 16
		return newTest(num, `point integer16 conversion test`, toTest(ToInt(&i)), ToIntValue(&i), int(i), nil)
	case 7:
		var i int32 = 32
		return newTest(num, `simple integer32 conversion test`, toTest(ToInt(i)), ToIntValue(i), int(i), nil)
	case 8:
		var i int32 = 32
		return newTest(num, `point integer32 conversion test`, toTest(ToInt(&i)), ToIntValue(&i), int(i), nil)
	case 9:
		var i int64 = 64
		return newTest(num, `simple integer64 conversion test`, toTest(ToInt(i)), ToIntValue(i), int(i), nil)
	case 10:
		var i int64 = 64
		return newTest(num, `point integer64 conversion test`, toTest(ToInt(&i)), ToIntValue(&i), int(i), nil)
	case 11:
		var ui uint = 51
		return newTest(num, `simple uint conversion test`, toTest(ToInt(ui)), ToIntValue(ui), int(ui), nil)
	case 12:
		var ui uint = 51
		return newTest(num, `point uint conversion test`, toTest(ToInt(&ui)), ToIntValue(&ui), int(ui), nil)
	case 13:
		var ui uint8 = 58
		return newTest(num, `simple uint8 conversion test`, toTest(ToInt(ui)), ToIntValue(ui), int(ui), nil)
	case 14:
		var ui uint8 = 58
		return newTest(num, `point uint8 conversion test`, toTest(ToInt(&ui)), ToIntValue(&ui), int(ui), nil)
	case 15:
		var ui uint16 = 116
		return newTest(num, `simple uint16 conversion test`, toTest(ToInt(ui)), ToIntValue(ui), int(ui), nil)
	case 16:
		var ui uint16 = 116
		return newTest(num, `point uint16 conversion test`, toTest(ToInt(&ui)), ToIntValue(&ui), int(ui), nil)
	case 17:
		var ui uint32 = 102
		return newTest(num, `simple uint32 conversion test`, toTest(ToInt(ui)), ToIntValue(ui), int(ui), nil)
	case 18:
		var ui uint32 = 102
		return newTest(num, `point uint32 conversion test`, toTest(ToInt(&ui)), ToIntValue(&ui), int(ui), nil)
	case 19:
		var ui uint64 = 104
		return newTest(num, `simple uint64 conversion test`, toTest(ToInt(ui)), ToIntValue(ui), int(ui), nil)
	case 20:
		var ui uint64 = 104
		return newTest(num, `point uint64 conversion test`, toTest(ToInt(&ui)), ToIntValue(&ui), int(ui), nil)
	case 21:
		var ui uintptr = 123
		return newTest(num, `simple uintptr conversion test`, toTest(ToInt(ui)), ToIntValue(ui), int(ui), nil)
	case 22:
		var ui uintptr = 123
		return newTest(num, `point uintptr conversion test`, toTest(ToInt(&ui)), ToIntValue(&ui), int(ui), nil)
	case 23:
		var f float32 = math.Pi
		return newTest(num, `simple float32 conversion test`, toTest(ToInt(f)), ToIntValue(f), int(f), nil)
	case 24:
		var f float32 = math.Pi
		return newTest(num, `pointer float32 conversion test`, toTest(ToInt(&f)), ToIntValue(&f), int(f), nil)
	case 25:
		var f float64 = math.Pi
		return newTest(num, `simple float64 conversion test`, toTest(ToInt(f)), ToIntValue(f), int(f), nil)
	case 26:
		var f float64 = math.Pi
		return newTest(num, `pointer float64 conversion test`, toTest(ToInt(&f)), ToIntValue(&f), int(f), nil)
	case 27:
		var elemNull *uint
		return newTest(num, `pointer pointer uint = nil conversion test`, toTest(ToInt(&elemNull)), ToIntValue(&elemNull), int(0), nil)
	case 28:
		b := true
		return newTest(num, `simple bool = true conversion test`, toTest(ToInt(b)), ToIntValue(b), int(1), nil)
	case 29:
		b := true
		return newTest(num, `pointer bool = true conversion test`, toTest(ToInt(&b)), ToIntValue(&b), int(1), nil)
	case 30:
		b := false
		return newTest(num, `simple bool = false conversion test`, toTest(ToInt(b)), ToIntValue(b), int(0), nil)
	case 31:
		b := false
		return newTest(num, `pointer bool = false conversion test`, toTest(ToInt(&b)), ToIntValue(&b), int(0), nil)
	case 32:
		return newTest(num, `pointer nil conversion test`, toTest(ToInt(nil)), ToIntValue(nil), int(0), nil)

	}

	return testStructure{}, false
}

func getElementForTestIntDefault(num int) (testStructure, bool) {
	switch num {
	case 0:
		return newTest(num, `simple 2147483647 conversion test`, toTest(ToInt(2147483647)), ToIntValue(2147483647), int(2147483647), nil)
	case 1:
		return newTest(num, `simple -2147483648 conversion test`, toTest(ToInt(-2147483648)), ToIntValue(-2147483648), int(-2147483648), nil)
	case 2:
		slice := []int{4, 3}
		return newTest(num, `simple []int{4, 3} conversion test`, toTest(ToInt(slice)), ToIntValue(slice), int(0), errors.New(`[]int format cannot be converted to int`))
	case 3:
		slice := []int{4, 3}
		return newTest(num, `pointer []int{4, 3} conversion test`, toTest(ToInt(slice)), ToIntValue(slice), int(0), errors.New(`[]int format cannot be converted to int`))

	}

	return testStructure{}, false
}

func TestToInt(t *testing.T) {
	testByLists(t, []testListStructure{
		{
			Name:           `General tests`,
			FuncionGetList: getElementForTestInt,
		},
		{
			Name:           `Personal tests`,
			FuncionGetList: getElementForTestIntDefault,
		},
	})
}

func getElementForTestInt8(num int) (testStructure, bool) {
	switch num {
	case 0:
		return newTest(num, `simple 127 conversion test`, toTest(ToInt8(127)), ToInt8Value(127), int8(127), nil)
	case 1:
		return newTest(num, `simple -128 conversion test`, toTest(ToInt8(-128)), ToInt8Value(-128), int8(-128), nil)
	case 2:
		arrayBool := [2]bool{true, false}
		return newTest(num, `simple []int{4, 3} conversion test`, toTest(ToInt8(arrayBool)), ToInt8Value(arrayBool), int8(0), errors.New(`[2]bool format cannot be converted to int8`))
	case 3:
		arrayBool := [2]bool{true, false}
		return newTest(num, `pointer []int{4, 3} conversion test`, toTest(ToInt8(arrayBool)), ToInt8Value(arrayBool), int8(0), errors.New(`[2]bool format cannot be converted to int8`))
	case 4:
		var f float64 = math.Pi + float64(987654318.0)
		return newTest(num, `simple pi+ conversion test`, toTest(ToInt8(f)), ToInt8Value(f), int8(f), nil)

	}

	return testStructure{}, false
}

func TestToInt8(t *testing.T) {
	testByLists(t, []testListStructure{
		{
			Name:           `Personal tests`,
			FuncionGetList: getElementForTestInt8,
		},
	})
}

func getElementForTestInt16(num int) (testStructure, bool) {
	switch num {
	case 0:
		return newTest(num, `simple 32767 conversion test`, toTest(ToInt16(32767)), ToInt16Value(32767), int16(32767), nil)
	case 1:
		return newTest(num, `simple -32768 conversion test`, toTest(ToInt16(-32768)), ToInt16Value(-32768), int16(-32768), nil)
	case 2:
		maper := map[int]int{1: 12, 3: 14}
		return newTest(num, `simple map[int]int conversion test`, toTest(ToInt16(maper)), ToInt16Value(maper), int16(0), errors.New(`map[int]int format cannot be converted to int16`))
	case 3:
		maper := map[int]int{1: 12, 3: 14}
		return newTest(num, `pointer map[int]int conversion test`, toTest(ToInt16(&maper)), ToInt16Value(&maper), int16(0), errors.New(`map[int]int format cannot be converted to int16`))

	}

	return testStructure{}, false
}

func TestToInt16(t *testing.T) {
	testByLists(t, []testListStructure{
		{
			Name:           `Personal tests`,
			FuncionGetList: getElementForTestInt16,
		},
	})
}

func getElementForTestInt32(num int) (testStructure, bool) {
	switch num {
	case 0:
		return newTest(num, `simple 2147483647 conversion test`, toTest(ToInt32(2147483647)), ToInt32Value(2147483647), int32(2147483647), nil)
	case 1:
		return newTest(num, `simple -2147483648 conversion test`, toTest(ToInt32(-2147483648)), ToInt32Value(-2147483648), int32(-2147483648), nil)
	case 4:
		var f float64 = math.Pi + float64(987654318.0)
		return newTest(num, `simple pi+ conversion test`, toTest(ToInt32(f)), ToInt32Value(f), int32(f), nil)
	case 5:
		return newTest(num, `simple []int{} conversion test`, toTest(ToInt32([]int{})), ToInt32Value([]int{}), int32(0), errors.New(`[]int format cannot be converted to int32`))
	case 6:
		arrayBool := [2]string{"true", "false"}
		return newTest(num, `simple [2]string{} conversion test`, toTest(ToInt32(arrayBool)), ToInt32Value(arrayBool), int32(0), errors.New(`[2]string format cannot be converted to int32`))
	case 7:
		arrayBool := [2]string{"true", "false"}
		return newTest(num, `pointer [2]string{} conversion test`, toTest(ToInt32(arrayBool)), ToInt32Value(arrayBool), int32(0), errors.New(`[2]string format cannot be converted to int32`))
	}

	return testStructure{}, false
}

func TestToInt32(t *testing.T) {
	testByLists(t, []testListStructure{
		{
			Name:           `Personal tests`,
			FuncionGetList: getElementForTestInt32,
		},
	})
}

func getElementForTestInt64(num int) (testStructure, bool) {
	switch num {
	case 0:
		return newTest(num, `simple 9223372036854775807 conversion test`, toTest(ToInt64(9223372036854775807)), ToInt64Value(9223372036854775807), int64(9223372036854775807), nil)
	case 1:
		return newTest(num, `simple -9223372036854775808 conversion test`, toTest(ToInt64(-9223372036854775808)), ToInt64Value(-9223372036854775808), int64(-9223372036854775808), nil)
	case 2:
		return newTest(num, `simple float32(-2.544) conversion test`, toTest(ToInt64(float32(-2.544))), ToInt64Value(float32(-2.544)), int64(-2), nil)
	case 3:
		return newTest(num, `simple string -2.544 conversion test`, toTest(ToInt64(`-2.544`)), ToInt64Value(`-2.544`), int64(-2), nil)
	case 4:
		return newTest(num, `simple string -2 conversion test`, toTest(ToInt64(`-2`)), ToInt64Value(`-2`), int64(-2), nil)
	case 5:
		var f float64 = math.Pi + float64(987654318.0)
		return newTest(num, `simple pi+ conversion test`, toTest(ToInt64(f)), ToInt64Value(f), int64(f), nil)
	case 6, 7:
		type structure struct {
			Val int `json:"val"`
		}
		structureItem := structure{Val: 43}
		if 6 == num {
			return newTest(num, `simple structure conversion test`, toTest(ToInt64(structureItem)), ToInt64Value(structureItem), int64(0), errors.New(`typo.structure format cannot be converted to int64`))
		} else {
			return newTest(num, `pointer structure conversion test`, toTest(ToInt64(&structureItem)), ToInt64Value(&structureItem), int64(0), errors.New(`typo.structure format cannot be converted to int64`))
		}
	case 8:
		return newTest(num, `simple []string{} conversion test`, toTest(ToInt64([]string{})), ToInt64Value([]string{}), int64(0), errors.New(`[]string format cannot be converted to int64`))
	case 9:
		return newTest(num, `simple []interface{}{1} conversion test`, toTest(ToInt64([]interface{}{1})), ToInt64Value([]interface{}{1}), int64(0), errors.New(`[]interface {} format cannot be converted to int64`))
	}

	return testStructure{}, false
}

func TestToInt64(t *testing.T) {
	testByLists(t, []testListStructure{
		{
			Name:           `Personal tests`,
			FuncionGetList: getElementForTestInt64,
		},
	})
}
