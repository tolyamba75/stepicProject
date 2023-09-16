package main

import "fmt"

// Структура
type Person struct {
	name string
	age  int
}

// Создание метода
func (p Person) aboutPerson() {
	fmt.Println(p.name, p.age)
}

//Создаем структуру tom, а затем обращаемся к методу
func main() {
	var tom = Person{name: "Том", age: 35}
	var andrey = Person{"Андрей", 27}
	var leha = Person{name: "Лёха", age: 23}

	tom.aboutPerson()
	andrey.aboutPerson()
	leha.aboutPerson()

} //ВЫВОД: Том 35
