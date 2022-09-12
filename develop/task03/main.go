package main

/*
Отсортировать строки в файле по аналогии с консольной утилитой sort: на входе подается файл из несортированными строками,
на выходе — файл с отсортированными.

Реализовать поддержку утилитой следующих ключей:
	-k — указание колонки для сортировки (слова в строке могут выступать в качестве колонок, по умолчанию разделитель — пробел)
	-n — сортировать по числовому значению
	-r — сортировать в обратном порядке
	-u — не выводить повторяющиеся строки
*/

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getDataForSort(path string) ([]string, error) {
	dataFromFile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return strings.Split(string(dataFromFile), "\n"), nil
}

//SortWithNFlag ...
func SortWithNFlag(data []string) []string {
	//Проверяем строки, если начинается с числа то парсим число и сравниваем по нему, если не число то строка 0
	sort.Slice(data, func(i, j int) bool {
		numberA := strings.Builder{}
		numberB := strings.Builder{}

		//Парсим первые цифры в строке - в число
		for _, v := range data[i] {
			//Если руна число
			if v >= 48 && v <= 57 {
				numberA.WriteRune(v)
			} else {
				break
			}
		}

		for _, v := range data[j] {
			//Если руна число
			if v >= 48 && v <= 57 {
				numberB.WriteRune(v)
			} else {
				break
			}
		}

		a, _ := strconv.Atoi(numberA.String())
		b, _ := strconv.Atoi(numberB.String())

		if a == b {
			slice := []string{
				data[i],
				data[j],
			}

			sort.Strings(slice)

			return data[i] == slice[0]
		}

		return a < b
	})

	return data
}

//SortWithKFlag ...
func SortWithKFlag(data []string, k int, n bool) []string {
	k-- //Поскольку пользователь передает k с единицы
	//Сортируем начиная с слова переданного в флаге k
	sort.Slice(data, func(i int, j int) bool {
		//Делим строки на колонки
		a := strings.Split(data[i], " ")
		b := strings.Split(data[j], " ")

		//Часть строки с слова с номером k
		var ak string
		if len(a)-1 < k {
			//Если строка короче чем k столбцов, сортируем как обычно, но такие строки выше
			//чем те которые сортируются по k
			ak = "!" + data[i]
		} else {
			ak = strings.Join(a[k:], " ")
		}
		//Часть строки с слова с номером k
		var bk string
		if len(b)-1 < k {
			//Если строка короче чем k столбцов, сортируем как обычно, но такие строки выше
			//чем те которые сортируются по k
			bk = "!" + data[j]
		} else {
			bk = strings.Join(b[k:], " ")
		}

		slice := []string{
			ak,
			bk,
		}
		//Сортируем так, как будто строки начинаются со слова k
		if !n {
			sort.Strings(slice)
		} else {
			slice = SortWithNFlag(slice)
		}

		return ak == slice[0] //i > j
	})

	return data
}

//RemoveDuplicates удалит повторяющиеся строки: флаг -u
func RemoveDuplicates(data []string) []string {
	m := make(map[string]struct{})
	ret := []string{}
	for _, v := range data {
		if _, ok := m[v]; !ok {
			m[v] = struct{}{}
			ret = append(ret, v)
		}
	}

	return ret
}

//Reverse перевернет массив: -r флаг
func Reverse(data []string) []string {
	end := len(data) - 1
	for i := 0; i < len(data)/2; i++ {
		data[i], data[end] = data[end], data[i]
		end--
	}

	return data
}

//WriteToFile ...
func WriteToFile(path string, data []string) error {
	dataForFile := []string{}

	for _, v := range data {
		dataForFile = append(dataForFile, v)
	}

	bytes := []byte(strings.Join(dataForFile, "\n"))

	err := ioutil.WriteFile(path, bytes, 0660)
	if err != nil {
		return err
	}

	return nil
}

//Sort ...
func Sort(data []string, rFlag bool, kFlag int, nFlag bool, uFlag bool) []string {
	if uFlag {
		data = RemoveDuplicates(data)
	}

	if kFlag == 1 && !nFlag {
		sort.Strings(data)
	} else {
		data = SortWithKFlag(data, kFlag, nFlag)
	}

	if rFlag {
		data = Reverse(data)
	}

	return data
}

func main() {
	path := flag.String("path", "", "file path to sort")
	r := flag.Bool("r", false, "reverse sort")
	k := flag.Int("k", 1, "num of column for sort")
	n := flag.Bool("n", false, "num sort")
	u := flag.Bool("u", false, "remove non-uniques elements")
	flag.Parse()

	data, err := getDataForSort(*path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	data = Sort(data, *r, *k, *n, *u)

	WriteToFile(*path, data)
}
