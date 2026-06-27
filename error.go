// Package error provides the ecosystem's sentinel-error mechanism: a single
// string-backed error type whose constants are matchable with errors.Is, never
// by string comparison. This library owns the mechanism; every consumer declares
// its own error values as constants of [Const] and keeps them in its own repo.
//
// Because the package is named error, import it bare in a dedicated sentinel file
// (which declares consts and uses no builtin error type), so declarations read
// error.Const; files that also use the builtin error type should alias the import.
package error

import (
	"fmt"
	"strings"
)

// Const is the sentinel-error type. Declare every error a package can emit as a
// const of this type so each path is matchable with errors.Is instead of by
// string comparison.
type Const string

// Error returns the constant's text, implementing the error interface.
func (e Const) Error() string { return string(e) }

// With wraps a cause and appends contextual args, returning a new error that
// still matches the sentinel (and the cause) under errors.Is. A non-nil cause is
// joined with %w so both are recoverable. Args render space-separated, so callers
// pass clean key/value pairs — err.With(cause, "key", value) — without baking
// separators into the key.
func (e Const) With(err error, args ...any) error {
	out := error(e)
	if err != nil {
		out = fmt.Errorf("%w: %w", e, err)
	}
	if len(args) > 0 {
		out = fmt.Errorf("%w: %s", out, strings.TrimSuffix(fmt.Sprintln(args...), "\n"))
	}
	return out
}

var _ error = Const("")
