package main

/*
Реализовать утилиту аналог консольной команды cut (man cut).
Утилита должна принимать строки через STDIN, разбивать по разделителю (TAB) на колонки и выводить запрошенные.

Реализовать поддержку утилитой следующих ключей:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем
*/

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func parseFieldsFlag(flag string) ([]int, error) {
	ret := []int{}

	fields := strings.Split(flag, ",")
	for _, v := range fields {
		v = strings.TrimSpace(v)
		field, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}

		ret = append(ret, field)
	}
	sort.Ints(ret)
	return ret, nil
}

//splitString делит строку на массив строк по разделителю
//также вернет true если разделителя в строке нет
func splitString(inp, sep string) ([]string, bool) {
	splited := strings.Split(inp, sep)
	return splited, splited[0] == inp
}

//Cut ...
func Cut(data string, dFlag string, fFlag []int, sFlag bool) string {
	if dFlag == "" {
		dFlag = "\t"
	}
	builder := []string{}
	splited, withoutDelimetr := splitString(data, dFlag)

	if sFlag && withoutDelimetr {
		return ""
	}
	for _, v := range fFlag {
		if len(splited) > v-1 {
			builder = append(builder, splited[v-1])
		}
	}

	return strings.Join(builder, " ")
}

func main() {
	delimiter := flag.String("d", "\t", "Разделитель")
	fFlag := flag.String("f", "", "Поля для выбора через ',' ")
	separated := flag.Bool("s", false, "Только строки с разделителем")

	flag.Parse()

	fields, err := parseFieldsFlag(*fFlag)
	if err != nil {
		fmt.Println(fmt.Sprintf("Некоррктные флаги: %s", err.Error()))
		os.Exit(1)
	}

	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() {
		line := sc.Text()

		fmt.Println(Cut(line, *delimiter, fields, *separated))
		fmt.Println()
	}
}
