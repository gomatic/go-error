# go-error

The ecosystem's sentinel-error mechanism: the `errs.Error` string type, its `Error()` method, and the `With(err, args…)` wrapper. **This module owns the mechanism only — never error values.** Every consumer declares its own `const ErrX errs.Error = "…"` in its own repo, matchable with `errors.Is`.

- CLI-agnostic, dependency-free (stdlib + testify for tests only). Nothing here may import a CLI or logging framework.
- Quality gate: gofumpt, `go vet`, staticcheck, govulncheck, gocognit ≤ 7, **100% coverage**.
- Consumed by `go-output`, `skytl`, `skykerneld`, `skym`, and any other Go program needing sentinels.
