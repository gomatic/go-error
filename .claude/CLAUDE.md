# go-error

The ecosystem's sentinel-error mechanism: the `errs.Const` type (a string-backed error), its `Error()` method, and the `With(err, args…)` wrapper. **Owns the mechanism only — never error values.** Every consumer declares its own `const ErrX errs.Const = "…"` in its own repo, matchable with `errors.Is`.

- Package is named `errs` (the suffix `error` is a predeclared identifier, so it cannot be the package name); the type is `Const` to avoid stutter (`errs.Const`). Generic — lives in `gomatic`.
- CLI-agnostic, dependency-free (testify for tests). Gate: gofumpt, vet, staticcheck, govulncheck, gocognit ≤ 7, 100% coverage.
