package ptr_test

import (
	"ptr"
	"testing"
)

func TestTo(t *testing.T) {
	assert := func(p *int) {
		if p == nil {
			t.Fatal("p is nil")
		}
		if ptr.From(p) != 1 {
			t.Fatal("p does not point to 1.")
		}
	}
	assert(ptr.To(1))
}
