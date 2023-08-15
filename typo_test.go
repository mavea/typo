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

type testListStructure struct {
	Name           string                          `json:"name"`
	FuncionGetList func(int) (testStructure, bool) `json:"funcion-get-list"`
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
