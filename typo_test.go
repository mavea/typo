package typo

import (
	"encoding/json"
	"reflect"
	"strconv"
	"testing"
)

type testStructure struct {
	Number         int
	Name           string
	ToValue        any
	ToValueAndErr  testResultFunction
	ExpectedResult any
	ExpectedErr    error
}

type testResultFunction struct {
	Result any
	Err    error
}

type testListStructure struct {
	Name           string                          `json:"name"`
	FuncionGetList func(int) (testStructure, bool) `json:"funcion-get-list"`
}

func toTest(value any, err error) testResultFunction {
	return testResultFunction{
		Result: value,
		Err:    err,
	}
}

func newTest(i int, name string, toValueAndErr testResultFunction, toValue any, expectedResult any, expectedErr error) (testStructure, bool) {
	return testStructure{
		Number:         i,
		Name:           name,
		ToValue:        toValue,
		ToValueAndErr:  toValueAndErr,
		ExpectedResult: expectedResult,
		ExpectedErr:    expectedErr,
	}, true
}

func echoAboutRunGroupTests(t *testing.T, text string) {
	t.Log("\x1b[0m\t\t\t" + ` [ ` + text + ` ] `)

}

func checkTest(t *testing.T, test testStructure) {
	num := strconv.FormatInt(int64(test.Number), 10)
	if nil != test.ExpectedErr {
		if nil == test.ToValueAndErr.Err {
			t.Error("\x1b[31m ✖ \x1b[0m" + ` [ ` + num + ` ] ` + test.Name + `: expected error message not received
expected error message: ` + writeNullInString + `
received error message: ` + test.ExpectedErr.Error())
			return
		} else if test.ToValueAndErr.Err.Error() != test.ExpectedErr.Error() {
			t.Error("\x1b[31m ✖ \x1b[0m" + ` [ ` + num + ` ] ` + test.Name + `: expected and received error message did not match
expected error message: ` + test.ToValueAndErr.Err.Error() + `
received error message: ` + test.ExpectedErr.Error())
			return
		}
	} else if nil != test.ToValueAndErr.Err {
		t.Error("\x1b[31m ✖ \x1b[0m" + ` [ ` + num + ` ] ` + test.Name + `: got an unexpected error message
received error message: ` + test.ToValueAndErr.Err.Error())
		return
	}
	expectedValue := reflect.ValueOf(test.ToValue)
	expectedValue2 := reflect.ValueOf(test.ToValueAndErr.Result)
	receivedValue := reflect.ValueOf(test.ExpectedResult)
	if expectedValue.Kind() != receivedValue.Kind() || expectedValue.Type() != receivedValue.Type() {
		t.Error("\x1b[31m ✖ \x1b[0m" + ` [ ` + num + ` ] ` + test.Name + `: data kind does not match
expected values: (` + expectedValue.Kind().String() + ` - ` + expectedValue.Type().String() + `) |` + ToStringValue(test.ToValue) + `|
received values: (` + receivedValue.Kind().String() + ` - ` + receivedValue.Type().String() + `) |` + ToStringValue(test.ExpectedResult) + `|`)
		return
	}

	if expectedValue2.Kind() != receivedValue.Kind() || expectedValue2.Type() != receivedValue.Type() {
		t.Error("\x1b[31m ✖ \x1b[0m" + ` [ ` + num + ` ] ` + test.Name + `: data kind does not match
expected values: (` + expectedValue2.Kind().String() + ` - ` + expectedValue2.Type().String() + `) |` + ToStringValue(test.ToValueAndErr.Result) + `|
received values: (` + receivedValue.Kind().String() + ` - ` + receivedValue.Type().String() + `) |` + ToStringValue(test.ExpectedResult) + `|`)
		return
	}

	switch receivedValue.Kind() {
	case reflect.Slice, reflect.Map, reflect.Struct:
		expectedJson, _ := json.Marshal(test.ToValue)
		expectedJson2, _ := json.Marshal(test.ToValueAndErr.Result)
		receivedJson, _ := json.Marshal(test.ExpectedResult)

		if !reflect.DeepEqual(test.ToValue, test.ExpectedResult) {
			t.Error("\x1b[31m ✖ \x1b[0m" + ` [ ` + num + ` ] ` + test.Name + `: expected and received values did not match
expected values: (` + expectedValue.Kind().String() + ` - ` + expectedValue.Type().String() + `) |` + string(expectedJson) + `|
received values: (` + receivedValue.Kind().String() + ` - ` + receivedValue.Type().String() + `) |` + string(receivedJson) + `|`)
			return
		}
		if !reflect.DeepEqual(test.ToValueAndErr.Result, test.ExpectedResult) {
			t.Error("\x1b[31m ✖ \x1b[0m" + ` [ ` + num + ` ] ` + test.Name + `: expected and received values did not match
expected values: (` + expectedValue2.Kind().String() + ` - ` + expectedValue2.Type().String() + `) |` + string(expectedJson2) + `|
received values: (` + receivedValue.Kind().String() + ` - ` + receivedValue.Type().String() + `) |` + string(receivedJson) + `|`)
			return
		}
		break
	default:
		if test.ToValue != test.ExpectedResult {
			t.Error("\x1b[31m ✖ \x1b[0m" + ` [ ` + num + ` ] ` + test.Name + `: expected and received values did not match
expected values: (` + expectedValue.Kind().String() + ` - ` + expectedValue.Type().String() + `) |` + ToStringValue(test.ToValue) + `|
received values: (` + receivedValue.Kind().String() + ` - ` + receivedValue.Type().String() + `) |` + ToStringValue(test.ExpectedResult) + `|`)
			return
		}
		if test.ToValueAndErr.Result != test.ExpectedResult {
			t.Error("\x1b[31m ✖ \x1b[0m" + ` [ ` + num + ` ] ` + test.Name + `: expected and received values did not match
expected values: (` + expectedValue2.Kind().String() + ` - ` + expectedValue2.Type().String() + `) |` + ToStringValue(test.ToValueAndErr.Result) + `|
received values: (` + receivedValue.Kind().String() + ` - ` + receivedValue.Type().String() + `) |` + ToStringValue(test.ExpectedResult) + `|`)
			return
		}
	}

	t.Log("\x1b[32m ✔ \x1b[0m" + ` [ ` + num + ` ] ` + test.Name)
}

func testByLists(t *testing.T, funcionsGetList []testListStructure) {
	var test testStructure
	if len(funcionsGetList) > 0 {
		for list := 0; len(funcionsGetList) > list; list++ {
			itemTest := funcionsGetList[list]
			echoAboutRunGroupTests(t, itemTest.Name)
			ok := true
			for i := 0; ok; i++ {
				test, ok = itemTest.FuncionGetList(i)
				if !ok {
					break
				}
				checkTest(t, test)
			}
		}
	} else {
		t.Error("\x1b[31m ✖ \x1b[0m" + ` functions that should produce test cases were not provided to the list tester`)
	}
}

func getElementForSetStringToWriteNull(num int) (testStructure, bool) {
	switch num {
	case 0:
		return newTest(num, `simple nil conversion test`, toTest(ToString([]*string{nil})), ToStringValue([]*string{nil}), `null`, nil)
	case 1:
		SetStringToWriteNull(`nils`)
		return newTest(num, `simple null conversion test`, toTest(ToString(nil)), ToStringValue(nil), `nils`, nil)
	case 2:
		SetStringToWriteNull(`null`)
		return newTest(num, `simple repeat nil conversion test`, toTest(To[string](nil)), ToValue[string](nil), `null`, nil)
	}

	return testStructure{}, false
}

func TestSetStringToWriteNull(t *testing.T) {
	testByLists(t, []testListStructure{
		{
			Name:           `Personal tests`,
			FuncionGetList: getElementForSetStringToWriteNull,
		},
	})
}

func getElementForSetStructureTagContainingNameTheElementJson(num int) (testStructure, bool) {
	type testsStruct struct {
		Name string `json:"name" tag:"newName"`
	}
	ms := map[string]string{`name`: `Josh`, `newName`: `Mike`}
	switch num {
	case 0:
		return newTest(num, `simple tag "json" conversion test`, toTest(To[testsStruct](ms)), ToValue[testsStruct](ms), testsStruct{Name: `Josh`}, nil)
	case 1:
		SetStructureTagContainingNameTheElementJson(`tag`)
		return newTest(num, `simple tag "tag" conversion test`, toTest(To[testsStruct](ms)), ToValue[testsStruct](ms), testsStruct{Name: `Mike`}, nil)
	case 2:
		SetStructureTagContainingNameTheElementJson(`json`)
		return newTest(num, `simple repeat tag "json" conversion test`, toTest(To[testsStruct](ms)), ToValue[testsStruct](ms), testsStruct{Name: `Josh`}, nil)
	}

	return testStructure{}, false
}

func TestSetStructureTagContainingNameTheElementJson(t *testing.T) {
	testByLists(t, []testListStructure{
		{
			Name:           `Personal tests`,
			FuncionGetList: getElementForSetStructureTagContainingNameTheElementJson,
		},
	})
}

func getElementForSetSpecialStringsInStructureTag(num int) (testStructure, bool) {
	type testsStruct struct {
		Name string `json:"omitempty,name,-,omitempty,newName,-"`
	}
	ms := map[string]string{`name`: `Josh`, `newName`: `Mike`}
	switch num {
	case 0:
		return newTest(num, `simple tag "json" conversion test`, toTest(To[testsStruct](ms)), ToValue[testsStruct](ms), testsStruct{Name: `Josh`}, nil)
	case 1:
		SetSpecialStringsInStructureTag([]string{`-`, `omitempty`, `name`})
		return newTest(num, `simple new tag "json" conversion test`, toTest(To[testsStruct](ms)), ToValue[testsStruct](ms), testsStruct{Name: `Mike`}, nil)
	case 2:
		SetSpecialStringsInStructureTag([]string{`-`, `omitempty`})
		return newTest(num, `simple repeat tag "json" conversion test`, toTest(To[testsStruct](ms)), ToValue[testsStruct](ms), testsStruct{Name: `Josh`}, nil)
	}

	return testStructure{}, false
}

func TestSetSpecialStringsInStructureTag(t *testing.T) {
	testByLists(t, []testListStructure{
		{
			Name:           `Personal tests`,
			FuncionGetList: getElementForSetSpecialStringsInStructureTag,
		},
	})
}
