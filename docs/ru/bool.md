
# Typo ToBool

---

[ ru ]/ [eng](..%2Feng%2Fbool.md)
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[/](..%2F..%2FREADME.ru.md) bool

---


## Функция преобразования в Bool

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

`ToBool(value any) (result bool, err error)`
<div><i>алиас:</i> `To[bool](value any) (result bool, err error)` </div>
        </td>
        <td>

`ToBoolValue(value any) (result bool)`
<div><i>алиас:</i> `ToValue[bool](value any) (result bool)` </div>
        </td>
    </tr>
    <tr>
        <th colspan="2">

#### Описание
</th>
    </tr>
    <tr>
        <td colspan="2">

Преобразует переданные данные `value` в переменную типа bool. Возвращает false и сообщение об ошибке в случае ошибки конвертации
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
    b, err := typo.ToBool(1)
    if nil != err {
        panic(err)
    }
    if b {
        print(`true`)
    } else {
        print(`false`)
    }
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
    if typo.ToBoolValue(`0`) {
        print(`true`)
    } else {
        print(`false`)
    }
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
true
```
</td>
        <td>

```
false
```
</td>
    </tr>
</table>