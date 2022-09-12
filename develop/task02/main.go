package main

/*
Создать Go-функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы/руны, например:
	● "a4bc2d5e" => "aaaabccddddde"
	● "abcd" => "abcd"
	● "45" => "" (некорректная строка)
	● "" => ""
Дополнительно:
Реализовать поддержку escape-последовательностей.
Например:
	● qwe\4\5 => qwe45 (*)
	● qwe\45 => qwe44444 (*)
	● qwe\\5 => qwe\\\\\ (*)
В случае если была передана некорректная строка, функция должна возвращать ошибку.
Написать unit-тесты.
*/

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

//ProcessEscape вызывается для обработки escape последовательности,
//возвращает обработанную последовательность для записи и изменяет указатель на количество обработанных символов
func ProcessEscape(runes []rune, point *int) string {
	var ret string

	if len(runes) >= 2 && unicode.IsDigit(runes[1]) {
		//Обработано 2 символа
		*point += 2
		// "\*number" записываем сивол * number раз
		ret = strings.Repeat(string(runes[0]), parseRuneToInt(runes[1]))
	} else {
		//Обработано 1 символ
		*point++
		// "\*" записываем символ * 1 раз
		ret = string(runes[0])
	}

	return ret
}

//unpack вызывается для распаковки части строки
//возвращает обработанную последовательность для записи и изменяет указатель на количество обработанных символов
func unpack(runes []rune, point *int) string {
	return strings.Repeat(string(runes[0]), parseRuneToInt(runes[1]) - 1)
}

func parseRuneToInt(r rune) int {
	//Тут игнорируем обработку чтобы не усложнять код и потому-что всегда проверяем с помощью isDigit
	num, _ := strconv.Atoi(string(r))
	return num
}

func unpackWithEscape(inp string) (string, error) {
	unpacked := strings.Builder{}
	runes := []rune(inp)

	if len(runes) != 0 && unicode.IsDigit(runes[0]){
		return "", fmt.Errorf("некорректные данные")
	}

	for i := 0; i < len(runes); i++ {
		if runes[i] == '\\' {
			//передаем для обработки escape
			unpacked.WriteString(ProcessEscape(runes[i+1:], &i))
		} else {
			if unicode.IsDigit(runes[i]) {
				//Передаем для распаковки
				unpacked.WriteString(unpack(runes[i-1:], &i))
			} else {
				//Просто записываем
				unpacked.WriteRune(runes[i])
			}
		}
	}

	return unpacked.String(), nil
}

func main() {
	str := `a4bc2d5e\34`
	fmt.Println(unpackWithEscape(str))
}
