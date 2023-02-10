package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// MacOS and Linux
func setFlags(NA, NB, NC *int, c, i, v, F, n *bool) {
	flag.IntVar(NA, "A", 0, "печатать +N строк после совпадения")
	flag.IntVar(NB, "B", 0, "печатать +N строк до совпадения")
	flag.IntVar(NC, "C", 0, "(A+B) печатать ±N строк вокруг совпадения")
	flag.BoolVar(c, "c", false, "количество строк")
	flag.BoolVar(i, "i", false, "игнорировать регистр")
	flag.BoolVar(v, "v", false, "вместо совпадения, исключать")
	flag.BoolVar(F, "F", false, "точное совпадение со строкой, не паттерн")
	flag.BoolVar(n, "n", false, "печатать номер строки")
}

func saveNLines(beforeMatchedLines []string, NB *int, line string) []string {
	if *NB != 0 {
		if len(beforeMatchedLines) == *NB+1 {
			beforeMatchedLines = append(beforeMatchedLines[:0], beforeMatchedLines[1:]...)
		}
		if len(beforeMatchedLines) < *NB+1 {
			beforeMatchedLines = append(beforeMatchedLines, line)
		}
	}

	return beforeMatchedLines
}

func separator(firstMatched, beforeMatched *bool, afterMatched *int, out *bufio.Writer) {
	if !*firstMatched && *afterMatched == 0 {
		fmt.Fprintln(out, "--")
	}

	if !*firstMatched && !*beforeMatched {
		fmt.Fprintln(out, "--")
	}
}

func main() {
	var NA, NB, NC, afterMatched, countMatched, numberLine int
	var c, i, v, F, n, firstMatched, beforeMatched bool
	firstMatched = true
	var withoutRegister string
	beforeMatchedLines := []string{}

	setFlags(&NA, &NB, &NC, &c, &i, &v, &F, &n)
	flag.Parse()

	var in io.Reader
	if filename := flag.Arg(1); filename != "" {
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

	search := flag.Arg(0)

	if NC > 0 {
		NA, NB = NC, NC
	}

	if i {
		withoutRegister += "(?i)"
	}

	for buf.Scan() {
		numberLine++

		var matched bool

		if F {
			matched = strings.Contains(buf.Text(), search)
		} else {
			matched, _ = regexp.MatchString(withoutRegister+search, buf.Text())
		}

		beforeMatchedLines = saveNLines(beforeMatchedLines, &NB, buf.Text())

		if v {
			matched = !matched
		}

		if matched {
			if c {
				countMatched++
				continue
			}

			if NA > 0 || NB > 0 {
				separator(&firstMatched, &beforeMatched, &afterMatched, out)
			}

			if n {
				fmt.Fprint(out, numberLine, ":")
			}

			if NB != 0 && !beforeMatched {
				for _, val := range beforeMatchedLines {
					fmt.Fprintln(out, val)
				}
			} else {
				fmt.Fprintln(out, buf.Text())
			}

			afterMatched = NA
			firstMatched = false
			beforeMatched = true
			beforeMatchedLines = nil
		} else if afterMatched > 0 {
			fmt.Fprintln(out, buf.Text())
			afterMatched--
			beforeMatched = false
		} else {
			beforeMatched = false
		}

	}

	if c {
		fmt.Fprintln(out, countMatched)
	}

}
