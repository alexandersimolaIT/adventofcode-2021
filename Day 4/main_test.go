package main

import "testing"

func Test_StringIsNumber(t *testing.T) {
	if StringIsNumber("a") {
		t.Error("StringIsNumber(\"a\") == true, when it should be false")
	}
	if StringIsNumber("") {
		t.Error("StringIsNumber(\"\") == true, when it should be false")
	}
	if !StringIsNumber("1") {
		t.Error("StringIsNumber(\"1\") == false, when it should be true")
	}
	if !StringIsNumber("123") {
		t.Error("StringIsNumber(\"123\") == false, when it should be true")
	}
	if StringIsNumber("123g") {
		t.Error("StringIsNumber(\"123g\") == true, when it should be false")
	}
}
