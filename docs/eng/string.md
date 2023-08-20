# Typo ToString

---

[ eng ] / [ ru ](..%2Fru%2Fstring.md)
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[/](..%2F..%2FREADME.md) string

---



### String conversion function

<table>
    <tr>
        <th>Returning errors</th>
        <th>No error return</th>
    </tr>
    <tr>
        <th colspan="2">

#### Functions
</th>
    </tr>
    <tr>
        <td>

`ToString(value any) (result string, err error)`
<div><i>alias:</i> `To[string](value any) (result string, err error)` </div>
        </td>
        <td>

`ToStringValue(value any) (result string)`
<div><i>alias:</i> `ToValue[string](value any) (result string)` </div>
        </td>
    </tr>
    <tr>
        <th colspan="2">

#### Description
</th>
    </tr>
    <tr>
        <td colspan="2">

Converts the given value to a string. In case of an error, an empty string and an error with the reasons for the 
conversion problem are returned. When passed to the `nil` function, it will return the value specified 
in [SetStringToWriteNull](settings.md)
</td>
    </tr>
    <tr>
        <th colspan="2">

#### Example
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

##### Execution result
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