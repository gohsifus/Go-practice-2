package main

import (
	"fmt"
	"time"
)

/*
Реализовать функцию, которая будет объединять один или более done-каналов в single-канал,
если один из его составляющих каналов закроется.

Определение функции:
var or func(channels ...<- chan interface{}) <- chan interface{}

Пример использования функции:
	sig := func(after time.Duration) <- chan interface{} {
	c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or (
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	fmt.Printf(“fone after %v”, time.Since(start))
*/

func sig(after time.Duration) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
	}()
	return c
}

func or(channels ...<-chan interface{}) <-chan interface{} {
	out := make(chan interface{})
	cl := make(chan int)
	quit := false

	for _, c := range channels {
		go func(c <-chan interface{}) {
			for !quit {
				select {
				//Если один из каналов закрылся
				case _, ok := <-cl:
					if !ok {
						break
					}
				case data, ok := <-c:
					//Если канал закрылся
					if !ok {
						//Уведомляем другие рутины
						quit = true
						close(cl)
						//Записываем значение в out и выходим
						out <- data
						break
					}
					//Если канал не закрыт кладем в канал
					out <- data
				}
			}
		}(c)
	}

	go func() {
		//Ждем в отдельной горутине, пока один из каналов не закроется
		for range cl {
		}
		close(out)
	}()

	return out
}

func main() {
	start := time.Now()
	<-or(
		sig(20*time.Second),
		sig(10*time.Second),
		sig(30*time.Second),
		sig(40*time.Second),
		sig(1*time.Second),
	)

	fmt.Printf("fone after %v", time.Since(start))
}
