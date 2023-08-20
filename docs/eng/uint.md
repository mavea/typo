# Typo ToUint

---

[ eng ] / [ ru ](..%2Fru%2Fuint.md)
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[/](..%2F..%2FREADME.md) uint

---



## Functions for converting to a positive integer

### uint conversion function

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

`ToUint(value any) (result uint, err error)`
<div><i>alias:</i> `To[uint](value any) (result uint, err error)` </div>
        </td>
        <td>

`ToUintValue(value any) (result uint)`
<div><i>alias:</i> `ToValue[uint](value any) (result uint)` </div>
        </td>
    </tr>
    <tr>
        <th colspan="2">

#### Description
</th>
    </tr>
    <tr>
        <td colspan="2">

Converts the given value to an integer. In case of an error, 0 is returned and an error with the causes of the problem during conversion
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

##### Execution result
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


### 8-bit positive integer conversion function

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

`ToUint8(value any) (result uint8, err error)`
<div><i>alias:</i> `To[uint8](value any) (result uint8, err error)` </div>
        </td>
        <td>

`ToUint8Value(value any) (result uint8)`
<div><i>alias:</i> `ToValue[uint8](value any) (result uint8)` </div>
        </td>
    </tr>
    <tr>
        <th colspan="2">

#### Description
</th>
    </tr>
    <tr>
        <td colspan="2">

Converts the given value to an 8-bit integer. In case of an error, 0 is returned and an error with the causes of the problem during conversion
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

##### Execution result
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


### 16-bit positive integer conversion function

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

`ToUint16(value any) (result uint16, err error)`
<div><i>alias:</i> `To[uint16](value any) (result uint16, err error)` </div>
        </td>
        <td>

`ToUint16Value(value any) (result uint16)`
<div><i>alias:</i> `ToValue[uint16](value any) (result uint16)` </div>
        </td>
    </tr>
    <tr>
        <th colspan="2">

#### Description
</th>
    </tr>
    <tr>
        <td colspan="2">

Converts the given value to a 16-bit integer. In case of an error, 0 is returned and an error with the causes of the problem during conversion
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

##### Execution result
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

### 32 bit positive integer conversion function

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

`ToUint32(value any) (result uint32, err error)`
<div><i>alias:</i> `To[uint32](value any) (result uint32, err error)` </div>
        </td>
        <td>

`ToUint32Value(value any) (result uint32)`
<div><i>alias:</i> `ToValue[uint32](value any) (result uint32)` </div>
        </td>
    </tr>
    <tr>
        <th colspan="2">

#### Description
</th>
    </tr>
    <tr>
        <td colspan="2">

Converts the given value to a 32-bit integer. In case of an error, 0 is returned and an error with the causes of the problem during conversion
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

##### Execution result
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


### 64-bit positive integer conversion function

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

`ToUint64(value any) (result uint64, err error)`
<div><i>alias:</i> `To[uint64](value any) (result uint64, err error)` </div>
        </td>
        <td>

`ToUint64Value(value any) (result uint64)`
<div><i>alias:</i> `ToValue[uint64](value any) (result uint64)` </div>
        </td>
    </tr>
    <tr>
        <th colspan="2">

#### Description
</th>
    </tr>
    <tr>
        <td colspan="2">

Converts the given value to a 64-bit integer. In case of an error, 0 is returned and an error with the causes of the problem during conversion
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

##### Execution result
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


### Convert function to Uintptr

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

`ToUintptr(value any) (result uintptr, err error)`
<div><i>alias:</i> `To[uintptr](value any) (result uintptr, err error)` </div>
        </td>
        <td>

`ToUintptrValue(value any) (result uintptr)`
<div><i>alias:</i> `ToValue[uintptr](value any) (result uintptr)` </div>
        </td>
    </tr>
    <tr>
        <th colspan="2">

#### Description
</th>
    </tr>
    <tr>
        <td colspan="2">

Converts the given value to a uintptr. In case of an error, 0 is returned and an error with the causes of the problem during conversion
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

##### Execution result
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



