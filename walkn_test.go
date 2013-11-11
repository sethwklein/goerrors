package errors_test

import (
	"fmt"
	"os"
	"sethwklein.net/go/errors"
)

func ExampleWalkN() {
	var list error
	for i := 1; i <= 1000; i++ {
		list = errors.Append(list, fmt.Errorf("number %v", i))
	}
	errors.WalkN(list, 3, func(e error) {
		// In real code, this should generally use os.Stderr, but
		// https://code.google.com/p/go/issues/detail?id=4550
		// broke that for examples.
		fmt.Fprintln(os.Stdout, e)
	})
	// Output:
	// number 1
	// number 2
	// number 3
}
