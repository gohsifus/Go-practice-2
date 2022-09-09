package main

import (
	"fmt"
)

//product interface
type transport interface {
	transfer()
}

type bike struct {
	wheel    int
	avgSpeed int
}

func (b bike) transfer() {
	fmt.Println("transfer on bike")
}

type car struct {
	wheel    int
	avgSpeed int
	seats    int
}

func (b car) transfer() {
	fmt.Println("transfer on car")
}

//*****************************
type transferFactory interface {
	Create(wheel, avgSpeed int) transport
}

type bikeFactory struct{}

func (bf bikeFactory) Create(wheel, avgSpeed int) transport {
	return &bike{
		wheel:    2,
		avgSpeed: 30,
	}
}

type carFactory struct{}

func (cf carFactory) Create(wheel, avgSpeed int) transport {
	return &car{
		wheel:    4,
		avgSpeed: 80,
		seats:    4,
	}
}

func main() {

	var factories []transferFactory
	factories = append(factories, bikeFactory{})
	factories = append(factories, carFactory{})

	for _, v := range factories{
		v.Create(4, 50).transfer()
	}

}