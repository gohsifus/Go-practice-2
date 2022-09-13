package main

import "fmt"

/*
	посетитель — это поведенческий шаблон проектирования, который позволяет добавлять поведение к структуре,
	фактически не изменяя структуру.
*/

//Структуры котором хотим добавить метод
type animal interface {
	say() string
	accept(visitor)
}

type visitor interface {
	visitForCat(*cat)
	visitForDog(*dog)
}

type dog struct {
	name string
	age  int
}

func (d dog) say() string {
	return "gaw"
}

func (d *dog) accept(v visitor) {
	v.visitForDog(d)
}

type cat struct {
	name string
	age  int
}

func (c cat) say() string {
	return "meow"
}

func (c *cat) accept(v visitor) {
	v.visitForCat(c)
}

//Реализует visitor добавляет методы к target структуре
type aging struct {
	addAge int
}

func (a aging) visitForCat(cat *cat) {
	cat.age += a.addAge
}

func (a aging) visitForDog(dog *dog) {
	dog.age += a.addAge
}

func main() {
	cat := cat{
		"boba",
		23,
	}

	fmt.Println(cat)
	cat.accept(aging{12})
	fmt.Println(cat)
}
