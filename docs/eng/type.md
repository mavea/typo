# Typo To[*]

---

[ eng ] / [ ru ](..%2Fru%2Ftype.md)
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[/](..%2F..%2FREADME.md) To[*]

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

## Basic type conversion function

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

`To[typeResult any](value any) (result typeResult, err error)`
<div><i>alias:</i> `ToType[typeResult any](value any) (result typeResult, err error)` </div>
</td>
        <td>

`ToValue[typeResult any](value any) (result typeResult)`
<div><i>alias:</i> `ToTypeValue[typeResult any](value any) (result typeResult)` </div>
</td>
  </tr>
  <tr>
    <th colspan="2">

#### Description
</th>
  </tr>
  <tr>
    <td colspan="2">

Converts the passed `value` data to the type specified in the generic `typeResult`. In case of an error during data 
conversion, the function will return the default value for the specified type, as well as the error with which the 
conversion ended.

This is the main function of the package, the rest are aliases of this function with predefined result data types
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

##### Execution result
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