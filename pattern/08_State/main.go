package main

import "fmt"

/*
	Используется в тех случаях, когда во время выполнения программы объект должен
	менять своё поведение в зависимости от своего состояния.
*/

type state interface {
	tempDown()
	tempUp()
}

type reactor struct {
	tooHot  state
	ok      state
	tooCold state
	off     state

	currentState state

	temperature int
}

func newReactor() *reactor {
	r := &reactor{
		temperature: 0,
	}

	thState := &tooHotState{r: r}
	tcState := &tooColdState{r: r}
	oState := &okState{r: r}
	ofState := &offState{r: r}

	r.tooHot = thState
	r.tooCold = tcState
	r.ok = oState
	r.off = ofState

	r.currentState = r.ok

	return r
}

//*******************************
type tooHotState struct {
	r *reactor
}

func (t *tooHotState) tempUp() {
	fmt.Println("Реактор перегрет - повышаем температуру температуру")
	t.r.temperature += 100
	t.r.currentState = t.r.off //Реактор выключается от перегрева
}

func (t *tooHotState) tempDown() {
	fmt.Println("Реактор перегрет - понижаем температуру")
	t.r.temperature -= 100
	t.r.currentState = t.r.ok
}

//*******************************
type okState struct {
	r *reactor
}

func (o *okState) tempUp() {
	fmt.Println("Реактор в норме - повышаем температуру")
	o.r.temperature += 100
	o.r.currentState = o.r.tooHot
}

func (o *okState) tempDown() {
	fmt.Println("Реактор в норме - понижаем температуру")
	o.r.temperature -= 100
	o.r.currentState = o.r.tooCold
}

//*******************************
type tooColdState struct {
	r *reactor
}

func (t *tooColdState) tempUp() {
	fmt.Println("Реактор сильно охлажден - повышаем температуру")
	t.r.temperature += 100
	t.r.currentState = t.r.ok
}

func (t *tooColdState) tempDown() {
	fmt.Println("Реактор сильно охлажден - понижаем температуру")
	t.r.temperature -= 100
	t.r.currentState = t.r.off
}

//*******************************
type offState struct {
	r *reactor
}

func (o *offState) tempUp() {
	fmt.Println("Реактор вышел из строя - пытаемся повысить температуру")
	o.r.temperature = 0
	o.r.currentState = o.r.ok
}

func (o *offState) tempDown() {
	fmt.Println("Реактор вышел из строя - пытаемся понизить температуру")
	o.r.temperature = 0
	o.r.currentState = o.r.ok
}

func main() {
	reactor := newReactor()

	reactor.currentState.tempUp()
	reactor.currentState.tempDown()
	reactor.currentState.tempDown()

}
