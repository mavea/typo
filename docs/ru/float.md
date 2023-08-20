# Typo ToFloat

---

[ ru ]/ [eng](..%2Feng%2Ffloat.md)
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[/](..%2F..%2FREADME.ru.md) float

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

## Функции преобразования в число с плавающей запятой

### Функция преобразования в число одинарной точности float32

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

`ToFloat32(value any) (result float32, err error)`
<div><i>алиас:</i> `To[float32](value any) (result float32, err error)` </div>
        </td>
        <td>

`ToFloat32Value(value any) (result float32)`
<div><i>алиас:</i> `ToValue[float32](value any) (result float32)` </div>
        </td>
    </tr>
    <tr>
        <th colspan="2">

#### Описание
</th>
    </tr>
    <tr>
        <td colspan="2">

Конвертирует полученное значение в значение одинарной точности (float32). В случае ошибки возвращается 0.0 и ошибка с причинами проблемы при конвертации
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
    f, err := typo.ToFloat32(1)
    if nil != err {
        panic(err)
    }
    print(typo.ToStringValue(f))
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
    print(typo.ToStringValue(typo.ToFloat32Value(`1`)))
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
1.0
```
</td>
        <td>

```
1.0
```
</td>
    </tr>
</table>


---


### Функция преобразования в число двойной точности float64

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

`ToFloat64(value any) (result float64, err error)`
<div><i>алиас:</i> `To[float64](value any) (result float64, err error)` </div>
        </td>
        <td>

`ToFloat64Value(value any) (result float64)`
<div><i>алиас:</i> `ToValue[float64](value any) (result float64)` </div>
        </td>
    </tr>
    <tr>
        <th colspan="2">

#### Описание
</th>
    </tr>
    <tr>
        <td colspan="2">

Конвертирует полученное значение в значение двойной точности (float64). В случае ошибки возвращается 0.0 и ошибка с причинами проблемы при конвертации
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
    f, err := typo.ToFloat64(5)
    if nil != err {
        panic(err)
    }
    print(typo.ToStringValue(f))
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
    print(typo.ToStringValue(typo.ToFloat64Value(`5.0`)))
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
5.0
```
</td>
        <td>

```
5.0
```
</td>
    </tr>
</table>