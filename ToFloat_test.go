package typo

import (
	"errors"
	"math"
	"testing"
)

func getElementForTestFloat(num int) (testStructure, bool) {
	switch num {
	case 0:
		return newTest(num, `simple integer conversion test`, toTest(ToFloat64(1)), ToFloat64Value(1), float64(1), nil)
	case 1:
		return newTest(num, `simple float conversion test`, toTest(ToFloat64(1.5)), ToFloat64Value(1.5), float64(1.5), nil)
	case 2:
		var i int = 2
		return newTest(num, `simple integer conversion test`, toTest(ToFloat64(i)), ToFloat64Value(i), float64(i), nil)
	case 3:
		var i int = 2
		return newTest(num, `pointer integer conversion test`, toTest(ToFloat64(&i)), ToFloat64Value(&i), float64(i), nil)
	case 4:
		var i int8 = 8
		return newTest(num, `simple integer8 conversion test`, toTest(ToFloat64(i)), ToFloat64Value(i), float64(i), nil)
	case 5:
		var i int8 = 8
		return newTest(num, `pointer integer8 conversion test`, toTest(ToFloat64(&i)), ToFloat64Value(&i), float64(i), nil)
	case 6:
		var i int16 = 16
		return newTest(num, `simple integer16 conversion test`, toTest(ToFloat64(i)), ToFloat64Value(i), float64(i), nil)
	case 7:
		var i int16 = 16
		return newTest(num, `pointer integer16 conversion test`, toTest(ToFloat64(&i)), ToFloat64Value(&i), float64(i), nil)
	case 8:
		var i int32 = 32
		return newTest(num, `simple integer32 conversion test`, toTest(ToFloat64(i)), ToFloat64Value(i), float64(i), nil)
	case 9:
		var i int32 = 32
		return newTest(num, `pointer integer32 conversion test`, toTest(ToFloat64(&i)), ToFloat64Value(&i), float64(i), nil)
	case 10:
		var i int64 = 64
		return newTest(num, `simple integer64 conversion test`, toTest(ToFloat64(i)), ToFloat64Value(i), float64(i), nil)
	case 11:
		var i int64 = 64
		return newTest(num, `pointer integer64 conversion test`, toTest(ToFloat64(&i)), ToFloat64Value(&i), float64(i), nil)
	case 12:
		var ui uint = 51
		return newTest(num, `simple uint conversion test`, toTest(ToFloat64(ui)), ToFloat64Value(ui), float64(ui), nil)
	case 13:
		var ui uint = 51
		return newTest(num, `pointer uint conversion test`, toTest(ToFloat64(&ui)), ToFloat64Value(&ui), float64(ui), nil)
	case 14:
		var ui uint8 = 58
		return newTest(num, `simple uint8 conversion test`, toTest(ToFloat64(ui)), ToFloat64Value(ui), float64(ui), nil)
	case 15:
		var ui uint8 = 58
		return newTest(num, `pointer uint8 conversion test`, toTest(ToFloat64(&ui)), ToFloat64Value(&ui), float64(ui), nil)
	case 16:
		var ui uint16 = 116
		return newTest(num, `simple uint16 conversion test`, toTest(ToFloat64(ui)), ToFloat64Value(ui), float64(ui), nil)
	case 17:
		var ui uint16 = 116
		return newTest(num, `pointer uint16 conversion test`, toTest(ToFloat64(&ui)), ToFloat64Value(&ui), float64(ui), nil)
	case 18:
		var ui uint32 = 102
		return newTest(num, `simple uint32 conversion test`, toTest(ToFloat64(ui)), ToFloat64Value(ui), float64(ui), nil)
	case 19:
		var ui uint32 = 102
		return newTest(num, `pointer uint32 conversion test`, toTest(ToFloat64(&ui)), ToFloat64Value(&ui), float64(ui), nil)
	case 20:
		var ui uint64 = 104
		return newTest(num, `simple uint64 conversion test`, toTest(ToFloat64(ui)), ToFloat64Value(ui), float64(ui), nil)
	case 21:
		var ui uint64 = 104
		return newTest(num, `pointer uint64 conversion test`, toTest(ToFloat64(&ui)), ToFloat64Value(&ui), float64(ui), nil)
	case 22:
		var ui uintptr = 123
		return newTest(num, `simple uintptr conversion test`, toTest(ToFloat64(ui)), ToFloat64Value(ui), float64(ui), nil)
	case 23:
		var ui uintptr = 123
		return newTest(num, `pointer uintptr conversion test`, toTest(ToFloat64(&ui)), ToFloat64Value(&ui), float64(ui), nil)
	case 24:
		var f float32 = math.Pi
		return newTest(num, `simple float32 conversion test`, toTest(ToFloat32(f)), ToFloat32Value(f), float32(f), nil)
	case 25:
		var f float32 = math.Pi
		return newTest(num, `pointer float32 conversion test`, toTest(ToFloat32(&f)), ToFloat32Value(&f), float32(f), nil)
	case 26:
		var elemNull *uint
		return newTest(num, `pointer uint = nil conversion test`, toTest(ToFloat64(elemNull)), ToFloat64Value(elemNull), float64(0), nil)
	case 27:
		var elemNull *uint
		return newTest(num, `pointer pointer uint = nil conversion test`, toTest(ToFloat32(&elemNull)), ToFloat32Value(&elemNull), float32(0), nil)
	case 28:
		b := true
		return newTest(num, `simple bool = true conversion test`, toTest(ToFloat64(b)), ToFloat64Value(b), float64(1), nil)
	case 29:
		b := true
		return newTest(num, `pointer bool = true conversion test`, toTest(ToFloat64(&b)), ToFloat64Value(&b), float64(1), nil)
	case 30:
		b := false
		return newTest(num, `simple bool = false conversion test`, toTest(ToFloat64(b)), ToFloat64Value(b), float64(0), nil)
	case 31:
		b := false
		return newTest(num, `pointer bool = false conversion test`, toTest(ToFloat64(&b)), ToFloat64Value(&b), float64(0), nil)
	case 32:
		return newTest(num, `pointer nil conversion test`, toTest(ToFloat64(nil)), ToFloat64Value(nil), float64(0), nil)

	}

	return testStructure{}, false
}

func getElementForTestFloat32(num int) (testStructure, bool) {
	switch num {
	case 0:
		i := int32(-2147483648)
		return newTest(num, `simple -2147483648 conversion test`, toTest(ToFloat32(i)), ToFloat32Value(i), float32(i), nil)
	case 1:
		return newTest(num, `simple []int{} conversion test`, toTest(ToFloat32([]int{})), ToFloat32Value([]int{}), float32(0), errors.New(`[]int format cannot be converted to float32`))
	case 2:
		arrayString := [2]string{"true", "false"}
		return newTest(num, `simple [2]string{"true", "false"} conversion test`, toTest(ToFloat32(arrayString)), ToFloat32Value(arrayString), float32(0), errors.New(`[2]string format cannot be converted to float32`))
	case 3:
		arrayString := [2]string{"true", "false"}
		return newTest(num, `pointer [2]string{"true", "false"} conversion test`, toTest(ToFloat32(&arrayString)), ToFloat32Value(&arrayString), float32(0), errors.New(`[2]string format cannot be converted to float32`))
	case 4:
		i := int64(9223372036854775807)
		return newTest(num, `simple 9223372036854775807 conversion test`, toTest(ToFloat32(i)), ToFloat32Value(i), float32(9223372036854775807), nil)
	case 5:
		i := int64(-9223372036854775808)
		return newTest(num, `simple -9223372036854775808 conversion test`, toTest(ToFloat32(i)), ToFloat32Value(i), float32(-9223372036854775808), nil)
	case 6, 7:
		type structure struct {
			Val int `json:"val"`
		}
		structureItem := structure{Val: 43}
		if num == 6 {
			return newTest(num, `simple structure conversion test`, toTest(ToFloat32(structureItem)), ToFloat32Value(structureItem), float32(0), errors.New(`typo.structure format cannot be converted to float32`))
		} else {
			return newTest(num, `pointer structure conversion test`, toTest(ToFloat32(&structureItem)), ToFloat32Value(&structureItem), float32(0), errors.New(`typo.structure format cannot be converted to float32`))
		}
	case 8:
		return newTest(num, `simple []string{} conversion test`, toTest(ToFloat32([]string{})), ToFloat32Value([]string{}), float32(0), errors.New(`[]string format cannot be converted to float32`))
	case 9:
		return newTest(num, `simple []interface{}{1} conversion test`, toTest(ToFloat32([]interface{}{1})), ToFloat32Value([]interface{}{1}), float32(0), errors.New(`[]interface {} format cannot be converted to float32`))
	case 10:
		s := `-2`
		return newTest(num, `simple string conversion test`, toTest(ToFloat32(s)), ToFloat32Value(s), float32(-2), nil)
	case 11:
		s := `-`
		return newTest(num, `simple string = - conversion test`, toTest(ToFloat32(s)), ToFloat32Value(s), float32(0), errors.New(`strconv.ParseFloat: parsing "-": invalid syntax`))
	case 12:
		s := `+`
		return newTest(num, `simple string = + conversion test`, toTest(ToFloat32(s)), ToFloat32Value(s), float32(0), errors.New(`strconv.ParseFloat: parsing "+": invalid syntax`))

	}

	return testStructure{}, false
}

func getElementForTestFloat64(num int) (testStructure, bool) {
	switch num {
	case 0:
		i := int(2147483647)
		return newTest(num, `simple 2147483647 conversion test`, toTest(ToFloat64(i)), ToFloat64Value(i), float64(i), nil)
	case 1:
		i := int(-2147483648)
		return newTest(num, `simple -2147483648 conversion test`, toTest(ToFloat64(i)), ToFloat64Value(i), float64(i), nil)
	case 2:
		slice := []int{4, 3}
		return newTest(num, `simple []int{4, 3} conversion test`, toTest(ToFloat64(slice)), ToFloat64Value(slice), float64(0), errors.New(`[]int format cannot be converted to float64`))
	case 3:
		slice := []int{4, 3}
		return newTest(num, `pointer []int{4, 3} conversion test`, toTest(ToFloat64(&slice)), ToFloat64Value(&slice), float64(0), errors.New(`[]int format cannot be converted to float64`))
	case 4:
		i := int8(127)
		return newTest(num, `simple int8 = 127 conversion test`, toTest(ToFloat64(i)), ToFloat64Value(i), float64(i), nil)
	case 5:
		i := int8(-128)
		return newTest(num, `simple int8 = -128 conversion test`, toTest(ToFloat64(i)), ToFloat64Value(i), float64(i), nil)
	case 6:
		var f float64 = math.Pi + float64(987654318.0)
		return newTest(num, `simple float pi+ conversion test`, toTest(ToFloat64(f)), ToFloat64Value(f), float64(f), nil)
	case 7:
		arrayBool := [2]bool{true, false}
		return newTest(num, `simple [2]bool conversion test`, toTest(ToFloat64(arrayBool)), ToFloat64Value(arrayBool), float64(0), errors.New(`[2]bool format cannot be converted to float64`))
	case 8:
		arrayBool := [2]bool{true, false}
		return newTest(num, `pointer [2]bool conversion test`, toTest(ToFloat64(&arrayBool)), ToFloat64Value(&arrayBool), float64(0), errors.New(`[2]bool format cannot be converted to float64`))
	case 9:
		i := int16(32767)
		return newTest(num, `simple int16 = 32767 conversion test`, toTest(ToFloat64(i)), ToFloat64Value(i), float64(i), nil)
	case 10:
		i := int16(-32768)
		return newTest(num, `simple int16 = -32768 conversion test`, toTest(ToFloat64(i)), ToFloat64Value(i), float64(i), nil)
	case 11:
		mapElem := map[int]int{1: 12, 3: 14}
		return newTest(num, `simple map[int]int conversion test`, toTest(ToFloat64(mapElem)), ToFloat64Value(mapElem), float64(0), errors.New(`map[int]int format cannot be converted to float64`))
	case 12:
		mapElem := map[int]int{1: 12, 3: 14}
		return newTest(num, `pointer map[int]int conversion test`, toTest(ToFloat64(&mapElem)), ToFloat64Value(&mapElem), float64(0), errors.New(`map[int]int format cannot be converted to float64`))
	case 13:
		s := `/`
		return newTest(num, `simple string = / conversion test`, toTest(ToFloat64(s)), ToFloat64Value(s), float64(0), errors.New(`strconv.ParseFloat: parsing "/": invalid syntax`))
	case 14:
		s := `2.544`
		return newTest(num, `simple string = 2.544 conversion test`, toTest(ToFloat64(s)), ToFloat64Value(s), float64(2.544), nil)
	case 15:
		s := `2`
		return newTest(num, `simple string = 2 conversion test`, toTest(ToFloat64(s)), ToFloat64Value(s), float64(2), nil)

	}

	return testStructure{}, false
}

func TestToFloat32(t *testing.T) {
	testByLists(t, []testListStructure{
		{
			Name:           `General tests`,
			FuncionGetList: getElementForTestFloat,
		},
		{
			Name:           `Personal tests`,
			FuncionGetList: getElementForTestFloat32,
		},
	})
}

func TestToFloat64(t *testing.T) {
	testByLists(t, []testListStructure{
		{
			Name:           `General tests`,
			FuncionGetList: getElementForTestFloat,
		},
		{
			Name:           `Personal tests`,
			FuncionGetList: getElementForTestFloat64,
		},
	})
}
