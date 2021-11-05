package main

import "testing"

func TestNum(t *testing.T) {
	n := []int{10, -15, 20, 25, 30}
	if len(n) != 5 {
		t.Errorf("Not enought numbers")
	}
}

func TestMinMax(t *testing.T) {
	n := []int{10, -15, 20, 25, 30}
	min, max := MinMax(n)
	for i, result := range n {
		if min != result {
			t.Errorf("Min function vith index %v not working", i)
		}
	}
	for i, result := range n {
		if max != result {
			t.Errorf("Max function with index %v not working", i)
		}
	}

}

func TestAvrg(t *testing.T)  {
	n := []int{10, -15, 20, 25, 30}
	k := []int{14, 70, 22, 0, -2}
	avrg, sum := Averg(n)
	for i, result := range k {
		if avrg != float64(result) {
			t.Errorf("Not correct avrg result in position %v", i)
		}
	}
	for i, result := range k {
		if sum != float64(result) {
			t.Errorf("Sum not correct in position %v", i)
		}
	}
}