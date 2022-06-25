# ordered-form
A form that maintains its order.

## Usage
```go
package main

import (
	"github.com/CrimsonAIO/ordered-form"
)

func main() {
    form := new(orderedform.Form)
	form.Add("foo", "a")
	form.Add("bar", "b")
	form.Add("baz", "a")
	fmt.Println(form.URLEncode())
	
	// Prints:
	// foo=a&bar=b&baz=a
}
```
