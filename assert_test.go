package assert

import (
	"testing"
)

func testAssert(t *testing.T, expected bool, block func(t *testing.T) bool) {
	t.Helper()

	t1 := &testing.T{}

	actual := block(t1)

	if actual != expected {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}
}

func Test_AssertEqual(t *testing.T) {
	testAssert(t, true, func(t *testing.T) bool {
		return Equal(t, 1, 1)
	})

	testAssert(t, false, func(t *testing.T) bool {
		return Equal(t, 1, 2)
	})

	testAssert(t, true, func(t *testing.T) bool {
		return Equal(t, 1, int(1))
	})

	testAssert(t, true, func(t *testing.T) bool {
		return Equal(t, 123, uint(123))
	})

	testAssert(t, true, func(t *testing.T) bool {
		return Equal(t, 321, uint16(321))
	})

	testAssert(t, true, func(t *testing.T) bool {
		return Equal(t, int32(100), int64(100))
	})

	testAssert(t, true, func(t *testing.T) bool {
		return Equal(t, 100, uint64(100))
	})
}

func Test_NotEqual(t *testing.T) {
	testAssert(t, true, func(t *testing.T) bool {
		return NotEqual(t, 2, 3)
	})
	testAssert(t, false, func(t *testing.T) bool {
		return NotEqual(t, 3, 3)
	})

	testAssert(t, true, func(t *testing.T) bool {
		return NotEqual(t, 2, int32(3))
	})

	testAssert(t, true, func(t *testing.T) bool {
		return NotEqual(t, int64(2), int32(3))
	})
}

func Test_StrictEqual(t *testing.T) {
	testAssert(t, true, func(t *testing.T) bool {
		return StrictEqual(t, 1, 1)
	})
	testAssert(t, false, func(t *testing.T) bool {
		return StrictEqual(t, 2, 1)
	})

	testAssert(t, false, func(t *testing.T) bool {
		return StrictEqual(t, int32(1), 1)
	})

	testAssert(t, false, func(t *testing.T) bool {
		return StrictEqual(t, 1, int64(1))
	})
}

func ptr(i int) *int {
	return &i
}

func Test_Same(t *testing.T) {
	testAssert(t, false, func(t *testing.T) bool {
		return Same(t, ptr(1), ptr(1))
	})
	testAssert(t, false, func(t *testing.T) bool {
		return Same(t, 1, 1)
	})
}

func Test_NotSame(t *testing.T) {
	testAssert(t, true, func(t *testing.T) bool {
		return NotSame(t, ptr(1), ptr(1))
	})
	testAssert(t, true, func(t *testing.T) bool {
		return NotSame(t, 1, 1)
	})
}

func Test_Contains_NotContains(t *testing.T) {
	testAssert(t, true, func(t *testing.T) bool {
		return Contains(t, []int{1, 2, 3}, 1)
	})
	testAssert(t, false, func(t *testing.T) bool {
		return Contains(t, []int{1, 2, 3}, 4)
	})
	testAssert(t, true, func(t *testing.T) bool {
		return Contains(t, []string{"Hello", "World"}, "Hello")
	})
	testAssert(t, true, func(t *testing.T) bool {
		return Contains(t, "Hello World", "World")
	})

	// NotContains
	testAssert(t, false, func(t *testing.T) bool {
		return NotContains(t, []int{1, 2, 3}, 1)
	})
	testAssert(t, true, func(t *testing.T) bool {
		return NotContains(t, []int{1, 2, 3}, 4)
	})
	testAssert(t, false, func(t *testing.T) bool {
		return NotContains(t, []string{"Hello", "World"}, "Hello")
	})
	testAssert(t, false, func(t *testing.T) bool {
		return NotContains(t, "Hello World", "World")
	})
}

func Test_ElementsMatch(t *testing.T) {
	testAssert(t, true, func(t *testing.T) bool {
		return ElementsMatch(t, []int{1, 2, 3}, []int{3, 1, 2})
	})
	testAssert(t, false, func(t *testing.T) bool {
		return ElementsMatch(t, []int{1, 2, 3}, []int{3, 1})
	})

	testAssert(t, true, func(t *testing.T) bool {
		return ElementsMatch(t, []string{"A", "B", "C"}, []string{"B", "C", "A"})
	})
	testAssert(t, false, func(t *testing.T) bool {
		return ElementsMatch(t, []string{"A", "B", "C"}, []string{"B", "D", "C", "A"})
	})
}

func Test_AssertNil(t *testing.T) {
	testAssert(t, true, func(t *testing.T) bool {
		return Nil(t, nil)
	})

	testAssert(t, false, func(t *testing.T) bool {
		return Nil(t, []int{})
	})

	testAssert(t, false, func(t *testing.T) bool {
		return Nil(t, 1)
	})

	testAssert(t, false, func(t *testing.T) bool {
		return Nil(t, "hello")
	})

	testAssert(t, false, func(t *testing.T) bool {
		return Nil(t, []int{1})
	})

	testAssert(t, false, func(t *testing.T) bool {
		return Nil(t, []string{"hello"})
	})

	testAssert(t, false, func(t *testing.T) bool {
		return Nil(t, []any{"hello"})
	})
}
