package main

import "testing"

func TestNum(t *testing.T) {
	n := newNum()
	if len(n) != 5 {
		t.Errorf("Not enought numbers")
	}
}

func TestMinMax(t *testing.T) {
	n := newNum()
	min, max := MinMax(n)
	if min != -15 {
		t.Errorf("Min function not working")
	}
	if max != 30 {
		t.Errorf("Max function not working")
	}
}

func TestAvrg(t *testing.T)  {
	n := newNum()
	avrg, sum := Averg(n)
	if avrg != 14 {
		t.Errorf("Avrg function not working")
	}
	if sum != 70 {
		t.Errorf("Sum function not working")
	}
}