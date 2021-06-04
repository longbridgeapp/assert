package assert

import "testing"

func TestEqualHTML(t *testing.T) {
	EqualHTML(t, "<p>Hello world</p><p>This is next line</p>", `
		<p>Hello world</p>
		<p>This is next line</p>
	`)
}
