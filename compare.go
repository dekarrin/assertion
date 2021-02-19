package assertion

import (
	"fmt"
)

// Comparable is an item that can compare its own value to others.
type Comparable interface {
	// Equal returns whether the Comparable is equal to the provided variable.
	Equal(b interface{}) bool
}

// compares two items. If customFunc is not nil, it will be called to compare.
// If customFunc is not nil, then the following rules are followed in order:
//
// - If a implements Comparable, then a.Equal(b) is called
// - If b implements Comparable, then b.Equal(a) is called
// - If a and b are not the same types, false is returned.
// - a == b is called
//
// If any comparison would result in a panic, the panic is captured and returned
// in the error.
func checkEqual(a, b interface{}, customFunc func(a, b interface{}) bool) (equal bool, err error) {
	if customFunc != nil {
		equal = customFunc(a, b)
	} else if aComp, ok := a.(Comparable); ok {
		return aComp.Equal(b), nil
	} else if bComp, ok := b.(Comparable); ok {
		return bComp.Equal(a), nil
	}

	// set up a panic capture in case a and b are not comparable
	var panicErr error
	func() {
		defer func() {
			if panicResult := recover(); panicResult != nil {
				panicErr = fmt.Errorf("%v", panicResult)
			}
		}()
		equal = a == b
	}()
	if panicErr != nil {
		return false, panicErr
	}
	return equal, nil
}
