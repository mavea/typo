# Typo ToComplex6

---

[ ru ]/ [eng](..%2Feng%2Fcomplex.md)
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[/](..%2F..%2FREADME.ru.md) complex 
---``

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

## Функции преобразования в комплексное число

### Функция преобразования complex64

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

`ToComplex64(value any) (result complex64, err error)`
<div><i>алиас:</i> `To[complex64](value any) (result complex64, err error)` </div>
        </td>
        <td>

`ToComplex64Value(value any) (result complex64)`
<div><i>алиас:</i> `ToValue[complex64](value any) (result complex64)` </div>
        </td>
    </tr>
    <tr>
        <th colspan="2">

#### Описание
</th>
    </tr>
    <tr>
        <td colspan="2">

Конвертирует полученное значение в комплексное число (complex64) на основании первых 2 элементов в переданном мапе, слайсе, массиве или структуре. В случае ошибки конвертации вернёт значение от 2 нулей, а так же описание произошедшей ошибки.
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
    c, err := typo.ToComplex64([2]int{1, 2})
    if nil != err {
        panic(err)
    }
    print(typo.ToStringValue(c))
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
    print(typo.ToStringValue(typo.ToComplex64Value([2]int{1, 2})))
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
(1+2i)
```
</td>
        <td>

```
(1+2i)
```
</td>
    </tr>
</table>


---


### Функция преобразования complex128

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

`ToComplex128(value any) (result complex128, err error)`
<div><i>алиас:</i> `To[complex128](value any) (result complex128, err error)` </div>
        </td>
        <td>

`ToComplex128Value(value any) (result complex128)`
<div><i>алиас:</i> `ToValue[complex128](value any) (result complex128)` </div>
        </td>
    </tr>
    <tr>
        <th colspan="2">

#### Описание
</th>
    </tr>
    <tr>
        <td colspan="2">

Конвертирует полученное значение в комплексное число (complex128) на основании первых 2 элементов в переданном мапе, слайсе, массиве или структуре. В случае ошибки конвертации вернёт значение от 2 нулей, а так же описание произошедшей ошибки.
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
    c, err := typo.ToComplex128([2]int{1, 2})
    if nil != err {
        panic(err)
    }
    print(typo.ToStringValue(c))
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
    print(typo.ToStringValue(typo.ToComplex128Value([2]int{1, 2})))
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
(1+2i)
```
</td>
        <td>

```
(1+2i)
```
</td>
    </tr>
</table>