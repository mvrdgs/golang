package main

import "fmt"

type Arquivo struct {
	nome       string
	tamanho    float64
	caracteres int
	palavras   int
	linhas     int
}

func (arquivo *Arquivo) TamanhoMedioDaPalavra() float64 {
	return float64(arquivo.caracteres) / float64(arquivo.palavras)
}

func (arquivo *Arquivo) MediaDePalavrasPorLinha() float64 {
	return float64(arquivo.palavras) / float64(arquivo.linhas)
}

func main() {
	arquivo := Arquivo{"artigo,txt", 12.68, 12986, 1862, 220}

	fmt.Println(arquivo)
	fmt.Printf("Média de palavras por linha: %.2f\n", arquivo.MediaDePalavrasPorLinha())
	fmt.Printf("Tamanho médio de palavra: %.2f\n", arquivo.TamanhoMedioDaPalavra())
}
