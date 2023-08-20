# Typo ToString

---

[ ru ]/ [eng](..%2Feng%2Fstring.md)
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[/](..%2F..%2FREADME.ru.md) string

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

### Функция преобразования в строку

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

`ToString(value any) (result string, err error)`
<div><i>алиас:</i> `To[string](value any) (result string, err error)` </div>
        </td>
        <td>

`ToStringValue(value any) (result string)`
<div><i>алиас:</i> `ToValue[string](value any) (result string)` </div>
        </td>
    </tr>
    <tr>
        <th colspan="2">

#### Описание
</th>
    </tr>
    <tr>
        <td colspan="2">

Конвертирует полученное значение в строку. В случае ошибки возвращается пустая строка и ошибка с причинами проблемы при конвертации. При передачах в функцию `nil`, вернёт значение, которое указано в [SetStringToWriteNull](settings.md)
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
    b := []bool{true, false}
    s, err := typo.ToString(b)
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
    print(typo.ToStringValue(map[int]int{1: 12, 3: 14}))
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
true false
```
</td>
        <td>

```
{"1":12,"3":14}
```
</td>
    </tr>
</table>