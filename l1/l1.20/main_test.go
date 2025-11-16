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

func TestReverseWords(t *testing.T) {
	data := []tests{
		{"1", "hello whole world", "world whole hello"},
		{"2", "кот смотрит на пустую коробку", "коробку пустую на смотрит кот"},
		{"3", "луна тихий город", "город тихий луна"},
		{"4", "ветер запах моря", "моря запах ветер"},
		{"empty", "", ""},
		{"one_word", "hello", "hello"},
		{"spaces", "a b", "b a"},
		{"unicode_mixed", "猫 看 月亮", "月亮 看 猫"},
	}
	for _, ts := range data {
		t.Run(ts.name, func(t *testing.T) {
			got := ReverseWords(ts.input)
			if got != ts.expected {
				t.Errorf("Неверный вывод %v != %v\n", got, ts.expected)
				return
			}
			fmt.Printf("%v -> %v\n", ts.input, ts.expected)
		})
	}
}
