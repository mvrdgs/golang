package main

import "fmt"

// closure é quando uma função manipula o contexto onde ela foi gerada, no caso a e b que foram geradas fora de seu escopo
func main() {
	a, b := 0, 1

	fib := func() int {
		a, b = b, a+b

		return a
	}

	for i := 0; i < 8; i++ {
		fmt.Printf("%d ", fib())
	}
}
