package main

import "fmt"

func PrecoFinal(precoCusto float64) (
	precoDolar float64,
	precoReal float64,
) {
	fatorLucro := 1.33
	taxaConversao := 2.34

	precoDolar = precoCusto * fatorLucro
	precoReal = precoDolar * taxaConversao

	// as variáveis não precisam ser especificadas no retorno já que foram declaradas como retorno na declaração
	// porém é uma boa prática para manter a legibilidade
	return precoDolar, precoReal
}

func main() {
	precoDolar, precoReal := PrecoFinal(34.99)

	fmt.Printf("Preço final em dólar: %.2f\nPreço final em reais: %.2f\n", precoDolar, precoReal)
}
