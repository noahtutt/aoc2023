package string

import "testing"

func TestReverse(t *testing.T) {
	str1 := "abcd"
	str1Reversed := "dcba"

	if str1 != Reverse(str1Reversed) {
		t.Fail()
	}

	str2 := "a𓀇"
	str2Reversed := "𓀇a"

	if str2 != Reverse(str2Reversed) {
		t.Fail()
	}
}
