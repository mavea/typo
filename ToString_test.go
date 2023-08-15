package typo

import (
	"math"
	"testing"
)

func getElementForTestString(num int) (testStructure, bool) {
	switch num {
	case 0:
		return newTest(num, `simple true conversion test`, toTest(ToString(true)), ToStringValue(true), `true`, nil)
	case 1:
		return newTest(num, `simple false conversion test`, toTest(ToString(false)), ToStringValue(false), `false`, nil)
	case 2:
		return newTest(num, `simple int 1 conversion test`, toTest(ToString(1)), ToStringValue(1), `1`, nil)
	case 3:
		return newTest(num, `simple int 0 conversion test`, toTest(ToString(0)), ToStringValue(0), `0`, nil)
	case 4:
		return newTest(num, `simple int -1 conversion test`, toTest(ToString(-1)), ToStringValue(-1), `-1`, nil)
	case 5:
		var i int = 2
		return newTest(num, `simple int variable conversion test`, toTest(ToString(i)), ToStringValue(i), `2`, nil)
	case 6:
		var i int = 2
		return newTest(num, `pointer int conversion test`, toTest(ToString(&i)), ToStringValue(&i), `2`, nil)
	case 7:
		var i int8 = 8
		return newTest(num, `simple int8 conversion test`, toTest(ToString(i)), ToStringValue(i), `8`, nil)
	case 8:
		var i int8 = 8
		return newTest(num, `pointer int8 conversion test`, toTest(ToString(&i)), ToStringValue(&i), `8`, nil)
	case 9:
		var i int16 = 16
		return newTest(num, `simple int16 conversion test`, toTest(ToString(i)), ToStringValue(i), `16`, nil)
	case 10:
		var i int16 = 16
		return newTest(num, `pointer int16 conversion test`, toTest(ToString(&i)), ToStringValue(&i), `16`, nil)
	case 11:
		var i int32 = 32
		return newTest(num, `simple int32 conversion test`, toTest(ToString(i)), ToStringValue(i), `32`, nil)
	case 12:
		var i int32 = 32
		return newTest(num, `pointer int32 conversion test`, toTest(ToString(&i)), ToStringValue(&i), `32`, nil)
	case 13:
		var i int64 = 64
		return newTest(num, `simple int64 conversion test`, toTest(ToString(i)), ToStringValue(i), `64`, nil)
	case 14:
		var i int64 = 64
		return newTest(num, `pointer int64 conversion test`, toTest(ToString(&i)), ToStringValue(&i), `64`, nil)
	case 15:
		var ui uint = 51
		return newTest(num, `simple uint conversion test`, toTest(ToString(ui)), ToStringValue(ui), `51`, nil)
	case 16:
		var ui uint = 51
		return newTest(num, `pointer uint conversion test`, toTest(ToString(&ui)), ToStringValue(&ui), `51`, nil)
	case 17:
		var ui uint8 = 58
		return newTest(num, `simple uint8 conversion test`, toTest(ToString(ui)), ToStringValue(ui), `58`, nil)
	case 18:
		var ui uint8 = 58
		return newTest(num, `pointer uint8 conversion test`, toTest(ToString(&ui)), ToStringValue(&ui), `58`, nil)
	case 19:
		var ui uint16 = 516
		return newTest(num, `simple uint16 conversion test`, toTest(ToString(ui)), ToStringValue(ui), `516`, nil)
	case 20:
		var ui uint16 = 516
		return newTest(num, `pointer uint16 conversion test`, toTest(ToString(&ui)), ToStringValue(&ui), `516`, nil)
	case 21:
		var ui uint32 = 532
		return newTest(num, `simple uint32 conversion test`, toTest(ToString(ui)), ToStringValue(ui), `532`, nil)
	case 22:
		var ui uint32 = 532
		return newTest(num, `pointer uint32 conversion test`, toTest(ToString(&ui)), ToStringValue(&ui), `532`, nil)
	case 23:
		var ui uint64 = 564
		return newTest(num, `simple uint64 conversion test`, toTest(ToString(ui)), ToStringValue(ui), `564`, nil)
	case 24:
		var ui uint64 = 564
		return newTest(num, `pointer uint64 conversion test`, toTest(ToString(&ui)), ToStringValue(&ui), `564`, nil)
	case 25:
		var ui uintptr = 123
		return newTest(num, `simple uintptr conversion test`, toTest(ToString(ui)), ToStringValue(ui), `123`, nil)
	case 26:
		var ui uintptr = 123
		return newTest(num, `pointer uintptr conversion test`, toTest(ToString(&ui)), ToStringValue(&ui), `123`, nil)
	case 27:
		var f float32 = math.Pi
		return newTest(num, `simple float32 pi conversion test`, toTest(ToString(f)), ToStringValue(f), `3.1415927410125732421875`, nil)
	case 28:
		var f float32 = math.Pi * -1
		return newTest(num, `pointer float32 pi conversion test`, toTest(ToString(&f)), ToStringValue(&f), `-3.1415927410125732421875`, nil)
	case 29:
		var f float64 = math.Pi
		return newTest(num, `simple float64 pi conversion test`, toTest(ToString(f)), ToStringValue(f), `3.141592653589793115997963468544185161590576171875`, nil)
	case 30:
		var f float64 = math.Pi
		return newTest(num, `pointer float64 pi conversion test`, toTest(ToString(&f)), ToStringValue(&f), `3.141592653589793115997963468544185161590576171875`, nil)
	case 31:
		var f float64 = math.Pi + float64(987654318.0)
		return newTest(num, `simple pi+ conversion test`, toTest(ToString(f)), ToStringValue(f), `987654321.14159262180328369140625`, nil)
	case 32:
		var f float64 = math.Pi + float64(987654318.0)
		return newTest(num, `pointer pi+ conversion test`, toTest(ToString(&f)), ToStringValue(&f), `987654321.14159262180328369140625`, nil)
	case 33:
		var elemNull *uint
		return newTest(num, `simple null variable conversion test`, toTest(ToString(elemNull)), ToStringValue(elemNull), `null`, nil)
	case 34:
		var elemNull *bool
		return newTest(num, `pointer null variable conversion test`, toTest(ToString(&elemNull)), ToStringValue(&elemNull), `null`, nil)
	case 35:
		return newTest(num, `simple nil test`, toTest(ToString(nil)), ToStringValue(nil), ``, nil)
	case 36:
		slice := []int{4, 3}
		return newTest(num, `simple []int{4, 3} conversion test`, toTest(ToString(slice)), ToStringValue(slice), `4 3`, nil)
	case 37:
		slice := []int{4, 3}
		return newTest(num, `pointer []int{4, 3} conversion test`, toTest(ToString(&slice)), ToStringValue(&slice), `4 3`, nil)
	case 38:
		b := true
		return newTest(num, `simple bool = true conversion test`, toTest(ToString(b)), ToStringValue(b), `true`, nil)
	case 39:
		b := true
		return newTest(num, `pointer bool = true conversion test`, toTest(ToString(&b)), ToStringValue(&b), `true`, nil)
	case 40:
		b := false
		return newTest(num, `simple bool = false conversion test`, toTest(ToString(b)), ToStringValue(b), `false`, nil)
	case 41:
		b := false
		return newTest(num, `pointer bool = false conversion test`, toTest(ToString(&b)), ToStringValue(&b), `false`, nil)
	case 42:
		maper := map[int]int{1: 12, 3: 14}
		return newTest(num, `simple map[int]int conversion test`, toTest(ToString(maper)), ToStringValue(maper), `{"1":12,"3":14}`, nil)
	case 43:
		maper := map[int]int{1: 12, 3: 14}
		return newTest(num, `pointer map[int]int conversion test`, toTest(ToString(&maper)), ToStringValue(&maper), `{"1":12,"3":14}`, nil)
	case 44, 45:
		type structure struct {
			Val int `json:"val"`
		}
		structureItem := structure{Val: 43}
		if 44 == num {
			return newTest(num, `simple structure conversion test`, toTest(ToString(structureItem)), ToStringValue(structureItem), `{"val":43}`, nil)
		} else {
			return newTest(num, `pointer structure conversion test`, toTest(ToString(structureItem)), ToStringValue(structureItem), `{"val":43}`, nil)
		}
	case 46:
		arrayString := [2]string{"true", "false"}
		return newTest(num, `simple [2]string conversion test`, toTest(ToString(arrayString)), ToStringValue(arrayString), `true false`, nil)
	case 47:
		arrayString := [2]string{"true", "false"}
		return newTest(num, `pointer [2]string conversion test`, toTest(ToString(&arrayString)), ToStringValue(&arrayString), `true false`, nil)
	case 48:
		arrayBool := [2]bool{false, true}
		return newTest(num, `simple [2]bool conversion test`, toTest(ToString(arrayBool)), ToStringValue(arrayBool), `false true`, nil)
	case 49:
		arrayBool := [2]bool{false, true}
		return newTest(num, `pointer [2]bool conversion test`, toTest(ToString(&arrayBool)), ToStringValue(&arrayBool), `false true`, nil)
	case 50:
		return newTest(num, `simple []int{} conversion test`, toTest(ToString([]int{})), ToStringValue([]int{}), ``, nil)
	case 51:
		return newTest(num, `simple []string{} conversion test`, toTest(ToString([]string{})), ToStringValue([]string{}), ``, nil)
	case 52:
		return newTest(num, `simple []interface{}{1, 2} conversion test`, toTest(ToString([]interface{}{1, 2})), ToStringValue([]interface{}{1, 2}), `1 2`, nil)
	case 53:
		return newTest(num, `simple -2.544 conversion test`, toTest(ToString(`-2.544`)), ToStringValue(`-2.544`), `-2.544`, nil)
	case 54:
		return newTest(num, `simple -2 conversion test`, toTest(ToString(`-2`)), ToStringValue(`-2`), `-2`, nil)
	case 55:
		return newTest(num, `simple - conversion test`, toTest(ToString(`-`)), ToStringValue(`-`), `-`, nil)
	case 56:
		return newTest(num, `simple + conversion test`, toTest(ToString(`+`)), ToStringValue(`+`), `+`, nil)
	case 57:
		return newTest(num, `simple / conversion test`, toTest(ToString(`/`)), ToStringValue(`/`), `/`, nil)
	case 58:
		return newTest(num, `simple 0 conversion test`, toTest(ToString(`0`)), ToStringValue(`0`), `0`, nil)
	case 59:
		return newTest(num, `simple string "nil" conversion test`, toTest(ToString(`nil`)), ToStringValue(`nil`), `nil`, nil)
	case 60:
		return newTest(num, `simple 2.544 conversion test`, toTest(ToString(`2.544`)), ToStringValue(`2.544`), `2.544`, nil)
	case 61:
		return newTest(num, `simple 2 conversion test`, toTest(ToString(`2`)), ToStringValue(`2`), `2`, nil)
	}
	return testStructure{}, false
}

func TestToString(t *testing.T) {
	testByLists(t, []testListStructure{
		{
			Name:           `General tests`,
			FuncionGetList: getElementForTestString,
		},
	})
}
