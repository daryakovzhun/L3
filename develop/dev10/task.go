package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

func main() {
	timeout := flag.String("timeout", "10s", "timeout")
	flag.Parse()

	if len(flag.Args()) == 0 {
		fmt.Println("Not enough arguments")
		return
	}

	exitCh := make(chan os.Signal, 1)
	signal.Notify(exitCh, syscall.SIGQUIT, syscall.SIGTERM)
	go Exit(exitCh)

	fmt.Println(*timeout)

	for {

	}

}

func Exit(exitCh chan os.Signal) {
	for {
		select {
		case <-exitCh:
			fmt.Println("Press Ctrl+D\nExit")
			os.Exit(0)
		default:
			fmt.Println("Hello")
		}
	}
}
