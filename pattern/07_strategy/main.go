package main

import "fmt"

/*
	Этот шаблон проектирования позволяет изменять поведение объекта во время выполнения
	без каких-либо изменений в классе этого объекта.

	Шаблон Strategy позволяет менять выбранный алгоритм независимо от объектов-клиентов, которые его используют.
*/

//strategy
type coolingAlgorithm interface {
	cool(r *reactor)
}

//context
type reactor struct {
	temp  int
	cAlgo coolingAlgorithm
}

func (r *reactor) cool(){
	r.cAlgo.cool(r)
}

type passiveCool struct{}

func (p *passiveCool) cool(r *reactor) {
	fmt.Println("Пассивное охлаждение")
}

type activeCool struct{}

func (a *activeCool) cool(r *reactor) {
	fmt.Println("Активное охлаждение")
}

func main() {
	r := reactor{
		temp: 28,
		cAlgo: &passiveCool{},
	}

	r.cool()

	r.cAlgo = &activeCool{}

	r.cool()
}
