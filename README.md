# Any

The interface type that specifies zero methods is known as the empty interface:

```
interface{}
```
An empty interface may hold values of any type. (Every type implements at least zero methods.)

Empty interfaces are used by code that handles values of unknown type. For example, fmt.Print takes any number of arguments of type `interface{}`.


## Example

```golang=
package main

import (
	"fmt"

	any "github.com/knowlet/go-any-type/pkg"
)

func main() {
	var i any.Any
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i any.Any) {
	fmt.Printf("(%v, %T)\n", i, i)
}

```

Output:
```
(<nil>, <nil>)
(42, int)
(hello, string)
```
