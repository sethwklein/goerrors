package errors_test

import (
	"fmt"
	"sethwklein.net/go/errors"
	"testing"
)

func TestWalkUntil(t *testing.T) {
	var list error
	for i := 1; i <= 1000; i++ {
		list = errors.Append(list, fmt.Errorf("number %v", i))
	}
	count := 0
	errors.WalkUntil(list, func(e error) bool {
		count++
		if count >= 3 {
			return false
		}
		return true
	})
	if count != 3 {
		t.Errorf("wrong count. expected 3 got %v", count)
	}
}
