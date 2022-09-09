package main

import "fmt"

/*
	посетитель — это поведенческий шаблон проектирования, который позволяет добавлять поведение к структуре,
	фактически не изменяя структуру.
*/

type animal interface {
	say() string
	accept(visitor)
}

type visitor interface{
	growOldForCat(*cat)
	growOldForDog(*dog)
}

type dog struct{
	name string
	age int
}

func (d dog) say() string{
	return "gaw"
}

func (d *dog) accept(v visitor) {
	v.growOldForDog(d)
}

type cat struct{
	name string
	age int
}

func (c cat) say() string{
	return "meow"
}

func (c *cat) accept(v visitor){
	v.growOldForCat(c)
}

type aging struct{
	addAge int
}

func (a aging) growOldForCat(cat *cat){
	cat.age += a.addAge
}

func (a aging) growOldForDog(dog *dog){
	dog.age += a.addAge
}

func main(){
	cat := cat{
		"boba",
		23,
	}

	fmt.Println(cat)
	cat.accept(aging{12})
	fmt.Println(cat)
}
