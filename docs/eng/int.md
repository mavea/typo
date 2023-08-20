# Typo ToInt

---

[ eng ] / [ ru ](..%2Fru%2Fint.md)
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[/](..%2F..%2FREADME.md) int

---



## Integer conversion functions

### Function to convert to integer int

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

`ToInt(value any) (result int, err error)`
<div><i>alias:</i> `To[int](value any) (result int, err error)` </div>
    </td>
<td>

`ToIntValue(value any) (result int)`
<div><i>alias:</i> `ToValue[int](value any) (result int)` </div>
        </td>
    </tr>
    <tr>
        <th colspan="2">

#### Description
</th>
    </tr>
    <tr>
        <td colspan="2">

Converts the given value to an integer. In case of an error, 0 is returned and an error with the causes of the problem 
during conversion
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


### 8-bit integer conversion function

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

`ToInt8(value any) (result int8, err error)`
<div><i>alias:</i> `To[int8](value any) (result int8, err error)` </div>
        </td>
        <td>

`ToInt8Value(value any) (result int8)`
<div><i>alias:</i> `ToValue[int8](value any) (result int8)` </div>
        </td>
    </tr>
    <tr>
        <th colspan="2">

#### Description
</th>
    </tr>
    <tr>
        <td colspan="2">

Converts the given value to an 8-bit integer. In case of an error, 0 is returned and an error with the causes of the 
problem during conversion
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


### ### 16-bit integer conversion function

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

`ToInt16(value any) (result int16, err error)`
<div><i>alias:</i> `To[int16](value any) (result int16, err error)` </div>
        </td>
        <td>

`ToInt16Value(value any) (result int16)`
<div><i>alias:</i> `ToValue[int16](value any) (result int16)` </div>
        </td>
    </tr>
    <tr>
        <th colspan="2">

#### Description
</th>
    </tr>
    <tr>
        <td colspan="2">

Converts the given value to a 16-bit integer. In case of an error, 0 is returned and an error with the causes of the 
problem during conversion
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

### Function to convert to 32 bit integer

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

`ToInt32(value any) (result int32, err error)`
<div><i>alias:</i> `To[int32](value any) (result int32, err error)` </div>
        </td>
        <td>

`ToInt32Value(value any) (result int32)`
<div><i>alias:</i> `ToValue[int32](value any) (result int32)` </div>
        </td>
    </tr>
    <tr>
        <th colspan="2">

#### Description
</th>
    </tr>
    <tr>
        <td colspan="2">

Converts the given value to a 32-bit integer. In case of an error, 0 is returned and an error with the causes of the 
problem during conversion
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


### Function to convert to 64 bit integer

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

`ToInt64(value any) (result int64, err error)`
<div><i>alias:</i> `To[int64](value any) (result int64, err error)` </div>
        </td>
        <td>

`ToInt64Value(value any) (result int64)`
<div><i>alias:</i> `ToValue[int64](value any) (result int64)` </div>
        </td>
    </tr>
    <tr>
        <th colspan="2">

#### Description
</th>
    </tr>
    <tr>
        <td colspan="2">

Converts the given value to a 64-bit integer. In case of an error, 0 is returned and an error with the causes of the 
problem during conversion
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

##### Execution result
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



