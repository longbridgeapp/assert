# Assert

[![Go](https://github.com/longbridgeapp/assert/actions/workflows/go.yml/badge.svg)](https://github.com/longbridgeapp/assert/actions/workflows/go.yml)

> Requirement Go 1.18+, lower Go version please use 0.x

Extends [stretchr/testify/assert](https://github.com/stretchr/testify/tree/master/assert) for add more useful methods.

- `assert.Equal` - asserts values equal, but ignore type.
- `assert.EqualHTML` - asserts HTML equal, ignore spaces.

## Installation

```bash
go get github.com/longbridgeapp/assert
```

## Usage

```go
package some_test

import (
  "github.com/longbridgeapp/assert"
)

func TestSomeMethod(t *testing.T) {
	assert.EqualHTML(t, "<p>Hello world</p><p>This is next line</p>", `
		<p>Hello world</p>
		<p>This is next line</p>
	`)
}
```

## License

MIT
