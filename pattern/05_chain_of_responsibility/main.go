package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Что делает:
	Цепочка обязанностей – позволяет передавать запросы последовательно по цепочке обработчиков - обьектов.
	Каждый последующий обработчик решает, может ли он обработать запрос сам и стоит ли передавать запрос дальше по цепи.

Применимость:
	1. Когда важно чтобы обработка выполнялась один за другим строго в определенном порядке.
	2. Когда программа должна обрабатывать разные запросы разными способами, но заранее не известно
		какие конкретно запросы будут приходить и какие обработчики для них понадобятся.
	3. Когда набор обьектов для способных обработать запрос должен задаваться динамически -
		(В любой момент можно вмешаться в существующую цепочку и переназначить связи так, чтобы убрать или добавить новое звено.).

Плюсы:
	1. Уменьшает зависимость между клиентом и обработчиками
	2. Обьект упрощается поскольку ему не нужно знать структуру цепочки
	3. Возможноть динамического добавления или удаления обязанностей посредством изменения элементов цепочки или их порядка

Минусы:
	1. Цепочки могут усложнять отслеживание запросов и их отладку
	2. Обработка запроса не гарантированна, если ни один обработчик не обработает запрос он просто вылетит из цепочки
		(хотя это может быть и плюсом)
*/

type request struct {
	resource string
	login    string
	password string
}

type handler interface {
	execute(*request)
	setNext(handler)
}

type identification struct {
	next handler
}

func (i identification) execute(r *request) {
	if r.login == "existingLogin" {
		i.next.execute(r)
	} else {
		fmt.Println("Прерывание цепочки: идентификация не пройдена")
	}
}

func (i *identification) setNext(h handler) {
	i.next = h
}

type authentication struct {
	next handler
}

func (a authentication) execute(r *request) {
	if r.password == "correctPassword" {
		a.next.execute(r)
	} else {
		fmt.Println("Прерывание цепочки: аунтетификация не пройдена")
	}
}

func (a *authentication) setNext(h handler) {
	a.next = h
}

//Двухфакторная аунтетификация, если подключена для аккаунта
type doubleAuth struct {
	next handler
}

func (d doubleAuth) execute(r *request){
	isDAuth := rand.Intn(3)
	//Если подключена
	if isDAuth == 1 {
		fmt.Println("Смс код отправлен")
		fmt.Println("Смс подтвержден")
		d.next.execute(r)
	} else {
		d.next.execute(r)
	}
}

func (d *doubleAuth) setNext(h handler){
	d.next = h
}

type authorization struct {
	next handler
}

func (a authorization) execute(r *request) {
	fmt.Println("Доступ к ресурсу предоставлен")
}

func (a *authorization) setNext(h handler) {
	a.next = h
}

func main() {
	rand.Seed(time.Now().UnixNano())

	request := &request{
		login:    "existingLogin",
		password: "correctPassword",
	}

	auth := &authentication{}
	ident := &identification{}
	ident.setNext(auth)
	dAuth := &doubleAuth{}
	auth.setNext(dAuth)
	authorization := &authorization{}
	dAuth.setNext(authorization)

	ident.execute(request)
}
