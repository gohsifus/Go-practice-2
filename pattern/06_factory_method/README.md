## Фабричный метод
Фабричный метод - это порождающий паттерн, который определяет общий интерфейс для создания объектов, но позволяет подклассам изменять тип создаваемого обьекта.

#### Применимость
Когда заранее не известны типы обьектов с которыми должен работать код (фабричный метод определяет общий интерфейс для создания и отделяет создание обьекта от остального кода), таким образом чтобы добавить поддержку нового продукта нужно создать новый подкласс создателя

#### Преимущества
* Соблюдается принцип открытости закрытости (расширяемость)
* Соблюдается принцип единой ответственности (можем переместить код создания продукта в одно место в программе, что упростит поддержку кода)

#### Недостатки 
* Код может стать более сложным из за введения множества классов создателей

#### Отличия
* Абстрактная фабрика создает семейства связанных обьектов (например комлпекты обувь + шорты, семейство nike и adidas)
* Фабрика - клиент работает только с фабричной структурой, которая создает продукт в зависимости от аргумента фабрики.
* Фабричный метод - создается структура классов в котором классы наследники переопределяют метод создания определенный в базовом классе.