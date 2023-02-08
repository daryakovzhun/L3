package main

import (
	"errors"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func Unpacking(s string) (string, error) {
	if len(s) == 0 {
		return s, nil
	}

	var res string
	for i := 0; i < len(s); i++ {
		if unicode.IsDigit(rune(s[i])) {
			return "", errors.New("некорректная строка")
		}
		if s[i] == '\\' {
			if i != len(s)-2 && unicode.IsDigit(rune(s[i+2])) {
				toDuplicate(&s, &res, i+2)
				i += 2
			} else {
				res += string(s[i+1])
				i++
			}
		} else {
			if i != len(s)-1 && unicode.IsDigit(rune(s[i+1])) {
				toDuplicate(&s, &res, i+1)
				i++
			} else {
				res += string(s[i])
			}
		}
	}
	return res, nil
}

func toDuplicate(s, res *string, k int) {
	count := int((*s)[k] - '0')
	for j := 0; j < count; j++ {
		*res += string((*s)[k-1])
	}
}
