module github.com/gomatic/go-error

go 1.26.4

require github.com/stretchr/testify v1.11.1

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rogpeppe/go-internal v1.14.1 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

// v0.3.0's tag was moved after the version was published: the module content
// no longer matches what the public checksum database recorded for it, so any
// fresh resolution fails verification with a SECURITY ERROR rather than a
// missing-dependency error. A released version can never be re-published, so
// the version is withdrawn instead of repaired. Every other released version
// was verified against the checksum database and is intact.
retract v0.3.0
