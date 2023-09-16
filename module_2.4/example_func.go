package main

import "fmt"

//Создаем структуру
type Person struct {
	name string
	age  int
}

// Создаем функцию, которая принимает строковое и целочисленное значение и выводит их
func aboutPerson(name string, age int) {
	fmt.Println(name, age)
}

//Создаем структуру tom, а затем обращаемся к функции, передавая параметры имени и возраста tom
func main() {
	var tom = Person{name: "Том", age: 35}
	aboutPerson(tom.name, tom.age)

} //ВЫВОД: Том 35
