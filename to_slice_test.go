package typo

import (
	"testing"
)

func getElementForTestSlice(num int) (testStructure, bool) {
	switch num {
	case 0:
		return newTest(num, `simple []int -> []string  conversion test`, toTest(ToSlice[string]([]int{1, 2, 9})), ToSliceValue[string]([]int{1, 2, 9}), []string{`1`, `2`, `9`}, nil)
	case 1:
		return newTest(num, `simple []int -> []int64 conversion test`, toTest(ToSlice[int64]([]int{1, 2, 9})), ToSliceValue[int64]([]int{1, 2, 9}), []int64{1, 2, 9}, nil)
	case 2:
		return newTest(num, `simple []int -> []int8 conversion test`, toTest(ToSlice[int8]([]int{1, 2, -9})), ToSliceValue[int8]([]int{1, 2, -9}), []int8{1, 2, -9}, nil)
	case 3:
		return newTest(num, `simple []string -> []int64 conversion test`, toTest(ToSlice[int64]([]string{`1`, `2`, `-9`})), ToSliceValue[int64]([]string{`1`, `2`, `-9`}), []int64{1, 2, -9}, nil)
	case 4:
		return newTest(num, `simple int -> []int8 conversion test`, toTest(ToSlice[int8](15)), ToSliceValue[int8](15), []int8{15}, nil)
	case 5:
		return newTest(num, `simple string -> []int8 conversion test`, toTest(ToSlice[int8](`18`)), ToSliceValue[int8](`18`), []int8{18}, nil)
	case 6:
		return newTest(num, `simple [][]int -> [][]int8 conversion test`, toTest(ToSlice[[]int8]([][]int{{1, 2}, {4, 3}})), ToSliceValue[[]int8]([][]int{{1, 2}, {4, 3}}), [][]int8{{1, 2}, {4, 3}}, nil)
	case 7:
		i := 42
		i8 := int8(42)
		return newTest(num, `simple int -> []*int8 conversion test`, toTest(ToSlice[*int8](i)), ToSliceValue[*int8](i), []*int8{&i8}, nil)
	case 8:
		i := 42
		i8 := int8(42)
		return newTest(num, `pointer *int -> []*int8 conversion test`, toTest(ToSlice[*int8](&i)), ToSliceValue[*int8](&i), []*int8{&i8}, nil)
	case 9:
		i := 17
		is := []string{`17`}
		isp := &is
		return newTest(num, `simple int -> []**[]string conversion test`, toTest(ToSlice[**[]string](i)), ToSliceValue[**[]string](i), []**[]string{&isp}, nil)
	case 10:
		i := 17
		ss := []string{`17`}
		pss := &ss
		return newTest(num, `pointer *int -> []**[]string conversion test`, toTest(ToSlice[**[]string](&i)), ToSliceValue[**[]string](&i), []**[]string{&pss}, nil)
	case 11:
		i := 88
		ss := []string{`88`}
		spss := []*[]string{&ss}
		return newTest(num, `simple int -> []**[]string conversion test`, toTest(ToSlice[*[]*[]string](i)), ToSliceValue[*[]*[]string](i), []*[]*[]string{&spss}, nil)
	case 12:
		i := 88
		ss := []string{`88`}
		spss := []*[]string{&ss}
		return newTest(num, `pointer *int -> []**[]string conversion test`, toTest(ToSlice[*[]*[]string](&i)), ToSliceValue[*[]*[]string](&i), []*[]*[]string{&spss}, nil)
	case 13:
		arrayBool := [2]bool{true, false}
		return newTest(num, `simple [2]bool -> []string conversion test`, toTest(ToSlice[string](arrayBool)), ToSliceValue[string](arrayBool), []string{`true`, `false`}, nil)
	case 14:
		arrayBool := [2]bool{true, false}
		return newTest(num, `simple [2]bool -> [][2]string conversion test`, toTest(ToSlice[[2]string](arrayBool)), ToSliceValue[[2]string](arrayBool), [][2]string{{`true`, ``}, {`false`, ``}}, nil)
	case 15:
		arrayBool := [2]bool{false, true}
		t := `true`
		f := `false`
		return newTest(num, `simple [2]bool -> [][2]*string conversion test`, toTest(ToSlice[[2]*string](arrayBool)), ToSliceValue[[2]*string](arrayBool), [][2]*string{{&f, nil}, {&t, nil}}, nil)
	case 16:
		arrayBool := [2]bool{true, false}
		return newTest(num, `simple [][2]bool -> [][2]string conversion test`, toTest(ToSlice[[2]string]([][2]bool{arrayBool})), ToSliceValue[[2]string]([][2]bool{arrayBool}), [][2]string{{`true`, `false`}}, nil)
	case 17:
		arrayBool := [2]bool{false, true}
		tf := `false true`
		return newTest(num, `simple [2]bool -> [][2]*string conversion test`, toTest(ToSlice[[2]*string](&arrayBool)), ToSliceValue[[2]*string](&arrayBool), [][2]*string{{&tf, nil}}, nil)
	case 18:
		arrayBool := [2]bool{true, false}
		return newTest(num, `simple *[2]bool -> [][2]string conversion test`, toTest(ToSlice[[2]string](&arrayBool)), ToSliceValue[[2]string](&arrayBool), [][2]string{{`true false`, ``}}, nil)
	case 19:
		var r [][2]string
		return newTest(num, `simple nil -> [][2]string conversion test`, toTest(ToSlice[[2]string](nil)), ToSliceValue[[2]string](nil), r, nil)
	}
	return testStructure{}, false
}

func TestToSlice(t *testing.T) {
	testByLists(t, []testListStructure{
		{
			Name:           `General tests`,
			FuncionGetList: getElementForTestSlice,
		},
	})
}
