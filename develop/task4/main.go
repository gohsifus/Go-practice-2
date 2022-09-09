package main

/*
Написать функцию поиска всех множеств анаграмм по словарю.

Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Требования:
	1. Входные данные для функции: ссылка на массив, каждый элемент которого - слово на русском языке в кодировке utf8
	2. Выходные данные: ссылка на мапу множеств анаграмм.
	3. Ключ - первое встретившееся в словаре слово из множества.
	Значение - ссылка на массив, каждый элемент которого, слово из множества.
	4. Массив должен быть отсортирован по возрастанию.
	5. Множества из одного элемента не должны попасть в результат.
	6. Все слова должны быть приведены к нижнему регистру.
	7. В результате каждое слово должно встречаться только один раз.
*/

import (
	"fmt"
	"sort"
	"strings"
)

//strToSet Приведет слово к форме для проверки анаграмм:  Абрикос -> срокиба
//приводит к нижнему регистру и сортирует
func strToSet(str string) string{
	str = strings.ToLower(str)
	runes := []rune(str)
	sort.Slice(runes, func(i, j int) bool{
		return runes[i] > runes[j]
	})
	return string(runes)
}

func getAnagramSet(data *[]string) *map[string]*[]string{
	anagrams := make(map[string]*[]string)

	for _, v := range *data{
		vs := strToSet(v)

		if len(vs) > 1 {
			if _, ok := anagrams[vs]; !ok{
				anagrams[vs] = &[]string{strings.ToLower(v)}
			} else {
				*anagrams[vs] = append(*anagrams[vs], strings.ToLower(v))
			}
		}
	}

	//Меняем ключ на первое слово из множества и сортируем массив
	for k, v := range anagrams {
		anagrams[(*v)[0]] = v
		delete(anagrams, k)
		sort.Strings(*v)
	}

	return &anagrams
}

func main() {
	arr := &[]string{"Привет", "Пятка", "Пятак", "Тяпка", "Кот", "Ток", "Листок", "Слиток", "Столик", "А"}

	anagrams := getAnagramSet(arr)
	for k, v := range *anagrams{
		fmt.Println(fmt.Sprintf("%v: %v", k, v))
	}
}