package main

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

type tests struct {
	name   string
	input  []string
	expect []string
}

func TestGenSet(t *testing.T) {
	stressSlice := make([]string, 0, 1_000_000)
	for i := 0; i < 1_000_000; i++ {
		stressSlice = append(stressSlice, "x")
	}
	data := []tests{
		{"basic", []string{"cat", "cat", "dog", "cat", "tree"}, []string{"cat", "dog", "tree"}},
		{"unique", []string{"a", "b", "c", "d"}, []string{"a", "b", "c", "d"}},
		{"dups", []string{"same", "same", "same", "same"}, []string{"same"}},
		{"register", []string{"Dog", "dog", "DOG"}, []string{"Dog", "dog", "DOG"}},
		{"unicode", []string{"кот", "кот", "dog", "🐶", "🐶", "🌳"}, []string{"кот", "dog", "🐶", "🌳"}},
		//{"stress", stressSlice, []string{"x"}},
	}
	for _, ts := range data {
		t.Run(ts.name, func(t *testing.T) {
			got := GenSet(ts.input)
			sort.Strings(got)
			expected := ts.expect
			sort.Strings(expected)
			if !reflect.DeepEqual(got, expected) {
				t.Errorf("Неверный вывод:\n%v != %v", got, expected)
				return
			}
			fmt.Printf("%v -> %v\n", ts.input, got)
		})
	}
}
