package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const CAMINHO string = ""

func lerTxt(nomeArquivo string) string {

	conteudo, _ := ioutil.ReadFile(CAMINHO + nomeArquivo)

	return string(conteudo)
}
func main() {

	linhas := strings.Split(string(lerTxt("NFe.txt")), ";")

	for _, v := range linhas {

		fmt.Println(v)

	}
}
