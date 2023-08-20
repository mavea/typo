# Typo ToFloat

---

[ eng ] / [ ru ](..%2Fru%2Ffloat.md)
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[/](..%2F..%2FREADME.md) float

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

## Functions for converting to floating point

### Function to convert to single precision float32

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

`ToFloat32(value any) (result float32, err error)`
<div><i>alias:</i> `To[float32](value any) (result float32, err error)` </div>
        </td>
        <td>

`ToFloat32Value(value any) (result float32)`
<div><i>alias:</i> `ToValue[float32](value any) (result float32)` </div>
        </td>
    </tr>
    <tr>
        <th colspan="2">

#### Description
</th>
    </tr>
    <tr>
        <td colspan="2">

Converts the given value to a single precision value (float32). In case of an error, 0.0 is returned and an error with 
the causes of the problem during conversion
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

##### Execution result
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


### Function to convert to double precision float64

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

`ToFloat64(value any) (result float64, err error)`
<div><i>alias:</i> `To[float64](value any) (result float64, err error)` </div>
        </td>
        <td>

`ToFloat64Value(value any) (result float64)`
<div><i>alias:</i> `ToValue[float64](value any) (result float64)` </div>
        </td>
    </tr>
    <tr>
        <th colspan="2">

#### Description
</th>
    </tr>
    <tr>
        <td colspan="2">

Converts the given value to a single precision value (float64). In case of an error, 0.0 is returned and an error with 
the causes of the problem during conversion
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

##### Execution result
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