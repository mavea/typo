# Typo
eng / [ru](README.ru.md)

---

A package for converting data from one type to another.

```go
package main

import (
    typo "github.com/mavea/typo"
)

func main() {
    s := `1`
    i, err := typo.To[int](s)
    if nil != err {
        panic(err)
    }
    i++
    s = typo.ToValue[int](i)
    print(s)
}
```
execution result
```
2
```

---

## Why is this needed?
One way or another, we constantly have to translate the type of variables from one to another, in addition, sometimes 
we do not know the type of incoming data, but only the type into which we need to convert this data.
An example is the data that came to us in the form of a json object. This library not only converts data to the required 
type, but also avoids golang-panic if the type suddenly turned out to be a non-digestible type, but formats an error in 
the error variable, which it returns as needed

---

## Main functions

### Settings
* `SetStringToWriteNull(str string)`
* `SetStructureTagContainingNameTheElementJson(name string)`
* `SetSpecialStringsInStructureTag(words []string)`

[description](docs%2Feng%2Fsettings.md)

---

### Main conversion function
* `To[typeResult any](value any) (result typeResult, err error)`
* `ToType[typeResult any](value any) (result typeResult, err error)`
* `ToValue[typeResult any](value any) (result typeResult)`
* `ToTypeValue[typeResult any](value any) (result typeResult)`

[description](docs%2Feng%2Ftype.md)

---

## Type Specified Functions

### Convert to bool
* `ToBool(value any) (result bool, err error)`
* `ToBoolValue(value any) (result bool)`

[description](docs%2Feng%2Fbool.md)

---

### Convert to complex
* `ToComplex64(value any) (result complex64, err error)`
* `ToComplex128(value any) (result complex128, err error)`
* `ToComplex64Value(value any) (result complex64)`
* `ToComplex128Value(value any) (result complex128)`

[description](docs%2Feng%2Fcomplex.md)

---

### Convert to float
* `ToFloat32(value any) (result float32, err error)`
* `ToFloat64(value any) (result float64, err error)`
* `ToFloat32Value(value any) (result float32)`
* `ToFloat64Value(value any) (result float64)`

[description](docs%2Feng%2Ffloat.md)

---

### Convert to int
* `ToInt(value any) (result int, err error)`
* `ToInt8(value any) (result int8, err error)`
* `ToInt16(value any) (result int16, err error)`
* `ToInt32(value any) (result int32, err error)`
* `ToInt64(value any) (result int64, err error)`
* `ToIntValue(value any) (result int)`
* `ToInt8Value(value any) (result int8)`
* `ToInt16Value(value any) (result int16)`
* `ToInt32Value(value any) (result int32)`
* `ToInt64Value(value any) (result int64)`

[description](docs%2Feng%2Fint.md)

---

### Convert to slice
* `ToSlice[typeResult any](value any) (result []typeResult, err error)`
* `ToSliceValue[typeResult any](value any) (result []typeResult)`

[description](docs%2Feng%2Fslice.md)

---

### Convert to string
* `ToString(value any) (result string, err error)`
* `ToStringValue(value any) (result string)`

[description](docs%2Feng%2Fstring.md)

---

### Convert to uint
* `ToUint(value any) (result uint, err error)`
* `ToUint8(value any) (result uint8, err error)`
* `ToUint16(value any) (result uint16, err error)`
* `ToUint32(value any) (result uint32, err error)`
* `ToUint64(value any) (result uint64, err error)`
* `ToUintptr(value any) (result uintptr, err error)`
* `ToUintValue(value any) (result uint)`
* `ToUint8Value(value any) (result uint8)`
* `ToUint16Value(value any) (result uint16)`
* `ToUint32Value(value any) (result uint32)`
* `ToUint64Value(value any) (result uint64)`
* `ToUintptrValue(value any) (result uintptr)`

[description](docs%2Feng%2Fuint.md)
