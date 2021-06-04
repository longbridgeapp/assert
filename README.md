# Assert

Extends [stretchr/testify/assert](github.com/stretchr/testify/assert) for add more useful methods.

- assert.Equal - asserts values equal, but ignore type.
- assert.EqualHTML - asserts HTML equal, ignore spaces.

## Installation

```bash
go get github.com/longbridgeapp/assert
```

## Usage

```
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
