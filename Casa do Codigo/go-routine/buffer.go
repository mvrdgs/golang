package main

import "fmt"

// c chan <- int controla o fluxo, só é possível receber valores
// c <-chan int possibilita apenas ler valores
func produzir(c chan<- int) {
	c <- 1
	c <- 2
	c <- 3

	close(c)
}

func main() {
	c := make(chan int, 3)
	go produzir(c)

	// for {
	// 	valor, ok := <-c

	// 	if ok {
	// 		fmt.Println(valor)
	// 	} else {
	// 		break
	// 	}
	// }

	for valor := range c {
		fmt.Println(valor)
	}
}
