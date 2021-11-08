package main

import "testing"

func TestNum(t *testing.T) {
	n := newNum()
	if len(n) != 5 {
		t.Errorf("Not correct number of items")
	}
}

func TestMin(t *testing.T) {
	testData := map[int][]int {
		10: []int{10, 15, 20, 25, 30},
		0: []int{0, 0, 0, 0, 0},
		-30: []int{-10, -15, -20, -25, -30},
	}
	for i, result := range testData {
		min, _ := MinMax(result)
		if min != int(i) {
			t.Errorf("Min function vith value %v not working", i)
		}
	}
}
func TestMax(t *testing.T) {
	testData := map[int][]int {
		30: []int{10, 15, 20, 25, 30},
		0: []int{0, 0, 0, 0, 0},
		-10: []int{-10, -15, -20, -25, -30},
	}
	for i, result := range testData {
		_, max := MinMax(result)
		if max != int(i) {
			t.Errorf("Max function vith value %v not working", i)
		}
	}
}

func TestAvrg(t *testing.T)  {
	testData := map[int][]int {
		20: []int{10, 15, 20, 25, 30},
		0: []int{0, 0, 0, 0, 0},
		-20: []int{-10, -15, -20, -25, -30},
	}
	for i, result := range testData {
		avrg, _ := Averg(result)
		if avrg != float64(i) {
			t.Errorf("Not correct avrg result in %v", i)
		}
	}
}
func TestSum(t *testing.T)  {
	testData := map[int][]int {
		100: []int{10, 15, 20, 25, 30},
		0: []int{0, 0, 0, 0, 0},
		-100: []int{-10, -15, -20, -25, -30},
	}
	for i, result := range testData {
		_, sum := Averg(result)
		if sum != float64(i) {
			t.Errorf("Not correct avrg result in %v", i)
		}
	}
}



