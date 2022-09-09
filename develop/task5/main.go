package main

/*
Реализовать утилиту фильтрации по аналогии с консольной утилитой (man grep — смотрим описание и основные параметры).
Реализовать поддержку утилитой следующих ключей:
	-A - "after" печатать +N строк после совпадения
	-B - "before" печатать +N строк до совпадения
	-C - "context" (A+B) печатать ±N строк вокруг совпадения
	-c - "count" (количество строк)
	-i - "ignore-case" (игнорировать регистр)
	-v - "invert" (вместо совпадения, исключать)
	-F - "fixed", точное совпадение со строкой, не паттерн
	-n - "line num", напечатать номер строки
*/

import (
	"flag"
	"fmt"
	console "grepUtil/colorConsole"
	"os"
	"regexp"
	"sort"
	"strings"
)

//getDataForSearch возвращает данные для поиска
func getDataForSearch(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return strings.Split(string(data), "\n"), nil
}

//prepareTarget подготавливает выражение поиска (F, i)
func prepareTarget(target string, F, i bool) string {
	if F {
		//F - fixedFlag точное совпадение а не регулярка
		target = regexp.QuoteMeta(target)
	}

	if i {
		//i - iFlag без учета регистра
		target = "(?i)" + target
	}

	return target
}

//grep (A, B, C, v)
func grep(data []string, target string, A, B, C int, c, i, v, F, n bool) map[int]string {
	found := make(map[int]string)
	numOfFound := -1 //Количество найденных, для флага c

	target = prepareTarget(target, F, i)

	regExp := regexp.MustCompile(target)

	for k, line := range data {
		match := regExp.Match([]byte(line))

		//Если строка удовлетворяет шаблону
		if match && !v {
			//Выбеляем найденные элементы
			result := regExp.ReplaceAllStringFunc(line, func(m string) string {
				return console.GetColorizeString(m)
			})

			//B и A имеют приоритет над C
			if B <= 0 {
				//Если B не задан используем C
				B = C
			}
			if A <= 0 {
				//Если A не задан используем C
				A = C
			}

			//Before Печатаем n строк перед найденной
			if B > 0 {
				for index := 1; index <= B; index++ {
					//Побочные строки не должны перезаписывать старые чтобы не стирать раскраску
					if _, ok := found[k-index]; !ok && k-index >= 0 {
						found[k-index] = data[k-index]
					}
				}
			}
			found[k] = result
			if c {
				numOfFound++
			}
			//After Печатаем n строк после найденной
			if A > 0 {
				for index := 1; index <= A; index++ {
					//Побочные строки не должны перезаписывать старые чтобы не стирать раскраску
					if _, ok := found[k+index]; !ok && k+index < len(data) {
						found[k+index] = data[k+index]
					}
				}
			}
		} else if v && !match {
			//Если инвертируем то записываем все кроме найденных
			found[k] = data[k]
			if c {
				numOfFound++
			}
		}
	}

	showResult(found, n, numOfFound)

	return found
}

//showResult распечатает результат в stdOut (n, c)
func showResult(m map[int]string, nFlag bool, numOfFound int) {
	if numOfFound >= 0 {
		fmt.Println(numOfFound + 1)
	} else {
		slice := []int{}
		for k := range m {
			slice = append(slice, k)
		}

		sort.Ints(slice)

		for _, v := range slice {
			if nFlag {
				fmt.Println(fmt.Sprintf("%v: %v", v+1, m[v]))
			} else {
				fmt.Println(m[v])
			}
		}
	}
}

//parseArgs распарсит аргументы не представленные флагами (пути к файлам и что искать)
func parseArgs(args *[]string) ([]string, string) {
	files := []string{}
	target := os.Args[1] //Выражение поиска всегда 1

	for k, v := range os.Args[2:] {
		if v == "-A" ||
			v == "-B" ||
			v == "-C" ||
			v == "-c" ||
			v == "-i" ||
			v == "-v" ||
			v == "-F" ||
			v == "-n" {
			//Так как flag перестает парсить флаги после аргументов без флага,
			//Меняем os.Args
			*args = append(os.Args[:1], os.Args[k+2:]...)
			break
		}
		files = append(files, v)
	}

	return files, target
}

func main() {
	files, target := parseArgs(&os.Args)

	AFlag := flag.Int("A", 0, "")
	BFlag := flag.Int("B", 0, "")
	CFlag := flag.Int("C", 0, "")
	cFlag := flag.Bool("c", false, "")
	iFlag := flag.Bool("i", false, "")
	vFlag := flag.Bool("v", false, "")
	fFlag := flag.Bool("F", false, "")
	nFlag := flag.Bool("n", false, "")

	flag.Parse()

	for _, path := range files {
		lines, err := getDataForSearch(path)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(path + ": ")
		grep(lines, target, *AFlag, *BFlag, *CFlag, *cFlag, *iFlag, *vFlag, *fFlag, *nFlag)
		fmt.Println()
	}
}
