package scanner

import (
	"testing"
)

func assertTrue(t *testing.T, b bool) {
	if !b {
		t.Errorf("Not true")
	}
}

func assertFalse(t *testing.T, b bool) {
	if b {
		t.Errorf("Not false")
	}
}

func TestR(t *testing.T) {
	l := r('c')

	assertTrue(t, l('c'))
	assertFalse(t, l('a'))
	assertFalse(t, l('b'))
}

func TestOr(t *testing.T) {
	l := or(r('a'), r('b'), r('c'))

	assertTrue(t, l('a'))
	assertTrue(t, l('b'))
	assertTrue(t, l('c'))
	assertFalse(t, l('d'))
	assertFalse(t, l('e'))
}
