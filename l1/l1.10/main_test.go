package main

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestGrouping(t *testing.T) {
	data := []float64{-0.1, -1, -9.9, -10, -19.9, -20, -29.9}
	got := Grouping(data)
	expected := map[int][]float64{
		-20: []float64{-29.9, -20},
		-10: []float64{-10, -19.9},
		0:   []float64{-1, -9.9, -0.1},
	}
	if len(got) != len(expected) {
		t.Errorf("Неверный результат: got - %v, expected - %v", got, expected)
	}
	for key := range expected {
		s1 := expected[key]
		sort.Float64s(s1)
		s2 := got[key]
		sort.Float64s(s2)
		if !reflect.DeepEqual(s1, s2) {
			fmt.Println(s1, s2)
			t.Errorf("Неверный результат: got - %v, expected - %v", got, expected)
			break
		}
	}
}
