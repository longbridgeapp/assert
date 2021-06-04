package assert

import (
	"testing"
)

func Test_AssertEqual(t *testing.T) {
	Equal(t, 1, int(1))
	Equal(t, 123, uint(123))
	Equal(t, 321, uint16(321))
	Equal(t, int32(100), int64(100))
	Equal(t, 100, uint64(100))
}

func Test_NotEqual(t *testing.T) {
	NotEqual(t, 2, 3)
	NotEqual(t, 2, int32(3))
}
