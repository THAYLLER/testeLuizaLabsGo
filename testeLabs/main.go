package main

import (
	"fmt"
	"io/ioutil"
)

const CAMINHO string = ""

func lerTxt(nomeArquivo string) string {

	conteudo, _ := ioutil.ReadFile(CAMINHO + nomeArquivo)

	return string(conteudo)
}
func main() {

	fmt.Print(lerTxt("NFe.txt"))
}
