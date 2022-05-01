package main

import "fmt"

type ListaDeCompras []string

func imprimirSlice(slice []string) {
	fmt.Println("Slice: ", slice)
}

func imprimirLista(lista ListaDeCompras) {
	fmt.Println("Lista: ", lista)
}

// Possivelmente era necessário fazer a conversão em versões anteriores do go
func main() {
	lista := ListaDeCompras{"Alface", "Atum", "Azeite"}
	slice := []string{"Alface", "Atum", "Azeite"}

	// imprimirSlice(lista)
	// imprimirSlice(slice)
	// imprimirLista(lista)
	// imprimirLista(slice)
	imprimirSlice([]string(lista))
	imprimirLista(ListaDeCompras(slice))
}

type ListaGenerica []interface{}

func (lista *ListaGenerica) RemoverIndice(indice int) interface{} {
	l := *lista
	removido := l[indice]
	*lista = append(l[0:indice], l[indice+1:]...)
	return removido
}
