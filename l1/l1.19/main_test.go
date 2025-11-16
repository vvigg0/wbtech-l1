package main

import (
	"fmt"
	"testing"
)

type tests struct {
	name     string
	input    string
	expected string
}

func TestReverseStr(t *testing.T) {
	data := []tests{
		{"ASCII", "Hello", "olleH"},
		{"Кириллица", "Привет", "тевирП"},
		{"Эмодзи", "😀😁😂🤣😃", "😃🤣😂😁😀"},
		{"ASCII + Эмодзи", "Go🐹Lang", "gnaL🐹oG"},
		{"Кириллица + Эмодзи", "Здра🚀вствуйте", "етйувтсв🚀ардЗ"},
		// string([]rune{....} это 👦👧👩👨 с zero-width joiner)
		{"Сложный эмодзи", "👨‍👩‍👧‍👦", string([]rune{128102, 8205, 128103, 8205, 128105, 8205, 128104})},
	}
	for _, ts := range data {
		t.Run(ts.name, func(t *testing.T) {
			got := ReverseStr(ts.input)
			if got != ts.expected {
				t.Errorf("Неверный вывод:\n%v != %v\n%v!=%v", got, ts.expected, []rune(got), []rune(ts.expected))
				return
			}
			fmt.Printf("%v -> %v\n", ts.input, got)
		})
	}
}
