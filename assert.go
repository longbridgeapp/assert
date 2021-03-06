package assert

import (
	"time"

	testifyAssert "github.com/stretchr/testify/assert"
)

type TestingT interface {
	Errorf(format string, args ...interface{})
}

type tHelper interface {
	Helper()
}

// Equal asserts two values are equal, this will ignore type. (int8, int32, int64, int, uint ...)
//
//    assert.Equal(t, 123, 123) // ok
//    assert.Equal(t, uint32(123), int64(123)) // ok
//    assert.Equal(t, 123, "123") // fail
//
// This method only detects whether the values are equal (understand by your brain).
// If you want to strictly check the type, please use:
//
//    assert.StrictEqual(t, 123, int64(123))
//
func Equal[A any, B any](t TestingT, expected A, actual B, msgAndArgs ...any) bool {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}

	return testifyAssert.EqualValues(t, expected, actual, msgAndArgs...)
}

// StrictEqual strictly asserts two values and type are equal.
//
//    assert.StrictEqual(t, 123, 123)
//    assert.StrictEqual(t, 123, int64(123)) // fail
//
func StrictEqual[A any, B any](t TestingT, expected A, actual B, msgAndArgs ...any) bool {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}

	return testifyAssert.Equal(t, expected, actual, msgAndArgs...)
}

// NotEqual asserts value not equal, ignore type.
//
//    assert.NotEqual(t, 12, 13) // ok
//    assert.NotEqual(t, 12, int32(12)) // fail
//    assert.NotEqual(t, 12, int32(13)) // ok
//
func NotEqual[A any, B any](t TestingT, expected A, actual B, msgAndArgs ...any) bool {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	return testifyAssert.NotEqualValues(t, expected, actual, msgAndArgs...)
}

// Same asserts that two pointers reference the same object.
//
//    assert.Same(t, ptr1, ptr2)
//
// Both arguments must be pointer variables. Pointer variable sameness is
// determined based on the equality of both type and value.
func Same[A any, B any](t TestingT, expected A, actual B, msgAndArgs ...any) bool {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	return testifyAssert.Same(t, expected, actual, msgAndArgs...)
}

// NotSame asserts that two pointers do not reference the same object.
//
//    assert.NotSame(t, ptr1, ptr2)
//
// Both arguments must be pointer variables. Pointer variable sameness is
// determined based on the equality of both type and value.
func NotSame[A any, B any](t TestingT, expected A, actual B, msgAndArgs ...any) bool {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	return testifyAssert.NotSame(t, expected, actual, msgAndArgs...)
}

// Contains asserts that the specified string, list(array, slice...) or map contains the
// specified substring or element.
//
//    assert.Contains(t, "Hello World", "World")
//    assert.Contains(t, ["Hello", "World"], "World")
//    assert.Contains(t, {"Hello": "World"}, "Hello")
//
func Contains[A any, B any](t TestingT, s A, contains B, msgAndArgs ...any) bool {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	return testifyAssert.Contains(t, s, contains, msgAndArgs...)
}

// NotContains asserts that the specified string, list(array, slice...) or map does NOT contain the
// specified substring or element.
//
//    assert.NotContains(t, "Hello World", "Earth")
//    assert.NotContains(t, ["Hello", "World"], "Earth")
//    assert.NotContains(t, {"Hello": "World"}, "Earth")
//
func NotContains[A any, B any](t TestingT, s A, contains B, msgAndArgs ...any) bool {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	return testifyAssert.NotContains(t, s, contains, msgAndArgs...)
}

// ElementsMatch asserts that the specified listA(array, slice...) is equal to specified
// listB(array, slice...) ignoring the order of the elements. If there are duplicate elements,
// the number of appearances of each of them in both lists should match.
//
// 	 assert.ElementsMatch(t, [1, 3, 2, 3], [1, 3, 3, 2])
//
func ElementsMatch[T any](t TestingT, listA, listB T, msgAndArgs ...any) (ok bool) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	return testifyAssert.ElementsMatch(t, listA, listB, msgAndArgs...)
}

// Panics asserts that the code inside the specified PanicTestFunc panics.
//
//   assert.Panics(t, func(){ GoCrazy() })
//
func Panics(t TestingT, f testifyAssert.PanicTestFunc, msgAndArgs ...any) bool {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	return testifyAssert.Panics(t, f, msgAndArgs...)
}

// NotPanics asserts that the code inside the specified PanicTestFunc panics.
//
//   assert.NotPanics(t, func(){ GoCrazy() })
//
func NotPanics(t TestingT, f testifyAssert.PanicTestFunc, msgAndArgs ...any) bool {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	return testifyAssert.NotPanics(t, f, msgAndArgs...)
}

// WithinDuration asserts that the two times are within duration delta of each other.
//
//   assert.WithinDuration(t, time.Now(), time.Now(), 10*time.Second)
//
func WithinDuration(t TestingT, expected, actual time.Time, delta time.Duration, msgAndArgs ...any) bool {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	return testifyAssert.WithinDuration(t, expected, actual, delta, msgAndArgs...)
}

// NoError asserts that a function returned no error (i.e. `nil`).
//
//   actualObj, err := SomeFunction()
//   if assert.NoError(t, err) {
//	   assert.Equal(t, expectedObj, actualObj)
//   }
//
func NoError(t TestingT, err error, msgAndArgs ...any) bool {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	return testifyAssert.NoError(t, err, msgAndArgs...)
}

// Error asserts that a function returned no error (i.e. `nil`).
//
//   actualObj, err := SomeFunction()
//   if assert.Error(t, err) {
//	   assert.Equal(t, expectedObj, actualObj)
//   }
//
func Error(t TestingT, err error, msgAndArgs ...any) bool {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	return testifyAssert.Error(t, err, msgAndArgs...)
}

// EqualError asserts that a function returned an error (i.e. not `nil`)
// and that it is equal to the provided error.
//
//   actualObj, err := SomeFunction()
//   assert.EqualError(t, err,  expectedErrorString)
//
func EqualError(t TestingT, theError error, errString string, msgAndArgs ...any) bool {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	return testifyAssert.EqualError(t, theError, errString, msgAndArgs...)
}

// True asserts that the specified value is true.
//
//    assert.True(t, myBool)
//
func True(t TestingT, value bool, msgAndArgs ...any) bool {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	return testifyAssert.True(t, value, msgAndArgs...)
}

// False asserts that the specified value is false.
//
//    assert.False(t, myBool)
//
func False(t TestingT, value bool, msgAndArgs ...any) bool {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	return testifyAssert.False(t, value, msgAndArgs...)
}

// Nil asserts that the specified object is nil.
//
//    assert.Nil(t, err)
//
func Nil(t TestingT, object any, msgAndArgs ...any) bool {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	return testifyAssert.Nil(t, object, msgAndArgs...)
}

// NotNil asserts that the specified object is not nil.
//
//    assert.NotNil(t, "hello")
//
func NotNil(t TestingT, object any, msgAndArgs ...any) bool {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	return testifyAssert.NotNil(t, object, msgAndArgs...)
}

// Empty asserts that the specified object is empty.  I.e. nil, "", false, 0 or either
// a slice or a channel with len == 0.
//
//    assert.Empty(t, "")
//
func Empty(t TestingT, object any, msgAndArgs ...any) bool {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	return testifyAssert.Empty(t, object, msgAndArgs...)
}

// Len asserts that the specified object has specific length.
// Len also fails if the object has a type that len() not accept.
//
//    assert.Len(t, mySlice, 3)
func Len(t TestingT, object any, length int, msgAndArgs ...any) bool {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	return testifyAssert.Len(t, object, length, msgAndArgs...)
}
