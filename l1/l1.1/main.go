package main

import "fmt"

type Human struct {
	name string
	age  int
}
type Action struct {
	actionName string
	Human
}

/*
2 метода с одинаковым названием - метод вызванный "ребёнком" перекрывает "родительский"
a.greeting1() выведет "Мое имя ..."
a.Human.greeting1() выведет "Привет,меня зовут ..."
*/
func (h *Human) greeting1() {
	fmt.Printf("Привет,меня зовут %v\n", h.name)
}
func (a *Action) greeting1() {
	fmt.Printf("Мое имя %v\n", a.name)
}

// метод для структуры "родителя",может вызываться "ребёнком"
func (h *Human) greeting2() {
	fmt.Printf("Мне %v\n", h.age)
}

// метод для структуры "ребёнка"
func (a *Action) greeting3() {
	fmt.Printf("И я люблю %v\n", a.actionName)
}

/*
	Action видит поля Human, но не наоборот - в таком случае надо явно передать структуру "ребёнок"
	"Ребёнок" знает про "родителя",а "родитель" ничего не знает о "ребёнке"
*/
func (h *Human) resume(a *Action) {
	fmt.Printf("Я %v, мне %v, люблю %v", h.name, h.age, a.actionName)
}

func main() {
	h := Human{name: "Артем", age: 18}
	a := Action{actionName: "бегать", Human: h}
	a.Human.greeting1()
	a.greeting1()
	a.greeting2()
	a.greeting3()
	h.resume(&a)
}
