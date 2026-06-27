# go-error

The ecosystem's sentinel-error mechanism: the `error.Const` type (a string-backed error), its `Error()` method, and the `With(err, args…)` wrapper. **Owns the mechanism only — never error values.** Every consumer declares its own `const ErrX error.Const = "…"` in its own repo, matchable with `errors.Is`.

- Package is named `error` (matching the `go-error` repo); import it **bare** in a dedicated sentinel file (consts only, no builtin `error` usage) so declarations read `error.Const`. Files that also use the builtin `error` type alias the import.
- CLI-agnostic, dependency-free (stdlib + testify for tests). Generic — lives in `gomatic` and is consumed by `template.cli` and the SkyKernel tools.
- Gate: gofumpt, vet, staticcheck, govulncheck, gocognit ≤ 7, 100% coverage.
