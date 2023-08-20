package typo

import (
	"errors"
	"testing"
)

func TestTo(t *testing.T) {
	testByLists(t, []testListStructure{
		{
			Name:           `General tests`,
			FuncionGetList: getElementForTestType,
		},
		{
			Name:           `Filling Easy structures`,
			FuncionGetList: getElementForTestTypeFillingEasyStructure,
		},
		{
			Name:           `Filling Hard structures`,
			FuncionGetList: getElementForTestTypeFillingHardStructure,
		},
		{
			Name:           `Nested structures`,
			FuncionGetList: getElementForTestTypeNestedStructures,
		},
	})
}
func TestToType(t *testing.T) {
	testByLists(t, []testListStructure{
		{
			Name:           `General tests`,
			FuncionGetList: getElementForTestType,
		},
		{
			Name:           `Filling Easy structures`,
			FuncionGetList: getElementForTestTypeFillingEasyStructure,
		},
		{
			Name:           `Filling Hard structures`,
			FuncionGetList: getElementForTestTypeFillingHardStructure,
		},
		{
			Name:           `Nested structures`,
			FuncionGetList: getElementForTestTypeNestedStructures,
		},
	})
}

func getElementForTestType(num int) (testStructure, bool) {
	switch num {
	case 0:
		return newTest(num, `simple []int{1, 2, 9} -> []string conversion test`, toTest(To[[]string]([]int{1, 2, 9})), ToValue[[]string]([]int{1, 2, 9}), []string{`1`, `2`, `9`}, nil)
	case 1:
		return newTest(num, `simple []int{1, 2, 9} -> []int64 conversion test`, toTest(To[[]int64]([]int{1, 2, 9})), ToValue[[]int64]([]int{1, 2, 9}), []int64{1, 2, 9}, nil)
	case 2:
		return newTest(num, `simple []int{1, 2, -9} -> []int8 conversion test`, toTest(To[[]int8]([]int{1, 2, -9})), ToValue[[]int8]([]int{1, 2, -9}), []int8{1, 2, -9}, nil)
	case 3:
		return newTest(num, `simple []string{"1", "2", "9"} -> []int64 conversion test`, toTest(To[[]int64]([]string{"1", "2", "9"})), ToValue[[]int64]([]string{"1", "2", "9"}), []int64{1, 2, 9}, nil)
	case 4:
		return newTest(num, `simple [][]int{{1, 2}, {4, 3}} -> [][]int8 conversion test`, toTest(To[[][]int8]([][]int{{1, 2}, {4, 3}})), ToValue[[][]int8]([][]int{{1, 2}, {4, 3}}), [][]int8{{1, 2}, {4, 3}}, nil)
	case 5:
		i := 25
		i8 := int8(i)
		return newTest(num, `simple int -> []*int8  conversion test`, toTest(To[[]*int8](i)), ToValue[[]*int8](i), []*int8{&i8}, nil)
	case 6:
		i := 27
		i8 := int8(i)
		return newTest(num, `simple *int -> []*int8 conversion test`, toTest(To[[]*int8](&i)), ToValue[[]*int8](&i), []*int8{&i8}, nil)
	case 7:
		i := 99
		ss := []string{`99`}
		pss := &ss
		return newTest(num, `simple int -> []**[]string conversion test`, toTest(To[[]**[]string](i)), ToValue[[]**[]string](i), []**[]string{&pss}, nil)
	case 8:
		i := 77
		ss := []string{`77`}
		pss := &ss
		return newTest(num, `simple *int -> []**[]string conversion test`, toTest(To[[]**[]string](&i)), ToValue[[]**[]string](&i), []**[]string{&pss}, nil)
	case 9:
		i := 199
		ss := []string{`199`}
		pss := []*[]string{&ss}
		return newTest(num, `simple int -> []*[]*[]string conversion test`, toTest(To[[]*[]*[]string](i)), ToValue[[]*[]*[]string](i), []*[]*[]string{&pss}, nil)
	case 10:
		i := 177
		ss := []string{`177`}
		pss := []*[]string{&ss}
		return newTest(num, `simple *int -> []*[]*[]string conversion test`, toTest(To[[]*[]*[]string](&i)), ToValue[[]*[]*[]string](&i), []*[]*[]string{&pss}, nil)
	case 11:
		arrayBool := [2]bool{true, false}
		return newTest(num, `simple [2]bool -> []string conversion test`, toTest(To[[]string](arrayBool)), ToValue[[]string](arrayBool), []string{`true`, `false`}, nil)
	case 12:
		arrayBool := [2]bool{true, false}
		return newTest(num, `simple [2]bool -> [][2]string conversion test`, toTest(To[[][2]string](arrayBool)), ToValue[[][2]string](arrayBool), [][2]string{{`true`, ``}, {`false`, ``}}, nil)
	case 13:
		arrayBool := [2]bool{true, false}
		t := `true`
		f := `false`
		return newTest(num, `simple [2]bool -> [][2]*string conversion test`, toTest(To[[][2]*string](arrayBool)), ToValue[[][2]*string](arrayBool), [][2]*string{{&t, nil}, {&f, nil}}, nil)
	case 14:
		arrayBool := [2]bool{true, false}
		return newTest(num, `simple [2]bool -> [][2]string conversion test`, toTest(To[[][2]string]([][2]bool{arrayBool})), ToValue[[][2]string]([][2]bool{arrayBool}), [][2]string{{`true`, `false`}}, nil)
	case 15:
		return newTest(num, `simple nil -> int conversion test`, toTest(To[int](nil)), ToValue[int](nil), 0, nil)
	case 16:
		arrayBool := [2]bool{true, false}
		s := `true false`
		return newTest(num, `simple *[2]bool -> [][2]*string conversion test`, toTest(To[[][2]*string](&arrayBool)), ToValue[[][2]*string](&arrayBool), [][2]*string{{&s, nil}}, nil)
	case 17:
		arrayBool := [2]bool{true, false}
		return newTest(num, `simple *[2]bool -> [][2]string conversion test`, toTest(To[[][2]string](&arrayBool)), ToValue[[][2]string](&arrayBool), [][2]string{{`true false`, ``}}, nil)
	case 18:
		arrayBool := [2]bool{true, false}
		return newTest(num, `simple [2]bool -> [][2]string conversion test`, toTest(To[[][2]string](arrayBool)), ToValue[[][2]string](arrayBool), [][2]string{{`true`, ``}, {`false`, ``}}, nil)
	case 19:
		arrayBool := [2]bool{true, false}
		return newTest(num, `simple [2]bool -> map[int]string conversion test`, toTest(To[map[int]string](arrayBool)), ToValue[map[int]string](arrayBool), map[int]string{0: `true`, 1: `false`}, nil)
	case 20:
		mapStringString := map[string]string{`a`: `b`, `c`: `2`}
		return newTest(num, `simple  map[string]string -> map[string]string  conversion test`, toTest(To[map[string]string](mapStringString)), ToValue[map[string]string](mapStringString), map[string]string{`a`: `b`, `c`: `2`}, nil)
	case 21:
		mapStringInterface := map[string]interface{}{`a`: `b`, `c`: 2}
		return newTest(num, `simple map[string]interface{} -> map[string]string conversion test`, toTest(To[map[string]string](mapStringInterface)), ToValue[map[string]string](mapStringInterface), map[string]string{`a`: `b`, `c`: `2`}, nil)
	case 22:
		mapNumericString := map[string]string{`1`: `b`, `2`: `d`}
		return newTest(num, `simple  map[string]string -> map[int]string conversion test`, toTest(To[map[int]string](mapNumericString)), ToValue[map[int]string](mapNumericString), map[int]string{1: `b`, 2: `d`}, nil)
	case 23:
		mapNumericString := map[string]string{`1`: `b`, `2`: `d`}
		return newTest(num, `simple map[string]string -> map[int8]string  conversion test`, toTest(To[map[int8]string](mapNumericString)), ToValue[map[int8]string](mapNumericString), map[int8]string{1: `b`, 2: `d`}, nil)
	case 24:
		mapNumericString := map[string]string{`1`: `b`, `2`: `d`}
		return newTest(num, `simple map[string]string -> map[int64]string  conversion test`, toTest(To[map[int64]string](mapNumericString)), ToValue[map[int64]string](mapNumericString), map[int64]string{1: `b`, 2: `d`}, nil)
	case 25:
		return newTest(num, `simple int -> [5]string  conversion test`, toTest(To[[5]string](5)), ToValue[[5]string](5), [5]string{`5`, ``, ``, ``, ``}, nil)
	case 26:
		return newTest(num, `simple int -> [5]int conversion test`, toTest(To[[5]int](6)), ToValue[[5]int](6), [5]int{6, 0, 0, 0, 0}, nil)
	case 27:
		return newTest(num, `simple int -> int64 conversion test`, toTest(To[int64](68)), ToValue[int64](68), int64(68), nil)
	case 28:
		return newTest(num, `simple string -> int64 conversion test`, toTest(To[int64](`69`)), ToValue[int64](`69`), int64(69), nil)
	case 29:
		return newTest(num, `simple string -> [5]int conversion test`, toTest(To[[5]int](`6`)), ToValue[[5]int](`6`), [5]int{6, 0, 0, 0, 0}, nil)
	case 30:
		return newTest(num, `simple char -> [5]int conversion test`, toTest(To[[5]int]('6')), ToValue[[5]int]('6'), [5]int{54, 0, 0, 0, 0}, nil)
	case 31:
		return newTest(num, `simple complex -> string conversion test`, toTest(To[string](complex(1, 2))), ToValue[string](complex(1, 2)), `(1+2i)`, nil)
	case 32:
		return newTest(num, `simple complex -> []string conversion test`, toTest(To[[]string](complex(1, 2))), ToValue[[]string](complex(1, 2)), []string{`(1+2i)`}, nil)
	case 33:
		return newTest(num, `simple complex -> [2]string conversion test`, toTest(To[[2]string](complex(1, 2))), ToValue[[2]string](complex(1, 2)), [2]string{`(1+2i)`, ``}, nil)
	case 34:
		var ss []string
		return newTest(num, `simple nil -> []string conversion test`, toTest(To[[]string](nil)), ToValue[[]string](nil), ss, nil)
	case 35:
		var ss [2]string
		return newTest(num, `simple nil -> [2]string conversion test`, toTest(To[[2]string](nil)), ToValue[[2]string](nil), ss, nil)
	case 36:
		var mis map[int]string
		return newTest(num, `simple nil -> map[int]string conversion test`, toTest(To[map[int]string](nil)), ToValue[map[int]string](nil), mis, nil)
	case 37:
		var mis *map[int]string
		return newTest(num, `simple nil -> *map[int]string conversion test`, toTest(To[*map[int]string](nil)), ToValue[*map[int]string](nil), mis, nil)
	case 38:
		var pi *int
		return newTest(num, `simple nil -> *int conversion test`, toTest(To[*int](nil)), ToValue[*int](nil), pi, nil)
	case 40:
		mapStringInterface := map[string]interface{}{`a`: `b`, `b`: 1, `ce`: 2, `D`: 3}
		return newTest(num, `simple map[string]interface{} -> [5]int conversion test`, toTest(To[[5]int](mapStringInterface)), ToValue[[5]int](mapStringInterface), [5]int{}, errors.New(`strconv.ParseFloat: parsing "b": invalid syntax`))
	case 41:
		mapStringInterface := map[string]interface{}{`a`: `b`, `b`: 1, `ce`: 2, `D`: 3}
		return newTest(num, `simple map[string]interface{} -> []string conversion test`, toTest(To[[]string](mapStringInterface)), ToValue[[]string](mapStringInterface), []string{"b", "1", "2", "3"}, nil) // @todo не постоянные тесты. лучше держать выключеными
	case 42:
		mapStringInterface := map[string]interface{}{`a`: `b`, `b`: 1, `ce`: 2, `D`: 3}
		return newTest(num, `simple map[string]interface{} -> [3]string conversion test`, toTest(To[[3]string](mapStringInterface)), ToValue[[3]string](mapStringInterface), [3]string{"b", "1", "2"}, nil) // @todo не постоянные тесты. лучше держать выключеными
	case 43:
		mapStringInterface := map[string]interface{}{`a`: `b`, `b`: 1, `ce`: 2, `D`: 3}
		return newTest(num, `simple map[string]interface{} -> [5]string conversion test`, toTest(To[[5]string](mapStringInterface)), ToValue[[5]string](mapStringInterface), [5]string{"b", "1", "2", "3", ""}, nil) // @todo не постоянные тесты. лучше держать выключеными
	case 44:
		return newTest(num, `simple []int -> complex64 conversion test`, toTest(To[complex64]([]int{1, 2})), ToValue[complex64]([]int{1, 2}), complex(1, 2), nil) //@todo json не работает с complex64
	case 45:
		return newTest(num, `simple ToType []int{1, 2, 9} -> []string conversion test`, toTest(ToType[[]string]([]int{1, 2, 9})), ToTypeValue[[]string]([]int{1, 2, 9}), []string{`1`, `2`, `9`}, nil)
	case 46:
		return newTest(num, `simple string -> int conversion test`, toTest(ToType[int](`123`)), ToTypeValue[int](`123`), 123, nil)
	}
	return testStructure{}, false
}
func getElementForTestTypeFillingEasyStructure(num int) (testStructure, bool) {
	type StructAB struct {
		A string `json:"a"`
		B int    `json:"b"`
	}
	type ListStructAB struct {
		AB1 StructAB `json:"ab1"`
		AB2 StructAB `json:"ab2"`
	}
	switch num {
	case 0:
		return newTest(num, `simple []string -> StructAB  conversion test`, toTest(To[StructAB]([]string{`b`, `2`})), ToValue[StructAB]([]string{`b`, `2`}), StructAB{A: `b`, B: 2}, nil)
	case 1:
		return newTest(num, `simple [3]string -> StructAB  conversion test`, toTest(To[StructAB]([3]string{`b`, `2`, `4`})), ToValue[StructAB]([3]string{`b`, `2`, `4`}), StructAB{A: `b`, B: 2}, nil)
	case 2:
		return newTest(num, `simple [1]string -> StructAB  conversion test`, toTest(To[StructAB]([1]string{`c`})), ToValue[StructAB]([1]string{`c`}), StructAB{A: `c`, B: 0}, nil)
	case 3:
		return newTest(num, `simple []map[string]string -> []StructAB  conversion test`, toTest(To[[]StructAB]([]map[string]string{{`a`: `aaaa`, `b`: `1`}, {`a`: `aaaa2`, `b`: `2`}})), ToValue[[]StructAB]([]map[string]string{{`a`: `aaaa`, `b`: `1`}, {`a`: `aaaa2`, `b`: `2`}}), []StructAB{{A: `aaaa`, B: 1}, {A: `aaaa2`, B: 2}}, nil)
	case 4:
		data := map[string]map[string]string{`ab1`: {`a`: `aaaa`, `b`: `1`}, `ab2`: {`a`: `aaaa2`, `b`: `2`}}
		return newTest(num, `simple map[string]map[string]string -> ListStructAB  conversion test`, toTest(To[ListStructAB](data)), ToValue[ListStructAB](data), ListStructAB{AB1: StructAB{A: `aaaa`, B: 1}, AB2: StructAB{A: `aaaa2`, B: 2}}, nil)
		/*case 5:
		data := map[string]map[string]string{`ab1`: {`a`: `aaaa`, `b`: `1`}, `ab2`: {`a`: `aaaa2`, `b`: `2`}}
		return newTest(num, `simple map[string]map[string]string -> []StructAB  conversion test`, toTest(To[[]StructAB](data)), ToValue[[]StructAB](data), []StructAB{{A: `aaaa`, B: 1}, {A: `aaaa2`, B: 2}}, nil) // @todo не постоянные тесты. лучше держать выключеными */
	}
	return testStructure{}, false
}
func getElementForTestTypeFillingHardStructure(num int) (testStructure, bool) {
	type structHard struct {
		A string  `json:"a,omitempty"`
		B *string `json:"b"`
		C int     `json:"ce"`
		D []int
	}
	str := `string data`
	structure := structHard{A: `ac`, B: &str, C: 3, D: []int{4, 5}}
	mapStringInterface := map[string]interface{}{`a`: `b`, `b`: 1, `ce`: 2, `D`: 3}
	mapStringString := map[string]string{`a`: `b`, `c`: `2`}
	type StructAB struct {
		A string `json:"a"`
		B int    `json:"b"`
	}
	type StructAC struct {
		A string `json:"a"`
		C int    `json:"b,omitempty"`
	}
	switch num {
	case 0:
		return newTest(num, `simple structHard -> map[string]string conversion test`, toTest(To[map[string]string](structure)), ToValue[map[string]string](structure), map[string]string{`D`: `4 5`, `a`: `ac`, `b`: `string data`, `ce`: `3`}, nil)
	case 1:
		return newTest(num, `simple structHard -> []string conversion test`, toTest(To[[]string](structure)), ToValue[[]string](structure), []string{`ac`, `string data`, `3`, `4 5`}, nil)
	case 2:
		return newTest(num, `simple structHard -> [4]string conversion test`, toTest(To[[4]string](structure)), ToValue[[4]string](structure), [4]string{`ac`, `string data`, `3`, `4 5`}, nil)
	case 3:
		return newTest(num, `simple structHard -> [5]string conversion test`, toTest(To[[5]string](structure)), ToValue[[5]string](structure), [5]string{`ac`, `string data`, `3`, `4 5`, ``}, nil)
	case 4:
		b := `1`
		return newTest(num, `simple map[string]interface{} -> structHard  conversion test`, toTest(To[structHard](mapStringInterface)), ToValue[structHard](mapStringInterface), structHard{A: `b`, B: &b, C: 2, D: []int{3}}, nil)
	case 5:
		return newTest(num, `simple map[string]string (a b) -> structHard  conversion test`, toTest(To[structHard](mapStringString)), ToValue[structHard](mapStringString), structHard{A: `b`, B: nil, C: 0, D: nil}, nil)
	case 6:
		b := `2`
		return newTest(num, `simple map[string]string (a c) -> structHard  conversion test`, toTest(To[structHard](map[string]string{`a`: `b`, `b`: `2`})), ToValue[structHard](map[string]string{`a`: `b`, `b`: `2`}), structHard{A: `b`, B: &b, C: 0, D: nil}, nil)
	case 7:
		return newTest(num, `simple nil -> structHard  conversion test`, toTest(To[structHard](nil)), ToValue[structHard](nil), structHard{A: ``, B: nil, C: 0, D: nil}, nil)
	case 8:
		b := `1223`
		return newTest(num, `simple StructAB -> structHard  conversion test`, toTest(To[structHard](StructAB{A: `aasd`, B: 1223})), ToValue[structHard](StructAB{A: `aasd`, B: 1223}), structHard{A: `aasd`, B: &b, C: 0, D: nil}, nil)
	case 9:
		b := `1223`
		return newTest(num, `simple StructAC -> structHard  conversion test`, toTest(To[structHard](StructAC{A: `absd`, C: 1223})), ToValue[structHard](StructAC{A: `absd`, C: 1223}), structHard{A: `absd`, B: &b, C: 1223, D: nil}, nil)
	}
	return testStructure{}, false
}
func getElementForTestTypeNestedStructures(num int) (testStructure, bool) {
	type structBase struct {
		A string  `json:"a"`
		B *string `json:"b"`
	}
	type structInherit struct {
		structBase
		C int `json:"ce"`
		D []int
	}
	str := `string data`
	mapStringInterface := map[string]interface{}{`a`: `b`, `b`: 1, `ce`: 2, `D`: 3}
	StructureBase := structBase{A: `ac`, B: &str}

	StructureInherit := structInherit{C: 3, D: []int{4, 5}}
	StructureInherit.A = `ac`
	StructureInherit.B = &str

	switch num {
	case 0:
		return newTest(num, `simple StructureBase -> []string conversion test`, toTest(To[[]string](StructureBase)), ToValue[[]string](StructureBase), []string{`ac`, `string data`}, nil)
	case 1:
		return newTest(num, `simple *StructureBase -> []string conversion test`, toTest(To[[]string](&StructureBase)), ToValue[[]string](&StructureBase), []string{`{"a":"ac","b":"string data"}`}, nil)
	case 2:
		return newTest(num, `simple StructureBase -> [2]string conversion test`, toTest(To[[2]string](StructureBase)), ToValue[[2]string](StructureBase), [2]string{`ac`, `string data`}, nil)
	case 3:
		return newTest(num, `simple structInherit -> map[string]string conversion test`, toTest(To[map[string]string](StructureInherit)), ToValue[map[string]string](StructureInherit), map[string]string{"D": "4 5", "a": "ac", "b": "string data", "ce": "3"}, nil)
	case 4:
		return newTest(num, `simple structInherit -> []string conversion test`, toTest(To[[]string](StructureInherit)), ToValue[[]string](StructureInherit), []string{"ac", "string data", "3", "4 5"}, nil)
	case 5:
		return newTest(num, `simple structInherit -> [5]string conversion test`, toTest(To[[5]string](StructureInherit)), ToValue[[5]string](StructureInherit), [5]string{"ac", "string data", "3", "4 5", ""}, nil)
	case 6:
		return newTest(num, `simple structInherit -> [4]string conversion test`, toTest(To[[4]string](StructureInherit)), ToValue[[4]string](StructureInherit), [4]string{"ac", "string data", "3", "4 5"}, nil)
	case 7:
		var ss [3]string
		return newTest(num, `simple structInherit -> [3]string conversion test`, toTest(To[[3]string](StructureInherit)), ToValue[[3]string](StructureInherit), ss, errors.New(`the passed variable or its part of type typo.structInherit cannot be placed in type [3]string because it is greater than it`))
	case 8:
		b := `1`
		return newTest(num, `simple map[string]interface{} -> structInherit conversion test`, toTest(To[structInherit](mapStringInterface)), ToValue[structInherit](mapStringInterface), structInherit{structBase: structBase{A: `b`, B: &b}, C: 2, D: []int{3}}, nil)
	case 9:
		b := `2`
		return newTest(num, `simple [5]string -> structInherit conversion test`, toTest(To[structInherit]([5]string{`b`, `2`, `4`, `8`, `12`})), ToValue[structInherit]([5]string{`b`, `2`, `4`, `8`, `12`}), structInherit{structBase: structBase{A: `b`, B: &b}, C: 4, D: []int{8}}, nil)
	case 10:
		b := `2`
		return newTest(num, `simple [5]interface{} -> structInherit conversion test`, toTest(To[structInherit]([5]interface{}{`b`, `2`, `4`, `8`, `12`})), ToValue[structInherit]([5]interface{}{`b`, `2`, `4`, `8`, `12`}), structInherit{structBase: structBase{A: `b`, B: &b}, C: 4, D: []int{8}}, nil)
	case 11:
		b := `2`
		return newTest(num, `simple [5]interface{} -> structInherit conversion test`, toTest(To[structInherit]([5]interface{}{`b`, `2`, `4`, []string{`8`, `12`}})), ToValue[structInherit]([5]interface{}{`b`, `2`, `4`, []string{`8`, `12`}}), structInherit{structBase: structBase{A: `b`, B: &b}, C: 4, D: []int{8, 12}}, nil)
	}
	return testStructure{}, false
}
