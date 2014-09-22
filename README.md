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
will produce:
```
+------+--------+--------+
| abc  | qwerty | bsod   |
+------+--------+--------+
| 123  | 12     | bsod   |
+------+--------+--------+
| bump | qwerty | kernel |
+------+--------+--------+
```
`gotable.NewTable` accepts additional arguments:
* `[]string` - table headers
* `[]rune` - separators (default is `[]rune{'+', '-', '|'}`)
* `bool` - enable/disable bold headers
Example:
```go
t := gotable.NewTable(
	// Table rows
	[]map[string]string{
		{"abc": "123", "qwerty": "12", "bsod": "bsod"},
		{"abc": "bump", "qwerty": "qwerty", "bsod": "kernel"},
	},
	// header is not bold
	false,
	// Table headers
	[]string{"qwerty", "bsod"},
	// Separators
	[]rune{'*', '=', ']'},
)
fmt.Println(t.GetTable())

```
will produce:
```
*========*========*
] qwerty ] bsod   ]
*========*========*
] 12     ] bsod   ]
*========*========*
] qwerty ] kernel ]
*========*========*
```
