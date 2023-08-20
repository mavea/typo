# Typo To[*]

---

[ ru ]/ [eng](..%2Feng%2Ftype.md)
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[/](..%2F..%2FREADME.ru.md) To[*]

---

<style>
table {
    padding: 0;
    margin: 0;
}
td {
    padding: 0;
    vertical-align: top;
}
th {
    vertical-align: top;
    text-align: left;
    padding: 0;
}
tr {
    padding: 0;
    margin: 0;
}
td div {
    text-align: right;
}
</style>

## Основная функция преобразования типов

<table>
  <tr>
    <th>С возвратом ошибок</th>
    <th>Без возврата ошибок</th>
  </tr>
  <tr>
    <th colspan="2">

#### Функции
</th>
  </tr>
  <tr>
    <td>

`To[typeResult any](value any) (result typeResult, err error)`
<div><i>алиас:</i> `ToType[typeResult any](value any) (result typeResult, err error)` </div>
</td>
    <td>

`ToValue[typeResult any](value any) (result typeResult)`
<div><i>алиас:</i> `ToTypeValue[typeResult any](value any) (result typeResult)` </div>
</td>
  </tr>
  <tr>
    <th colspan="2">

#### Описание
</th>
  </tr>
  <tr>
    <td colspan="2">

Преобразует переданные данные `value` в указанный в дженерике тип `typeResult`. В случае ошибки при конвертации данных, 
функция вернёт значение по умолчанию для указанного типа, а так же ошибку с которой закончилась конвертация.

Это основная функция пакета, остальные являются алиасами данной функции с предустановленными типами данных результата
вместо дженерика.
</td>
  </tr>
  <tr>
    <th colspan="2">

#### Пример
</th>
  </tr>
  <tr>
    <td>

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
    s, err = typo.To[int](i)
    if nil != err {
        panic(err)
    }
    print(s)
}
```
</td>
        <td>

```go
package main

import (
typo "github.com/mavea/typo"
)

func main() {
i := typo.ToValue[int](`1`)
i++
print(typo.ToValue[int](i))
}
```
</td>
  </tr>
  <tr>
    <th colspan="2">

##### Результат исполнения
</th>
  </tr>
  <tr>
    <td>

```
2
```
</td>
    <td>

```
2
```
</td>
  </tr>
</table>