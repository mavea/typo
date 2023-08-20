# Typo 
ru / [eng](README.md)

---

Пакет перевода данных из одного типа в другой.

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
результат исполнения
```
2
```

---

## Зачем это нужно?
Так или иначе, нам постоянно приходится переводить тип переменных из одного в другой, кроме того, при этом иногда мы не знаем тип
входящих данных, а только тип, в который нам эти данные нужно преобразовать.
Примером могут выступать данные, пришедшие к нам в виде json объекта. Данная библиотека не только переводит данные в необходимый тип, но и
позволяет избежать golang-паники, если тип оказался внезапно неприемлемого типа, а оформляет ошибку в error переменную, которую возвращает по необходимости

---

## Основные функции

### Настройки
* `SetStringToWriteNull(str string)`
* `SetStructureTagContainingNameTheElementJson(name string)`
* `SetSpecialStringsInStructureTag(words []string)`

[описание](docs%2Fru%2Fsettings.md)

---

### Основная функция конвертации
* `To[typeResult any](value any) (result typeResult, err error)`
* `ToType[typeResult any](value any) (result typeResult, err error)`
* `ToValue[typeResult any](value any) (result typeResult)`
* `ToTypeValue[typeResult any](value any) (result typeResult)`

[описание](docs%2Fru%2Ftype.md)

---

## Функции с указанием типа

### Конвертация в bool
* `ToBool(value any) (result bool, err error)`
* `ToBoolValue(value any) (result bool)`

[описание](docs%2Fru%2Fbool.md)

---

### Конвертация в complex
* `ToComplex64(value any) (result complex64, err error)`
* `ToComplex128(value any) (result complex128, err error)`
* `ToComplex64Value(value any) (result complex64)`
* `ToComplex128Value(value any) (result complex128)`

[описание](docs%2Fru%2Fcomplex.md)

---

### Конвертация в float
* `ToFloat32(value any) (result float32, err error)`
* `ToFloat64(value any) (result float64, err error)`
* `ToFloat32Value(value any) (result float32)`
* `ToFloat64Value(value any) (result float64)`

[описание](docs%2Fru%2Ffloat.md)

---

### Конвертация в int
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

[описание](docs%2Fru%2Fint.md)

---

### Конвертация в slice
* `ToSlice[typeResult any](value any) (result []typeResult, err error)`
* `ToSliceValue[typeResult any](value any) (result []typeResult)`

[описание](docs%2Fru%2Fslice.md)

---

### Конвертация в string
* `ToString(value any) (result string, err error)`
* `ToStringValue(value any) (result string)`

[описание](docs%2Fru%2Fstring.md)

---

### Конвертация в uint
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

[описание](docs%2Fru%2Fuint.md)
