package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	buf := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(">")
		cmd, err := buf.ReadString('\n')
		cmd = strings.TrimSuffix(cmd, "\n")
		cmd = strings.TrimSpace(cmd)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		if err = command(cmd); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func command(cmd string) error {
	cmdArr := strings.Split(cmd, " ")
	switch cmdArr[0] {
	case "cd":
	case "echo":
	case "pwd":
	case "kill":
	case "ps":
	case "quit":
		os.Exit(0)
	}
	return nil
}
