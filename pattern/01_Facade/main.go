package main

import "fmt"

/*
	Паттерн фасад предоставляет упрощенный интерфейс для взаимодействия с сложной системой.
	Фасад уменьшает число объектов, с которыми клиентам приходится иметь дело, упрощая работу с подсистемой.

	Сложная система - Db, ValidationSystem, EGURLApi, NotificationSystem
	Фасад - orderFacade

	Клиенты использующие фасад могут создать заявку не работая напрямую с большим количеством компонентов сложной системы.
*/

//Db База данных
type Db struct{}

func (d Db) save(data string) {
	fmt.Println("Сохранение заявки в базе")
}

// EGRULApi система для проверки юрлиц
type EGRULApi struct{}

func (e EGRULApi) isLegalEntity() bool {
	fmt.Println("Запрос в ЕГРЮЛ")
	return true
}

//ValidationSystem система валидации данных
type ValidationSystem struct{}

func (v ValidationSystem) Validate(data string) bool {
	fmt.Println("Проверка данных")
	return true
}

//NotificationSystem ...
type NotificationSystem struct{}

func (n NotificationSystem) sendNotify() {
	fmt.Println("Отправление уведомления о созданной заявке менеджеру")
	fmt.Println("Отправление уведомления заказчику")
}

//Фасад
type orderFacade struct {
	Store            Db
	ValidationSystem ValidationSystem
	EApi             EGRULApi
	NotifyS          NotificationSystem
}

//CreateOrder заявка создается с помощью одного метода
func (o orderFacade) CreateOrder(dataForOrder string) {
	o.validationSystem.Validate(dataForOrder)
	o.eApi.isLegalEntity()
	o.store.save(dataForOrder)
	o.notifyS.sendNotify()
}

func main() {
	facade := orderFacade{}
	facade.CreateOrder("order data")
}
