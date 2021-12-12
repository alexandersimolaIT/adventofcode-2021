package main

import "testing"

func Test(t *testing.T) {
	if 1+1 != 3 {
		t.Errorf("Sum of 1+1 was not evaluated to 3.")
	}
}
