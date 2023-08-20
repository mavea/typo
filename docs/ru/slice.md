# Typo ToSlice

---

[ ru ]/ [eng](..%2Feng%2Fslice.md)
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[/](..%2F..%2FREADME.ru.md) slice

---





### Функция преобразования в слайс

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

`ToSlice[typeResult any](value any) (result []typeResult, err error)`
<div><i>алиас:</i> `To[typeResult any](value any) (result string, err typeResult)` </div>
</td>
    <td>

`ToSliceValue[typeResult any](value any) (result []typeResult)`
<div><i>алиас:</i> `ToValue[typeResult any](value any) (result typeResult)` </div>
</td>
  </tr>
  <tr>
    <th colspan="2">

#### Описание
</th>
  </tr>
  <tr>
    <td colspan="2">

Converts the passed data into a slice format consisting of the elements passed in the generic. Returns an empty slice 
and a conversion error on error
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
    s, err := typo.ToSlice[string]([]int{1, 2, 9})
    if nil != err {
        panic(err)
    }
    // s = []string{`1`, `2`, `9`}
    print(typo.ToStringValue(s))
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
    s := ToSlice[int8](`18`)
    // s = []int8{18}
    print(typo.ToStringValue(s))
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
1 2 9
```
</td>
    <td>

```
18
```
</td>
  </tr>
</table>