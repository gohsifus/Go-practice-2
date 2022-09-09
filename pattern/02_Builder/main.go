package main

import "fmt"

/*
	builder предоставляет способ создания составного обьекта
	builder выносит конструирование обьекта за пределы собственного класса

	Может быть несколько разных ConcreteBuilder-классов,
	каждый из которых реализует различный способ создания продукта
*/

//Интерфейс строителя - определяет шаги построения обьекта, реализуемые в конкретных строителях
type iArmyBuilder interface {
	setWeaponType() iArmyBuilder
	setClothesType() iArmyBuilder
	setNumOfMembers() iArmyBuilder
}

//product создаваемый объект, это класс, который определяет наш сложный объект, который мы пытаемся собрать
type army struct{
	weaponType string
	clothesType string
	numOfMembers int
}

//director директор определяет порядок вызова строительных шагов
//директор не обязательный класс, если порядок шагов не имеет значения - строить обьект можно непосредственно с помощью builder
type director struct{
	builder iArmyBuilder
}

func (d *director) Build(){
	d.builder.setWeaponType().setClothesType().setNumOfMembers()
}

//*********************************
//Конкретный строитель - дешевая армия
type cheapArmy struct {
	weaponType   string
	clothesType  string
	numOfMembers int
}

func (c *cheapArmy) setWeaponType() iArmyBuilder{
	c.weaponType = "Копье"
	return c
}

func (c *cheapArmy) setClothesType() iArmyBuilder{
	c.clothesType = "Туника"
	return c
}

func (c *cheapArmy) setNumOfMembers() iArmyBuilder{
	c.numOfMembers = 18000
	return c
}

func (c cheapArmy) getProduct() *army{
	return &army{
		weaponType: c.weaponType,
		clothesType: c.clothesType,
		numOfMembers: c.numOfMembers,
	}
}

//*********************************
//Конкретный строитель - дорогая армия
type expensiveArmy struct {
	weaponType   string
	clothesType  string
	numOfMembers int
}

func (e *expensiveArmy) setWeaponType() iArmyBuilder{
	e.weaponType = "Меч"
	return e
}

func (e *expensiveArmy) setClothesType() iArmyBuilder{
	e.weaponType = "Доспехи"
	return e
}

func (e *expensiveArmy) setNumOfMembers() iArmyBuilder{
	e.numOfMembers = 50000
	return e
}

func (e expensiveArmy) getProduct() *army{
	return &army{
		weaponType: e.weaponType,
		clothesType: e.clothesType,
		numOfMembers: e.numOfMembers,
	}
}

//*********************************

func main() {
	army1 := &cheapArmy{}
	director := director{
		builder: army1,
	}
	director.Build()

	fmt.Println(army1)
}
