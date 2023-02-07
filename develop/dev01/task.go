package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"io"
	"os"
	"time"
)

/*
=== Базовая задача ===

Создать программу печатающую точное время с использованием NTP библиотеки.Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и golint.
*/

func getTime() (time.Time, error) {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")

	if err != nil {
		_, errorString := io.WriteString(os.Stderr, err.Error())
		if errorString != nil {
			fmt.Printf("Не удалось записать ошибку в Stderr: %s", errorString)
		}
		return time, err
	}
	return time, nil
}
