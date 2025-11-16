package main

import "fmt"

func ReverseWords(str string) string {
	// исключение строк, которые нет смысла переворачивать
	if len(str) == 0 || len(str) == 1 {
		return str
	}
	// 1.переворот строки
	r := []rune(str)
	left := 0
	right := len(r) - 1
	for left < right {
		r[left], r[right] = r[right], r[left]
		left++
		right--
	}
	/*
		сокращенный вид(тяжело читается)
		for left, right := 0, len(r)-1; left < right; left, right = left+1, right-1 {
			r[left], r[right] = r[right], r[left]
		}
	*/

	// 2. переворот слов в перевернутой строке
	left = 0
	for i := range r {
		// ловим на пробеле между словами или конце строки
		if r[i] == ' ' || i == len(r)-1 {
			if i == len(r)-1 {
				right = i
			} else {
				// смещаем правый "указатель" на  правую границу слова
				right = i - 1
			}
			// меняем буквы местами и смещаем "указатели"
			for left < right {
				r[left], r[right] = r[right], r[left]
				left++
				right--
			}
			// смещаем левый "указатель" на левую границу слова
			left = i + 1
		}
	}
	return string(r)
}

func main() {
	str := "sun dog snow"
	fmt.Println(ReverseWords(str))
}
