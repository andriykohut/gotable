gotable
=======

ASCII tables from slices of maps

## Installation
`go get github.com/andriykohut/gotable`
## Usage
```go
package main

import (
	"fmt"
	"github.com/andriykohut/gotable"
)

func main() {
	t := gotable.NewTable(
	  // Table rows as slice of maps
		[]map[string]string{
			{"abc": "123", "qwerty": "12", "bsod": "bsod"},
			{"abc": "bump", "qwerty": "qwerty", "bsod": "kernel"},
		},
	)
	fmt.Println(t.GetTable())
}
```
will produce
```
+------+--------+--------+
| abc  | qwerty | bsod   |
+------+--------+--------+
| 123  | 12     | bsod   |
+------+--------+--------+
| bump | qwerty | kernel |
+------+--------+--------+
```
