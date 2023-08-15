package typo

import (
	"errors"
	"math"
	"testing"
)

func getElementForTestBool(num int) (testStructure, bool) {
	switch num {
	case 0:
		return newTest(num, `simple true conversion test`, toTest(ToBool(true)), ToBoolValue(true), true, nil)
	case 1:
		return newTest(num, `simple false conversion test`, toTest(ToBool(false)), ToBoolValue(false), false, nil)
	case 2:
		return newTest(num, `simple integer conversion test`, toTest(ToBool(1)), ToBoolValue(1), true, nil)
	case 3:
		var i int = 2
		return newTest(num, `pointer integer conversion test`, toTest(ToBool(&i)), ToBoolValue(&i), true, nil)
	case 4:
		var i int8 = 8
		return newTest(num, `simple integer 8 in variable conversion test`, toTest(ToBool(i)), ToBoolValue(i), true, nil)
	case 5:
		var i int8 = 8
		return newTest(num, `pointer integer 8 conversion test`, toTest(ToBool(&i)), ToBoolValue(&i), true, nil)
	case 6:
		var i int16 = 16
		return newTest(num, `simple integer 16 in variable conversion test`, toTest(ToBool(i)), ToBoolValue(i), true, nil)
	case 7:
		var i int16 = 16
		return newTest(num, `pointer integer 16 conversion test`, toTest(ToBool(&i)), ToBoolValue(&i), true, nil)
	case 8:
		var i int32 = 32
		return newTest(num, `simple integer 32 in variable conversion test`, toTest(ToBool(i)), ToBoolValue(i), true, nil)
	case 9:
		var i int32 = 32
		return newTest(num, `pointer integer 32 conversion test`, toTest(ToBool(&i)), ToBoolValue(&i), true, nil)
	case 10:
		var i int64 = 64
		return newTest(num, `simple integer 64 in variable conversion test`, toTest(ToBool(i)), ToBoolValue(i), true, nil)
	case 11:
		var i int64 = 64
		return newTest(num, `pointer integer 64 conversion test`, toTest(ToBool(&i)), ToBoolValue(&i), true, nil)
	case 12:
		var i int = 0
		return newTest(num, `simple integer = 0 conversion test`, toTest(ToBool(i)), ToBoolValue(i), false, nil)
	case 13:
		var i int = 0
		return newTest(num, `pointer integer = 0 conversion test`, toTest(ToBool(i)), ToBoolValue(i), false, nil)
	case 14:
		var i int = -1
		return newTest(num, `simple integer = -1 conversion test`, toTest(ToBool(i)), ToBoolValue(i), true, nil)
	case 15:
		var i int = -1
		return newTest(num, `pointer integer = -1 conversion test`, toTest(ToBool(i)), ToBoolValue(i), true, nil)
	case 16:
		var ui uint = 51
		return newTest(num, `simple uint in variable conversion test`, toTest(ToBool(ui)), ToBoolValue(ui), true, nil)
	case 17:
		var ui uint = 51
		return newTest(num, `pointer uint conversion test`, toTest(ToBool(&ui)), ToBoolValue(&ui), true, nil)
	case 18:
		var ui uint8 = 58
		return newTest(num, `simple uint8 in variable conversion test`, toTest(ToBool(ui)), ToBoolValue(ui), true, nil)
	case 19:
		var ui uint8 = 58
		return newTest(num, `pointer uint8 conversion test`, toTest(ToBool(&ui)), ToBoolValue(&ui), true, nil)
	case 20:
		var ui uint16 = 516
		return newTest(num, `simple uint16 in variable conversion test`, toTest(ToBool(ui)), ToBoolValue(ui), true, nil)
	case 21:
		var ui uint16 = 516
		return newTest(num, `pointer uint16 conversion test`, toTest(ToBool(&ui)), ToBoolValue(&ui), true, nil)
	case 22:
		var ui uint32 = 532
		return newTest(num, `simple uint32 in variable conversion test`, toTest(ToBool(ui)), ToBoolValue(ui), true, nil)
	case 23:
		var ui uint32 = 532
		return newTest(num, `pointer uint32 conversion test`, toTest(ToBool(&ui)), ToBoolValue(&ui), true, nil)
	case 24:
		var ui uint64 = 564
		return newTest(num, `simple uint64 in variable conversion test`, toTest(ToBool(ui)), ToBoolValue(ui), true, nil)
	case 25:
		var ui uint64 = 564
		return newTest(num, `pointer uint64 conversion test`, toTest(ToBool(&ui)), ToBoolValue(&ui), true, nil)
	case 26:
		var ui uintptr = 123
		return newTest(num, `simple uintptr in variable conversion test`, toTest(ToBool(ui)), ToBoolValue(ui), true, nil)
	case 27:
		var ui uintptr = 123
		return newTest(num, `pointer uintptr conversion test`, toTest(ToBool(&ui)), ToBoolValue(&ui), true, nil)
	case 28:
		var f float32 = math.Pi
		return newTest(num, `simple float32 in variable conversion test`, toTest(ToBool(f)), ToBoolValue(f), true, nil)
	case 29:
		var f float32 = math.Pi
		return newTest(num, `pointer float32 conversion test`, toTest(ToBool(&f)), ToBoolValue(&f), true, nil)
	case 30:
		var f float64 = math.Pi
		return newTest(num, `simple float64 in variable conversion test`, toTest(ToBool(f)), ToBoolValue(f), true, nil)
	case 31:
		var f float64 = math.Pi
		return newTest(num, `pointer float64 conversion test`, toTest(ToBool(&f)), ToBoolValue(&f), true, nil)
	case 32:
		var f float64 = math.Pi + float64(987654318.0)
		return newTest(num, `simple float64 in variable conversion test 2`, toTest(ToBool(f)), ToBoolValue(f), true, nil)
	case 33:
		var f float64 = math.Pi + float64(987654318.0)
		return newTest(num, `pointer float64 conversion test 2`, toTest(ToBool(&f)), ToBoolValue(&f), true, nil)
	case 34:
		var iNull *int
		return newTest(num, `simple float64 in variable conversion test 2`, toTest(ToBool(iNull)), ToBoolValue(iNull), false, nil)
	case 35:
		var iNull *int
		return newTest(num, `pointer float64 conversion test 2`, toTest(ToBool(&iNull)), ToBoolValue(&iNull), false, nil)
	case 36:
		slice := []int{4, 3}
		return newTest(num, `[]int in variable conversion test`, toTest(ToBool(slice)), ToBoolValue(slice), true, nil)
	case 37:
		slice := []int{7, 8}
		return newTest(num, `pointer []int conversion test`, toTest(ToBool(&slice)), ToBoolValue(&slice), true, nil)
	case 38:
		b := true
		return newTest(num, `simple bool = true in variable conversion test`, toTest(ToBool(b)), ToBoolValue(b), true, nil)
	case 39:
		b := true
		return newTest(num, `pointer bool = true conversion test`, toTest(ToBool(&b)), ToBoolValue(&b), true, nil)
	case 40:
		b := false
		return newTest(num, `simple bool = false in variable conversion test`, toTest(ToBool(b)), ToBoolValue(b), false, nil)
	case 41:
		b := false
		return newTest(num, `pointer bool = false conversion test`, toTest(ToBool(&b)), ToBoolValue(&b), false, nil)
	case 42:
		maper := map[int]int{1: 12, 3: 14}
		return newTest(num, `simple map[int]int in variable conversion test`, toTest(ToBool(maper)), ToBoolValue(maper), true, nil)
	case 43:
		maper := map[int]int{1: 16, 3: 15}
		return newTest(num, `pointer map[int]int conversion test`, toTest(ToBool(&maper)), ToBoolValue(&maper), true, nil)
	case 44, 45:
		type structure struct {
			Val int `json:"val"`
		}
		structureItem := structure{Val: 43}
		if 44 == num {
			return newTest(num, `simple structure in variable conversion test`, toTest(ToBool(structureItem)), ToBoolValue(structureItem), false, errors.New(`typo.structure format cannot be converted to bool`))
		} else {
			return newTest(num, `pointer structure conversion test`, toTest(ToBool(&structureItem)), ToBoolValue(&structureItem), false, errors.New(`typo.structure format cannot be converted to bool`))
		}
	case 46:
		return newTest(num, `simple []int{} conversion test`, toTest(ToBool([]int{})), ToBoolValue([]int{}), false, nil)
	case 47:
		return newTest(num, `simple []string{} conversion test`, toTest(ToBool([]string{})), ToBoolValue([]string{}), false, nil)
	case 48:
		return newTest(num, `simple []interface{}{1} conversion test`, toTest(ToBool([]interface{}{1})), ToBoolValue([]interface{}{1}), true, nil)
	case 49:
		return newTest(num, `simple []interface{}{0} conversion test`, toTest(ToBool([]interface{}{0})), ToBoolValue([]interface{}{0}), true, nil)
	case 50:
		return newTest(num, `simple []interface{}{-1} conversion test`, toTest(ToBool([]interface{}{-1})), ToBoolValue([]interface{}{-1}), true, nil)
	case 51:
		arrayString := [2]string{"true", "false"}
		return newTest(num, `simple [2]string in variable conversion test`, toTest(ToBool(arrayString)), ToBoolValue(arrayString), false, errors.New(`[2]string format cannot be converted to bool`))
	case 52:
		arrayString := [2]string{"true", "false"}
		return newTest(num, `pointer [2]string conversion test`, toTest(ToBool(&arrayString)), ToBoolValue(&arrayString), false, errors.New(`[2]string format cannot be converted to bool`))
	case 53:
		arrayBool := [2]bool{true, false}
		return newTest(num, `simple [2]bool in variable conversion test`, toTest(ToBool(arrayBool)), ToBoolValue(arrayBool), false, errors.New(`[2]bool format cannot be converted to bool`))
	case 54:
		arrayBool := [2]bool{true, false}
		return newTest(num, `pointer [2]bool conversion test`, toTest(ToBool(&arrayBool)), ToBoolValue(&arrayBool), false, errors.New(`[2]bool format cannot be converted to bool`))
	case 55:
		return newTest(num, `simple string = -2.544 conversion test`, toTest(ToBool(`-2.544`)), ToBoolValue(`-2.544`), true, nil)
	case 56:
		return newTest(num, `simple string = -2 conversion test`, toTest(ToBool(`-2`)), ToBoolValue(`-2`), true, nil)
	case 57:
		return newTest(num, `simple string = 0 conversion test`, toTest(ToBool(`0`)), ToBoolValue(`0`), false, nil)
	case 58:
		return newTest(num, `simple string = '' conversion test`, toTest(ToBool(``)), ToBoolValue(``), false, nil)
	case 59:
		return newTest(num, `simple string = - conversion test`, toTest(ToBool(`-`)), ToBoolValue(`-`), false, nil)
	case 60:
		return newTest(num, `simple string = + conversion test`, toTest(ToBool(`+`)), ToBoolValue(`+`), true, nil)
	case 61:
		return newTest(num, `simple string = 1 conversion test`, toTest(ToBool(`1`)), ToBoolValue(`1`), true, nil)
	case 62:
		return newTest(num, `simple string = qwe conversion test`, toTest(ToBool(`qwe`)), ToBoolValue(`qwe`), true, nil)
	case 63:
		return newTest(num, `simple string = / conversion test`, toTest(ToBool(`/`)), ToBoolValue(`/`), true, nil)
	case 64:
		return newTest(num, `simple string = 2 conversion test`, toTest(ToBool(`2`)), ToBoolValue(`2`), true, nil)
	case 65:
		return newTest(num, `simple string = 2.544 conversion test`, toTest(ToBool(`2.544`)), ToBoolValue(`2.544`), true, nil)
	case 66:
		return newTest(num, `simple map[string]interface{}{} conversion test`, toTest(ToBool(map[string]interface{}{})), ToBoolValue(map[string]interface{}{}), false, nil)
	case 67:
		var mapNil map[int]int = nil
		return newTest(num, `simple map[int]int = nil conversion test`, toTest(ToBool(mapNil)), ToBoolValue(mapNil), false, nil)
	case 68:
		return newTest(num, `simple nil conversion test`, toTest(ToBool(nil)), ToBoolValue(nil), false, nil)
	case 69:
		type structure struct {
			Val int `json:"val"`
		}
		var structureItem *structure = nil
		return newTest(num, `simple nil conversion test`, toTest(ToBool(structureItem)), ToBoolValue(structureItem), false, nil)
	case 70:
		var complexItem = complex(1, 0)
		return newTest(num, `simple complex(1, 0) in variable conversion test`, toTest(ToBool(complexItem)), ToBoolValue(complexItem), true, nil)
	case 71:
		var complexItem = complex(1, 0)
		return newTest(num, `pointer complex(1, 0) conversion test`, toTest(ToBool(&complexItem)), ToBoolValue(&complexItem), true, nil)
	case 72:
		var complexItem = complex(0, 0)
		return newTest(num, `simple complex(0, 0) in variable conversion test`, toTest(ToBool(complexItem)), ToBoolValue(complexItem), false, nil)
	case 73:
		var complexItem = complex(0, 0)
		return newTest(num, `pointer complex(0, 0) conversion test`, toTest(ToBool(&complexItem)), ToBoolValue(&complexItem), false, nil)
	case 74:
		var complexItem = complex(0, 1)
		return newTest(num, `simple complex(0, 1) in variable conversion test`, toTest(ToBool(complexItem)), ToBoolValue(complexItem), true, nil)
	case 75:
		var complexItem = complex(0, 1)
		return newTest(num, `pointer complex(0, 1) conversion test`, toTest(ToBool(&complexItem)), ToBoolValue(&complexItem), true, nil)
	}
	return testStructure{}, false
}

func TestToBool(t *testing.T) {
	testByLists(t, []testListStructure{
		{
			Name:           `General tests`,
			FuncionGetList: getElementForTestBool,
		},
	})
}
