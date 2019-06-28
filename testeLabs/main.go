package main

import (
	"bufio"
	"fmt"
	"os"
)

const CAMINHO string = ""

func lerTxt(nomeArquivo string) ([]string, error) {

	file, err := os.Open(CAMINHO + nomeArquivo)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var literalLines []string

	scanner := bufio.NewScanner(file)

	// guarda cada linha em indice diferente do slice
	for scanner.Scan() {
		literalLines = append(literalLines, scanner.Text())
	}

	return literalLines, scanner.Err()
}

func main() {

	var linha []string

	linha = nil
	linha, _ = lerTxt("NFe.txt")

	for _, line := range linha {

		if line != "Data;Tipo;CnpjCpf;Numero;Serie;Modelo;Chave;ValorTotal;ValorProd;ValorICMS;ValorIPI;Status" {
			fmt.Println(line)
		}
	}
	/*
		transformar o conteudo do txt em json e  criar as rotas
		estruturar melhor o projeto
	*/
	//router := mux.NewRouter()
	//router.HandleFunc("/contato", GetPeople).Methods("GET")
}
