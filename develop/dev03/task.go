package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	var column int
	var n, r, u bool
	flag.IntVar(&column, "k", 1, "указание колонки для сортировки")
	flag.BoolVar(&n, "n", false, "сортировать по числовому значению")
	flag.BoolVar(&r, "r", false, "сортировать в обратном порядке")
	flag.BoolVar(&u, "u", false, "не выводить повторяющиеся строки")
	flag.Parse()

	var in io.Reader
	if filename := flag.Arg(0); filename != "" {
		f, err := os.Open(filename)
		if err != nil {
			fmt.Println("error opening file: err:", err)
			os.Exit(1)
		}
		defer f.Close()

		in = f
	} else {
		in = os.Stdin
	}

	buf := bufio.NewScanner(in)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	fileContents := [][]string{}
	for buf.Scan() {
		fileContents = append(fileContents, strings.Split(buf.Text(), " "))
	}

	if n {
		sort.Slice(fileContents, func(i, j int) bool {
			a, _ := strconv.Atoi(fileContents[i][column-1])
			b, _ := strconv.Atoi(fileContents[j][column-1])
			return a < b
		})
	} else {
		sort.Slice(fileContents, func(i, j int) bool {
			return fileContents[i][column-1] < fileContents[j][column-1]
		})
	}

	if u {
		withoutDuplicate := [][]string{}
		for i := 0; i < len(fileContents); i++ {
			if i != len(fileContents)-1 &&
				strings.Join(fileContents[i], " ") == strings.Join(fileContents[i+1], " ") {
				continue
			}
			withoutDuplicate = append(withoutDuplicate, fileContents[i])
		}
		fileContents = nil
		fileContents = withoutDuplicate
	}

	if r {
		for i := len(fileContents) - 1; i >= 0; i-- {
			fmt.Fprintln(out, strings.Join(fileContents[i], " "))
		}
	} else {
		for i := 0; i < len(fileContents); i++ {
			fmt.Fprintln(out, strings.Join(fileContents[i], " "))
		}
	}

	if err := buf.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error reading: err:", err)
	}
}
