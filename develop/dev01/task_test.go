package main

import (
	"fmt"
	"testing"
	"time"
)

func TestGetTime(t *testing.T) {
	testCount := 5

	for i := 0; i < testCount; i++ {
		tm, err := getTime()
		if err != nil {
			t.Errorf("Ошибка при вызове функции: %s", err)
			continue
		}
		fmt.Println(tm)
		time.Sleep(time.Millisecond * 500)
	}
}
