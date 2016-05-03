package kmp

import "testing"

func TestKMP(t *testing.T) {
	res := KMP("cocabc", "coccocabcabcocabc")
	if len(res) != 2 && res[0] != 3 && res[1] != 11 {
		t.Error("KMP matching error")
	}
}

func TestKMP2(t *testing.T) {
	res := KMP("coc", "cococ")
	if len(res) != 2 && res[0] != 0 && res[1] != 2 {
		t.Error("KMP matching error")
	}
}
