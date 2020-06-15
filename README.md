# Unik

A UUID without hypens

## Getting Started

A UUID package for Go. It generated a UUID without hyphens.

### Prerequisites

Go language.

### Installing

Use the go command:

```go
$ go get github.com/amitgunjal/unik
```

```go
package main

import (
	"fmt"
	"github.com/amitgunjal/unik"
)

func main() {
	v := unik.NewUUID()
	fmt.Println(v)
}
```

## Built With

* [go.uuid](https://github.com/satori/go.uuid) - The Go package used.


## Authors

* **Amit Gunjal** follow at [amitgunjal](https://github.com/amitgunjal)

## Acknowledgments

* [go.uuid](https://github.com/satori/go.uuid)
* Inspiration
* etc
