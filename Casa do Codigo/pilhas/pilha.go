package main

import (
	"errors"
	"fmt"
)

type Pilha struct {
	valores []interface{}
}

func (pilha Pilha) Tamanho() int {
	return len(pilha.valores)
}

func (pilha Pilha) Vazia() bool {
	return pilha.Tamanho() == 0
}

func (pilha *Pilha) Empilhar(valor interface{}) {
	pilha.valores = append(pilha.valores, valor)
}

func (pilha *Pilha) Desempilhar() (interface{}, error) {
	if pilha.Vazia() {
		return nil, errors.New("Pilha vazia!")
	}

	valor := pilha.valores[pilha.Tamanho()-1]
	pilha.valores = pilha.valores[:pilha.Tamanho()-1]
	return valor, nil
}

func main() {
	pilha := Pilha{}

	pilha.Empilhar("Go")
	pilha.Empilhar(2022)
	pilha.Empilhar(3.14)
	pilha.Empilhar("Fim")

	fmt.Println("Tamanho da pilha ", pilha.Tamanho())

	for !pilha.Vazia() {
		valor, _ := pilha.Desempilhar()
		fmt.Println("Desempilhando", valor)
		fmt.Println("Tamanho: ", pilha.Tamanho())
		fmt.Println("Vazia?", pilha.Vazia())
	}
}
