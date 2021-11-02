package main

import "testing"

func TestNum(t *testing.T) {
	n := newNum()
	if len(n) != 5 {
		t.Errorf("Not enought numbers")
	}

}