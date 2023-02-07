package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	testCount := 5

	for i := 0; i < testCount; i++ {
		tm, err := getTime()
		if err != nil {
			log.Println("Ошибка при вызове функции: ", err)
			continue
		}
		fmt.Println(tm)
		time.Sleep(time.Millisecond * 500)
	}
}
