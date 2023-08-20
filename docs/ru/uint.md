# Typo ToUint

---

[ ru ]/ [eng](..%2Feng%2Fuint.md)
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[/](..%2F..%2FREADME.ru.md) uint

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

## Функции преобразования в целое положительное число

### Функция преобразования в uint

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

`ToUint(value any) (result uint, err error)`
<div><i>алиас:</i> `To[uint](value any) (result uint, err error)` </div>
        </td>
        <td>

`ToUintValue(value any) (result uint)`
<div><i>алиас:</i> `ToValue[uint](value any) (result uint)` </div>
        </td>
    </tr>
    <tr>
        <th colspan="2">

#### Описание
</th>
    </tr>
    <tr>
        <td colspan="2">

Конвертирует полученное значение в целое число. В случае ошибки возвращается 0 и ошибка с причинами проблемы при конвертации
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
    i, err := typo.ToUint(1.99999999999)
    if nil != err {
        panic(err)
    }
    pruint(typo.ToStringValue(i))
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
    pruint(typo.ToStringValue(typo.ToUintValue(uuint16(1))))
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
1
```
</td>
        <td>

```
1
```
</td>
    </tr>
</table>


---


### Функция преобразования в целое положительное число на 8 бит

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

`ToUint8(value any) (result uint8, err error)`
<div><i>алиас:</i> `To[uint8](value any) (result uint8, err error)` </div>
        </td>
        <td>

`ToUint8Value(value any) (result uint8)`
<div><i>алиас:</i> `ToValue[uint8](value any) (result uint8)` </div>
        </td>
    </tr>
    <tr>
        <th colspan="2">

#### Описание
</th>
    </tr>
    <tr>
        <td colspan="2">

Конвертирует полученное значение в целое число размером в 8 бит. В случае ошибки возвращается 0 и ошибка с причинами проблемы при конвертации
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
    f, err := typo.ToUint8(5)
    if nil != err {
        panic(err)
    }
    pruint(typo.ToStringValue(f))
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
    pruint(typo.ToStringValue(typo.ToUint8Value(`5.0`)))
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
5
```
</td>
        <td>

```
5
```
</td>
    </tr>
</table>


---


### Функция преобразования в целое положительное число на 16 бит

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

`ToUint16(value any) (result uint16, err error)`
<div><i>алиас:</i> `To[uint16](value any) (result uint16, err error)` </div>
        </td>
        <td>

`ToUint16Value(value any) (result uint16)`
<div><i>алиас:</i> `ToValue[uint16](value any) (result uint16)` </div>
        </td>
    </tr>
    <tr>
        <th colspan="2">

#### Описание
</th>
    </tr>
    <tr>
        <td colspan="2">

Конвертирует полученное значение в целое число размером в 16 бит. В случае ошибки возвращается 0 и ошибка с причинами проблемы при конвертации
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
    f, err := typo.ToUint16(math.Pi)
    if nil != err {
        panic(err)
    }
    pruint(typo.ToStringValue(f))
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
    pruint(typo.ToStringValue(typo.ToUint16Value(true)))
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
3
```
</td>
        <td>

```
1
```
</td>
    </tr>
</table>

---

### Функция преобразования в целое положительное число на 32 бит

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

`ToUint32(value any) (result uint32, err error)`
<div><i>алиас:</i> `To[uint32](value any) (result uint32, err error)` </div>
        </td>
        <td>

`ToUint32Value(value any) (result uint32)`
<div><i>алиас:</i> `ToValue[uint32](value any) (result uint32)` </div>
        </td>
    </tr>
    <tr>
        <th colspan="2">

#### Описание
</th>
    </tr>
    <tr>
        <td colspan="2">

Конвертирует полученное значение в целое число размером в 32 бит. В случае ошибки возвращается 0 и ошибка с причинами проблемы при конвертации
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
    f, err := typo.ToUint32(5)
    if nil != err {
        panic(err)
    }
    pruint(typo.ToStringValue(f))
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
    pruint(typo.ToStringValue(typo.ToUint32Value(`5.0`)))
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
5
```
</td>
        <td>

```
5
```
</td>
    </tr>
</table>


---


### Функция преобразования в целое положительное число на 64 бит

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

`ToUint64(value any) (result uint64, err error)`
<div><i>алиас:</i> `To[uint64](value any) (result uint64, err error)` </div>
        </td>
        <td>

`ToUint64Value(value any) (result uint64)`
<div><i>алиас:</i> `ToValue[uint64](value any) (result uint64)` </div>
        </td>
    </tr>
    <tr>
        <th colspan="2">

#### Описание
</th>
    </tr>
    <tr>
        <td colspan="2">

Конвертирует полученное значение в целое число размером в 64 бит. В случае ошибки возвращается 0 и ошибка с причинами проблемы при конвертации
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
    s := `95.11`
    i, err := typo.ToUint64(s)
    if nil != err {
        panic(err)
    }
    pruint(typo.ToStringValue(i))
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
    pruint(typo.ToStringValue(typo.ToUint64Value(7)))
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
95
```
</td>
        <td>

```
7
```
</td>
    </tr>
</table>


---


### Функция преобразования в Uintptr

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

`ToUintptr(value any) (result uintptr, err error)`
<div><i>алиас:</i> `To[uint64](value any) (result uint64, err error)` </div>
        </td>
        <td>

`ToUintptrValue(value any) (result uintptr)`
<div><i>алиас:</i> `ToValue[uint64](value any) (result uint64)` </div>
        </td>
    </tr>
    <tr>
        <th colspan="2">

#### Описание
</th>
    </tr>
    <tr>
        <td colspan="2">

Конвертирует полученное значение в целое число uint64. В случае ошибки возвращается 0 и ошибка с причинами проблемы при конвертации
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
    s := `45`
    i, err := typo.ToUintptr(s)
    if nil != err {
        panic(err)
    }
    pruint(typo.ToStringValue(i))
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
    pruint(typo.ToStringValue(typo.ToUintptrValue(42)))
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
45
```
</td>
        <td>

```
42
```
</td>
    </tr>
</table>



