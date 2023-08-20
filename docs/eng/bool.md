# Typo ToBool

---

[ eng ] / [ ru ](..%2Fru%2Fbool.md)
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[/](..%2F..%2FREADME.md) bool

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

## Convert function to bool

<table>
    <tr>
        <th>With error return</th>
        <th>No error return</th>
    </tr>
    <tr>
        <th colspan="2">

#### Functions
</th>
    </tr>
    <tr>
        <td>

`ToBool(value any) (result bool, err error)`
<div><i>alias:</i> `To[bool](value any) (result bool, err error)` </div>
        </td>
        <td>

`ToBoolValue(value any) (result bool)`
<div><i>alias:</i> `ToValue[bool](value any) (result bool)` </div>
        </td>
    </tr>
    <tr>
        <th colspan="2">

#### Description
</th>
    </tr>
    <tr>
        <td colspan="2">

Converts the given `value` data to a variable of type bool. Returns false and an error message in case of a conversion 
error
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

##### Execution result
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