[![Go Reference](https://pkg.go.dev/badge/github.com/dekarrin/assertion.svg)](https://pkg.go.dev/github.com/dekarrin/assertion)

# assertion
This is a library that is used for testing in golang.

All assertions are performed by an Asserter which wraps a *testing.T and calls
failure functions on it when assertions made by the Asserter fail.

There are not currently any tests for the assertion library itself.

## Usage
To use, `import "github.com/dekarrin/assertion"` and then create an `Asserter`
with `New`. The Asserter can then have assertions called on it.
