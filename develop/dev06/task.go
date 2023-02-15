package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func setFlags(f *int, d *string, s *bool) {
	flag.IntVar(f, "f", 0, "выбрать поля (колонки)")
	flag.StringVar(d, "d", "\t", "спользовать другой разделитель")
	flag.BoolVar(s, "s", false, "только строки с разделителем")
}

func Cut(str, d *string, f *int, s *bool) string {
	if *s && !strings.Contains(*str, *d) {
		return ""
	}

	strArr := strings.Split(*str, *d)

	if *f > len(strArr) {
		return *str + "\n"
	} else {
		return strArr[*f-1] + "\n"
	}
}

func main() {

	var f int
	var d string
	var s bool

	setFlags(&f, &d, &s)
	flag.Parse()

	if f == 0 {
		log.Fatalln("cut: you must specify a list of bytes, characters, or fields")
	}

	for {
		buf, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			break
		}

		str := strings.Trim(buf, "\n")
		res := Cut(&str, &d, &f, &s)
		fmt.Print(res)
	}

}
