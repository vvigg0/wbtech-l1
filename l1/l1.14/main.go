package main

import (
	"fmt"
	"reflect"
)

func defineType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Это int: %v\n", v)
	case string:
		fmt.Printf("Это string: %v\n", v)
	case bool:
		fmt.Printf("Это bool: %v\n", v)
	/*
		case chan string,chan int,chan struct{},chan bool ... ... ...:
			fmt.Printf("Это chan: %v\n", v)
	*/
	default:
		t := reflect.TypeOf(i)
		if t != nil && t.Kind() == reflect.Chan {
			fmt.Printf("Это канал: %v\n", t)
			return
		}
		fmt.Printf("Неизвестный тип: %v\n", v)
	}
}
func main() {
	defineType(35)
	defineType("Hello world")
	defineType(false)
	defineType(make(chan string))
	defineType(make(chan bool))
	defineType(35.43)
	defineType([]int{1, 2})
}
