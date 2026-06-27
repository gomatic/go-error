# go-error

A constant/sentinel error helper for Go: `errs.Const` is a string newtype that implements `error`, and its `.With(err, args...)` method wraps a cause with `%w` so the result stays matchable with `errors.Is` against both the sentinel and the cause.

## Install

```sh
go get github.com/gomatic/go-error
```

## Usage

```go
package main

import (
	"errors"
	"fmt"

	errs "github.com/gomatic/go-error"
)

// Declare every error a package can emit as a const of errs.Const.
const ErrFoo errs.Const = "foo failed"

func do(cause error) error {
	// Wrap the cause; the result still matches ErrFoo (and cause) under errors.Is.
	return ErrFoo.With(cause)
}

func main() {
	err := do(errors.New("disk full"))
	fmt.Println(errors.Is(err, ErrFoo)) // true
}
```

The library owns the mechanism only — it ships no error values. Every consumer declares its own `const ErrX errs.Const = "…"` in its own repo.
