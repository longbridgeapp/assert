package assert

import (
	"regexp"
	"strings"
)

var (
	htmlSpaceRe = regexp.MustCompile(`(?m)>[\s]+<`)
)

// EqualHTML asserts two HTML equal, will ignore spaces / breaks between >  <
//
//   assert.EqualHTML(t, "<p>Hello</p> <p>world<p>", "<p>Hello</p>    <p>world<p>") // ok
//
func EqualHTML(t TestingT, exptected, actual string) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}

	exptected = strings.TrimSpace(exptected)
	actual = strings.TrimSpace(actual)

	if htmlSpaceRe.ReplaceAllString(exptected, "><") != htmlSpaceRe.ReplaceAllString(actual, "><") {
		t.Errorf("\nexptected:\n%s\nactual   :\n%s", exptected, htmlSpaceRe.ReplaceAllString(actual, "><"))
	}
}
