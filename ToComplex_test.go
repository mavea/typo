package typo

import (
	"math"
	"testing"
)

func getElementForTestComplex64(num int) (testStructure, bool) {
	switch num {
	case 0:
		return newTest(num, `simple []int{1, 2} conversion test`, toTest(ToComplex64([]int{1, 2})), ToComplex64Value([]int{1, 2}), complex64(complex(1, 2)), nil)
	case 1:
		return newTest(num, `simple complex(3, 2) conversion test`, toTest(ToComplex64(complex(3, 2))), ToComplex64Value(complex(3, 2)), complex64(complex(3, 2)), nil)
	case 2:
		return newTest(num, `simple [2]int{1, 2} conversion test`, toTest(ToComplex64([2]int{1, 2})), ToComplex64Value([2]int{1, 2}), complex64(complex(1, 2)), nil)
	case 3:
		return newTest(num, `simple [3]int{1, 2, 3} conversion test`, toTest(ToComplex64([3]int{1, 2, 3})), ToComplex64Value([3]int{1, 2, 3}), complex64(complex(1, 2)), nil)
	case 4:
		return newTest(num, `simple []string{"1", "2"} conversion test`, toTest(ToComplex64([]string{`1`, `2`})), ToComplex64Value([]string{`1`, `2`}), complex64(complex(1, 2)), nil)
	case 5:
		return newTest(num, `simple [2]string{"1", "5"} conversion test`, toTest(ToComplex64([2]string{`1`, `5`})), ToComplex64Value([2]string{`1`, `5`}), complex64(complex(1, 5)), nil)
	case 6:
		p := math.Pi
		return newTest(num, `simple []interface{}{&pi, 1.345345} conversion test`, toTest(ToComplex64([]interface{}{&p, 1.345345})), ToComplex64Value([]interface{}{&p, 1.345345}), complex64(complex(math.Pi, 1.345345)), nil)
	case 7:
		type testStruct struct {
			A string `json:"a"`
			B int32  `json:"b"`
		}
		return newTest(num, `simple testStruct conversion test`, toTest(ToComplex64(testStruct{A: `1.9876543`, B: 8})), ToComplex64Value(testStruct{A: `1.9876543`, B: 8}), complex64(complex(1.9876543, 8)), nil)
	case 8:
		return newTest(num, `simple nil conversion test`, toTest(ToComplex64(nil)), ToComplex64Value(nil), complex64(complex(0, 0)), nil)
	case 9:
		p := math.Pi
		s := []interface{}{&p, 1.345345}
		return newTest(num, `pointer []interface{}{&pi, 1.345345} conversion test`, toTest(ToComplex64(&s)), ToComplex64Value(&s), complex64(complex(math.Pi, 1.345345)), nil)
		/*
							// tests are not stable due to the unstable behavior of map
			case 10:
				return newTest(num, `simple map[string]string conversion test`, toTest(ToComplex64(map[string]string{`f`: `1`, `t`: `2`})), ToComplex64Value(map[string]string{`f`: `1`, `t`: `2`}), complex64(complex(1, 2)), nil)
			case 11:
				m := map[string]int64{`f`: 5, `t`: 2, `tr`: 3}
				return newTest(num, `simple map[string]int64 conversion test`, toTest(ToComplex64(m)), ToComplex64Value(m), complex64(complex(5, 2)), nil)
			case 12:
				m := map[string]int64{`f`: 5, `t`: 2, `tr`: 3}
				return newTest(num, `simple map[string]int64 conversion test`, toTest(ToComplex64(&m)), ToComplex64Value(&m), complex64(complex(5, 2)), nil)
			case 13:
				m := map[string]float64{`f`: math.Pi, `t`: 1.345345, `tr`: 3}
				return newTest(num, `simple map[string]float64 conversion test`, toTest(ToComplex64(&m)), ToComplex64Value(&m), complex64(complex(math.Pi, 1.345345)), nil)/**/

	}

	return testStructure{}, false
}

func getElementForTestComplex128(num int) (testStructure, bool) {
	switch num {
	case 0:
		return newTest(num, `simple []int{1, 2} conversion test`, toTest(ToComplex128([]int{1, 2})), ToComplex128Value([]int{1, 2}), complex128(complex(1, 2)), nil)
	case 1:
		return newTest(num, `simple complex(3, 2) conversion test`, toTest(ToComplex128(complex(3, 2))), ToComplex128Value(complex(3, 2)), complex128(complex(3, 2)), nil)
	case 2:
		return newTest(num, `simple [2]int{1, 2} conversion test`, toTest(ToComplex128([2]int{1, 2})), ToComplex128Value([2]int{1, 2}), complex128(complex(1, 2)), nil)
	case 3:
		return newTest(num, `simple [3]int{1, 2, 3} conversion test`, toTest(ToComplex128([3]int{1, 2, 3})), ToComplex128Value([3]int{1, 2, 3}), complex128(complex(1, 2)), nil)
	case 4:
		return newTest(num, `simple []string{"1", "2"} conversion test`, toTest(ToComplex128([]string{`1`, `2`})), ToComplex128Value([]string{`1`, `2`}), complex128(complex(1, 2)), nil)
	case 5:
		return newTest(num, `simple [2]string{"1", "5"} conversion test`, toTest(ToComplex128([2]string{`1`, `5`})), ToComplex128Value([2]string{`1`, `5`}), complex128(complex(1, 5)), nil)
	case 6:
		p := math.Pi
		return newTest(num, `simple []interface{}{&pi, 1.345345} conversion test`, toTest(ToComplex128([]interface{}{&p, 1.345345})), ToComplex128Value([]interface{}{&p, 1.345345}), complex128(complex(math.Pi, 1.345345)), nil)
	case 7:
		type testStruct struct {
			A string `json:"a"`
			B int32  `json:"b"`
		}
		return newTest(num, `simple testStruct conversion test`, toTest(ToComplex128(testStruct{A: `1.9876543`, B: 8})), ToComplex128Value(testStruct{A: `1.9876543`, B: 8}), complex128(complex(1.9876543, 8)), nil)
	case 8:
		return newTest(num, `simple nil conversion test`, toTest(ToComplex128(nil)), ToComplex128Value(nil), complex128(complex(0, 0)), nil)
	case 9:
		p := math.Pi
		s := []interface{}{&p, 1.345345}
		return newTest(num, `pointer []interface{}{&pi, 1.345345} conversion test`, toTest(ToComplex128(&s)), ToComplex128Value(&s), complex128(complex(math.Pi, 1.345345)), nil)
		/*
							// tests are not stable due to the unstable behavior of map
			case 10:
				return newTest(num, `simple map[string]string conversion test`, toTest(ToComplex128(map[string]string{`f`: `1`, `t`: `2`})), ToComplex128Value(map[string]string{`f`: `1`, `t`: `2`}), complex128(complex(1, 2)), nil)
			case 11:
				m := map[string]int64{`f`: 5, `t`: 2, `tr`: 3}
				return newTest(num, `simple map[string]int64 conversion test`, toTest(ToComplex128(m)), ToComplex128Value(m), complex128(complex(5, 2)), nil)
			case 12:
				m := map[string]int64{`f`: 5, `t`: 2, `tr`: 3}
				return newTest(num, `simple map[string]int64 conversion test`, toTest(ToComplex128(&m)), ToComplex128Value(&m), complex128(complex(5, 2)), nil)
			case 13:
				m := map[string]float64{`f`: math.Pi, `t`: 1.345345, `tr`: 3}
				return newTest(num, `simple map[string]float64 conversion test`, toTest(ToComplex128(&m)), ToComplex128Value(&m), complex128(complex(math.Pi, 1.345345)), nil)/**/

	}

	return testStructure{}, false
}

func TestToComplex64(t *testing.T) {
	testByLists(t, []testListStructure{
		{
			Name:           `Personal tests`,
			FuncionGetList: getElementForTestComplex64,
		},
	})
}

func TestToComplex128(t *testing.T) {
	testByLists(t, []testListStructure{
		{
			Name:           `Personal tests`,
			FuncionGetList: getElementForTestComplex128,
		},
	})
}
