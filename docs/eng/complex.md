# Typo ToComplex6

---

[ eng ] / [ ru ](..%2Fru%2Fcomplex.md)
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[/](..%2F..%2FREADME.md) complex
---``



## Functions for converting to a complex number

### complex64 conversion function

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

`ToComplex64(value any) (result complex64, err error)`
<div><i>alias:</i> `To[complex64](value any) (result complex64, err error)` </div>
        </td>
        <td>

`ToComplex64Value(value any) (result complex64)`
<div><i>alias:</i> `ToValue[complex64](value any) (result complex64)` </div>
        </td>
    </tr>
    <tr>
        <th colspan="2">

#### Description
</th>
    </tr>
    <tr>
        <td colspan="2">

Converts the given value to a complex number (complex64) based on the first 2 elements in the given map, slice, array, 
or structure. In case of a conversion error, it will return a value from 2 zeros, as well as a description of the error 
that occurred.
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

##### Execution result
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


### complex128 conversion function

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

`ToComplex128(value any) (result complex128, err error)`
<div><i>alias:</i> `To[complex128](value any) (result complex128, err error)` </div>
        </td>
        <td>

`ToComplex128Value(value any) (result complex128)`
<div><i>alias:</i> `ToValue[complex128](value any) (result complex128)` </div>
        </td>
    </tr>
    <tr>
        <th colspan="2">

#### Description
</th>
    </tr>
    <tr>
        <td colspan="2">

Converts the given value to a complex number (complex128) based on the first 2 elements in the given map, slice, array, 
or structure. In case of a conversion error, it will return a value from 2 zeros, as well as a description of the error 
that occurred.
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

##### Execution result
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