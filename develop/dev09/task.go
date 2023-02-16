package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	m := flag.Bool("m", false, "mirror - download site")
	flag.Parse()

	url := flag.Arg(0)

	filename := "download/" + "1.txt"
	if !*m {
		resp, err := http.Get(url)
		if err != nil {
			log.Fatalln(err)
		}
		defer resp.Body.Close()

		file, err := os.Create(filename)
		if err != nil {
			log.Fatalln(err)
		}
		defer file.Close()

		_, err = io.Copy(file, resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
	}

}
