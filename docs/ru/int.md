# Typo ToInt

---

[ ru ]/ [eng](..%2Feng%2Fint.md)
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[/](..%2F..%2FREADME.ru.md) int

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

## Функции преобразования в целое число

### Функция преобразования в целое число int

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

`ToInt(value any) (result int, err error)`
<div><i>алиас:</i> `To[int](value any) (result int, err error)` </div>
        </td>
        <td>

`ToIntValue(value any) (result int)`
<div><i>алиас:</i> `ToValue[int](value any) (result int)` </div>
        </td>
    </tr>
    <tr>
        <th colspan="2">

#### Описание
</th>
    </tr>
    <tr>
        <td colspan="2">

Конвертирует полученное значение в целое число. В случае ошибки возвращается 0 и ошибку с причинами проблемы при конвертации
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
    i, err := typo.ToInt(1.99999999999)
    if nil != err {
        panic(err)
    }
    print(typo.ToStringValue(i))
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
    print(typo.ToStringValue(typo.ToIntValue(uint16(1))))
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


### Функция преобразования в целое число на 8 бит

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

`ToInt8(value any) (result int8, err error)`
<div><i>алиас:</i> `To[int8](value any) (result int8, err error)` </div>
        </td>
        <td>

`ToInt8Value(value any) (result int8)`
<div><i>алиас:</i> `ToValue[int8](value any) (result int8)` </div>
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
    f, err := typo.ToInt8(5)
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
    print(typo.ToStringValue(typo.ToInt8Value(`5.0`)))
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


### Функция преобразования в целое число на 16 бит

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

`ToInt16(value any) (result int16, err error)`
<div><i>алиас:</i> `To[int16](value any) (result int16, err error)` </div>
        </td>
        <td>

`ToInt16Value(value any) (result int16)`
<div><i>алиас:</i> `ToValue[int16](value any) (result int16)` </div>
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
    f, err := typo.ToInt16(math.Pi)
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
    print(typo.ToStringValue(typo.ToInt16Value(true)))
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

### Функция преобразования в целое число на 32 бит

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

`ToInt32(value any) (result int32, err error)`
<div><i>алиас:</i> `To[int32](value any) (result int32, err error)` </div>
        </td>
        <td>

`ToInt32Value(value any) (result int32)`
<div><i>алиас:</i> `ToValue[int32](value any) (result int32)` </div>
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
    f, err := typo.ToInt32(5)
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
    print(typo.ToStringValue(typo.ToInt32Value(`5.0`)))
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


### Функция преобразования в целое число на 64 бит

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

`ToInt64(value any) (result int64, err error)`
<div><i>алиас:</i> `To[int64](value any) (result int64, err error)` </div>
        </td>
        <td>

`ToInt64Value(value any) (result int64)`
<div><i>алиас:</i> `ToValue[int64](value any) (result int64)` </div>
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
    s := `-95.11`
    i, err := typo.ToInt64(s)
    if nil != err {
        panic(err)
    }
    print(typo.ToStringValue(i))
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
    print(typo.ToStringValue(typo.ToInt64Value(-7)))
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
-95
```
</td>
        <td>

```
-7
```
</td>
    </tr>
</table>



