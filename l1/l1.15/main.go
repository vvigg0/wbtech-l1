package main

import (
	"fmt"
	"strings"
)

/*
Проблема была в
justString = v[:100]
происходила утечка памяти,justString ссылался на
изначальный слайс байтов строки v(большой строки)
из за чего сборщик мусора не смог бы "собрать" строку v
Подстроки следует копировать, чтобы избежать удержания всей исходной строки в памяти.
*/
var justString string

func createHugeString(i int) string {
	return strings.Repeat("x", i)
}
func someFunc() {
	v := createHugeString(1 << 30)
	justString = strings.Clone(v[:100]) // или // justString = string([]byte(v[:100]))
	fmt.Println(justString)
}

func main() {
	someFunc()
}
