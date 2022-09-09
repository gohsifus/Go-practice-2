package main

/*
	Команда — это поведенческий шаблон проектирования. Он предлагает инкапсулировать запрос как отдельный объект.
	Созданный объект имеет всю информацию о запросе и, таким образом, может выполнять его самостоятельно.
*/

import "fmt"

type cooker interface{
	makePizza(name string)
}

//receiver
type cook struct{

}

func (c cook) makePizza(name string){
	fmt.Println("mk pizza - ", name)
}

type command interface {
	execute()
}

//concreteCommand
type makePizzaCommand struct{
	name string
	cook cook
}

func (m makePizzaCommand) execute(){
	m.cook.makePizza(m.name)
}

//invoker
type order struct{
	command
}

func (o order) invoke(){
	o.execute()
}


func main(){
	cook := cook{}
	pizza := makePizzaCommand{
		name: "california",
		cook: cook,
	}

	order := order{pizza}
	order.invoke()
}
