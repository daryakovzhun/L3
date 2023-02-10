package main

import (
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func FindAnagrams(words *[]string) *map[string][]string {
	res := make(map[string][]string)
	wordsSortKey := make(map[string]string)

	for _, word := range *words {
		originW := strings.ToLower(word)
		sortW := []rune(originW)
		sort.Slice(sortW, func(i, j int) bool {
			return sortW[i] < sortW[j]
		})

		if val, ex := wordsSortKey[string(sortW)]; ex {
			res[val] = append(res[val], originW)
		} else {
			wordsSortKey[string(sortW)] = originW
			res[originW] = []string{}
		}
	}

	for key, val := range res {
		if len(val) == 0 {
			delete(res, key)
			continue
		}
		sort.Strings(res[key])
	}

	return &res
}

//func main() {
//	x := []string{"листок", "пятак", "пятка", "столик", "слиток", "тяпка"}
//	y := FindAnagrams(&x)
//	fmt.Println(*y)
//}
